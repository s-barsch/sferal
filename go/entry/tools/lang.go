package tools

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Langs = map[string]string{
	"de": "Deutsch",
	"en": "English",
}

var AboutName = map[string]string{
	"de": "über",
	"en": "about",
}

var KineName = map[string]string{
	"de": "reels",
	"en": "reels",
}

func Title(str string) string {
	return cases.Title(language.German).String(str)
}
