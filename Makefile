GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get



APP_NAME := server_db
BUILD_DIR := build

LINUX_BUILD := $(BUILD_DIR)/linux
WINDOWS_BUILD := $(BUILD_DIR)/windows


TARGETS = linux windows

all: clean $(TARGETS)


windows:
	@echo "Compiling for Windows amd64..."
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BUILD)/$(APP_NAME).exe

linux:
	@echo "Compiling for Linux amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BUILD)/$(APP_NAME)


clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	mkdir $(BUILD_DIR)
