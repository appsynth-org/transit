package android_writer

import (
	"fmt"

	"github.com/beevik/etree"
)

type Document struct {
	Lang    string
	Doc     *etree.Document
	Element *etree.Element
}

func (d *Document) WriteComment(comment string) {
	d.Element.CreateComment(comment)
}

func (d *Document) WriteAttribute(key string, value string) {
	element := d.Element.CreateElement("string")
	element.CreateAttr("name", key)
	element.CreateText(value)
}

func (d *Document) Close() {
	filePath := fmt.Sprintf("./output/Android/%s.xml", d.Lang)
	d.Doc.Indent(2)
	d.Doc.WriteToFile(filePath)
}

func NewDocument(lang string) *Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	element := doc.CreateElement("resources")
	element.CreateComment(fmt.Sprintf("lang %s", lang))
	return &Document{
		Lang:    lang,
		Doc:     doc,
		Element: element,
	}
}
