protoc:
	@protoc \
		-I ./protobuf \
		--go_out=plugins=grpc:./pb \
		--grpc-gateway_out=./pb \
		./protobuf/ping.proto

ping:
	@echo '{}' | evans --path ./protobuf/ -p 9998 --service PingAPI --call Ping ./protobuf/*.proto
