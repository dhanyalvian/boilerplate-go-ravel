package company

import (
	"goravel/config/tables"
	"testing"

	"github.com/goravel/framework/database/orm"
	"github.com/stretchr/testify/assert"
)

func TestUser_TableName(t *testing.T) {
	// Create a new instance of Employee
	user := User{}

	// Check if the TableName() method returns the correct table name
	expectedTableName := tables.TABLE_USER
	actualTableName := user.TableName()

	// Assert that the table name is correct
	assert.Equal(t, expectedTableName, actualTableName, "Table name should be TABLE_USER")
}

func TestUser_Fields(t *testing.T) {
	tableName := tables.TABLE_USER

	type fields struct {
		Model    orm.Model
		Uid      string
		Username string
		Name     string
		IsActive bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test with all fields populated",
			fields: fields{
				Model:    orm.Model{ID: 1},
				Uid:      "0420a006-daee-4cd8-84b9-932121224248",
				Username: "john.doe",
				Name:     "John Doe",
				IsActive: true,
			},
			want: tableName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &User{
				Model:    tt.fields.Model,
				Uid:      tt.fields.Uid,
				Username: tt.fields.Username,
				Name:     tt.fields.Name,
				IsActive: tt.fields.IsActive,
			}
			if got := r.TableName(); got != tt.want {
				t.Errorf("User.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
