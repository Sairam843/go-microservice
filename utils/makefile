.PHONY: protogen setupGrpc gqlgen

GOPATH_BIN := /usr/local/go/bin

protogen:
	@if [ -z "$(PROTO_DIR)" ]; then \
		echo "❌ Usage: make protogen PROTO_DIR=./proto"; \
		exit 1; \
	fi
	@echo "📁 Using PROTO_DIR=$(PROTO_DIR)"
	@echo "📁 Generating Go code to ./../cache/internal/pb"
	protoc \
		-I$(shell realpath $(PROTO_DIR)) \
		--plugin=protoc-gen-go=$(GOPATH_BIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(GOPATH_BIN)/protoc-gen-go-grpc \
		--go_out=./../cache/internal/pb \
		--go-grpc_out=./../cache/internal/pb \
		$(shell realpath $(PROTO_DIR))/*.proto
	@echo "✅ protobuf generated"

setupGrpc:
	@echo "📦 Installing protoc & Go plugins..."
	wget -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v30.2/protoc-30.2-linux-x86_64.zip && \
	sudo apt install -y unzip && \
	unzip -o protoc.zip -d /tmp/protoc && \
	sudo rm -rf /usr/local/include/google && \
	sudo mv -f /tmp/protoc/bin/protoc /usr/local/bin/ && \
	sudo mv -f /tmp/protoc/include/* /usr/local/include/ && \
	rm protoc.zip && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "🎉 protoc and plugins installed like a legend"

setupGqlgen:
	@echo "📦 Installing gqlgen..."
	@go install github.com/99designs/gqlgen@latest
	@echo "🔗 Adding $(shell go env GOPATH)/bin to PATH via ~/.bashrc"
	@grep -qxF 'export PATH="$$PATH:$(shell go env GOPATH)/bin"' ~/.bashrc || \
		echo 'export PATH="$$PATH:$(shell go env GOPATH)/bin"' >> ~/.bashrc
	@echo "🎉 gqlgen installed"
	@echo "✅ PATH updated. Run: source ~/.bashrc"

gqlgen:
	@if [ -z "$(GQL_DIR)" ]; then \
		echo "❌ Usage: make gqlgen GQL_DIR=./gateway"; \
		exit 1; \
	fi
	@echo "📁 Using GQL_DIR=$(GQL_DIR)"
	cd $(GQL_DIR) && \
	go mod tidy && \
	gqlgen generate
