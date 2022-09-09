package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	input := strings.NewReader(`{
        "id": "bcajbavajbvajbvbjlaknLHLKUCTG",
        "category_name": "categoria 1",
        "description": "description 1"
    }`)

	req, _ := http.NewRequest("POST", "http://localhost:8080/go-rest-api/categories", input)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/go-rest-api/categories", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetActiveCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/go-rest-api/categories/active", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	input := strings.NewReader(`{
        "id": "bcajbavajbvajbvbjlaknLHLKUCTG",
        "description": "description updated"
    }`)

	req, _ := http.NewRequest("PATCH", "http://localhost:8080/go-rest-api/categories/bcajbavajbvajbvbjlaknLHLKUCTG", input)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "http://localhost:8080/go-rest-api/categories/bcajbavajbvajbvbjlaknLHLKUCTG", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
