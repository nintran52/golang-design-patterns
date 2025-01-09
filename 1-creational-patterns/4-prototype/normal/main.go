package main

import "fmt"

// File struct
// Đối tượng File đại diện cho tệp tin trong hệ thống.
type File struct {
	name string
}

// print hiển thị tên File với thụt lề.
func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

// Folder struct
// Đối tượng Folder đại diện cho thư mục chứa các đối tượng File hoặc Folder khác.
type Folder struct {
	children []interface{} // Không có giao diện chung, sử dụng interface{}
	name     string
}

// print hiển thị tên Folder và các con của nó với thụt lề.
func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		switch v := child.(type) {
		case *File:
			v.print(indentation + indentation)
		case *Folder:
			v.print(indentation + indentation)
		}
	}
}

// main function
// Mã client quản lý logic sao chép thủ công mà không dùng Prototype Pattern.
func main() {
	// Tạo các đối tượng File
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	// Tạo thư mục Folder1 chứa File1
	folder1 := &Folder{
		children: []interface{}{file1},
		name:     "Folder1",
	}

	// Tạo thư mục Folder2 chứa Folder1, File2, và File3
	folder2 := &Folder{
		children: []interface{}{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	// Sao chép thủ công Folder2
	cloneFolder := &Folder{name: folder2.name + "_clone"}
	var clonedChildren []interface{}
	for _, child := range folder2.children {
		switch v := child.(type) {
		case *File:
			clonedChildren = append(clonedChildren, &File{name: v.name + "_clone"})
		case *Folder:
			cloneSubFolder := &Folder{name: v.name + "_clone"}
			var subFolderChildren []interface{}
			for _, subChild := range v.children {
				switch sv := subChild.(type) {
				case *File:
					subFolderChildren = append(subFolderChildren, &File{name: sv.name + "_clone"})
				}
			}
			cloneSubFolder.children = subFolderChildren
			clonedChildren = append(clonedChildren, cloneSubFolder)
		}
	}
	cloneFolder.children = clonedChildren

	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

// === Nhược Điểm ===
// 1. Không tuân theo nguyên tắc Mở/Đóng (Open/Closed Principle):
//    - Khi thêm loại đối tượng mới (vd: Symbol), cần cập nhật logic trong tất cả switch-case.
//    - Điều này dẫn đến mã khó bảo trì và dễ lỗi.

// 2. Lặp lại mã (Code Duplication):
//    - Logic sao chép phải được viết lại cho từng loại đối tượng như File và Folder.

// 3. Phụ thuộc mạnh vào switch-case:
//    - Với mỗi loại đối tượng mới, cần sửa đổi các switch-case hiện tại, làm mã phức tạp hơn.

// 4. Khả năng mở rộng hạn chế:
//    - Không thể tái sử dụng logic sao chép một cách dễ dàng khi có nhiều loại đối tượng khác nhau.

// === Khó Khăn ===
// 1. Khó bảo trì: Hệ thống lớn lên, việc chỉnh sửa logic sao chép tại nhiều nơi dễ gây lỗi.
// 2. Tăng độ phức tạp: Mã xử lý switch-case làm mã khó đọc và quản lý.
// 3. Không đảm bảo tính nhất quán: Không có giao diện chung để đảm bảo chuẩn hóa giữa các đối tượng.
