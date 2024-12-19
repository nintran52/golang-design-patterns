// Factory Method Pattern in Go
package main

import "fmt"

// Bước 1: Tạo một giao diện chung cho tất cả các sản phẩm
// Định nghĩa một giao diện chung cho tất cả các sản phẩm
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Cấu trúc Gun triển khai giao diện IGun
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Bước 2: Thêm một phương thức nhà máy (factory method) trống vào lớp tạo
// Lớp cơ sở chứa phương thức nhà máy
type GunFactory interface {
	CreateGun(gunType string) (IGun, error)
}

// Cấu trúc BaseGunFactory với một triển khai mặc định của phương thức nhà máy
type BaseGunFactory struct{}

func (f *BaseGunFactory) CreateGun(gunType string) (IGun, error) {
	// Triển khai mặc định để các lớp con ghi đè
	return nil, fmt.Errorf("Factory method not implemented")
}

// Bước 3: Thay thế các constructor bằng các lời gọi tới factory method
// Các sản phẩm cụ thể (Ak47 và Musket)
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// Bước 4: Tạo một tập hợp các lớp tạo (creator subclass)
// Cấu trúc GunFactory đóng vai trò như một nhà máy để tạo các loại súng cụ thể
type ConcreteGunFactory struct {
	BaseGunFactory
}

func (f *ConcreteGunFactory) CreateGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Unknown gun type: %s", gunType)
	}
}

// Bước 5: Tái sử dụng tham số điều khiển nếu có quá nhiều loại sản phẩm
// Đã được triển khai trong lệnh switch trong phương thức nhà máy

// Bước 6: Biến phương thức nhà máy cơ sở thành trừu tượng hoặc mặc định
// Lớp BaseGunFactory cung cấp hành vi mặc định, và các lớp con có thể ghi đè.

// Mã client
func main() {
	// Tạo một instance của nhà máy
	factory := &ConcreteGunFactory{}

	// Sử dụng nhà máy để tạo sản phẩm
	ak47, err := factory.CreateGun("ak47")
	if err != nil {
		fmt.Println(err)
		return
	}

	musket, err := factory.CreateGun("musket")
	if err != nil {
		fmt.Println(err)
		return
	}

	// In chi tiết của các sản phẩm
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
