# Makefile for inflation
#
# TODO Limitations
# fmt and validate currently require valid AWS logged-in, since setting
# AWS_ROOT_ACCOUNT and ACCOUNTS use aws cli - even thought not used - my
# Makefile-fu not good enough to fix RN - see $(findstring) and ifneq ...

SHELL := /bin/bash
AWS_REGION ?= ap-southeast-2
BUCKET ?= cconf
SAML2AWS_BIN ?= saml2aws
AWS_PROFILE ?= default
IDP_ACCOUNT ?= default
DOCKER_IMAGE_TERRAFORM = 566134440840.dkr.ecr.ap-southeast-2.amazonaws.com/terraform-provider-cloudconformity:d66c3844056f0465ccd4cc3b7c38967e40a7aacb
AWS_PROVIDER ?= aws


CONTAINER_NAME = "terraform-$(shell date '+%Y-%m-%d-%H.%m.%S')"
AWS_ACCESS_KEY_ID ?= $(shell $(SAML2AWS_BIN) script --idp-account=$(IDP_ACCOUNT) --profile=$(AWS_PROFILE) | grep 'AWS_ACCESS_KEY_ID=".*"' | cut -d\" -f2)
AWS_SECRET_ACCESS_KEY ?= $(shell $(SAML2AWS_BIN) script --idp-account=$(IDP_ACCOUNT) --profile=$(AWS_PROFILE)| grep 'AWS_SECRET_ACCESS_KEY=".*"' | cut -d\" -f2)
AWS_SESSION_TOKEN ?= $(shell $(SAML2AWS_BIN) script --idp-account=$(IDP_ACCOUNT) --profile=$(AWS_PROFILE)| grep 'AWS_SESSION_TOKEN=".*"' | cut -d\" -f2)
AWS_SECURITY_TOKEN ?= $(shell $(SAML2AWS_BIN) script --idp-account=$(IDP_ACCOUNT) --profile=$(AWS_PROFILE)| grep 'AWS_SECURITY_TOKEN=".*"' | cut -d\" -f2)

TF_FOLDER ?= $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)

# Append these to Terraform flags
TERRAFORM_FLAGS := -v $(TF_FOLDER):/usr/src --workdir /usr/src/test \
	-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
	-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
	-e AWS_SESSION_TOKEN=${AWS_SESSION_TOKEN} \
	-e AWS_SECURITY_TOKEN=${AWS_SECURITY_TOKEN}

# if this session isn't interactive, then we don't want to allocate a
# TTY, which would fail, but if it is interactive, we do want to attach
# so that the user can send e.g. ^C through.
INTERACTIVE := $(shell [ -t 0 ] && echo 1 || echo 0)
ifeq ($(INTERACTIVE), 1)
	DOCKER_FLAGS += -t
endif

.PHONY: aws-init
aws-init:
	docker run --rm -i $(DOCKER_FLAGS) \
		--name $(CONTAINER_NAME) \
		${TERRAFORM_FLAGS} \
		$(DOCKER_IMAGE_TERRAFORM) \
		init \
		-backend-config "bucket=$(BUCKET)" \
		-backend-config "region=$(AWS_REGION)" \

.PHONY: aws-plan
aws-plan:
		docker run \
			--rm \
			-i $(DOCKER_FLAGS) \
			$(TERRAFORM_FLAGS) \
			$(DOCKER_IMAGE_TERRAFORM) \
			plan \

.PHONY: aws-apply
aws-apply:
		docker run \
			--rm \
			-i $(DOCKER_FLAGS) \
			$(TERRAFORM_FLAGS) \
			$(DOCKER_IMAGE_TERRAFORM) \
			apply \
