BINARY_NAME=tree2

INSTALL_DIR=/usr/local/bin

all: build

build:
	go build -o $(BINARY_NAME) main.go

install: build
	sudo mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
