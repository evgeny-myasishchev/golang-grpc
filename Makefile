.PHONY: proto clean

.DEFAULT_GOAL=proto

DETECTED_OS := $(shell uname)

ifeq ($(DETECTED_OS),Darwin)
    PLATFORM=osx-x86_64
endif
ifeq ($(DETECTED_OS),Linux)
	PLATFORM=linux-x86_64
endif

PROTOC_VERSI0ON=3.14.0
PROTOC_ZIP_URL=https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSI0ON)/protoc-$(PROTOC_VERSI0ON)-$(PLATFORM).zip
PROTOC_ZIP=opt/tmp/protoc-$(PROTOC_VERSI0ON)-$(PLATFORM).zip
PROTOC=opt/bin/protoc

opt/bin:
	mkdir -p opt/bin

opt/tmp:
	mkdir -p opt/tmp

$(PROTOC_ZIP): opt/tmp
	curl -L $(PROTOC_ZIP_URL) -o $(PROTOC_ZIP)

$(PROTOC): $(PROTOC_ZIP)
	unzip $(PROTOC_ZIP) -d opt
	touch $(PROTOC)

clean:
	rm -r -f opt

%.pb.go: %.proto $(PROTOC)
	$(PROTOC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $<

proto: pkg/*/*.pb.go