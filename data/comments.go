package data

// Comment defines the structure for an API comment
type Comment struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
	CreatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}
