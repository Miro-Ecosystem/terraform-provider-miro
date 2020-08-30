version = 99.0.0
provider_macos_path = registry.terraform.io/Miro-Ecosystem/miro/$(version)/darwin_amd64/

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test ./... -v -timeout 120s

testacc: fmtcheck
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120s

cassettes: fmtcheck
	RECORD=true TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

install_macos:
	@go build -o terraform-provider-miro_$(version)
	@mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
	@mv terraform-provider-miro_$(version)  ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)

.PHONY: build test testacc fmt cassettes vet fmtcheck errcheck install_macos
