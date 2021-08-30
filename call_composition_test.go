package invoker

import (
	"reflect"
	"strconv"
	"testing"
)

type MyMath struct {
	Operand1 int
	Operand2 int
}

func (myMath *MyMath) Sum() int {
	return myMath.Operand1 + myMath.Operand2
}

func (myMath *MyMath) SumWithExtras(a, b, c int) int {
	return myMath.Sum() + a + b + c
}

func (myMath *MyMath) Average() float64 {
	return float64(myMath.Sum()) / 2.0
}

func (myMath *MyMath) ConvertSumToStr() string {
	return strconv.Itoa(myMath.Sum())
}

func (myMath *MyMath) AreOperandsEqual() bool {
	return myMath.Operand1 == myMath.Operand2
}

func (myMath *MyMath) PutOperandsInSlice() []int {
	return []int{myMath.Operand1, myMath.Operand2}
}

func (myMath *MyMath) PutOperandsBetweenXs() []string {
	return []string{"X" + strconv.Itoa(myMath.Operand1) + "X", "X" + strconv.Itoa(myMath.Operand2) + "X"}
}

func (myMath *MyMath) TransformOperandsToTenPercent() []float64 {
	return []float64{(float64(myMath.Operand1) / 10.0), (float64(myMath.Operand2) / 10.0)}
}

func (myMath *MyMath) AreOperandsEven() []bool {
	return []bool{((myMath.Operand1 % 2) == 0), ((myMath.Operand2 % 2) == 0)}
}

func TestCallMethod(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	value, _ := CallMethod(testMyMath, "Sum")
	intValue := value[0].Interface().(int)
	if intValue == 5 {
		t.Log("CallMethod(testMyMath, \"Sum\") PASSED, returned value is 5")
	} else {
		t.Errorf("CallMethod(testMyMath, \"Sum\") FAILED, returned value expected to be 5, but is %d", intValue)
	}
}

func TestCallMethodAndReturnInt(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	intValue, _ := CallMethodAndReturnInt(testMyMath, "Sum")
	if intValue == 5 {
		t.Log("CallMethodAndReturnInt(testMyMath, \"Sum\") PASSED, returned value is 5")
	} else {
		t.Errorf("CallMethodAndReturnInt(testMyMath, \"Sum\") FAILED, returned value expected to be 5, but is %d", intValue)
	}
}

func TestCallMethodWithParamsAndReturnInt(t *testing.T) {
	var args []interface{}
	args = append(args, 1)
	args = append(args, 2)
	args = append(args, 3)

	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	intValue, _ := CallMethodAndReturnInt(testMyMath, "SumWithExtras", 1, 2, 3)
	if intValue == 11 {
		t.Log("CallMethodAndReturnInt(testMyMath, \"SumWithExtras\") PASSED, returned value is 11")
	} else {
		t.Errorf("CallMethodAndReturnInt(testMyMath, \"SumWithExtras\") FAILED, returned value expected to be 11, but is %d", intValue)
	}
}

func TestCallMethodAndReturnBool(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 2
	boolValue, _ := CallMethodAndReturnBool(testMyMath, "AreOperandsEqual")
	if boolValue {
		t.Log("CallMethodAndReturnBool(testMyMath, \"AreOperandsEqual\") PASSED, returned value is TRUE")
	} else {
		t.Error("CallMethodAndReturnBool(testMyMath, \"AreOperandsEqual\") FAILED, returned value expected to be TRUE, but is FALSE")
	}
}

func TestCallMethodAndReturnFloat64(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 4
	testMyMath.Operand2 = 2
	floatValue, _ := CallMethodAndReturnFloat64(testMyMath, "Average")
	if floatValue == 3.0 {
		t.Log("CallMethodAndReturnFloat64(testMyMath, \"Average\") PASSED, returned value is 3.0")
	} else {
		t.Errorf("CallMethodAndReturnFloat64(testMyMath, \"Average\") FAILED, returned value expected to be 3.0, but is %f", floatValue)
	}
}

func TestCallMethodAndReturnString(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 9
	testMyMath.Operand2 = 9
	stringValue, _ := CallMethodAndReturnString(testMyMath, "ConvertSumToStr")
	if stringValue == "18" {
		t.Log("CallMethodAndReturnString(testMyMath, \"ConvertSumToStr\") PASSED, returned value is \"18\"")
	} else {
		t.Errorf("CallMethodAndReturnString(testMyMath, \"ConvertSumToStr\") FAILED, returned value expected to be \"18\", but is \"%s\"", stringValue)
	}
}

func TestCallMethodAndReturnSliceOfInt(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	intValues, _ := CallMethodAndReturnSliceOfInt(testMyMath, "PutOperandsInSlice")
	if len(intValues) == 2 {
		t.Log("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") PASSED, length of the returned value is 2")

		if intValues[0] == 2 {
			t.Log("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") PASSED, value of fisrt element in slice is 2")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") FAILED, value of fisrt element in slice should be 2, but is %d", intValues[0])
		}

		if intValues[1] == 3 {
			t.Log("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") PASSED, value of second element in slice is 3")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") FAILED, value of second element in slice should be 3, but is %d", intValues[1])
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfInt(testMyMath, \"PutOperandsInSlice\") FAILED, length of the returned value expected to be 2, but is %d", len(intValues))
	}
}

func TestCallMethodAndReturnSliceOfString(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	stringValues, _ := CallMethodAndReturnSliceOfString(testMyMath, "PutOperandsBetweenXs")
	if len(stringValues) == 2 {
		t.Log("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") PASSED, length of the returned value is 2")

		if stringValues[0] == "X2X" {
			t.Log("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") PASSED, value of fisrt element in slice is \"X2X\"")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") FAILED, value of fisrt element in slice should be \"X2X\", but is \"X%sX\"", stringValues[0])
		}

		if stringValues[1] == "X3X" {
			t.Log("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") PASSED, value of second element in slice is \"X3X\"")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") FAILED, value of second element in slice should be 3, but is %s", stringValues[1])
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfString(testMyMath, \"PutOperandsBetweenXs\") FAILED, length of the returned value expected to be 2, but is %d", len(stringValues))
	}
}

func TestCallMethodAndReturnSliceOfFloat64(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	float64Values, _ := CallMethodAndReturnSliceOfFloat64(testMyMath, "TransformOperandsToTenPercent")
	if len(float64Values) == 2 {
		t.Log("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") PASSED, length of the returned value is 2")

		if float64Values[0] == 0.2 {
			t.Log("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") PASSED, value of fisrt element in slice is 0.2")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") FAILED, value of fisrt element in slice should be 0.2, but is %f", float64Values[0])
		}

		if float64Values[1] == 0.3 {
			t.Log("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") PASSED, value of second element in slice is 0.3")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") FAILED, value of second element in slice should be 0.3, but is %f", float64Values[1])
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfFloat64(testMyMath, \"TransformOperandsToTenPercent\") FAILED, length of the returned value expected to be 2, but is %d", len(float64Values))
	}
}

func TestCallMethodAndReturnSliceOfBool(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	boolValues, _ := CallMethodAndReturnSliceOfBool(testMyMath, "AreOperandsEven")
	if len(boolValues) == 2 {
		t.Log("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") PASSED, length of the returned value is 2")

		if boolValues[0] {
			t.Log("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") PASSED, value of fisrt element in slice is TRUE")
		} else {
			t.Error("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") FAILED, value of fisrt element in slice should be TRUE, but is FALSE")
		}

		if !boolValues[1] {
			t.Log("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") PASSED, value of second element in slice is FALSE")
		} else {
			t.Error("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") FAILED, value of second element in slice should be FALSE, but is TRUE")
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfBool(testMyMath, \"AreOperandsEven\") FAILED, length of the returned value expected to be 2, but is %d", len(boolValues))
	}
}

func TestInitializeValueAndItsPointer(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMathValue, testMyMathPointer := initializeValueAndItsPointer(testMyMath)

	switch v := testMyMathValue.Interface().(type) {
	case TestMyMath:
		t.Log("initializeValueAndItsPointer(testMyMath) PASSED, first returned type value is a TestMyMath")
	default:
		t.Errorf("initializeValueAndItsPointer(testMyMath) FAILED, first returned type value is a %v", v)
	}

	switch v := testMyMathPointer.Interface().(type) {
	case *TestMyMath:
		t.Log("initializeValueAndItsPointer(testMyMath, \"sum\") PASSED, second returned type value is a *TestMyMath")
	default:
		t.Errorf("initializeValueAndItsPointer(testMyMath, \"sum\") FAILED, second returned type value is a %v", v)
	}

	testMyMath1 := &TestMyMath{}
	testMyMathValue1, testMyMathPointer1 := initializeValueAndItsPointer(testMyMath1)

	switch v := testMyMathValue1.Interface().(type) {
	case TestMyMath:
		t.Log("initializeValueAndItsPointer(testMyMath, \"sum\") PASSED, first returned type value is a TestMyMath")
	default:
		t.Errorf("initializeValueAndItsPointer(testMyMath, \"sum\") FAILED, first returned type value is a %v", v)
	}

	switch v := testMyMathPointer1.Interface().(type) {
	case *TestMyMath:
		t.Log("initializeValueAndItsPointer(testMyMath, \"sum\") PASSED, second returned type value is a *TestMyMath")
	default:
		t.Errorf("initializeValueAndItsPointer(testMyMath, \"sum\") FAILED, second returned type value is a %v", v)
	}
}

func TestConvertInterfaceArgsToReflectValue(t *testing.T) {
	var iValues interface{}
	var args []interface{}

	args = append(args, 10)
	args = append(args, "abc")
	args = append(args, 2.3)
	args = append(args, true)

	iValues = convertInterfaceArgsToReflectValue(args)
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
