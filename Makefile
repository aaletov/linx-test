include variables.mk

binary:
	go build -o $(BUILD_DIR)/linx.out .
