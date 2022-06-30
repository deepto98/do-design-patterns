package main

import (
	"fmt"
	"strings"
)

/*  BUILDER:
Simpler objects can be created in a single constructor call
For bigger objects, step-by-step construction is suited
Builder pattern provides an API for constructing an obj
step-by-step
*/
type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func main() {

	//We want to wrap a string around html tags, without having to
	// manually append strings (whichc is inefficient, coz a brand new string needs to be allocated
	//in memory, since strings are immutable in go)
	text := "hello"
	//Builder is an api which can be used to efficiently and easily build or concatenate strings
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(text)
	sb.WriteString("</p>")

	//to convert back from builder obj to string
	str := sb.String()
	fmt.Println(str)
}
