package main

type Int int

func (t Int) Add(v Int) Int {
	return Int(int(t) + int(v))
}

func AGlobalFunction(a, b int) int {
	return a + b
}

func FunctionWithMultipleReturnTypes() (int, string) {
	return 3, "hello"
}

func FunctionWithNamedReturnValues() (a int, b string) {
	a = 4
	b = "bye"
	return
}

func TestFunctionWithNamedReturnValues(t *T) {
	intValue, stringValue := FunctionWithNamedReturnValues()

	t.AssertEqualInt(4, intValue)
	t.AssertTrue("bye" == stringValue)
}

func TestFunctionWithMultipleReturnTypes(t *T) {
	intValue, stringValue := FunctionWithMultipleReturnTypes()

	t.AssertEqualInt(3, intValue)
	t.AssertTrue("hello" == stringValue)
}

func TestCallingAGlobalFunction(t *T) {
	result := AGlobalFunction(3, 4)
	t.AssertEqualInt(7, int(result))
}

func TestEveryTypeCanHaveMethods(t *T) {
	v := Int(4)
	result := v.Add(3)
	t.AssertEqualInt(7, int(result))
}

type Interface interface {
	ReturnValue() int
}

type Implementation struct {
}

// The declaration of this function makes
// only a pointer to Implementation satisfy Interface.
func (t *Implementation) ReturnValue() int {
	return 2
}

func TestInterfaces(t *T) {
	var i Interface
	v := Implementation{}
	i = &v //Only a pointer to Implementation satisfies Interface.
	value := i.ReturnValue()

	t.AssertEquals(2, value)
}
