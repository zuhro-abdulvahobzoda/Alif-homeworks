package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func addLineNumbers(src, dst string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer dstF.Close()

	scanner := bufio.NewScanner(srcF)
	writer := bufio.NewWriter(dstF)

	defer writer.Flush()

	lineNum := 1

	for scanner.Scan() {
		line := fmt.Sprintf("%d: %s\n", lineNum, scanner.Text())

		_, err := writer.WriteString(line)
		if err != nil {
			return fmt.Errorf("ошибка записи в буфер: %w", err)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return nil
}

func mergeFiles(file1, file2, dst string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ошибка создания целевого файла: %w", err)
	}
	defer dstFile.Close()

	appendFromFile := func(filePath string) error {
		srcFile, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("ошибка открытия файла %s: %w", filePath, err)
		}
		defer srcFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return fmt.Errorf("ошибка копирования файла %s: %w", filePath, err)
		}
		return nil
	}

	if err := appendFromFile(file1); err != nil {
		return err
	}

	if err := appendFromFile(file2); err != nil {
		return err
	}

	return nil
}

func main() {
	addLineNumbers("source.txt", "copied.txt")
	mergeFiles("source.txt", "copied.txt", "second_copy.txt")
}
