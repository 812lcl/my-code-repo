package leetcode

import "testing"

func TestConstructor_InsertFront(t *testing.T) {
	obj := Constructor(3)
	if !obj.InsertFront(9) {
		t.Errorf("InsertFront() = false, want true")
	}
	if obj.GetRear() != 9 {
		t.Errorf("GetRear() = %d, want 9", obj.GetRear())
	}
	if !obj.InsertFront(9) {
		t.Errorf("InsertFront() = false, want true")
	}
	if obj.GetRear() != 9 {
		t.Errorf("GetRear() = %d, want 9", obj.GetRear())
	}
	if !obj.InsertLast(5) {
		t.Errorf("InsertLast() = false, want true")
	}
	if obj.GetFront() != 9 {
		t.Errorf("GetFront() = %d, want 9", obj.GetFront())
	}
	if obj.GetRear() != 5 {
		t.Errorf("GetRear() = %d, want 5", obj.GetRear())
	}
	if obj.GetFront() != 9 {
		t.Errorf("GetFront() = %d, want 9", obj.GetFront())
	}
	if obj.InsertLast(8) {
		t.Errorf("InsertLast() = true, want false")
	}
	if !obj.DeleteLast() {
		t.Errorf("InsertLast() = false, want true")
	}
	if obj.GetFront() != 9 {
		t.Errorf("GetFront() = %d, want 9", obj.GetFront())
	}
}
