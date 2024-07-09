# Initialize the environment
init: generate
	go mod tidy
	go mod vendor

generate: generate_api \
	generate_mocks 

# Generates OpenAPI interface file
generate_api: generated/api/api.server.gen.go generated/api/api.types.gen.go
	mkdir -p generated
generated/api/api.server.gen.go: api/api.yaml
	mkdir -p generated/api
	oapi-codegen --package api -generate server,spec $< > $@
generated/api/api.types.gen.go: api/api.yaml
	mkdir -p generated/api
	oapi-codegen --package api -generate types $< > $@

# Generates mocks for interfaces
INTERFACES_GO_FILES := $(shell find internal external -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

# Generate mocks for interfaces
generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))


clean:
	rm -rf generated build vendor
	find . -name "*.mock.gen.go" -type f -delete
	find . -name "*.gen.go" -type f -not -path "./internal/db/gormgen/*" -delete
	find . -name "*.out" -type f -delete
	find . -name "wire_gen.go" -type f -delete