package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

// Comment defines the structure for an API comment
type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"postId" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Text      string `json:"text" validate:"required"`
	CreatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (c *Comment) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

func (c *Comment) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

type Comments []*Comment

func (c *Comments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

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

func AddComment(c *Comment) {
	c.ID = getNextID()
	commentList = append(commentList, c)
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
