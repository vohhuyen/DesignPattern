// originator.go: Originator
package main

import "fmt"

// Step 1: Xác định lớp Originator
// Đây là lớp mà bạn muốn lưu trạng thái.
// Originator có các phương thức để tạo Memento và khôi phục trạng thái từ Memento.
type Originator struct {
	state string
}

// Tạo Memento từ trạng thái hiện tại
func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

// Khôi phục trạng thái từ Memento
func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

// Đặt trạng thái mới
func (e *Originator) setState(state string) {
	e.state = state
}

// Lấy trạng thái hiện tại
func (e *Originator) getState() string {
	return e.state
}

// memento.go: Memento
// Step 2: Tạo lớp Memento
// Lớp này sẽ lưu trạng thái của Originator.
// Memento thường được thiết kế immutable, không có phương thức setter.
type Memento struct {
	state string
}

// Lấy trạng thái đã lưu trong Memento
func (m *Memento) getSavedState() string {
	return m.state
}

// caretaker.go: Caretaker
// Step 3: Tạo lớp Caretaker
// Caretaker quản lý các Memento và cung cấp khả năng lưu trữ và khôi phục lịch sử.
type Caretaker struct {
	mementoArray []*Memento // Danh sách các Memento
}

// Thêm Memento vào danh sách
func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

// Lấy Memento từ danh sách theo index
func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

// main.go: Client code
func main() {
	// Step 4: Triển khai Caretaker và Originator
	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0), // Khởi tạo danh sách Memento
	}

	originator := &Originator{
		state: "A", // Khởi tạo trạng thái ban đầu
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // Lưu trạng thái "A"

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // Lưu trạng thái "B"

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento()) // Lưu trạng thái "C"

	// Khôi phục trạng thái
	originator.restoreMemento(caretaker.getMemento(1)) // Khôi phục về trạng thái "B"
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0)) // Khôi phục về trạng thái "A"
	fmt.Printf("Restored to State: %s\n", originator.getState())
}