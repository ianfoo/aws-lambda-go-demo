BUILD=main
DEPLOYMENT=deployment.zip
FUNCTION=HelloFunction
REGION=us-west-1

build: $(BUILD)
	GOOS=linux go build -o $(BUILD) .

package: build
	zip $(DEPLOYMENT) $(BUILD)

# Note that if you are deploying the function for the first time, you will need
# to use the create-function subcommand. update-function-code only works after
# the function has been created.
deploy: package
	aws lambda update-function-code \
		--region $(REGION) \
		--function-name $(FUNCTION) \
		--zip-file fileb://./$(DEPLOYMENT)

.PHONY: clean
clean:
	rm -f $(BUILD) $(DEPLOYMENT)
