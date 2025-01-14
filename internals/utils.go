package internals

import (
	"bytes"
	"drawio2go/models"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getFieldName(value string) string {
	normalizedValue := strings.Replace(value, "\u00a0", " ", -1)
	fieldInfo := strings.Split(normalizedValue, " ")
	if len(fieldInfo) > 0 {
		fieldName := fieldInfo[0]
		return strings.Trim(fieldName, "")
	}

	return ""
}

func getFieldType(value string) string {
	normalizedValue := strings.Replace(value, "\u00a0", " ", -1)
	fieldValues := strings.Split(normalizedValue, " ")
	if len(fieldValues) > 1 {
		// first position is the field name
		originalFieldType := strings.Trim(fieldValues[1], "")
		if strings.Contains(originalFieldType, "char") {
			return "string"
		} else if strings.Contains(originalFieldType, "bool") {
			return "bool"
		} else if strings.Contains(originalFieldType, "float") || strings.Contains(originalFieldType, "double") {
			return "float64"
		} else if strings.Contains(originalFieldType, "int") {
			return "int"
		} else if strings.Contains(originalFieldType, "date") {
			return "time.Time"
		}
		return ""
	}

	return ""
}

func getDatabaseFieldType(value string) string {
	if strings.Contains(value, "string") {
		return "text"
	} else if strings.Contains(value, "bool") {
		return "bool"
	} else if strings.Contains(value, "float") {
		return "double"
	} else if strings.Contains(value, "time") {
		return "timestamp"
	} else if strings.Contains(value, "int") {
		return "integer"
	}
	return ""
}

func hasTimeField(tables []*models.Table) bool {
	for _, table := range tables {
		for _, field := range table.Fields {
			if field.Type == "time.Time" {
				return true
			}
		}
	}

	return false
}

func title(s string) string {
	caser := cases.Title(language.Und)
	return caser.String(s)
}

func normalizeFieldName(field string) string {
	var sb bytes.Buffer
	skipUnderscore := false
	for pos, letter := range field {
		if pos == 0 {
			sb.WriteString(strings.ToUpper(string(letter)))
		} else {
			if string(letter) == "_" {
				skipUnderscore = true
				continue
			}

			if skipUnderscore {
				sb.WriteString(strings.ToUpper(string(letter)))
				skipUnderscore = false
			} else {
				sb.WriteString(string(letter))
			}
		}
	}

	return sb.String()
}

func minus(a int, b int) int {
	return a - b
}
