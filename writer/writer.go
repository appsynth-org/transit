package writer

import (
	"github.com/appsynth-org/transit/reader"
	android "github.com/appsynth-org/transit/writer/android"
	ios "github.com/appsynth-org/transit/writer/ios"
)

func GenerateLocaleFiles(groups []reader.LocalizeGroup) {
	iosWriterEN := ios.NewDocument("en")
	iosWriterTH := ios.NewDocument("th")

	androidWriterEN := android.NewDocument("en")
	androidWriterTH := android.NewDocument("th")

	for _, group := range groups {
		iosWriterEN.WriteComment(group.GroupName)
		iosWriterTH.WriteComment(group.GroupName)

		androidWriterEN.WriteComment(group.GroupName)
		androidWriterTH.WriteComment(group.GroupName)

		for _, translation := range group.Keys {
			// Skip Android empty key
			if len(translation.AndroidKey) > 0 {
				androidWriterEN.WriteAttribute(translation.AndroidKey, translation.Translation.En)
				androidWriterTH.WriteAttribute(translation.AndroidKey, translation.Translation.Th)
			}

			// Skip iOS empty key
			if (len(translation.IosKey)) > 0 {
				iosWriterEN.WriteAttribute(translation.IosKey, translation.Translation.En)
				iosWriterTH.WriteAttribute(translation.IosKey, translation.Translation.Th)
			}

		}
	}

	iosWriterEN.Close()
	iosWriterTH.Close()

	androidWriterEN.Close()
	androidWriterTH.Close()
}
