package main

import (
	"encoding/json"
	"os"

	"github.com/invopop/jsonschema"

	"nba-player-go/player"
)

func main() {
	playerSchema := jsonschema.Reflect(&player.Player{})
	file, err := os.Create("./schema/schemas.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(playerSchema); err != nil {
		panic(err)
	}
}
