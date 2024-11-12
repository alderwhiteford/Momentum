package storage

import (
	"fmt"
	"momentum/utilities"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func InitializeSelectionOnAll[T any](db *PostgresDB, tableName string, entity []T) (error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	
	if err := db.Select(&entity, query); err != nil {
		return utilities.BadRequest("failed finding items")
	}

	return nil
}

func InitializeSingleSelectionOnEntity[T any](db *PostgresDB, tableName string, entity T, id uuid.UUID) (error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	if err := db.Get(&entity, query, id); err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed finding item with id: %s", id))
	}

	return nil
}

func InitializeMultiSelectionOnEntity[T any](db *PostgresDB, tableName string, entity []T, id uuid.UUID, idFieldName string) (error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", tableName, idFieldName)

	if err := db.Select(&entity, query, id); err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed finding items with id: %s", id))
	}

	return nil
}

func InitializeDeletion(db *PostgresDB, tableName string, id uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	
	result, err := db.Exec(query, id)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("error executing delete: %s", err.Error()))
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return utilities.InternalServerError(fmt.Sprintf("something went wrong: %s", err.Error()))
	}
	if rows == 0 {
		return utilities.BadRequest(fmt.Sprintf("item does not exist: %s", id))
	}

	return nil
}

func InitializeCreation(db *PostgresDB, query string, entity interface{}) error {
	_, err := db.NamedExec(query, entity)
	if err != nil {
		return err
	}

	return nil
}

func InitializeMutation(db *PostgresDB, query string, args []interface{}) error {
	result, err := db.Exec(query, args...)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("error executing update: %s", err.Error()))
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return utilities.InternalServerError(fmt.Sprintf("something went wrong: %s", err.Error()))
	}
	if rows == 0 {
		return utilities.BadRequest("failed to find entity in db")
	}

	return nil
}

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
