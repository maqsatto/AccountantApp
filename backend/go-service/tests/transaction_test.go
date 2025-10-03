package tests

import (
	"accountantapp/go-service/internal/controllers"
	"accountantapp/go-service/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/transactions", controllers.CreateTransaction)
	r.GET("/transactions", controllers.GetAllTransactions)
	r.GET("/transactions/:id", controllers.GetTransactionByID)
	return r
}

func TestCreateTransaction(t *testing.T) {
	r := SetupRouter()

	tx := models.Transaction{
		UserID:      1,
		AccountID:   1,
		Amount:      100,
		Description: "Test transaction",
		Date:        time.Now(),
	}

	jsonValue, _ := json.Marshal(tx)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetAllTransactions(t *testing.T) {
	r := SetupRouter()
	req, _ := http.NewRequest("GET", "/transactions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetTransactionByID(t *testing.T) {
	r := SetupRouter()
	req, _ := http.NewRequest("GET", "/transactions/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
