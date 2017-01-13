package parameters

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

func ToDBName(str string) string {
	return gorm.ToDBName(str)
}

func ToDBNames(strs []string) []string {
	list := make([]string, len(strs))
	for i, s := range strs {
		list[i] = gorm.ToDBName(s)
	}
	return list
}

func String(str string, valids []string) string {
	for _, s := range valids {
		if s == str {
			return s
		}
	}
	return ""
}

type Order struct {
	Field string
	Asc   bool
}

func FieldsOf(model interface{}) (fields []string) {
	tp := reflect.TypeOf(model)
	len := tp.NumField()
	for i := 0; i < len; i++ {
		fields = append(fields, tp.Field(i).Name)
	}
	return
}

// orders like  +name,-create
func StrToOrders(str string, model interface{}) []Order {
	if str == "" {
		return nil
	}
	list := strings.Split(str, ",")
	orders := make([]Order, 0)
	fields := ToDBNames(FieldsOf(model))
	for _, item := range list {
		order := StrToOrder(item, fields)
		if order != nil {
			orders = append(orders, *order)
		}
	}
	return orders
}

func StrToOrder(str string, fields []string) *Order {
	str = strings.TrimSpace(str)

	field := ""
	isAsc := false

	if str == "" {
		return nil
	}

	firstLetter := str[0:1]

	if firstLetter == "+" {
		isAsc = true
		field = String(ToDBName(str[1:]), fields)
	} else if firstLetter == "-" {
		isAsc = false
		field = String(ToDBName(str[1:]), fields)
	} else {
		isAsc = true
		field = String(ToDBName(str), fields)
	}

	if field == "" {
		return nil
	}

	return &Order{field, isAsc}
}

func OrderToStr(orders []Order) string {
	if len(orders) == 0 {
		return ""
	}

	order := orders[0]
	str := order.Field
	if !order.Asc {
		str += " desc"
	}

	if len(orders) == 1 {
		return str
	}

	for _, order := range orders[1:] {
		str += ", " + order.Field
		if !order.Asc {
			str += " desc"
		}
	}
	return str
}
