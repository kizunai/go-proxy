# init project path
HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output
TMPDIR  := $(HOMEDIR)/tmp
APPDIR  := $(HOMEDIR)/tmp/apps/GoProxy
APP="GoProxy"

# init command params
GOROOT  := /root/software/go19
GO      := $(GOROOT)/bin/go
GOPATH  := $(GO) env GOPATH
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test -gcflags="-N -l"
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")

# make, make all
all: compile package

# set proxy env
set-env:
		$(GO) env -w GO111MODULE=on
		$(GO) env -w GONOSUMDB=\*

# make compile
compile: build
build:
		$(GOBUILD) -o $(HOMEDIR)/$(APP)

# make package
package: package-bin
package-bin:
		mkdir -p $(OUTDIR)
		mkdir -p $(TMPDIR)
		mkdir -p $(APPDIR)
		mkdir -p $(APPDIR)/bin

		mv $(APP) $(APPDIR)/bin
		cp -r conf $(APPDIR)/bin
		cp control.sh $(APPDIR)/
		tar -zcvf $(OUTDIR)/source.tar.gz -C $(TMPDIR) .
		tm -rf $(TMPDIR)
