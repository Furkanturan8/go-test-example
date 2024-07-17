package test

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoAppWithTest/router"
)

func TestCreateTodoHandler(t *testing.T) {
	app := fiber.New()
	router.SetupRoutes(app)

	// yeni bir todos oluştur
	todo := `{"title":"Test Todo","completed":false}`
	req := httptest.NewRequest("POST", "/api/todos", strings.NewReader(todo)) // istek at
	req.Header.Set("Content-Type", "application/json")

	// -1 parametresi, fiber framework'ünün Test fonksiyonunun bir parametresidir
	// ve bu parametre, fiber uygulamasının çeşitli test senaryolarını çalıştırmak için kullanılır.
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdTodo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&createdTodo)

	// beklenen değer ile test edeceğimiz değeri eşit mi diye kontrol/test eder
	assert.Equal(t, "Test Todo", createdTodo["title"])
	assert.Equal(t, false, createdTodo["completed"])
}

func TestGetTodosHandler(t *testing.T) {
	app := fiber.New()
	router.SetupRoutes(app)

	// Öncelikle bir todos oluşturalım
	todo := `{"title":"Test Todo","completed":false}`
	createReq := httptest.NewRequest("POST", "/api/todos", strings.NewReader(todo)) // 3.parametre HTTP isteğinin gövdesine (body) eklenmesi gereken veriyi sağlar
	createReq.Header.Set("Content-Type", "application/json")
	app.Test(createReq, -1)

	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var todos []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&todos)
	assert.GreaterOrEqual(t, len(todos), 1) // todos dizisinin en az 1 eleman içerip içermediğini doğrular.
}

func TestUpdateTodoHandler(t *testing.T) {
	app := fiber.New()
	router.SetupRoutes(app)

	// Öncelikle bir todos oluşturalım
	todo := `{"title":"Test Todo","completed":false}`
	createReq := httptest.NewRequest("POST", "/api/todos", strings.NewReader(todo))
	createReq.Header.Set("Content-Type", "application/json")
	createResp, _ := app.Test(createReq, -1)

	var createdTodo map[string]interface{}
	json.NewDecoder(createResp.Body).Decode(&createdTodo)
	todoID := createdTodo["id"].(float64)

	// Şimdi todos'u güncelle
	updatedTodo := `{"title":"Updated Todo","completed":true}`
	updateReq := httptest.NewRequest("PUT", fmt.Sprintf("/api/todos/%v", int(todoID)), strings.NewReader(updatedTodo))
	updateReq.Header.Set("Content-Type", "application/json")
	updateResp, _ := app.Test(updateReq, -1)

	// Durum kodunu kontrol et
	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	var updatedTodoResp map[string]interface{}
	json.NewDecoder(updateResp.Body).Decode(&updatedTodoResp)

	assert.Equal(t, "Updated Todo", updatedTodoResp["title"])
	assert.Equal(t, true, updatedTodoResp["completed"])
}

func TestDeleteTodoHandler(t *testing.T) {
	app := fiber.New()
	router.SetupRoutes(app)

	todo := `{"title":"Test Todo","completed":false}`
	createReq := httptest.NewRequest("POST", "/api/todos", strings.NewReader(todo))
	createReq.Header.Set("Content-Type", "application/json")
	app.Test(createReq, -1)

	req := httptest.NewRequest("DELETE", "/api/todos/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
