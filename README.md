# Welcome to Invoker

## What's Invoker?

_Invoker_ is a _Golang_ package which calls methods from _Compositions_ or _Interfaces_ passing to it the object (or interface), the name of the method as a string and the parameters dynamically. Also, you can handle the return of the method regardless of the type.

## First example

Let's take a look at the most general function: _CallMethod_. It return an slice of _reflect.Value_ so you need to convert to the right type, as you can see in the example below.

```go
package main

import (
  "fmt"
  "github.com/adilsonchacon/invoker"
)

type ThisMath struct {
  Operand1 int
  Operand2 int
}

func (thisMath *ThisMath) Sum() int {
  return thisMath.Operand1 + thisMath.Operand2
}

func (thisMath *ThisMath) Average() float64 {
  return (float64(thisMath.Operand1) + float64(thisMath.Operand2)) / 2.0
}

func main() {
  thisMath := ThisMath{}
  thisMath.Operand1 = 2
  thisMath.Operand1 = 3

  sumValue, _ := invoker.CallMethod(thisMath, "Sum")
  intValue := sumValue[0].Interface().(int)
  fmt.Printf("Sum is %d\n", intValue)

  avgValue, _ := invoker.CallMethod(thisMath, "Average")
  floatValue := avgValue[0].Interface().(float64)
  fmt.Printf("Average is %f\n", floatValue)
}
```

Output:
```
Sum is 3
Average is 1.500000
```

If you realized, the lines `intValue := sumValue[0].Interface().(int)` and `floatValue := avgValue[0].Interface().(float64)` try to convert the values to _int_ and to _float64_, respectively.

If your method returns a _string_ you would need to use _string_ between parentheses:

```go
  concatValue, _ := invoker.CallMethod(thisMath, "Concat")
  strValue := strValue[0].Interface().(string)
```

Or, if your method returns a type that you created _MyType_ you would need to use _MyType_ between parentheses:

```go
  someMethodValue, _ := invoker.CallMethod(thisMath, "SomeMethod")
  myTypeValue := someMethodValue[0].Interface().(MyType)
```

and so on.

### You Don't Need To Convert Always

The _invoker_ package has functions that converts from _[]reflect.value_ to many primitive types as you can see in the list below:

|Function Name|Description|
|---|---|
|CallMethodAndReturnInt|Call method and converts the result to int|
|CallMethodAndReturnFloat64|Call method and converts the result to float64|
|CallMethodAndReturnString|Call method and converts the result to string|
|CallMethodAndReturnBool|Call method and converts the result to bool|
|CallMethodAndReturnSliceOfInt|Call method and converts the result to slice of int|
|CallMethodAndReturnSliceOfFloat64|Call method and converts the result to slice of float64|
|CallMethodAndReturnSliceOfString|Call method and converts the result to slice of string|
|CallMethodAndReturnSliceOfBool|Call method and converts the result to slice of bool|

__OBS:__ All those methods above work with the same parameters that `CallMethod`

Example:

```go
sumValue, _ := invoker.CallMethodAndReturnInt(thisMath, "Sum")
fmt.Printf("Sum is %d\n", sumValue)
```

## Passing Parameters

There are two ways to pass the parameters. Let's take a look at them:

```go
package main

import (
  "fmt"
  "github.com/adilsonchacon/invoker"
)

type MyAverage struct {
  Operand1 int
  Operand2 int
}

func (myAvg *MyAverage) WeightedAverage(weight1 float64, weight2 float64) float64 {
  return ((float64(myAvg.Operand1) * weight1) + (float64(myAvg.Operand2) * weight2)) / (weight1 + weight2)
}

func main() {
  myAvg := MyAverage{}
  myAvg.Operand1 = 2
  myAvg.Operand2 = 3

  avgValueSample1, _ := invoker.CallMethod(myAvg, "WeightedAverage", 4.0, 6.0)
  floatValueSample1 := avgValueSample1[0].Interface().(float64)
  fmt.Printf("Weighted Average (sample 1) is %f\n", floatValueSample1)

  var args []interface{}
  args = append(args, 5.0)
  args = append(args, 7.0)
  avgValueSample2, _ := invoker.CallMethodWithArgsAsSliceOfInterface(myAvg, "WeightedAverage", args)
  floatValueSample2 := avgValueSample2[0].Interface().(float64)
  fmt.Printf("Weighted Average (sample 2) is %f\n", floatValueSample2)
}
```

Output:
```
Weighted Average (sample 1) is 2.600000
Weighted Average (sample 2) is 2.583333
```

__OBS:__ In both samples we can substitute _CallMethod_ and _CallMethodWithArgsAsSliceOfInterface_ for _CallMethodAndReturnFloat64_ and _CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64_, respectively, and suppress the conversion parts.

### Again, You Don't Need To Convert Always

Same stuff: the _invoker_ package has functions that converts from _[]reflect.value_ to many primitive types as you can see in the list below:

|Function Name|Description|
|---|---|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnInt|Call method and converts the result to int|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnFloat64|Call method and converts the result to float64|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnString|Call method and converts the result to string|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnBool|Call method and converts the result to bool|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfInt|Call method and converts the result to slice of int|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfFloat64|Call method and converts the result to slice of float64|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfString|Call method and converts the result to slice of string|
|CallMethodWithArgsAsSliceOfInterfaceAndReturnSliceOfBool|Call method and converts the result to slice of bool|

__OBS:__ All those methods above work with the same parameters that _CallMethodWithArgsAsSliceOfInterface_

## Object As An Interface Value

You can pass the object as an interface that invoker will handle it for you.

```go
package main

import (
  "fmt"
  "github.com/adilsonchacon/invoker"
  "sort"
)

type Manager struct {
  Skills []string
}

func (manager *Manager) SortSkills() []string {
  sort.Strings(manager.Skills)
  return manager.Skills
}

type Developer struct {
  Skills []string
}

func (developer *Developer) SortSkills() []string {
  sort.Strings(developer.Skills)
  return developer.Skills
}

func SortSkills(iObject interface{}) []string {
  sliceOfStringValue, _ := invoker.CallMethodAndReturnSliceOfString(iObject, "SortSkills")
  return sliceOfStringValue
}

func main() {
  var anyObject interface{}

  manager := Manager{}
  manager.Skills = []string{"Scrum", "PMBOK", "ITIL", "COBIT"}

  developer := Developer{}
  developer.Skills = []string{"Go", "Postgresql", "Design Patterns"}

  anyObject = &manager
  fmt.Println(SortSkills(anyObject))

  anyObject = &developer
  fmt.Println(SortSkills(anyObject))
}
```

Output:
```
[COBIT ITIL PMBOK Scrum]
[Design Patterns Go Postgresql]
```

## Work With Interfaces

For _invoker_, _Interfaces_ work just like _Composition_

```go
package main

import (
  "fmt"
  "github.com/adilsonchacon/invoker"
)

type MySetInterface interface {
  Add(value int)
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

func main() {
  mySet := &MySet{}

  invoker.CallMethod(mySet, "Add", 2)
  invoker.CallMethod(mySet, "Add", 5)

  boolValueSample1, _ := invoker.CallMethodAndReturnBool(mySet, "Contains", 2)
  fmt.Printf("Does mySet contain 2? %t\n", boolValueSample1)

  boolValueSample2, _ := invoker.CallMethodAndReturnBool(mySet, "Contains", 8)
  fmt.Printf("Does mySet contain 8? %t\n", boolValueSample2)
}
```

Output:
```
Does mySet contain 2? true
Does mySet contain 8? false
```

## Error Handling

All functions of the _invoker_ package has error handling. The example below shows an error for trying to call method that does not exist:

```go
package main

import (
  "fmt"
  "github.com/adilsonchacon/invoker"
)

type ThisMath struct {
  Operand1 int
  Operand2 int
}

func (thisMath *ThisMath) Sum() int {
  return thisMath.Operand1 + thisMath.Operand2
}

func main() {
  thisMath := ThisMath{}
	thisMath.Operand1 = 2
	thisMath.Operand2 = 3
	_, err := invoker.CallMethod(thisMath, "InvalidMethod")

  if err != nil {
    fmt.Printf("Error: %s", err)
  }
}
```

Output
```
Error: method not found
```

## License

Invoker is released under the [MIT License](https://opensource.org/licenses/MIT).
