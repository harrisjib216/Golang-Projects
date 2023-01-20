package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = SetupRouter()

func Test_PageNotFound(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dfjklsda", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 404, res.Code)
	assert.Equal(t, "Opps! The page you're looking for does not exist", res.Body.String())
}

func Test_HomePage(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "Welcome to the album api", res.Body.String())
}

// todo
func Test_GetAllAlbums_Pass(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, Database, res.Body)
}

// todo
func Test_GetOneAlbum_Fail(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album/500", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "Album could not be found", res.Body.String())
}

// todo
func Test_GetOneAlbum_Pass(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album/1234", nil)

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, Database[0], res.Result().Body)
}
