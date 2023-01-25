NAME		:= nature-remo-exporter

GOOS		:= $(shell go env GOOS)
GOARCH		:= $(shell go env GOARCH)

OUTPUT		:= out/$(NAME)
LDFLAGS		:=
EXTLDFLAGS	:=
TAGS		:=

# https://github.com/golang/go/issues/26492#issuecomment-435462350
ifeq ($(GOOS),windows)
OUTPUT		:= $(addsuffix .exe,$(NAME))
LDFLAGS		:= $(LDFLAGS) -H=windowsgui
EXTLDFLAGS	:= $(EXTLDFLAGS) -static
endif
ifneq (,$(filter $(GOOS),linux freebsd netbsd openbsd dragonfly))
TAGS		:= $(TAGS) netgo
EXTLDFLAGS	:= $(EXTLDFLAGS) -static
endif
ifeq ($(GOOS),darwin)
TAGS		:= $(TAGS) netgo
LDFLAGS		:= $(LDFLAGS) -s
EXTLDFLAGS	:= $(EXTLDFLAGS) -sectcreate __TEXT __info_plist Info.plist
endif
ifeq ($(GOOS),android)
LDFLAGS		:= $(LDFLAGS) -s
endif

GO.build	:= GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -tags '$(TAGS)' -ldflags '$(LDFLAGS) -extldflags "$(EXTLDFLAGS)"'
GO.test		:= go test -v -cover -coverprofile=coverage.out -covermode=atomic

.PHONY: all
all: $(OUTPUT)

.PHONY: test
test:
	$(GO.test) ./...

out/%: FORCE
	$(GO.build) -o $@ ./cmd/$(notdir $@)

.PHONY: FORCE
FORCE:
