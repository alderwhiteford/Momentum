package utilities

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func BuildUpdateQuery(dbTable string, id uuid.UUID, updates interface{}) (*string, []interface{}, error) {
	val := reflect.ValueOf(updates)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	var dbSetClauses []string
	var args []interface{}
	argPosition := 1

	// For i in number of fields in struct
	for i := 0 ; i < val.NumField() ; i++ {
		field := val.Field(i)

		if !field.CanInterface() {
			continue
		}

		// Ignore the ID field
		if typ.Field(i).Name == "ID" {
			continue
		}

		// Get the db tag
		dbTag := typ.Field(i).Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue
		}

		// Ignore zero values
		if field.IsZero() {
			continue
		}

		dbSetClauses = append(dbSetClauses, fmt.Sprintf("%s = $%d", dbTag, argPosition))
		args = append(args, field.Interface())
		argPosition += 1
	}

	if len(dbSetClauses) == 0 {
		return nil, nil, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		dbTable,
		strings.Join(dbSetClauses, ", "),
		argPosition,
	)

	args = append(args, id)

	return &query, args, nil
}
