protoc:
	@protoc -I ./protobuf --go_out=plugins=grpc:./pb ./protobuf/*.proto

ping:
	@echo '{}' | evans --path ../protobuf/ -p 9998 --service PingAPI --call Ping ./protobuf/*.proto
