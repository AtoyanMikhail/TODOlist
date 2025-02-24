package model

import (
	"context"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	AddNote(ctx context.Context, note string) error
	EditNote(ctx context.Context, id int, updatedNote string) error
	GetNotes(ctx context.Context) ([]Note, error)
	GetNoteByID(ctx context.Context, id int) (Note, error)
	DeleteNote(ctx context.Context, id string) error
}

type Service interface {
	AddNote(ctx context.Context, note string) error
	EditNote(ctx context.Context, id int, updatedNote string) error
	GetNotes(ctx context.Context) ([]Note, error)
	GetNoteByID(ctx context.Context, id int) (Note, error)
	DeleteNote(ctx context.Context, id string) error
}

type AddNoteReq struct {
	Text string `json:"text"`
}

type EditNoteReq struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type DeleteNoteReq struct {
    ID string `json:"id"`
}
