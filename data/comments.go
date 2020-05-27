package data

import (
	"context"
	"encoding/json"
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

// Validate data
func (c *Comment) Validate() error {
	validate := validator.New()

	return validate.Struct(c)
}

// FromJSON : Decode reader as Comment struct
func (c *Comment) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

// Comments : Type for slice of comments
type Comments []*Comment

// ToJSON : Encode writer as Comment struct
func (c *Comments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func FindComments(postId int) Comments {
	var results Comments
	cur, err := Collection.Find(context.TODO(), bson.D{{"postid", postId}})
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

// GetComments : Find every comment in database
func GetComments(postID int) Comments {
	comments := FindComments(postID)

	return comments
}

// AddComment : Add comment to database
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
