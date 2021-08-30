package invoker

import "testing"

type MySetInterface interface {
	Add(value int)
	GetValues() []int
	Contains(value int) bool
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

func (mySet *MySet) Contains(value int) bool {
	if mySet.Values != nil {
		for currentValue := range mySet.Values {
			if currentValue == value {
				return true
			}
		}
	}

	return false
}

func TestMySetGetValues(t *testing.T) {
	var mySet MySet
	mySet.Add(2)
	mySet.Add(3)

	intValues, _ := CallMethodAndReturnSliceOfInt(mySet, "GetValues")
	if len(intValues) == 2 {
		t.Log("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") PASSED, length of the returned value is 2")

		if intValues[0] == 2 || intValues[0] == 3 {
			t.Log("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") PASSED, value of fisrt element in slice is 2 or 3")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") FAILED, value of fisrt element in slice should be 2 or 3, but is %d", intValues[0])
		}

		if intValues[1] == 2 || intValues[1] == 3 {
			t.Log("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") PASSED, value of second element in slice is 2 or 3")
		} else {
			t.Errorf("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") FAILED, value of second element in slice should be 2 or 3, but is %d", intValues[1])
		}

	} else {
		t.Errorf("CallMethodAndReturnSliceOfInt(mySet, \"GetValues\") FAILED, length of the returned value expected to be 2, but is %d", len(intValues))
	}
}

func TestMySetContains(t *testing.T) {
	var mySet MySet
	var boolValue bool
	mySet.Add(2)
	mySet.Add(3)

	boolValue, _ = CallMethodAndReturnBool(mySet, "Contains", 2)
	if boolValue {
		t.Log("CallMethodAndReturnBool(mySet, \"Contains\", 2) PASSED, returned value is TRUE")
	} else {
		t.Error("CallMethodAndReturnBool(mySet, \"Contains\", 2) FAILED, returned value expected to be TRUE, but is FALSE")
	}

	boolValue, _ = CallMethodAndReturnBool(mySet, "Contains", 4)
	if !boolValue {
		t.Log("CallMethodAndReturnBool(mySet, \"Contains\", 4) PASSED, returned value is FALSE")
	} else {
		t.Error("CallMethodAndReturnBool(mySet, \"Contains\", 4) FAILED, returned value expected to be TRUE, but is TRUE")
	}
}
