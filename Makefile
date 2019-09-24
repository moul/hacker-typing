GOPKG ?=	moul.io/hacker-typing
DOCKER_IMAGE ?=	moul/hacker-typing
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
