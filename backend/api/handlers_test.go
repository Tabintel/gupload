// backend/api/handlers_test.go

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUpload(t *testing.T) {
	// Create a test Gin context
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/upload", HandleUpload)

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/upload", nil)
	req.Header.Add("Content-Type", "multipart/form-data")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Image uploaded successfully!")
	assert.Contains(t, w.Body.String(), "https://example.com/uploads/")
}
