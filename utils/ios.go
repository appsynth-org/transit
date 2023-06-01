package utils

import (
	"fmt"
	"log"
	"os"
)

type Document struct {
	File         *os.File
	FileLocation string
}

func NewIosDocument(lang string) *Document {
	filePath := fmt.Sprintf("./output/iOS/%s.strings", lang)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	return &Document{
		File:         file,
		FileLocation: fmt.Sprintf("./output/iOS/%s.strings", lang),
	}
}

func (d *Document) WriteComment(comment string) {
	_, err := d.File.WriteString(fmt.Sprintf("/* %s */\n", comment))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func (d *Document) WriteToFile(key string, value string) {
	_, err := d.File.WriteString(fmt.Sprintf("\"%s\" = \"%s\";\n", key, value))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}
