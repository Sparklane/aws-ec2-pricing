NAME = aws-ec2-pricing
VERSION ?= latest
SSH_PRIVATE_KEY ?= ~/.ssh/id_rsa
REGISTRY ?= ""

all: build publish

# ### Lock dependencies ###
# This is used at dev time for building glide.lock file
# Must be run whenever you need to refresh dependecies
lock-dep:
	@go get github.com/ngdinhtoan/glide
	@go get github.com/ngdinhtoan/glide-cleanup
	@~/go/bin/glide-cleanup && ~/go/bin/glide update

build:
	@docker build -t $(NAME):$(VERSION) \
	--rm=true .

publish: build
	@docker tag $(NAME):$(VERSION) $(REGISTRY)/$(NAME):$(VERSION)
	@docker push $(REGISTRY)/$(NAME):$(VERSION)
