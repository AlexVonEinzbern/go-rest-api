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

func TestPostLogin(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	input := strings.NewReader(`{
		"CustomerID": "a0ec0ec0-60f9-441c-91e3-8ecfc7174bf5"
    }`)

	req, _ := http.NewRequest("POST", "http://localhost:8080/go-rest-api/login", input)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetLogin(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/go-rest-api/login/3e8806dd-bb77-4525-9906-6f3c97df3758", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetLoginDate(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/go-rest-api/login/date/2022-09-11", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}