package test

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"

	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

// GormHelper struct definition
type GormHelper struct {
}

func (h *GormHelper) FixedFullRe(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

// MockRowBuilder type definition
type MockRowBuilder struct {
	rows *sqlmock.Rows
}

// Of creates rows for the struct
func (rb *MockRowBuilder) Of(v interface{}) *MockRowBuilder {
	names, _ := rb.listFields(v)
	if rb.rows == nil {
		rb.rows = sqlmock.NewRows(names)
	}
	return rb
}

// Add will add new row for returning.
func (rb *MockRowBuilder) Add(v interface{}) *MockRowBuilder {
	names, fields := rb.listFields(v)
	if rb.rows == nil {
		rb.rows = sqlmock.NewRows(names)
	}
	rb.rows = rb.rows.AddRow(fields[:]...)
	return rb
}

// Build will return mocked sql rows.
func (rb *MockRowBuilder) Build() *sqlmock.Rows {
	return rb.rows
}

func (rb *MockRowBuilder) listFields(a interface{}) ([]string, []driver.Value) {
	elements := reflect.ValueOf(a)
	var names []string
	var values []driver.Value

	for i := 0; i < elements.NumField(); i++ {
		name := gorm.ToDBName(elements.Type().Field(i).Name)
		if name == "db_model" {
			continue
		}
		names = append(names, name)
		values = append(values, elements.Field(i).Interface())
	}

	return names, values
}
