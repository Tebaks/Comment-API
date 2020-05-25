package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Comment defines the structure for an API comment
type Comment struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
	CreatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (c *Comment) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Comment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

type Comments []*Comment

func GetComments() Comments {
	return commentList
}

var ErrCommentNotFound = fmt.Errorf("Comment not found.")

func findComment(id int) (*Comment, int, error) {
	for i, c := range commentList {
		if c.ID == id {
			return c, i, nil
		}
	}

	return nil, -1, ErrCommentNotFound
}

func getNextID() int {
	cl := commentList[len(commentList)-1]
	return cl.ID + 1
}

var commentList = []*Comment{
	&Comment{
		ID:        1,
		Author:    "Tebaks",
		Text:      "This is a nice profile",
		CreatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
	&Comment{
		ID:        2,
		Author:    "Ejorange",
		Text:      "Git Gud",
		CreatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
}
