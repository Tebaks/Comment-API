package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Comment-API/data"
)

type Comments struct {
	l *log.Logger
}

func NewComments(l *log.Logger) *Comments {
	return &Comments{l}
}

func (c *Comments) GetComments(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET Comments")

	cl := data.GetComments()

	err := cl.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

type KeyComment struct{}

func (c *Comments) AddComment(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST Comment")

	comment := r.Context().Value(KeyComment{}).(data.Comment)
	data.AddComment(&comment)
}

func (c Comments) MiddlewareValidateComment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		comment := data.Comment{}

		err := comment.FromJSON(r.Body)
		if err != nil {
			c.l.Println("[ERROR] deserializing comment", err)
			http.Error(rw, "Error reading comment", http.StatusBadRequest)
			return
		}

		err = comment.Validate()
		if err != nil {
			c.l.Println("[ERROR] validating comment ", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating comment: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyComment{}, comment)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
