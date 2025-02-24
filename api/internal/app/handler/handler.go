package handler

import (
	"net/http"
	"strconv"
	"todolist/internal/app/model"

	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	service model.Service
}

func New(service model.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.AddNoteReq

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		if req.Text == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "attempt to add an empty note"})
			return
		}

		if err := h.service.AddNote(ctx.Request.Context(), req.Text); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not add note "})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"mesage": "created successfully"})
	}
}

func (h *Handler) EditNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.EditNoteReq

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		if req.Text == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "attempt to set empty text"})
			return
		}

		if err := h.service.EditNote(ctx.Request.Context(), req.ID, req.Text); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not edit note"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
	}
}

func (h *Handler) GetNotes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := h.service.GetNotes(ctx.Request.Context())
		if err != nil {
			log.Printf("Error getting notes %s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not get notes"})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
  func (h *Handler) GetNoteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || id <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid note ID"})
			return
		}

		note, err := h.service.GetNoteByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "note not found"})
			return
		}

		ctx.JSON(http.StatusOK, note)
	}
}

func (h *Handler) DeleteNote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || id <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid note ID"})
			return
		}

		if err := h.service.DeleteNote(ctx.Request.Context(), strconv.Itoa(id)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete note"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
	}
}
