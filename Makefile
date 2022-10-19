include variables.mk

binary:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/linx.out .

image:
	docker build . -t aapozd/lynx-test:$(COMMIT_HASH)

run:
	docker run \
	-v $(PROJECT_DIR)/resources:/app/resources \
	aapozd/lynx-test:$(COMMIT_HASH) \
	--path="/app/resources/$(FILE)"

