APP_NAME=test-case-ndi
MAIN_FILE=cmd/main.go
BUILD_DIR=bin

# default command ketika make dijalankan tanpa argumen
.PHONY: all
all: run

# membersihkan direktori build
.PHONY: clean
clean:
	@echo "cleaning..."
	@rm -rf $(BUILD_DIR)

# build aplikasi
.PHONY: build
build:
	@echo "building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# build dan jalankan aplikasi
.PHONY: run
run: build
	@echo "running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

# install dependensi
.PHONY: deps
deps:
	@echo "installing dependencies..."
	@go mod tidy
	@go mod download

# help command
.PHONY: help
help:
	@echo "available commands:"
	@echo "  make build  - build the application"
	@echo "  make run    - build and run the application"
	@echo "  make clean  - remove build artifacts"
	@echo "  make deps   - install dependencies"
	@echo "  make help   - show this help message"