default: build

BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=go-mvc-srv

export GOTPATH=$(CURDIR)

makedir:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

get:
	$(GOGET)

build:
	$(GOBUILD)

install:
	$(GOINSTALL)

clean:
	@rm -fr $(BUILDPATH)/bin/$(EXENAME)
	@rm -fr $(BUILDPATH)/pkg

all: makedir get build
