package service

import (
	"context"
	"time"
	"todolist/internal/app/model"
)

type service struct {
	repository model.Repository
	timeout    time.Duration
}

func New(repository model.Repository) *service {
	return &service{
		repository: repository,
		timeout:    time.Second * 2,
	}
}

func (s *service) AddNote(c context.Context, note string) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.AddNote(ctx, note)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) EditNote(c context.Context, id int, updatedNote string) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.EditNote(ctx, id, updatedNote)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetNotes(c context.Context) ([]model.Note, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	notes, err := s.repository.GetNotes(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil

}

func (s *service) GetNoteByID(c context.Context, id int) (model.Note, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	note, err := s.repository.GetNoteByID(ctx, id)
	if err != nil {
		return model.Note{}, err
	}
	return note, nil
}

func (s *service) DeleteNote(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.repository.DeleteNote(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
