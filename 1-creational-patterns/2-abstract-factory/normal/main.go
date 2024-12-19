package main

import "fmt"

// Không sử dụng pattern: Code sẽ trực tiếp tạo các sản phẩm mà không có các lớp abstract hoặc factory.

// Sản phẩm cụ thể: AdidasShoe, NikeShoe, AdidasShirt, NikeShirt

type Shoe struct {
	logo string
	size int
}

type Shirt struct {
	logo string
	size int
}

// Tạo trực tiếp các sản phẩm Adidas
func newAdidasShoe() *Shoe {
	return &Shoe{
		logo: "adidas",
		size: 14,
	}
}

func newAdidasShirt() *Shirt {
	return &Shirt{
		logo: "adidas",
		size: 14,
	}
}

// Tạo trực tiếp các sản phẩm Nike
func newNikeShoe() *Shoe {
	return &Shoe{
		logo: "nike",
		size: 14,
	}
}

func newNikeShirt() *Shirt {
	return &Shirt{
		logo: "nike",
		size: 14,
	}
}

// Mã client
func main() {
	// Tạo sản phẩm trực tiếp mà không thông qua factory
	adidasShoe := newAdidasShoe()
	adidasShirt := newAdidasShirt()

	nikeShoe := newNikeShoe()
	nikeShirt := newNikeShirt()

	// In thông tin sản phẩm
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s *Shoe) {
	fmt.Printf("Shoe Logo: %s\n", s.logo)
	fmt.Printf("Shoe Size: %d\n", s.size)
}

func printShirtDetails(s *Shirt) {
	fmt.Printf("Shirt Logo: %s\n", s.logo)
	fmt.Printf("Shirt Size: %d\n", s.size)
}

// Khó khăn khi không sử dụng pattern:
// 1. Mã kém linh hoạt: Thêm thương hiệu hoặc sản phẩm mới đòi hỏi phải viết thêm nhiều hàm khởi tạo.
// 2. Nhân bản mã: Logic khởi tạo bị lặp lại, khó bảo trì và dễ xảy ra lỗi.
// 3. Khó mở rộng: Nếu cần thêm loại sản phẩm mới (ví dụ: quần thể thao), mã sẽ trở nên phức tạp hơn.
// 4. Thiếu kiểm soát tập trung: Không có nơi tập trung để quản lý việc tạo sản phẩm, gây khó khăn khi cần sửa đổi hoặc mở rộng.
