package main

import "fmt"

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Info(msg string)  { fmt.Println("INFO:", msg) }
func (c ConsoleLogger) Error(msg string) { fmt.Println("ERROR:", msg) }
func (c ConsoleLogger) Debug(msg string) { fmt.Println("DEBUG:", msg) }

type FileLogger struct{}

func (f FileLogger) Info(msg string)  { fmt.Println("FILE INFO:", msg) }
func (f FileLogger) Error(msg string) { fmt.Println("FILE ERROR:", msg) }
func (f FileLogger) Debug(msg string) { fmt.Println("FILE DEBUG:", msg) }

func main() {
	var logger Logger = ConsoleLogger{}
	logger.Info("Приложение запущено")
	logger.Error("Ошибка соединения")
	logger.Debug("Отладочная информация")
}
