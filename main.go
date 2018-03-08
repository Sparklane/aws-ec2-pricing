package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cristim/ec2-instances-info"
	"github.com/namsral/flag"
)

var ec2Conn *ec2.EC2
var region string

type instancePrice struct {
	Ondemand float64            `json:"ondemand"`
	Spot     map[string]float64 `json:"spot"`
}

func main() {
	var instanceType string
	var formated bool

	// Params
	flag.StringVar(&instanceType, "instance-type", "", "(Optional) Instance type")
	flag.BoolVar(&formated, "formated", false, "(Optional) if set, it will format output")
	flag.Parse()

	// Init EC2 connexion
	sess := session.Must(session.NewSession())
	ec2Conn = ec2.New(sess)
	region = ec2Conn.SigningRegion

	// Get OnDemand info
	data, err := ec2instancesinfo.Data()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// For each instance get the spot prices
	instances := make(map[string]instancePrice)
	for _, i := range *data {
		// Filter out other instance type "instance-type" param is set
		if instanceType != "" && instanceType != i.InstanceType {
			continue
		}

		instances[i.InstanceType] = instancePrice{
			Ondemand: i.Pricing[region].Linux.OnDemand,
			Spot:     map[string]float64{},
		}

		// get current instance spot price
		spotPrices := getCurrentSpotPrice(i.InstanceType)
		for _, spotPrice := range spotPrices {
			price, err := strconv.ParseFloat(*spotPrice.SpotPrice, 64)
			if err != nil {
				log.Panic(err)
			}
			instances[i.InstanceType].Spot[*spotPrice.AvailabilityZone] = price
		}
	}

	// Print JSON response
	if instanceType != "" {
		fmt.Println(printJSON(instances[instanceType], formated))
	} else {
		fmt.Println(printJSON(instances, formated))
	}

}

// get current instance spot price
func getCurrentSpotPrice(instanceType string) []*ec2.SpotPrice {

	params := &ec2.DescribeSpotPriceHistoryInput{
		ProductDescriptions: []*string{
			aws.String("Linux/UNIX (Amazon VPC)"),
		},
		InstanceTypes: []*string{
			aws.String(instanceType),
		},
		StartTime: aws.Time(time.Now().Add(-1 * 1)),
		EndTime:   aws.Time(time.Now()),
	}

	resp, err := ec2Conn.DescribeSpotPriceHistory(params)

	if err != nil {
		log.Panic("Failed requesting spot prices:", err.Error())
	}

	return resp.SpotPriceHistory
}

// print object in json
func printJSON(object interface{}, formated bool) string {
	var err error
	var j []byte

	if !formated {
		j, err = json.Marshal(object)
	} else {
		j, err = json.MarshalIndent(object, "", "   ")
	}

	if err != nil {
		log.Panic(err)
	}
	return string(j)
}
