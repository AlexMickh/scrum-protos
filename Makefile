FILE = auth

gen:
	@protoc --go_out=./pkg/api/${FILE} --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/api/${FILE} --go-grpc_opt=paths=source_relative \
		./proto/${FILE}/*proto
