package invoker

import "testing"

func TestCatchErrorOnCallMethodAndReturnInt(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnInt(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnInt(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnInt(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodAndReturnFloat64(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnFloat64(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnFloat64(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnFloat64(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodAndReturnString(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnString(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnString(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnString(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodAndReturnBool(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnBool(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnBool(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnBool(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodAndReturnSliceOfInt(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnSliceOfInt(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnSliceOfInt(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnSliceOfInt(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodAndReturnSliceOfFloat64(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnSliceOfFloat64(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnSliceOfFloat64(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnSliceOfFloat64(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodAndReturnSliceOfString(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnSliceOfString(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnSliceOfString(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnSliceOfString(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodAndReturnSliceOfBool(t *testing.T) {
	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3
	_, err := CallMethodAndReturnSliceOfBool(testMyMath, "InvalidMethod")

	if err != nil {
		t.Log("Error caught on CallMethodAndReturnSliceOfBool(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodAndReturnSliceOfBool(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnInt(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnInt(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64(testMyMath, \"InvalidMethod\") FAILED")
	}
}
func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnString(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnString(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnBool(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnBool(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnOfInt(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfInterfaceAndReturnOfInt(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfFloat64(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfFloat64(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfString(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfString(testMyMath, \"InvalidMethod\") FAILED")
	}
}

func TestCatchErrorOnCallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(t *testing.T) {
	var args []interface{}

	testMyMath := MyMath{}
	testMyMath.Operand1 = 2
	testMyMath.Operand2 = 3

	_, err := CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool(testMyMath, "InvalidMethod", args)

	if err != nil {
		t.Log("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfBool(testMyMath, \"InvalidMethod\") PASSED")
	} else {
		t.Error("Error caught on CallMethodWithArgsAsSliceOfFloaterfaceAndReturnOfBool(testMyMath, \"InvalidMethod\") FAILED")
	}
}
