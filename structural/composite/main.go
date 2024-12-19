package main

import "fmt"

// Bước 2: Định nghĩa giao diện Component
// Đây là giao diện chung cho cả các phần tử đơn giản (File) và phần tử phức tạp (Folder).
// Interface này khai báo phương thức `search`.
type Component interface {
	search(keyword string)
}

// Bước 3: Tạo lớp Leaf để đại diện cho các phần tử đơn giản
// `File` là một phần tử đơn giản và triển khai giao diện `Component`.
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Đang tìm từ khóa '%s' trong file '%s'\n", keyword, f.name)
}

// Bước 4: Tạo lớp Container để đại diện cho các phần tử phức tạp
// `Folder` là một phần tử phức tạp và có thể chứa các File hoặc Folder khác.
type Folder struct {
	name       string
	components []Component // Danh sách các phần tử con, có thể là File hoặc Folder
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Đang tìm từ khóa '%s' trong folder '%s'\n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword) // Ủy thác việc tìm kiếm cho các phần tử con
	}
}

// Bước 5: Định nghĩa các phương thức để thêm phần tử con vào container
// Các phương thức này cho phép thêm động các phần tử con vào Folder.
func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

func main() {
	// Bước 1: Đảm bảo mô hình có thể biểu diễn dưới dạng cấu trúc cây
	// Tạo các phần tử đơn giản (File) đại diện cho các node lá
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	// Tạo các phần tử phức tạp (Folder) đại diện cho các containers
	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1) // Thêm một File vào Folder1

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)   // Thêm một File vào Folder2
	folder2.add(file3)   // Thêm một File khác vào Folder2
	folder2.add(folder1) // Thêm Folder1 làm phần tử con của Folder2

	// Thực hiện tìm kiếm từ khóa trong toàn bộ cấu trúc cây
	folder2.search("hoa hồng")
}