package main

type Record struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type     string `json:"type"`
			Optional bool   `json:"optional"`
			Field    string `json:"field"`
			Name     string `json:"name,omitempty"`
			Version  int    `json:"version,omitempty"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
	} `json:"schema"`
	Values map[string]interface{} `json:"payload"`
}

type Schema struct {
	Type   string `json:"type"`
	Fields []struct {
		Type     string `json:"type"`
		Optional bool   `json:"optional"`
		Field    string `json:"field"`
		Name     string `json:"name,omitempty"`
		Version  int    `json:"version,omitempty"`
	} `json:"fields"`
	Optional bool   `json:"optional"`
	Name     string `json:"name"`
}
