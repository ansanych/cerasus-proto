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

genSettings:
	protoc -I ./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/settings.proto

genProducts:
	protoc -I ./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/products.proto

genSales:
	protoc -I ./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/sales.proto

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

genYM:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/ym.proto

genNews:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/new.proto

genServices:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/services.proto

genPricer:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/pricer.proto

genCounter:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/counter.proto

genAvailabler:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/availabler.proto

genProfile:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/profile.proto

genBrand:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/brand.proto

genLogger:
	protoc -I ./api_v1 \
			--proto_path=./api_v1 \
			--go_out ./api_v1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v1 \
			--go-grpc_opt paths=source_relative \
			./api_v1/logger.proto

#v2
genCerasus2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/cerasus_v2.proto

genAuth2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/auth_v2.proto

genSettings2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/settings_v2.proto

genWB2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/wb_v2.proto

genOZ2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/oz_v2.proto

genYM2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/ym_v2.proto

genProducts2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/products_v2.proto

genCounter2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/counter_v2.proto

genServices2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/services_v2.proto

genPricer2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/pricer_v2.proto

genBoard2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/board_v2.proto

genBrand2:
	protoc -I ./api_v2 \
			--go_out ./api_v2 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_v2 \
			--go-grpc_opt paths=source_relative \
			./api_v2/brand_v2.proto