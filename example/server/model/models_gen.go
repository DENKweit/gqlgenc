// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Book interface {
	IsBook()
}

type Media interface {
	IsMedia()
}

type Chatroom struct {
	Name     string     `json:"name"`
	Messages []*Message `json:"messages"`
}

type ColoringBook struct {
	Title  string   `json:"title"`
	Colors []string `json:"colors"`
}

func (ColoringBook) IsBook() {}

type Image struct {
	Size int `json:"size"`
}

func (Image) IsMedia() {}

type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostCreateInput struct {
	Text string `json:"text"`
}

type SomeExtraType struct {
	Child *SomeExtraTypeChild `json:"child"`
}

type SomeExtraTypeChild struct {
	ID string `json:"id"`
}

type Textbook struct {
	Title   string   `json:"title"`
	Courses []string `json:"courses"`
}

func (Textbook) IsBook() {}

type UploadData struct {
	Size int `json:"size"`
}

type UploadFilesMap struct {
	Somefile *UploadData `json:"somefile"`
}

type UploadFilesMapInput struct {
	Somefile graphql.Upload `json:"somefile"`
}

type Video struct {
	Duration int `json:"duration"`
}

func (Video) IsMedia() {}
