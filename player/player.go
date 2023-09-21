package player

type School struct {
	Name string `json:"name" jsonschema:"required,description=School name"`
	City string `json:"city" jsonschema:"required,description=School city"`
}

type Player struct {
	Name     string  `json:"name" jsonschema:"required,description=Player name"`
	Team     string  `json:"team" jsonschema:"required,description=Player team"`
	Position string  `json:"position" jsonschema:"required,description=Player position"`
	Number   int32   `json:"number" jsonschema:"required,description=Player number"`
	School   *School `json:"school" jsonschema:"required,description=Player school"`
}
