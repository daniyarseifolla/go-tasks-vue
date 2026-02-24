APP_NAME=finance-tracker
BUILD_DIR=bin

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) .

run: build
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
