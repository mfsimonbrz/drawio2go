package internals

import (
	"drawio2go/models"
	"strings"
)

func GetFieldName(value string) string {
	normalizedValue := strings.Replace(value, "\u00a0", " ", -1)
	fieldInfo := strings.Split(normalizedValue, " ")
	if len(fieldInfo) > 0 {
		fieldName := fieldInfo[0]
		return strings.Trim(fieldName, "")
	}

	return ""
}

func GetFieldType(value string) string {
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

func GetDatabaseFieldType(value string) string {
	if strings.Contains(value, "string") {
		return "text"
	} else if strings.Contains(value, "bool") {
		return "bool"
	} else if strings.Contains(value, "float") {
		return "double"
	} else if strings.Contains(value, "time") {
		return "timestamp"
	}
	return ""
}

func HasTimeField(tables []*models.Table) bool {
	for _, table := range tables {
		for _, field := range table.Fields {
			if field.Type == "time.Time" {
				return true
			}
		}
	}

	return false
}
