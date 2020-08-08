package main

import (
	"fmt"
	"io"
	"os"
)

type MyInt int

var i int
var j MyInt

type I interface {
	M1(string) string
	M2()
	M3(int)
	M4() string
}
type K interface {
	K1()
}
type T1 struct {
	name string
	age  int
}
type T2 struct {
	name string
	age  int
}

func (t T1) M1(s string) string {
	t.name = s
	return fmt.Sprintf("Name is set to %s", t.name)
}
func (t T1) M2() {

}
func (t T1) M3(i int) {
	t.age = i
}
func (t T1) M4() string {
	return fmt.Sprintf("The age of person is %d", t.age)
}

func (k T2) M1(s string) string {
	k.name = s
	return fmt.Sprintf("Name is set to %s", k.name)
}
func (k T2) M2() {

}
func (k T2) M3(i int) {
	k.age = i
}
func (k T2) M4() string {
	return fmt.Sprintf("The age of person is %d", k.age)
}
func (k T2) K1() {
}
func main() {
	//Question: If the dynamic type works for interfaces, then why not for the Values
	// i = j => Error
	//Type T1 has implemented the interface I
	var t *T1
	// Type T2 implements the interface I and interface K rest the same
	var k T2
	var ii I
	ii = t // dynamic type assignment
	/*Even if t is nil ii is not nil as it has 2 things type and value ; type is main.*T and so
	ii is not nil. Below statement prints that ii nil is false
	*/
	fmt.Printf(" Dynamic type of I is %T and the ii nil is %v even if t is %v\n", ii, ii == nil, t == nil)
	ii = k
	fmt.Printf(" Dynamic type of I is %T and the value is %v \n", ii, ii)
	/*Answer: t = k will not work as they are concrete implementations of  I and K
	However, ii = t or ii =k will work as ii is not a concrete implementation but  it is an interface
	For the i =j statement both i and j are concerete implementations
	*/

	/*Example 2
	A variable of interface type stores a pair: the concrete value assigned to the variable,
	and that value's type descriptor.
	To be more precise, the value is the underlying concrete data item that implements the interface and
	the type describes the full type of that item. For instance, after
	r contains, schematically, the (value, type) pair, (tty, *os.File). Notice that the type *os.File implements methods other than Read;
	even though the interface value provides access only to the Read method, the value inside carries all the type information about that value
	*/
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		//handle the error
	}
	r = tty
	fmt.Println(r)

	var w io.Writer
	w = r.(io.Writer)
	/*
		The expression in this assignment is a type assertion; what it asserts is that
		the item inside r also implements io.Writer, and so we can assign it to w.
		After the assignment, w will contain the pair (tty, *os.File).
		That's the same pair as was held in r.
		The static type of the interface determines what methods may be invoked with an interface variable,
		even though the concrete value inside may have a larger set of methods.
	*/
	var empty interface{}
	empty = w

}
