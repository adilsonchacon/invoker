package invoker

import (
	"reflect"
	"testing"
)

func TestConvertInterfaceArgsToReflectValue(t *testing.T) {
	var iValues interface{}
	var args [4]interface{}

	args[0] = 10
	args[1] = "abc"
	args[2] = 2.3
	args[3] = true

	iValues = convertInterfaceArgsToReflectValue(args[0], args[1], args[2], args[3])
	switch v := iValues.(type) {
	case []reflect.Value:
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, returned type is []reflect.Value")
	default:
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected return []reflect.Value, got %v", v)
	}

	values := iValues.([]reflect.Value)
	if len(values) == 4 {
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, length of []reflect.Value is equal to 4")
	} else {
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected length of []reflect.Value if equal to 4, got %d", len(values))
	}

	firstElement := values[0]
	switch v := firstElement.Interface().(type) {
	case int:
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, first returned element is int")
	default:
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FALIED, expected first returned element was int, got %v", v)
	}

	firstValue := firstElement.Interface().(int)
	if firstValue == 10 {
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, first returned element value is equal to 10")
	} else {
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected first returned element value equals to 10, got %d", firstValue)
	}

	secondElement := values[1]
	switch v := secondElement.Interface().(type) {
	case string:
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, second returned element is string")
	default:
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FALIED, expected second returned element is string, got %v", v)
	}

	secondValue := secondElement.Interface().(string)
	if secondValue == "abc" {
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, second returned element value equals to \"abc\"")
	} else {
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected second returned element value equals to \"abc\", got %s", secondValue)
	}

	thirdElement := values[2]
	switch v := thirdElement.Interface().(type) {
	case float64:
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, third returned element type is float64")
	default:
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FALIED, expected third returned element type was float64, got %v", v)
	}

	thirdValue := thirdElement.Interface().(float64)
	if thirdValue == 2.3 {
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, third returned element value equals to 2.3")
	} else {
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected third returned element value equals to 2.3, got %f", thirdValue)
	}

	fourthElement := values[3]
	switch v := fourthElement.Interface().(type) {
	case bool:
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, fourth returned element type is boolean")
	default:
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FALIED, expected fourth returned element type was boolean, got %v", v)
	}

	fourthValue := fourthElement.Interface().(bool)
	if fourthValue {
		t.Log("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) PASSED, fourth returned element value is true")
	} else {
		t.Errorf("convertInterfaceArgsToReflectValue(10, \"abc\", 2.3, true) FAILED, expected fourth returned element value was true, got false")
	}
}
