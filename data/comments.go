package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
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
	var results []*Comment

	cur, err := Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var cmnt Comment
		err := cur.Decode(&cmnt)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &cmnt)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return results
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
	c.CreatedOn = time.Now().UTC().String()
	_, err := Collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	dbSize, err := Collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return (int)(dbSize + 1)
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
