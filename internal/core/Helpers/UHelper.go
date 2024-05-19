package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func TextParse(str string, separating byte) []string {
	length := len(str)
	parsedString := make([]string, 0)
	last := 0
	for i := 0; i < length; i++ {

		if str[i] == separating {
			parsedString = append(parsedString, str[last:i])
			last = i + 1
		}
	}
	if str[last:] != "" {
		parsedString = append(parsedString, str[last:])
	}
	return parsedString
}
func ArrayPrint(arr []string) {
	for i, v := range arr {
		fmt.Printf("%d = %v \n", i, v)
	}
}
func StrToInt(str string) int {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Dönüştürme hatası:", err)
		return 0
	}
	return int(number)
}
func InterfaceToStruct(data interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, result); err != nil {
		return err
	}

	return nil
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
func LastIndex(arr interface{}, fieldName string) interface{} {
	arrValue := reflect.ValueOf(arr)
	if arrValue.Kind() != reflect.Slice || arrValue.Len() == 0 {
		return nil
	}
	lastElem := arrValue.Index(arrValue.Len() - 1)
	elemType := lastElem.Type()
	for i := 0; i < elemType.NumField(); i++ {
		if elemType.Field(i).Name == fieldName {

			return lastElem.Field(i).Interface()
		}
	}
	return nil
}
func GetField[T any](v *T, field string) reflect.Value {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}
func SetFieldValue(obj interface{}, fieldName string, value interface{}) error {
	structValue := reflect.ValueOf(obj)
	if structValue.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer to a struct")
	}

	fieldValue := structValue.Elem().FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("cannot set field %s value", fieldName)
	}

	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(value.(string))
	case reflect.Int:
		intValue, ok := value.(int)
		if !ok {
			return fmt.Errorf("value %v is not an int", value)
		}
		fieldValue.SetInt(int64(intValue))
	case reflect.Bool:
		boolValue, ok := value.(bool)
		if !ok {
			return fmt.Errorf("value %v is not a bool", value)
		}
		fieldValue.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported field type %s", fieldValue.Kind())
	}

	return nil
}
