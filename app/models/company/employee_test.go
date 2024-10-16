package company

import (
	"goravel/app/models"
	"reflect"

	"goravel/config/responses"
	"goravel/config/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_TableName(t *testing.T) {
	// Create a new instance of Employee
	employee := Employee{}

	// Check if the TableName() method returns the correct table name
	expectedTableName := tables.TABLE_EMPLOYEE
	actualTableName := employee.TableName()

	// Assert that the table name is correct
	assert.Equal(t, expectedTableName, actualTableName, "Table name should be TABLE_EMPLOYEE")
}

func TestEmployee_Fields(t *testing.T) {
	tableName := tables.TABLE_EMPLOYEE

	type fields struct {
		BaseStruct models.BaseStruct
		Nik        string
		Name       string
		Gender     string
		IsActive   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test with all fields populated",
			fields: fields{
				BaseStruct: models.BaseStruct{ID: 1},
				Nik:        "12345",
				Name:       "John Doe",
				Gender:     "M",
				IsActive:   true,
			},
			want: tableName,
		},
		{
			name: "Test with empty Nik",
			fields: fields{
				BaseStruct: models.BaseStruct{ID: 2},
				Nik:        "",
				Name:       "Jane Doe",
				Gender:     "F",
				IsActive:   false,
			},
			want: tableName,
		},
		{
			name: "Test with empty Name",
			fields: fields{
				BaseStruct: models.BaseStruct{ID: 3},
				Nik:        "67890",
				Name:       "",
				Gender:     "M",
				IsActive:   true,
			},
			want: tableName,
		},
		{
			name: "Test with invalid Gender",
			fields: fields{
				BaseStruct: models.BaseStruct{ID: 4},
				Nik:        "54321",
				Name:       "Chris Smith",
				Gender:     "X", // Assuming valid gender is 'M' or 'F'
				IsActive:   false,
			},
			want: tableName,
		},
		{
			name: "Test with all fields empty",
			fields: fields{
				BaseStruct: models.BaseStruct{ID: 0},
				Nik:        "",
				Name:       "",
				Gender:     "",
				IsActive:   false,
			},
			want: tableName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Employee{
				BaseStruct: tt.fields.BaseStruct,
				Nik:        tt.fields.Nik,
				Name:       tt.fields.Name,
				Gender:     tt.fields.Gender,
				IsActive:   tt.fields.IsActive,
			}
			if got := r.TableName(); got != tt.want {
				t.Errorf("Employee.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetRecords(t *testing.T) {
	type fields struct {
		BaseStruct models.BaseStruct
		Nik        string
		Name       string
		Gender     string
		IsActive   bool
	}
	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   responses.Pagination
		want1  []Employee
	}{
		{
			name: "Get first page with limit 2",
			fields: fields{
				BaseStruct: models.BaseStruct{},
				Nik:        "123456",
				Name:       "John Doe",
				Gender:     "M",
				IsActive:   true,
			},
			args: args{
				page:  1,
				limit: 2,
			},
			want: responses.Pagination{
				Page:         1,
				Limit:        2,
				TotalRecords: 3,
				TotalPage:    2,
			},
			want1: []Employee{
				{Nik: "123456", Name: "John Doe", Gender: "M", IsActive: true},
				{Nik: "654321", Name: "Jane Doe", Gender: "F", IsActive: true},
			},
		},
		{
			name: "Get second page with limit 2",
			fields: fields{
				BaseStruct: models.BaseStruct{},
				Nik:        "123456",
				Name:       "John Doe",
				Gender:     "M",
				IsActive:   true,
			},
			args: args{
				page:  2,
				limit: 2,
			},
			want: responses.Pagination{
				Page:         2,
				Limit:        2,
				TotalRecords: 3,
				TotalPage:    2,
			},
			want1: []Employee{
				{Nik: "789012", Name: "Alice", Gender: "F", IsActive: false},
			},
		},
		{
			name: "Request limit higher than available records",
			fields: fields{
				BaseStruct: models.BaseStruct{},
				Nik:        "123456",
				Name:       "John Doe",
				Gender:     "M",
				IsActive:   true,
			},
			args: args{
				page:  1,
				limit: 5,
			},
			want: responses.Pagination{
				Page:         1,
				Limit:        5,
				TotalRecords: 3,
				TotalPage:    1,
			},
			want1: []Employee{
				{Nik: "123456", Name: "John Doe", Gender: "M", IsActive: true},
				{Nik: "654321", Name: "Jane Doe", Gender: "F", IsActive: true},
				{Nik: "789012", Name: "Alice", Gender: "F", IsActive: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Employee{
				BaseStruct: tt.fields.BaseStruct,
				Nik:        tt.fields.Nik,
				Name:       tt.fields.Name,
				Gender:     tt.fields.Gender,
				IsActive:   tt.fields.IsActive,
			}
			got, got1 := r.GetRecords(tt.args.page, tt.args.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Employee.GetRecords() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Employee.GetRecords() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
