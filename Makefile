NAME:=$(shell basename `git rev-parse --show-toplevel`)

# go or swift are currently supported
lang?=go

# add any components here after their directory has been created with .proto files
COMPONENTS:=common connector device 

all: go swift kotlin

go:
	make docker-run lang=go

swift:
	make docker-run lang=swift

kotlin:
	make docker-run lang=kotlin

docker-run: docker-build
	@for component in $(COMPONENTS); do \
		docker run \
			-v $(PWD)/common:/app/shared/common \
			-v $(PWD)/$$component:/app/proto \
			-v $(PWD)/gen:/app/gen \
			$(NAME)-builder-$(lang); \
	done

docker-build:
	docker build -f docker/$(lang).Dockerfile -t $(NAME)-builder-$(lang) .
