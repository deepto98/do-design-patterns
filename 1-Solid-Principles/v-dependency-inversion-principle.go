package main

import "fmt"

//Not the same as Dependency Injection, although that is derived from this
//DIP : Higher Level Modules shouldn't depend on
// Lower Lever Modules
//Both should depend on abstractions

type Relationship int

const (
	Parent  Relationship = iota //0
	Child                       //1
	Sibling                     //2
)

type Person struct {
	name string
}
type Info struct {
	from         *Person
	to           *Person
	relationship Relationship
}

//Low Level Module - storing all relationships
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentChild(parent, child *Person) {
	parentChildRel := Info{parent, child, Parent}
	r.relations = append(r.relations, parentChildRel)
}

//High Level Module
type Research struct {
	//breaking DIP coz Research depends on Relationships
	relationships Relationships
}

//Investigate to find children of a Name
func (r *Research) Investigate(name string) {
	relations := r.relationships.relations

	for _, rel := range relations {
		if rel.from.name == name && rel.relationship == Parent {
			fmt.Printf("%s's child is %v\n", name, rel.to.name)
		}
	}
}

//2. Compying with DIP
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

//Replacing High Level Module Research, it now has an abstraction
// called RelationshipBrowser
type ResearchNew struct {
	browser RelationshipBrowser
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *ResearchNew) InvestigateNew(name string) {
	children := r.browser.FindAllChildrenOf(name)
	for _, v := range children {
		fmt.Printf("%s's child is %s\n", name, v.name)
	}
}

func main() {
	parent := Person{"Jim"}
	child1 := Person{"A"}
	child2 := Person{"B"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentChild(&parent, &child1)
	relationships.AddParentChild(&parent, &child2)
	// high-level module
	researchNew := ResearchNew{&relationships}
	researchNew.InvestigateNew("Jim")
}
