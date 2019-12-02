protoc:
	@protoc \
		-I ./protobuf \
		-I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		--go_out=plugins=grpc:./pb \
		--grpc-gateway_out=./pb \
		./protobuf/ping.proto

ping:
	@echo '{}' | evans --path ../protobuf/ -p 9998 --service PingAPI --call Ping ./protobuf/*.proto
