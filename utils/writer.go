package utils

import (
	"github.com/appsynth-org/transit/reader"
	android "github.com/appsynth-org/transit/writer/android"
	ios "github.com/appsynth-org/transit/writer/ios"
)

func GenerateLocale(groups []reader.LocalizeGroup) {
	iosWriterEN := ios.NewDocument("en")
	iosWriterTH := ios.NewDocument("th")

	androidWriterEN := android.NewDocument("en")
	androidWriterTH := android.NewDocument("th")

	for _, group := range groups {
		// EN
		iosWriterEN.WriteComment(group.GroupName)
		androidWriterEN.WriteComment(group.GroupName)
		// TH
		iosWriterTH.WriteComment(group.GroupName)
		androidWriterTH.WriteComment(group.GroupName)

		for _, translation := range group.Keys {
			// Skip Android empty key
			if len(translation.AndroidKey) > 0 {
				// EN
				androidWriterEN.WriteAttribute(translation.AndroidKey, translation.Translation.En)
				// TH
				androidWriterTH.WriteAttribute(translation.AndroidKey, translation.Translation.Th)
			}

			// Skip iOS empty key
			if (len(translation.IosKey)) > 0 {
				// EN
				iosWriterEN.WriteAttribute(translation.IosKey, translation.Translation.En)
				// TH
				iosWriterTH.WriteAttribute(translation.IosKey, translation.Translation.Th)
			}

		}
	}

	iosWriterEN.Close()
	iosWriterTH.Close()
	androidWriterEN.Close()
	androidWriterTH.Close()
}
