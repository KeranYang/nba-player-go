package player

// swagger:model
type Player struct {
	// the player's name
	Name string `json:"name,omitempty"`
	// the player's team name
	Team string `json:"team,omitempty"`
	// the player's position
	Position string `json:"position,omitempty"`
	// the player's number
	Number int32 `json:"number,omitempty"`
	// the player's school
	School School `json:"school,omitempty"`
}

// swagger:model
type School struct {
	// the school's name
	Name string `json:"name,omitempty"`
	// the school's city
	City string `json:"city,omitempty"`
}
