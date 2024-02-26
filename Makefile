genCerasus:
	protoc -I ./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/cerasus.proto

genAuth:
	protoc -I ./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/auth.proto

genOzon:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/ozon.proto

genWB:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/wb.proto