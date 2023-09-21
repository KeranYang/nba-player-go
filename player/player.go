package player

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
