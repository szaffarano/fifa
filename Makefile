REPO 		?= github.com/szaffarano
PACKAGE  	?= fifa
DATE    	?= $(shell date +%FT%T%z)
VERSION 	?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

PLATFORMS 	:= linux/amd64 windows/amd64 windows/386
DIST_DIR	:= dist

OS 		= $(shell echo $@ | cut -d"/" -f1)
ARCH 	= $(shell echo $@ | cut -d"/" -f2)
EXT 	= $(shell [ "$(OS)" = "windows" ] && echo -n ".exe")

FLAG_VER	= -X $(REPO)/$(PACKAGE)/cmd.Version=$(VERSION)
FLAG_DATE	= -X $(REPO)/$(PACKAGE)/cmd.BuildDate=$(DATE)

compile: clean
	go build \
		-ldflags '$(FLAG_VER) $(FLAG_DATE)' \
		-o '$(DIST_DIR)/$(PACKAGE)'

release: $(PLATFORMS)
	@echo "Construcción exitosa"
	
clean:
	rm -rf $(DIST_DIR)

$(PLATFORMS): clean
	GOOS=$(OS) GOARCH=$(ARCH) go build \
		-tags release \
		-ldflags '$(FLAG_VER) $(FLAG_DATE)' \
		-o '$(DIST_DIR)/$(OS)-$(ARCH)/$(PACKAGE)${EXT}' \
		main.go
	cd $(DIST_DIR); \
	  zip -r $(PACKAGE)-$(VERSION)-$(OS)-$(ARCH).zip $(OS)-$(ARCH); \
	  rm -rf $(OS)-$(ARCH)

.PHONY: release $(PLATFORMS) clean