package main

import (
	"encoding/json"
	"os"

	"github.com/alecthomas/jsonschema"
)

type School struct {
	Name string `json:"name,omitempty"`
	City string `json:"city,omitempty"`
}

type Player struct {
	Name     string  `json:"name,omitempty"`
	Team     string  `json:"team,omitempty"`
	Position string  `json:"position,omitempty"`
	Number   int32   `json:"number,omitempty"`
	School   *School `json:"school,omitempty"`
}

func main() {
	playerSchema := jsonschema.Reflect(&Player{})
	schoolSchema := jsonschema.Reflect(&School{})

	schemas := map[string]interface{}{
		"Player": playerSchema,
		"School": schoolSchema,
	}

	file, err := os.Create("schemas.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(schemas); err != nil {
		panic(err)
	}
}
