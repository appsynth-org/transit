package utils

import (
	"fmt"
	"log"

	ss "github.com/appsynth-org/transit/service"
	"github.com/beevik/etree"
)

func GenerateLocale(groups []ss.LocalizeGroup) {
	// iOS
	docIosEN := NewIosDocument("en")
	docIosTH := NewIosDocument("th")
	docIosEN.WriteComment("lang en")
	docIosTH.WriteComment("lang th")

	defer docIosEN.File.Close()
	defer docIosTH.File.Close()

	// Android
	docAndroidEN, elementAndroidEN := newAndroidDocument("en")
	docAndroidTH, elementAndroidTH := newAndroidDocument("th")

	for _, group := range groups {
		// EN
		elementAndroidEN.CreateComment(group.GroupName)
		docIosEN.WriteComment(group.GroupName)
		// TH
		elementAndroidTH.CreateComment(group.GroupName)
		docIosTH.WriteComment(group.GroupName)

		for _, translation := range group.Keys {
			// Skip Android empty key
			if len(translation.AndroidKey) > 0 {
				// EN
				eachEN := elementAndroidEN.CreateElement("string")
				eachEN.CreateAttr("name", translation.AndroidKey)
				eachEN.CreateText(translation.Translation.En)
				// TH
				eachTH := elementAndroidTH.CreateElement("string")
				eachTH.CreateAttr("name", translation.AndroidKey)
				eachTH.CreateText(translation.Translation.Th)
			}

			// Skip iOS empty key
			if (len(translation.IosKey)) > 0 {
				// EN
				docIosEN.WriteToFile(translation.IosKey, translation.Translation.En)
				// TH
				docIosTH.WriteToFile(translation.IosKey, translation.Translation.Th)
			}

		}
	}

	docAndroidEN.Indent(2)
	docAndroidTH.Indent(2)

	err := docAndroidEN.WriteToFile("./output/Android/en.xml")
	if err != nil {
		log.Fatal(err)
	}
	err = docAndroidTH.WriteToFile("./output/Android/th.xml")
	if err != nil {
		log.Fatal(err)
	}
}

func newAndroidDocument(lang string) (*etree.Document, *etree.Element) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	element := doc.CreateElement("resources")
	element.CreateComment(fmt.Sprintf("lang %s", lang))
	return doc, element
}
