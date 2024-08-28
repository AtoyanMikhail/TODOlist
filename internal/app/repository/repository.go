package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"todolist/internal/app/model"

	_ "github.com/mattn/go-sqlite3"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func New(db DBTX) model.Repository {
	return &repository{db: db}
}

func (r *repository) AddNote(ctx context.Context, text string) error {
	query := "INSERT INTO notes (text, created_at) VALUES ($1, $2)"

	created_at := time.Now().Local().UTC()
	if _, err := r.db.ExecContext(ctx, query, text, created_at); err != nil {
		return err
	}

	log.Printf("Note \"%s\" created successfully at %v", text, created_at)
	return nil
}

func (r *repository) EditNote(ctx context.Context, id int, updatedNote string) error {
	query := "UPDATE notes SET text = $1 WHERE id = $2"

	if _, err := r.db.ExecContext(ctx, query, updatedNote, id); err != nil {
		return err
	}

	log.Printf("Note with ID %d updated successfully to \"%s\"", id, updatedNote)
	return nil
}

func (r *repository) GetNotes(ctx context.Context) ([]model.Note, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM notes")
	if err != nil {
		return []model.Note{}, err
	}
	defer rows.Close()

	var notes []model.Note

	for rows.Next() {
		var note model.Note
		rows.Scan(&note.ID, &note.Text, &note.CreatedAt)
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *repository) GetNoteByID(ctx context.Context, id int) (model.Note, error) {
	var note model.Note

	err := r.db.QueryRowContext(ctx, "SELECT id, text, created_at FROM notes WHERE id = $1", id).Scan(&note.ID, &note.Text, &note.CreatedAt)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (r *repository) DeleteNote(ctx context.Context, id string) error {
	query := "DELETE FROM notes WHERE id = $1"

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return err
	}

	log.Printf("Note with ID %s deleted successfully", id)
	return nil
}
