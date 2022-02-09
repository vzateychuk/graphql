package meta

type Metadata struct {
	Id   string                 `json:"id"`
	Name string                 `json:"name"`
	Type string                 `json:"type"`
	Meta map[string]interface{} `json:"meta"`
}
