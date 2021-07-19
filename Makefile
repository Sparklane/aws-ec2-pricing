NAME = aws-ec2-pricing
VERSION ?= latest
SSH_PRIVATE_KEY ?= ~/.ssh/id_rsa
REGISTRY ?= ""

all: build

build:
	@docker build -t $(NAME):$(VERSION) \
	--rm=true .

publish: build
	@docker tag $(NAME):$(VERSION) $(REGISTRY)/$(NAME):$(VERSION)
	@docker push $(REGISTRY)/$(NAME):$(VERSION)
