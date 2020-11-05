package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/assert"
	"reflect"
	l "service_templated/pkg/libservice_template"
	"testing"
)

type Mango struct {
	*l.DefaultEntity
	Color string `name:"color" sql:"VARCHAR(32)"`
}

func (m Mango) Name() string {
	return "mango"
}

func (m Mango) Type() reflect.Type {
	return reflect.TypeOf(m)
}

func (m Mango) Value(key string) (interface{}, error) {
	switch key {
	case "id":
		return m.ID, nil
	case "color":
		return m.Color, nil
	}
	return nil, l.FieldNotFound
}

func TestSave(t *testing.T) {
	repository := PostgresRepository{}
	db, err := sqlx.Open("ramsql", "TestSave")
	assert.Nil(t, err)
	repository.DB = db
	mango := Mango{
		Color: "#ffffff",
	}
	err = repository.CreateTable(mango, true)
	assert.Nil(t, err)
	mango.DefaultEntity = &l.DefaultEntity{}
	err = repository.Save(mango)
	assert.Nil(t, err)
}
