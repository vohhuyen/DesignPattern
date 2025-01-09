// collection.go: Collection
package main

import "fmt"

// iterator.go: Iterator
// Bước 1: Khai báo interface iterator với các phương thức để duyệt qua.
// Phải có các phương thức kiểm tra phần tử tiếp theo và lấy phần tử đó.
type Iterator interface {
	hasNext() bool
	getNext() *User
}

// Bước 2: Khai báo interface collection và mô tả phương thức tạo iterator.
// Kiểu trả về của phương thức phải là interface iterator.
type Collection interface {
	createIterator() Iterator
}

// userIterator.go: Concrete iterator
// Bước 3: Triển khai lớp iterator cụ thể cho các collection mà bạn muốn duyệt qua.
// Một đối tượng iterator phải liên kết với một instance cụ thể của collection.
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// userCollection.go: Concrete collection
// Bước 4: Triển khai interface collection trong các lớp collection cụ thể.
// Đối tượng collection phải chuyển chính nó vào constructor của iterator để tạo liên kết giữa chúng.
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// user.go: User struct (Client code sử dụng struct này)
// Một struct User đơn giản đại diện cho các phần tử trong collection.
type User struct {
	name string
	age  int
}

// main.go: Client code
func main() {
	// Bước 5: Thay thế toàn bộ code duyệt qua collection bằng iterator.
	// Client tạo một iterator mới mỗi khi cần duyệt qua các phần tử trong collection.

	// Tạo một số đối tượng User
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	// Tạo một collection User và thêm các đối tượng User
	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	// Tạo iterator cho collection User
	iterator := userCollection.createIterator()

	// Sử dụng iterator để duyệt qua các phần tử trong collection
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
