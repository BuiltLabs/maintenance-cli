VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: dev

tools:
	go get -u github.com/aws/aws-sdk-go/service/dynamodb
	go get -u gopkg.in/urfave/cli.v1

# bin generates all binaries
bin: fmtcheck
	sh -c "'$(CURDIR)/scripts/build.sh'"

# dev creates binaries for testing locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: fmtcheck
	@GO_DEV=1 sh -c "'$(CURDIR)/scripts/build.sh'"


# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: bin