.PHONY: genproto
genproto:
	protoc -I. --go_out=plugins=grpc:./pd/auth ./auth.proto
	protoc -I. --go_out=plugins=grpc:./pd/fight ./fight.proto
