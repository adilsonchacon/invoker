package invoker

import "reflect"

func convertInterfaceArgsToReflectValue(args ...interface{}) []reflect.Value {
	var inputs []reflect.Value
	for i, _ := range args {
		inputs = append(inputs, reflect.ValueOf(args[i]))
	}

	return inputs
}
