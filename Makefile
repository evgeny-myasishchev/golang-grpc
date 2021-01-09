.PHONY: protoc clean

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

clean:
	rm -r -f opt

$(PROTOC_ZIP): opt/tmp
	curl -L $(PROTOC_ZIP_URL) -o $(PROTOC_ZIP)

$(PROTOC): $(PROTOC_ZIP)
	unzip $(PROTOC_ZIP) -d opt

protoc: $(PROTOC)