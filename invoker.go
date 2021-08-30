package invoker

import (
	"errors"
	"fmt"
	"reflect"
)

func CallMethod(iValue interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	inputs := convertInterfaceArgsToReflectValue(args)

	value, ptr := initializeValueAndItsPointer(iValue)

	methodValue := value.MethodByName(methodName)
	methodPtr := ptr.MethodByName(methodName)

	if methodValue.IsValid() {
		return methodValue.Call(inputs), nil
	} else if methodPtr.IsValid() {
		return methodPtr.Call(inputs), nil
	}

	return []reflect.Value{}, errors.New("method not found")
}

func CallMethodAndReturnBool(iValue interface{}, methodName string, args ...interface{}) (bool, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(bool), nil
	} else {
		return false, err
	}
}

func CallMethodAndReturnInt(iValue interface{}, methodName string, args ...interface{}) (int, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(int), nil
	} else {
		return 0, err
	}
}

func CallMethodAndReturnFloat64(iValue interface{}, methodName string, args ...interface{}) (float64, error) {
	reflectValues, error := CallMethod(iValue, methodName, args...)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(float64), nil
	} else {
		return 0.0, error
	}
}

func CallMethodAndReturnString(iValue interface{}, methodName string, args ...interface{}) (string, error) {
	reflectValues, error := CallMethod(iValue, methodName, args...)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(string), nil
	} else {
		return "", error
	}
}

func CallMethodAndReturnSliceOfInt(iValue interface{}, methodName string, args ...interface{}) ([]int, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]int), nil
	} else {
		return []int{}, err
	}
}

func CallMethodAndReturnSliceOfBool(iValue interface{}, methodName string, args ...interface{}) ([]bool, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]bool), nil
	} else {
		return []bool{}, err
	}
}

func CallMethodAndReturnSliceOfFloat64(iValue interface{}, methodName string, args ...interface{}) ([]float64, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]float64), nil
	} else {
		return []float64{}, err
	}
}

func CallMethodAndReturnSliceOfString(iValue interface{}, methodName string, args ...interface{}) ([]string, error) {
	reflectValues, err := CallMethod(iValue, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]string), nil
	} else {
		return []string{}, err
	}
}

func initializeValueAndItsPointer(iValue interface{}) (reflect.Value, reflect.Value) {
	var ptr reflect.Value
	var value reflect.Value

	switch v := iValue.(type) {
	case reflect.Value:
		value = v
	default:
		value = reflect.ValueOf(iValue)
	}

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(iValue))
		temp := ptr.Elem()
		temp.Set(value)
	}

	return value, ptr
}

func convertInterfaceArgsToReflectValue(args []interface{}) []reflect.Value {
	fmt.Println("convertInterfaceArgsToReflectValue: ")
	var inputs []reflect.Value

	for i := range args {
		inputs = append(inputs, reflect.ValueOf(args[i]))
	}

	return inputs
}
