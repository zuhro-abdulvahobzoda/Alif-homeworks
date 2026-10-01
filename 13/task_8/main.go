package main

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	filename string
}

func NewLogger(filename string) *Logger {
	return &Logger{filename: filename}
}

func (l *Logger) write(level, msg string) error {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("[%s] %-5s %s\n",
		time.Now().Format("2006-01-02 15:04:05"), level, msg)
	_, err = f.WriteString(line)

	return err
}

func (l *Logger) Info(msg string)  { l.write("INFO", msg) }
func (l *Logger) Warn(msg string)  { l.write("WARN", msg) }
func (l *Logger) Error(msg string) { l.write("ERROR", msg) }

func main() {
	log := NewLogger("app.log")

	log.Info("The Program started up")
	log.Info("The user logged into system")
	log.Warn("Fast response(ig): 0.5с")
	log.Info("Connection to data base has gone successfully!")
	log.Info("The user added some data")
	log.Error("Something had gone completely wrong and all data gor wiped")
	log.Warn("The user must face the consequences of his actions...")
	log.Info("Oh, the user already left from the crime scene")
	log.Warn("Developers developers, the data base had got wiped")
	log.Info("Program shut down")

	content, _ := os.ReadFile("app.log")
	fmt.Print(string(content))
}
