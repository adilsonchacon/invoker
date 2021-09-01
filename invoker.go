package invoker

import (
	"errors"
	"reflect"
)

func CallMethod(iObject interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	inputs := convertInterfaceArgsToReflectValue(args)

	value, ptr := initializeValueAndItsPointer(iObject)

	methodValue := value.MethodByName(methodName)
	methodPtr := ptr.MethodByName(methodName)

	if methodValue.IsValid() {
		return methodValue.Call(inputs), nil
	} else if methodPtr.IsValid() {
		return methodPtr.Call(inputs), nil
	}

	return []reflect.Value{}, errors.New("method not found")
}

func CallMethodAndReturnInt(iObject interface{}, methodName string, args ...interface{}) (int, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(int), nil
	} else {
		return 0, err
	}
}

func CallMethodAndReturnFloat64(iObject interface{}, methodName string, args ...interface{}) (float64, error) {
	reflectValues, error := CallMethod(iObject, methodName, args...)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(float64), nil
	} else {
		return 0.0, error
	}
}

func CallMethodAndReturnString(iObject interface{}, methodName string, args ...interface{}) (string, error) {
	reflectValues, error := CallMethod(iObject, methodName, args...)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(string), nil
	} else {
		return "", error
	}
}

func CallMethodAndReturnBool(iObject interface{}, methodName string, args ...interface{}) (bool, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(bool), nil
	} else {
		return false, err
	}
}

func CallMethodAndReturnSliceOfInt(iObject interface{}, methodName string, args ...interface{}) ([]int, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]int), nil
	} else {
		return []int{}, err
	}
}

func CallMethodAndReturnSliceOfFloat64(iObject interface{}, methodName string, args ...interface{}) ([]float64, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]float64), nil
	} else {
		return []float64{}, err
	}
}

func CallMethodAndReturnSliceOfString(iObject interface{}, methodName string, args ...interface{}) ([]string, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]string), nil
	} else {
		return []string{}, err
	}
}

func CallMethodAndReturnSliceOfBool(iObject interface{}, methodName string, args ...interface{}) ([]bool, error) {
	reflectValues, err := CallMethod(iObject, methodName, args...)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]bool), nil
	} else {
		return []bool{}, err
	}
}

func CallMethodWithArgsAsSliceOfInterface(iObject interface{}, methodName string, args []interface{}) ([]reflect.Value, error) {
	return CallMethod(iObject, methodName, args...)
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(iObject interface{}, methodName string, args []interface{}) (int, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(int), nil
	} else {
		return 0, err
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(iObject interface{}, methodName string, args []interface{}) (float64, error) {
	reflectValues, error := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(float64), nil
	} else {
		return 0.0, error
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnString(iObject interface{}, methodName string, args []interface{}) (string, error) {
	reflectValues, error := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if error == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(string), nil
	} else {
		return "", error
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(iObject interface{}, methodName string, args []interface{}) (bool, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().(bool), nil
	} else {
		return false, err
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(iObject interface{}, methodName string, args []interface{}) ([]int, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]int), nil
	} else {
		return []int{}, err
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(iObject interface{}, methodName string, args []interface{}) ([]float64, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]float64), nil
	} else {
		return []float64{}, err
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(iObject interface{}, methodName string, args []interface{}) ([]string, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]string), nil
	} else {
		return []string{}, err
	}
}

func CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(iObject interface{}, methodName string, args []interface{}) ([]bool, error) {
	reflectValues, err := CallMethodWithArgsAsSliceOfInterface(iObject, methodName, args)

	if err == nil && len(reflectValues) > 0 {
		return reflectValues[0].Interface().([]bool), nil
	} else {
		return []bool{}, err
	}
}

func initializeValueAndItsPointer(iObject interface{}) (reflect.Value, reflect.Value) {
	var ptr reflect.Value
	var value reflect.Value

	switch v := iObject.(type) {
	case reflect.Value:
		value = v
	default:
		value = reflect.ValueOf(iObject)
	}

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(iObject))
		temp := ptr.Elem()
		temp.Set(value)
	}

	return value, ptr
}

func convertInterfaceArgsToReflectValue(args []interface{}) []reflect.Value {
	var inputs []reflect.Value

	for i := range args {
		inputs = append(inputs, reflect.ValueOf(args[i]))
	}

	return inputs
}
