package main

import "fmt"

type Product struct {
	name  string
	color Color
	size  Size
}
type Color int

// this format assigns 0,1,2
const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

//OCP - open for extension, closed for modification

/* 1.
the first implementation (separate filter methods on Filter Struct) breaks OCP

Because, if later we want FilterBySize we add:

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
result := make([]*Product, 0)

for i, product := range products {
	if product.size == size {
			result = append(result, &products[i])
	}
}
return result
}

THIS VIOLATES THE OPEN CLOSED PRINCIPLE, BECAUSE WE'VE MODIFIED THE FILTER TYPE
AND ADDED A NEW METHOD
*/

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, product := range products {
		if product.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// func main() {
// 	apple := Product{"apple", red, small}
// 	car := Product{"car", green, large}

// 	products := []Product{apple, car}
// 	fmt.Printf("Only red products\n")

// 	f := Filter{}
// 	redProducts := f.FilterByColor(products, red)

// 	for _, v := range redProducts {
// 		fmt.Printf("%v\n", v.name)
// 	}

// }

//2. THE FOLLOWING IMPLEMENTATION SATISFIES OCP
//Using Specification Pattern - one basic interface with a validator
//The interface is open for extension, i.e. more Specifications can be added,
// but not for modification i.e we won't change BetterFilter
type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

//Color Specification implements Specification with the method isSatisfied
func (c ColorSpecification) isSatisfied(p *Product) bool {
	if p.color == c.color {
		return true
	}
	return false
}

//Size Specification
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	if s.size == p.size {
		return true
	}
	return false
}

//Spec with combinations
type andSpecification struct {
	color, size Specification
}

func (s andSpecification) isSatisfied(p *Product) bool {
	if s.size.isSatisfied(p) && s.color.isSatisfied(p) {
		return true
	}
	return false
}

//Struct to filter (not necessary, but good to have to make a method)
type BetterFilter struct {
}

//Universal Filter function which accepts Specification interface
func (filter *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	filteredProducts := make([]*Product, 0)

	for i, product := range products {
		if spec.isSatisfied(&product) {
			filteredProducts = append(filteredProducts, &products[i])
		}
	}
	return filteredProducts
}

func main() {
	//Initializing products
	apple := Product{"apple", red, small}
	car := Product{"car", green, large}
	smallCar := Product{"small Car", green, small}
	products := []Product{apple, car, smallCar}

	f := BetterFilter{}

	//Filtering red products
	redSpec := ColorSpecification{red}
	redProducts := f.Filter(products, redSpec)
	fmt.Printf("Only red products\n")
	for _, v := range redProducts {
		fmt.Printf("%v\n", v.name)
	}

	//Composite filter spec for Green+Small Products
	smallGreenSpec := andSpecification{
		ColorSpecification{green}, SizeSpecification{small},
	}
	smallGreenProducts := f.Filter(products, smallGreenSpec)

	fmt.Printf("Only Small and Green Products\n")
	for _, v := range smallGreenProducts {
		fmt.Printf("%v\n", v.name)
	}
}
