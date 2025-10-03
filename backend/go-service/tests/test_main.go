package tests

import (
	"accountantapp/go-service/internal/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Подключаем базу
	database.Connect() // правильное имя пакета

	// Можно очистить таблицы перед тестами, если нужно
	// database.DB.Exec("DELETE FROM transactions")
	// database.DB.Exec("DELETE FROM accounts")
	// database.DB.Exec("DELETE FROM users")

	// Запуск всех тестов
	code := m.Run()

	os.Exit(code)
}
