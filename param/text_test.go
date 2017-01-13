package parameters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModel struct {
	Name   string
	Create string
}

func TestToDBNames(t *testing.T) {
	fields := ToDBNames([]string{"Name", "Create"})
	assertFields := []string{"name", "create"}
	for i, s := range fields {
		assert.Equal(t, assertFields[i], s)
	}
}

func TestFieldsOf(t *testing.T) {
	fields := FieldsOf(testModel{})
	t.Log(fields)
}

func testOrder(s string, fields []string, field string, isAsc bool, assert *assert.Assertions) {
	order := StrToOrder(s, fields)
	assert.NotNil(order)
	assert.Equal(field, order.Field)
	assert.Equal(isAsc, order.Asc)
}

func TestStrToOrder(t *testing.T) {
	assert := assert.New(t)
	fields := ToDBNames(FieldsOf(testModel{}))

	testOrder("name", fields, "name", true, assert)
	testOrder("+name", fields, "name", true, assert)
	testOrder("-name", fields, "name", false, assert)

	assert.Nil(StrToOrder("", fields))
}

func TestStrToOrders(t *testing.T) {
	model := testModel{}

	assert.Nil(t, StrToOrders("", model))

	assert.Equal(t, 1, len(StrToOrders("name", model)))
	assert.Equal(t, 1, len(StrToOrders("+name", model)))
	assert.Equal(t, 1, len(StrToOrders("-name", model)))

	assert.Equal(t, 2, len(StrToOrders("name, -create", model)))
	assert.Equal(t, 2, len(StrToOrders("+name, +create", model)))
	assert.Equal(t, 2, len(StrToOrders("-name, -create", model)))

	assert.Equal(t, 2, len(StrToOrders("-name, +create", model)))

}

func TestOrderToStr(t *testing.T) {
	str := ""
	str = OrderToStr([]Order{
		Order{"name", false},
	})
	assert.Equal(t, "name desc", str)

	str = OrderToStr([]Order{
		Order{"name", true},
	})
	assert.Equal(t, "name", str)

	str = OrderToStr([]Order{
		Order{"name", false},
		Order{"create", true},
	})
	assert.Equal(t, "name desc, create", str)

	str = OrderToStr([]Order{
		Order{"name", true},
		Order{"create", false},
	})
	assert.Equal(t, "name, create desc", str)

	str = OrderToStr([]Order{})
	assert.Equal(t, "", str)

	str = OrderToStr(nil)
	assert.Equal(t, "", str)
}
