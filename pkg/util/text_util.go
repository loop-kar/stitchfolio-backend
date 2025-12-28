package util

import "strings"

func ReplaceTemplateValues(content []byte, templateValues map[string]string) []byte {
	htmlString := string(content)

	for k, v := range templateValues {
		htmlString = strings.Replace(htmlString, k, v, 10)
	}

	return []byte(htmlString)
}
