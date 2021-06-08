.POSIX:

.PHONY: all
all: install build e2e

.PHONY: install
install:
	command -v terraform >/dev/null 2>&1 || GO111MODULE="on" go get github.com/hashicorp/terraform@v0.15.5
	command -v terraform-docs >/dev/null 2>&1 || GO111MODULE="on" go get github.com/segmentio/terraform-docs@v0.14.1
	command -v tfsec >/dev/null 2>&1 || GO111MODULE="on" go get github.com/liamg/tfsec/cmd/tfsec@v0.39.42
	command -v tflint >/dev/null 2>&1 || GO111MODULE="on" go get github.com/terraform-linters/tflint@v0.29.0

.PHONY: build
build:
	rake build

.PHONY: e2e
e2e:
	rake e2e