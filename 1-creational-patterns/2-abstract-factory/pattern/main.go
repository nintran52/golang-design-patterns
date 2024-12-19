package main

import "fmt"

// Step 1: Map out a matrix of distinct product types versus variants of these products
// Sản phẩm: Shoe và Shirt
// Biến thể: Adidas và Nike

// Step 2: Declare abstract product interfaces
// Khai báo các abstract product interfaces cho từng loại sản phẩm

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Step 3: Declare the abstract factory interface
// Khai báo abstract factory interface để tạo tất cả các loại sản phẩm

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// Step 4: Implement a set of concrete factory classes
// Tạo các concrete factory class, mỗi class tương ứng với một biến thể sản phẩm

type Adidas struct{}

type Nike struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

// Step 2 Continued: Implement abstract products
// Khai báo và triển khai các abstract products

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// AdidasShoe implements IShoe
type AdidasShoe struct {
	Shoe
}

// NikeShoe implements IShoe
type NikeShoe struct {
	Shoe
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

// AdidasShirt implements IShirt
type AdidasShirt struct {
	Shirt
}

// NikeShirt implements IShirt
type NikeShirt struct {
	Shirt
}

// Step 5: Create factory initialization code
// Tạo mã khởi tạo factory để khởi tạo concrete factory class dựa trên cấu hình hoặc môi trường

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	} else if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Invalid brand type")
}

// Step 6: Replace direct product constructors with factory methods
// Tìm tất cả các lời gọi trực tiếp đến product constructors và thay thế chúng bằng các factory methods

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Shoe Logo: %s\n", s.getLogo())
	fmt.Printf("Shoe Size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Shirt Logo: %s\n", s.getLogo())
	fmt.Printf("Shirt Size: %d\n", s.getSize())
}
