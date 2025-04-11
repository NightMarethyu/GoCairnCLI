package game

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Uses        int    `json:"uses"`
	Magical     bool   `json:"magical"`

	// Optional:
	DamageDie  int `json:"damage_die,omitempty"`
	HealingDie int `json:"healing_die,omitempty"`
	ArmorValue int `json:"armor_value,omitempty"`
}

type Background struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StartingItemIDs []string `json:"starting_items"`
}

type Enemy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HP          int    `json:"hp"`
	DamageDie   int    `json:"damage_die"`
}
