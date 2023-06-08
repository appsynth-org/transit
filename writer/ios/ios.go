package ios_writer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Document struct {
	Lang string
	File *os.File
}

func (d *Document) WriteComment(comment string) {
	_, err := d.File.WriteString(fmt.Sprintf("/* %s */\n", comment))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func (d *Document) WriteAttribute(key string, value string) {
	_, err := d.File.WriteString(fmt.Sprintf("\"%s\" = \"%s\";\n", key, value))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func (d *Document) Close() {
	d.File.Close()
}

func createDirectoryIfNotExist(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}
	return nil
}

func NewDocument(lang string) *Document {
	filePath := fmt.Sprintf("../../output/iOS/%s.strings", lang)
	createDirectoryIfNotExist(filePath)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	return &Document{
		Lang: lang,
		File: file,
	}
}
