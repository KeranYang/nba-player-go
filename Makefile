.PHONY: swagger-gen
swagger-gen:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o ./swagger-schema/player.json --scan-models



