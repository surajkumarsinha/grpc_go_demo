build:
	mkdir pb
	mkdir pb/messages
	mkdir pb/services
	make messages 
	make services  

messages:
	protoc --proto_path=protos/messages --go_out=pb/messages --go_opt=paths=source_relative protos/messages/*.proto

services:
	# protoc --proto_path=protos/services --proto_path=protos/messages --go_out=pb/services --go-grpc_out=. --go_opt=paths=source_relative protos/services/*.proto
	protoc --proto_path=protos/services --proto_path=protos/messages --go_out=pb/services --go-grpc_out=. --go_opt=paths=source_relative \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out . \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt generate_unbound_methods=true \
	protos/services/*.proto

gateway:
	# protoc -I . --grpc-gateway_out ./pb/services \
	#   --proto_path=protos/services --proto_path=protos/messages \
  #   --grpc-gateway_opt logtostderr=true \
  #   --grpc-gateway_opt paths=source_relative \
  #   protos/services/*.proto

  
clean:
	rm -rf pb
	# rm pb/messages/*go 
	# rm pb/services/*go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

rest:
	go run cmd/gw/main.go 
 
test:
	go test -cover -race ./serializer