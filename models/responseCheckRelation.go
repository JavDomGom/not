package models

/*ResponseCheckRelation returns true or false after consulting
the relation between two users. */
type ResponseCheckRelation struct {
	Status bool `json:"status"`
}
