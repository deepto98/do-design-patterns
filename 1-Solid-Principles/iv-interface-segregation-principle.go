package main

/*
 interface segregation principle (ISP) states that no code
 should be forced to depend on methods it does not use.
*/
type Document struct{}

//1. Implementation which breaks ISP
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultifunctionPrinter struct{}

func (m *MultifunctionPrinter) Print(d Document) {}
func (m *MultifunctionPrinter) Scan(d Document)  {}
func (m *MultifunctionPrinter) Fax(d Document)   {}

//For an OlderPrinter, we have to prevent Scan and Fax
type OlderPrinter struct{}

func (o *OlderPrinter) Print(d Document) {}

//Deprecated:
func (o *OlderPrinter) Scan(d Document) {
	panic("")
}

//Deprecated:
func (o *OlderPrinter) Fax(d Document) {
	panic("")
}

func main() {
	op := OlderPrinter{}

	//The ability to still Fax with an OlderPrinter breaks
	//Interface Segregation Principle
	op.Fax(Document{})
}

//2.a. ISP compliant architcture - breakup into smaller interfaces

type Scanner interface {
	Scan(d Document)
}
type Faxer interface {
	Fax(d Document)
}
type Printer interface {
	Print(d Document)
}

type OnlyPrinter struct{}

//only implements Printer interface
func (p *OnlyPrinter) Print(d Document) {}

type Photocopier struct{}

//only both Printer and Scanner interface

func (p *Photocopier) Print(d Document) {}
func (p *Photocopier) Scan(d Document)  {}

//2.b. We can also combine multiple interfaces into another interface
type PhotocopierInterface interface {
	Printer
	Scanner
}

//Decorator Design Pattern
type PhotocopierMachine struct {
	printer Printer
	scanner Scanner
}

func (p *PhotocopierMachine) Print(d Document) {
	p.printer.Print(d)
}

func (p *PhotocopierMachine) Scan(d Document) {
	p.scanner.Scan(d)
}
