package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/adapter/storage"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// MediaHandler handles media upload and deletion endpoints.
type MediaHandler struct {
	storage storage.StorageProvider
}

// NewMediaHandler creates a new MediaHandler.
func NewMediaHandler(storageProvider storage.StorageProvider) *MediaHandler {
	return &MediaHandler{storage: storageProvider}
}

// uploadURLRequest is the request body for POST /media/upload-url.
type uploadURLRequest struct {
	Filename    string `json:"filename" validate:"required"`
	ContentType string `json:"content_type" validate:"required"`
}

// GetUploadURL handles POST /media/upload-url.
// Returns a presigned URL for the client to upload a file directly to object storage.
func (h *MediaHandler) GetUploadURL(c *fiber.Ctx) error {
	var req uploadURLRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(req); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	if h.storage == nil {
		return response.Error(c, fiber.StatusServiceUnavailable, "storage_unavailable", "storage service is not configured")
	}

	// Generate a unique object key.
	key := fmt.Sprintf("uploads/%s/%s", uuid.New().String(), req.Filename)

	// Generate a presigned URL valid for 15 minutes.
	url, err := h.storage.GetPresignedURL(c.Context(), key, 15*time.Minute)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "upload_url_failed", err.Error())
	}

	return response.Success(c, fiber.Map{
		"upload_url":   url,
		"key":          key,
		"content_type": req.ContentType,
		"expires_in":   900, // 15 minutes in seconds
	})
}

// Delete handles DELETE /media/:id.
// Removes a file from object storage by its key.
func (h *MediaHandler) Delete(c *fiber.Ctx) error {
	key := c.Params("id")
	if key == "" {
		return response.Error(c, fiber.StatusBadRequest, "missing_key", "media key is required")
	}

	if h.storage == nil {
		return response.Error(c, fiber.StatusServiceUnavailable, "storage_unavailable", "storage service is not configured")
	}

	if err := h.storage.Delete(c.Context(), key); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "delete_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "file deleted"})
}
