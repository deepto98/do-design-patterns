package main

import (
	"fmt"
	"strings"
)

// https://dave.cheney.net/2016/08/20/solid-go-design
/*
	Single Responsibility Pattern advocates the separation of concern,
	splitting different functionalities into different constructs

	With this Journal struct, we only deal with reading, adding, removing entries
	Other more generic functionalities such as persisting to a file, loading from a file
	should be implemented as functions in a separate package, and not as methods on Journal

	Antipattern:God Object (putting everything into one construct)

*/
var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}
func (j *Journal) addEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%s", text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) removeEntry(index int) bool {
	/*
	   ... is syntax for variadic arguments in Go.

	   Basically, when defining a function it puts all the arguments that you pass into one slice of that type. By doing that, you can pass as many arguments as you want (for example, fmt.Println can take as many arguments as you want).

	   Now, when calling a function, ... does the opposite: it unpacks a slice and passes them as separate arguments to a variadic function.

	   So what this line does:

	   a = append(a[:0], a[1:]...)
	   is essentially:

	   a = append(a[:0], a[1], a[2])
	*/
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
	entryCount--

	return true
}

func main() {
	journal := Journal{}
	journal.addEntry("Woke up")
	journal.addEntry("Brushed my teeth")
	journal.addEntry("Went for a walk")
	journal.removeEntry(1)
	fmt.Printf("%s\n", journal.String())
}
