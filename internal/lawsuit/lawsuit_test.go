package lawsuit_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/raninho/stiuswal/internal/lawsuit"
	"github.com/raninho/stiuswal/test"
)

func TestLawsuitStructs(t *testing.T) {
	type Fields struct {
		fieldName string
		fieldType string
		tags      string
	}

	var tt = []struct {
		modelName   string
		modelStruct interface{}
		fields      []Fields
	}{
		{"Lawsuit", &lawsuit.Lawsuit{}, []Fields{
			{"ID", "uint", `gorm:"primary_key"`},
			{"OrderID", "string", `gorm:"column:order_id"`},
			{"LawsuitNumber", "string", `gorm:"column:lawsuit_number"`},
			{"Processed", "bool", `gorm:"column:processed"`},
			{"Criticized", "bool", `gorm:"column:criticized"`},
			{"CreatedAt", "time.Time", `gorm:"column:created_at"`},
			{"UpdatedAt", "time.Time", `gorm:"column:updated_at"`},
			{"Output", "postgres.Jsonb", `gorm:"column:output"`},
		}},
	}
	for _, tc := range tt {

		reflectionElements := reflect.ValueOf(tc.modelStruct).Elem()
		for i, field := range tc.fields {

			testName := fmt.Sprintf("%s %s", tc.modelName, field.fieldName)
			t.Run(testName, func(t *testing.T) {

				// Check field name
				v := reflectionElements.FieldByName(field.fieldName)
				assert.True(t, v.IsValid(), "Field '%s' not found in the struct!", field.fieldName)

				// Check field type
				ts := fmt.Sprintf("%v", v.Type())
				assert.Equal(t, ts, field.fieldType,
					"Field '%s' should be of type '%s', but received '%s'",
					field.fieldName, field.fieldType, ts)

				// Check field tag
				tag := fmt.Sprintf("%s", reflectionElements.Type().Field(i).Tag)
				assert.Equal(t, tag, field.tags,
					"Field '%s' tag changed. Tag expected: %s",
					field.fieldName, tag)
			})

		}

		assert.Equal(t, len(tc.fields), reflectionElements.NumField(),
			"Found %d elements in struct (%s) while waiting for %d",
			len(tt), tc.modelName, reflectionElements.NumField())
	}
}

func TestLawsuit_TableName(t *testing.T) {
	ls := lawsuit.Lawsuit{}

	if ls.TableName() != lawsuit.LawsuitTableName {
		t.Fatal(errors.New("LawsuitTableName is invalid"))
	}
}

func TestIsLawsuitNumber(t *testing.T) {
	if !lawsuit.IsLawsuitNumber("0019528-16.2005.8.02.0001") {
		t.Fatal("IsLawsuitNumber is invalid")
	}

	if !lawsuit.IsLawsuitNumber("0055534-96.2012.8.12.0001") {
		t.Fatal("IsLawsuitNumber is invalid")
	}

	if lawsuit.IsLawsuitNumber("0055534-96.2012.8.120001") {
		t.Fatal("IsLawsuitNumber is invalid")
	}

	if lawsuit.IsLawsuitNumber("0055534-9620128120001") {
		t.Fatal("IsLawsuitNumber is invalid")
	}

	if lawsuit.IsLawsuitNumber("00555349620128120001") {
		t.Fatal("IsLawsuitNumber is invalid")
	}
}

func TestLawsuit_Ok(t *testing.T) {
	ls := lawsuit.Lawsuit{}

	ls.LawsuitNumber = "0019528-16.2005.8.02.0001"
	if err := ls.Ok(); err != nil {
		t.Fatal("Ok is invalid")
	}

	ls.LawsuitNumber = "001952816.2005.8.02.0001"
	if err := ls.Ok(); err == nil {
		t.Fatal("Ok is invalid")
	}
}

func TestLawsuit_Create(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	ls := lawsuit.New("0019528-16.2005.8.02.0001")

	if _, err := lawsuit.GetByOrderID(db, ls.OrderID); err == nil {
		t.Fatal("lawsuit exists")
	}

	if err := lawsuit.Create(db, ls); err != nil {
		t.Fatal(err.Error())
	}

	if _, err := lawsuit.GetByOrderID(db, ls.OrderID); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLawsuit_Update(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	ls := lawsuit.New("0019528-16.2005.8.02.0001")

	if _, err := lawsuit.GetByOrderID(db, ls.OrderID); err == nil {
		t.Fatal("lawsuit exists")
	}

	if err := lawsuit.Create(db, ls); err != nil {
		t.Fatal(err.Error())
	}

	ls2, err := lawsuit.GetByOrderID(db, ls.OrderID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ls2.Processed {
		t.Fatal("Update with error")
	}

	ls.Processed = true

	if err := lawsuit.Update(db, ls); err != nil {
		t.Fatal(err.Error())
	}

	ls3, err := lawsuit.GetByOrderID(db, ls.OrderID)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !ls3.Processed {
		t.Fatal("Update with error")
	}
}
