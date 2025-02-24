package server

import (
	"fmt"
	"todolist/internal/app/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() Server {
	r := gin.New()

	return Server{router: r}
}

func (s *Server) Run(address, port string) error {
	err := s.router.Run(fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) UseMiddleware(middleware ...gin.HandlerFunc) {
	s.router.Use(middleware...)
}

func (s *Server) SetupRoutes(handler handler.Handler) {
	r := s.router.Group("/api/v1")
	{
		// Add note
		r.POST("/addnote", handler.AddNote())

		// Edit note
		r.PATCH("/editnote", handler.EditNote())

		// Get all notes
		r.GET("/notes", handler.GetNotes())

		// Get note by ID
		r.GET("/notes/:id", handler.GetNoteByID())

		// Delete note
		r.DELETE("/deletenote/:id", handler.DeleteNote())
	}
}
