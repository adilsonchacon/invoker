package invoker

import "reflect"

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

func convertInterfaceArgsToReflectValue(args ...interface{}) []reflect.Value {
	var inputs []reflect.Value
	for i, _ := range args {
		inputs = append(inputs, reflect.ValueOf(args[i]))
	}

	return inputs
}
