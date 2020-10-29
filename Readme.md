# AWS EC2 Pricing

Will get current spot and ondemand prices for the current AWS Region

## build
```bash
make build
```

## publish
```bash
REGISTRY="..." make publish
```

Usage: 
* `formated`: (Optional) if set, it will format output
* `instance-type` string (Optional) Instance type (ex: m3.large)


Example for getting all prices:

```
docker run -it \
  -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
  -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
  -e AWS_REGION=eu-west-1 \
  aws-ec2-pricing --instance-type m4.xlarge
```

Returns: 

```
{
  "ondemand": 0.222,
  "spot": {
    "eu-west-1a": 0.0642,
    "eu-west-1b": 0.0642,
    "eu-west-1c": 0.0642
  }
}
```


Example for getting prices for one instance tyes with a formated output:
```
docker run -it aws-ec2-pricing --instance-type m4.xlarge --formated
```
