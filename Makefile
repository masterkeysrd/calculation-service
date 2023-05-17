# initialize the go development environment
dev/init: gvm/init

# init installs the go version and initializes the go environment
gvm/init:
	$(MAKE) -C scripts/gvm gvm/init

# use sets the go version and initializes the go environment
gvm/use:
	$(MAKE) -C scripts/gvm gvm/use

.PHONY: dev/init gvm/init gvm/use