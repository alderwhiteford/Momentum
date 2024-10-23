package utilities

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func BuildCreateQuery(dbTable string, item interface{}) (*string, error) {
	val, typ := extractValAndType(item)

	var dbColumnClauses []string
	
	for i := 0 ; i < val.NumField() ; i++ {
		field := val.Field(i)

		if !field.CanInterface() {
			continue
		}

		dbTag := typ.Field(i).Tag.Get("db")
		if dbTag == "" || dbTag == "-" {
			continue
		}

		if field.IsZero() {
			continue
		}

		dbColumnClauses = append(dbColumnClauses, dbTag)
	}

	if len(dbColumnClauses) == 0 {
		return nil, fmt.Errorf("there are no fields to insert into the database")
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		dbTable,
		strings.Join(dbColumnClauses, ", "),
		fmt.Sprintf(":%s", strings.Join(dbColumnClauses, ", :")),
	)

	return &query, nil
}

func BuildUpdateQuery(dbTable string, id uuid.UUID, updates interface{}) (*string, []interface{}, error) {
	val, typ := extractValAndType(updates)

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

func extractValAndType(item interface{}) (reflect.Value, reflect.Type) {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val, val.Type()
}
