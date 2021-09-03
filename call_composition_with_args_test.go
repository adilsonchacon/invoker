package invoker

import "testing"

func TestCallMethodWithArgsAsSliceOfInterface(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, 1)
	args = append(args, "Total: ")

	value, _ := CallMethodWithArgsAsSliceOfInterface(testMyMath, "SumWithExtraIntAndPrefixString", args)
	stringValue := value[0].Interface().(string)
	if stringValue == "Total: 6" {
		t.Log("CallMethod(testMyMath, \"Sum\") PASSED, returned value is \"Total: 6\"")
	} else {
		t.Errorf("CallMethod(testMyMath, \"Sum\") FAILED, returned value expected to be \"Total: 6\", but is %s", stringValue)
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnInt(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, 1)
	args = append(args, 2)
	args = append(args, 3)

	intValue, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, "SumWithExtras", args)
	if intValue == 11 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, \"SumWithExtras\") PASSED, returned value is 11")
	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, \"SumWithExtras\") FAILED, returned value expected to be 11, but is %d", intValue)
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, 3.0)
	args = append(args, 2.0)

	float64Value, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, "WeightedAverage", args)
	if float64Value == 2.4 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, \"WeightedAverage\") PASSED, returned value is 2.4")
	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, \"WeightedAverage\") FAILED, returned value expected to be 2.4, but is %f", float64Value)
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnString(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, 1)
	args = append(args, "Total: ")

	stringValue, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, "SumWithExtraIntAndPrefixString", args)
	if stringValue == "Total: 6" {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, \"SumWithExtraIntAndPrefixString\") PASSED, returned value is \"Total: 6\"")
	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, \"SumWithExtraIntAndPrefixString\") FAILED, returned value expected to be \"Total: 6\", but is \"%s\"", stringValue)
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnBool(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 7
	testMyMath.Operand2 = 10

	args = append(args, 2)

	boolValue, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, "IsFirstOperandBiggerThanArgValue", args)
	if boolValue {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, \"IsFirstOperandBiggerThanArgValue\") PASSED, returned value is TRUE")
	} else {
		t.Error("CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, \"IsFirstOperandBiggerThanArgValue\") FAILED, returned value expected to be TRUE, but is FALSE")
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, 10)

	intValues, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, "MultiplyOperandsByArgValue", args)
	if len(intValues) == 2 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") PASSED, length of the returned value is 2")

		if intValues[0] == 20 {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") PASSED, value of fisrt element in slice is 20")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") FAILED, value of fisrt element in slice should be 20, but is %d", intValues[0])
		}

		if intValues[1] == 30 {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") PASSED, value of second element in slice is 30")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") FAILED, value of second element in slice should be 30, but is %d", intValues[1])
		}

	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, \"MultiplyOperandsByArgValue\") FAILED, length of the returned value expected to be 2, but is %d", len(intValues))
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	args = append(args, "X")

	stringValues, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, "PutOperandsBetweenString", args)
	if len(stringValues) == 2 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") PASSED, length of the returned value is 2")

		if stringValues[0] == "X2X" {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") PASSED, value of fisrt element in slice is \"X2X\"")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") FAILED, value of fisrt element in slice should be \"X2X\", but is \"X%sX\"", stringValues[0])
		}

		if stringValues[1] == "X3X" {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") PASSED, value of second element in slice is \"X3X\"")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") FAILED, value of second element in slice should be 3, but is %s", stringValues[1])
		}

	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenString\") FAILED, length of the returned value expected to be 2, but is %d", len(stringValues))
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 4

	args = append(args, 2.0)

	float64Values, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, "DivideOperandsByArgValue", args)
	if len(float64Values) == 2 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") PASSED, length of the returned value is 2")

		if float64Values[0] == 1.0 {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") PASSED, value of fisrt element in slice is 1.0")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") FAILED, value of fisrt element in slice should be 1.0, but is %f", float64Values[0])
		}

		if float64Values[1] == 2.0 {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") PASSED, value of second element in slice is 2.0")
		} else {
			t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") FAILED, value of second element in slice should be 2.0, but is %f", float64Values[1])
		}

	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, \"DivideOperandsByArgValue\") FAILED, length of the returned value expected to be 2, but is %d", len(float64Values))
	}
}

func TestCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 7
	testMyMath.Operand2 = 12

	args = append(args, 5)

	boolValues, _ := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, "AreOperandsBiggerThanArgValue", args)
	if len(boolValues) == 2 {
		t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") PASSED, length of the returned value is 2")

		if boolValues[0] {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") PASSED, value of fisrt element in slice is TRUE")
		} else {
			t.Error("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") FAILED, value of fisrt element in slice should be TRUE, but is FALSE")
		}

		if boolValues[1] {
			t.Log("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") PASSED, value of second element in slice is TRUE")
		} else {
			t.Error("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") FAILED, value of second element in slice should be TRUE, but is FALSE")
		}

	} else {
		t.Errorf("CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, \"AreOperandsBiggerThanArgValue\") FAILED, length of the returned value expected to be 2, but is %d", len(boolValues))
	}
}
