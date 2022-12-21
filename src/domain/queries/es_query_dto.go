package queries

type EsQuery struct {
	Equeals []FieldValue `json:"equeals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
