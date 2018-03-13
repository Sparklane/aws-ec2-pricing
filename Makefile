.PHONY: all
VERSION ?= latest
APP ?= aws-ec2-pricing
NAME ?= $(APP)
SHELL=/bin/bash -O extglob -c
SSH_PRIVATE_KEY ?= ~/.ssh/id_rsa

all: build publish

# ### Lock dependencies ###
# This is used at dev time for building glide.lock file
# Must be run whenever you need to refresh dependecies
lock-dep:
	@go get github.com/ngdinhtoan/glide-cleanup
	@glide cleanup \
		&& glide update

# ### Build ####
build:
	@docker build -t sparklane/$(NAME) .
