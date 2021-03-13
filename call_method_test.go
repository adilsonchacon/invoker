package invoker

import (
	"reflect"
	"strconv"
	"testing"
)

// ----------- Start of Interface -----------
type MySetInterface interface {
	Add(value int)
	GetValues() []int
}

type MySet struct {
	Values map[int]bool
}

func (mySet *MySet) Add(value int) {
	if mySet.Values == nil {
		mySet.Values = make(map[int]bool)
	}

	mySet.Values[value] = true
}

func (mySet *MySet) GetValues() []int {
	var values []int

	if mySet.Values != nil {
		for value := range mySet.Values {
			values = append(values, value)
		}
	}

	return values
}

// ----------- End of Interface -----------

// ----------- Start of Composition -----------
type MyMath struct {
	Operand1 int
	Operand2 int
}

func (myMath *MyMath) Sum() int {
	return myMath.Operand1 + myMath.Operand2
}

func (myMath *MyMath) IsSumZero() bool {
	return myMath.Sum() == 0
}

func (myMath *MyMath) Average() float64 {
	return float64(myMath.Sum()) / 2.0
}

func (myMath *MyMath) ConvertSumToStr() string {
	return strconv.Itoa(myMath.Sum())
}

func (myMath *MyMath) GetOperandsAsInt() []int {
	return []int{myMath.Operand1, myMath.Operand2}
}

func (myMath *MyMath) GetOperandsAsString() []string {
	return []string{strconv.Itoa(myMath.Operand1), strconv.Itoa(myMath.Operand2)}
}

func (myMath *MyMath) GetOperandsAsFloat64() []float64 {
	return []float64{float64(myMath.Operand1), float64(myMath.Operand2)}
}

func (myMath *MyMath) AreOperandsEven() []bool {
	return []bool{(myMath.Operand1%2 == 0), (myMath.Operand2%2 == 0)}
}

// ----------- End of Composition -----------

func TestInitializeValueAndItsPointer(t *testing.T) {
	type TestMyMath struct {
		MyMath
	}

	testMyMath := TestMyMath{}
	testMyMathValue, testMyMathPointer := initializeValueAndItsPointer(testMyMath)

	switch v := testMyMathValue.Interface().(type) {
	case TestMyMath:
		t.Log("initializeValueAndItsPointer(testMyMath, \"sum\") PASSED, first returned type value is a TestMyMath")
	default:
		t.Errorf("initializeValueAndItsPointer(testMyMath, \"sum\") FAILED, first returned type value is a %v", v)
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
