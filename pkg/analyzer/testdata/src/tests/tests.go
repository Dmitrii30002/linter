package tests

import (
	"log"
	"log/slog"
)

func testFunc() {
	//1. Лог-сообщения должны начинаться со строчной буквы
	//❌Неправильно
	log.Fatal("Starting server on port 8080")   // want `^[a-z]+\.[a-zA-Z]+: log messages must start with a lowercase letter \(message: .+\)$`
	slog.Error("Failed to connect to database") // want `^[a-z]+\.[a-zA-Z]+: log messages must start with a lowercase letter \(message: .+\)$`

	//✅Правильно
	log.Fatal("starting server on port 8080")
	slog.Error("failed to connect to database")

	//2. Лог-сообщения должны быть только на английском языке
	//❌Неправильно
	log.Print("zапуск сервера")                   // want `^[a-z]+\.[a-zA-Z]+: log message should be in English only \(message: .+\)$`
	log.Fatal("oшибка подключения к базе данных") // want `^[a-z]+\.[a-zA-Z]+: log message should be in English only \(message: .+\)$`

	//✅Правильно
	log.Print("starting server")
	log.Fatal("failed to connect to database")

	//3. Лог-сообщения не должны содержать спецсимволы или эмодзи
	//❌Неправильно
	log.Print("server started!🚀")                 // want `^[a-z]+\.[a-zA-Z]+: log messages should not contain special characters or emojis \(message: .+\)$`
	log.Fatal("connection failed!!!")             // want `^[a-z]+\.[a-zA-Z]+: log messages should not contain special characters or emojis \(message: .+\)$`
	log.Fatal("warning: something went wrong...") // want `^[a-z]+\.[a-zA-Z]+: log messages should not contain special characters or emojis \(message: .+\)$`

	//✅Правильно
	log.Print("server started")
	log.Fatal("connection failed")
	log.Fatal("something went wrong")

	//4. Лог-сообщения не должны содержать потенциально чувствительные данные
	//❌Неправильно
	password := "a"
	apiKey := "apiKey"
	token := "a"
	log.Print("user password ", password) // want `^[a-z]+\.[a-zA-Z]+: log message should not contain sensitive data \(message: .+\)$`
	log.Print("apiKey" + apiKey)          // want `^[a-z]+\.[a-zA-Z]+: log message should not contain sensitive data \(message: .+\)$`
	log.Print("token " + token)           // want `^[a-z]+\.[a-zA-Z]+: log message should not contain sensitive data \(message: .+\)$`

	//✅Правильно
	log.Print("user authenticated successfully")
	log.Print("api request completed")
	log.Print("token validated")
}
