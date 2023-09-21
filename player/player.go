package player

type School struct {
	Name string `json:"name" jsonschema:"required"`
	City string `json:"city" jsonschema:"required"`
}

type Player struct {
	Name     string  `json:"name" jsonschema:"required"`
	Team     string  `json:"team" jsonschema:"required"`
	Position string  `json:"position" jsonschema:"required"`
	Number   int32   `json:"number" jsonschema:"required"`
	School   *School `json:"school" jsonschema:"required"`
}
