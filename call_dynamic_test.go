package invoker

import (
	"reflect"
	"strconv"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Info MySet
}

func (person *Person) Stringfy() string {
	return person.Name + " (" + strconv.Itoa(person.Age) + ")"
}

func TestCallMethodFromAnInterface(t *testing.T) {
	tPerson := Person{}
	tPerson.Name = "John"
	tPerson.Age = 20

	var myInterface interface{}
	myInterface = &tPerson

	iInfo := reflect.ValueOf(myInterface).Interface()
	stringValue, _ := CallMethodAndReturnString(iInfo, "Stringfy")

	if stringValue == "John (20)" {
		t.Log("CallMethodAndReturnString(iInfo, \"Stringfy\") PASSED, returned value is \"John (20)\"")
	} else {
		t.Errorf("CallMethodAndReturnString(iInfo, \"Stringfy\") FAILED, returned value expected to be \"John (20)\", but is \"%s\"", stringValue)
	}
}

func TestCallMethodSubClassFromAnAInterface(t *testing.T) {
	tPerson := Person{}
	tPerson.Info.Add(2)
	tPerson.Info.Add(3)

	var myInterface interface{}
	myInterface = &tPerson

	iInfo := reflect.ValueOf(myInterface).Elem().FieldByName("Info").Interface()
	intValues, _ := CallMethodAndReturnSliceOfInt(iInfo, "GetValues")

	if len(intValues) == 2 {
		t.Log("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") PASSED, length of the returned value is 2")

		if intValues[0] == 2 {
			t.Log("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") PASSED, value of fisrt element in slice is 2")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") FAILED, value of fisrt element in slice should be 2, but is %d", intValues[0])
		}

		if intValues[1] == 3 {
			t.Log("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") PASSED, value of second element in slice is 3")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") FAILED, value of second element in slice should be 3, but is %d", intValues[1])
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfInt(iInfo, \"GetValues\") FAILED, length of the returned value expected to be 2, but is %d", len(intValues))
	}
}
