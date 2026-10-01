package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Write() {
	tasks := "Make some dinner\nDraw something\nPlay Minecraft\nSleep bruh\nQuestion yourself why you do not understand Go\n"
	file, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.Write([]byte(tasks))

}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	counter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func copyFile(src, dst string) error {
	scrFile, err := os.Open(src)

	if err != nil {
		return err
	}

	defer scrFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, scrFile)
	if err != nil {
		return err
	}

	return dstFile.Sync()
}

func grep(filename, word string) ([]string, error) {
	words := []string{} // creating a slice for storing words
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lowerWord := strings.ToLower(word)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(strings.ToLower(line), lowerWord) {
			words = append(words, line)
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}

	}
	return words, nil
}

func main() {
	Write()
	fmt.Println(countLines("tasks.txt"))

	lines, err := countLines("tasks.txt")
	if err != nil {
		log.Println("Error counting lines:", err)
	} else {
		fmt.Println("Line count:", lines)
	}

	copyFile("tasks.txt", "tasks_copy.txt")
	if err := copyFile("tasks.txt", "tasks_copy.txt"); err != nil {
		log.Println("Error copying file:", err)
	}

	res, err1 := grep("minecraft_facts.txt", "mob")
	if err1 != nil {
		log.Println("An error occured: ", err)
	}
	fmt.Println(res)
}
