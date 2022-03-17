package i18n

const (
	ENGLISH = "en"
	FRENCH  = "fr"
)

var selectedLanguage = ENGLISH

func ShowText(tag string) string {
	var text string
	switch selectedLanguage {
	case ENGLISH:
		text = findInLanguage(EN, tag)
	case FRENCH:
		text = findInLanguage(FR, tag)
	}
	return text
}

func findInLanguage(languageObject []map[string]string, tag string) string {
	var result string
	for _, item := range languageObject {
		if item[tag] != "" {
			result = item[tag]
		}
	}
	return result
}
