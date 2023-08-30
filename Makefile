.PHONY: run
run:
	docker-compose -f ./build/docker-compose.yml rm && \
	docker-compose -f ./build/docker-compose.yml build --no-cache && \
 	docker-compose -f ./build/docker-compose.yml up


#.PHONY: pgup
#pgup:
#	docker-compose rm -f && docker-compose build --no-cache && docker-compose up




.PHONY: generate
generate: generate-pb generate-gw generate-openapi

generate-pb:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	./proto/contract.proto

generate-gw:
	protoc -I . --grpc-gateway_out=logtostderr=true:. --grpc-gateway_opt paths=source_relative \
        ./proto/contract.proto

generate-openapi:
	protoc -I. \
		--openapiv2_out . \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt use_go_templates=true \
		./proto/contract.proto