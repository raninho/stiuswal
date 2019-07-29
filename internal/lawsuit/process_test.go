package lawsuit_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/raninho/stiuswal/internal/lawsuit"
	"github.com/raninho/stiuswal/test"
	"github.com/stretchr/testify/assert"
)

func TestProcessFinishedInputStructs(t *testing.T) {
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
		{"ProcessFinishedInput", &lawsuit.ProcessFinishedInput{}, []Fields{
			{"OrderID", "string", `json:"order-id"`},
			{"LawsuitNumber", "string", `json:"lawsuit-number"`},
			{"Criticized", "bool", `json:"criticized"`},
			{"Output", "json.RawMessage", `json:"output"`},
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

func TestProcessFinishedInput_Ok(t *testing.T) {
	i := lawsuit.ProcessFinishedInput{}

	i.LawsuitNumber = "0019528-16.2005.8.02.0001"
	i.OrderID = "e9f7f07c-b805-45c6-90bc-21f273b27d63"
	if err := i.Ok(); err != nil {
		t.Fatal("Ok is invalid")
	}

	i.LawsuitNumber = "001952816.2005.8.02.0001"
	i.OrderID = "e9f7f07c-b805-45c6-90bc-21f273b27d63"
	if err := i.Ok(); err == nil {
		t.Fatal("Ok is valid")
	}

	i.LawsuitNumber = "0019528-16.2005.8.02.0001"
	i.OrderID = "e9f7f07cb805-45c6-90bc-21f273b27d63"
	if err := i.Ok(); err == nil {
		t.Fatal("Ok is valid")
	}
}

func TestProcessFinished(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	db.AutoMigrate(lawsuit.Lawsuit{})

	ls := lawsuit.New("0019528-16.2005.8.02.0001")
	err := lawsuit.Create(db, ls)
	if err != nil {
		t.Fatal(err.Error())
	}

	input := &lawsuit.ProcessFinishedInput{}
	input.OrderID = ls.OrderID
	input.LawsuitNumber = ls.LawsuitNumber

	ls2, err := lawsuit.ProcessFinished(db, input)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !ls2.Processed {
		t.Fatal(errors.New("ProcessFinished is wrong"))
	}

	ls3, err := lawsuit.GetByOrderIDAndLawsuitNumber(db, ls2.OrderID, ls2.LawsuitNumber)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !ls3.Processed {
		t.Fatal(errors.New("ProcessFinished is wrong"))
	}
}

func TestProcess(t *testing.T) {
	db := test.NewDBTest()
	defer db.Close()
	defer db.DropTable(lawsuit.Lawsuit{}.TableName())

	queue := test.NewQueueTest()
	defer queue.Close()

	db.AutoMigrate(lawsuit.Lawsuit{})

	lawsuitNumber := "0019528-16.2005.8.02.0001"

	ls, err := lawsuit.Process(db, queue, lawsuitNumber)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ls.LawsuitNumber != lawsuitNumber {
		t.Fatal("ls.LawsuitNumber != lawsuitNumber ")
	}
}
