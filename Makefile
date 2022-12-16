
all: clean prepare convert apply-fixes generate

convert:
	go run cmd/convert/main.go

prepare:
	git submodule update --recursive --remote
	cd internal/gosdk-codegen && go build -o ../../gosdk-codegen ./cmd/gosdk-codegen/main.go

generate:
	./gencode.sh

apply-fixes:
	sed -i '' -e 's/\"type\":\"number\"/ \"type\":\"integer\"/g' ./internal/specs-v3/services.json

	# TODO:https://github.com/amzn/selling-partner-api-docs/issues/480
	node -p "var j = require('./internal/specs-v3/ordersV0.json'); j[\"components\"][\"schemas\"][\"ProductInfoDetail\"][\"properties\"][\"NumberOfItems\"][\"type\"]=\"string\"; JSON.stringify({...j})" > tmp.$$.json && mv -f tmp.$$.json ./internal/specs-v3/ordersV0.json
	node -p "var j = require('./internal/specs-v3/ordersV0.json'); j[\"components\"][\"schemas\"][\"OrderItem\"][\"properties\"][\"IsGift\"][\"type\"]=\"boolean\"; JSON.stringify({...j})" > tmp.$$.json && mv -f tmp.$$.json ./internal/specs-v3/ordersV0.json
	node -p "var j = require('./internal/specs-v3/ordersV0.json'); j[\"components\"][\"schemas\"][\"BuyerRequestedCancel\"][\"properties\"][\"IsBuyerRequestedCancel\"][\"type\"]=\"string\"; JSON.stringify({...j})" > tmp.$$.json && mv -f tmp.$$.json ./internal/specs-v3/ordersV0.json

clean:
	@rm -rf generated/*
	@rm gosdk-codegen 2> /dev/null || true
	@rm internal/specs-v3/*.json 2> /dev/null || true
