.PHONY: swagger-gen
swagger-gen:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o ./swagger-schema/helper/player-swagger.json --scan-models
	swagger mixin ./swagger-schema/helper/base.json ./swagger-schema/helper/player-swagger.json -o ./swagger-schema/player.json

.PHONY: gen-standard-schema
gen-standard-schema:
	go run main.go


