package strings

import "database/sql"

// NullStringToString nullstring转换成string
func NullStringToString(nullStrings []sql.NullString) []string {
	var strings []string
	for _, nullString := range nullStrings {
		if nullString.Valid {
			strings = append(strings, nullString.String)
		}
	}
	return strings
}
