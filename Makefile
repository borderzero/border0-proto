NAME:=$(shell basename `git rev-parse --show-toplevel`)

# go or swift are currently supported
lang?=go

# add any components here after their directory has been created with .proto files
COMPONENTS:=common connector device 

all: go swift

go:
	make docker-run lang=go

swift:
	make docker-run lang=swift

docker-run: docker-build
	$(foreach component, $(COMPONENTS), docker run -v $(PWD)/common:/app/shared/common -v $(PWD)/$(component):/app/proto $(NAME)-builder-$(lang);)

docker-build:
	docker build -f Dockerfile.$(lang) -t $(NAME)-builder-$(lang) .
