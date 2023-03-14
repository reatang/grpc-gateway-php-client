
curl_test:
	@go build
	protoc --proto_path=. --proto_path=./third_party \
	  --php_out=. --php_opt=paths=source_relative \
	  --php-grpc-gateway-curl_out=. \
	  testdata/*.proto
