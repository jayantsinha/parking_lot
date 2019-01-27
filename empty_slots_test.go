package main

import (
	"testing"
)

func TestNewEmptySlots(t *testing.T) {
	emptySlots := *NewEmptySlots()
	if emptySlots.set == nil {
		t.Errorf("Expected EmptySlots to be initialized but failed!")
	}
}

func TestEmptySlots_Add(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{2, true},
		{2, true},
	}
	for _, tt := range tests {
		t.Run("EmptySlots.Add", func(t *testing.T) {
			set := *NewEmptySlots()
			if got := set.Add(tt.n); got != tt.want {
				t.Errorf("EmptySlots.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptySlots_Contains(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{2, true},
		{3, false},
	}
	for idx, tt := range tests {
		t.Run("EmptySlots.Contains", func(t *testing.T) {
			set := *NewEmptySlots()
			if idx < 2 {
				set.Add(tt.n)
			}
			if got := set.Contains(tt.n); got != tt.want {
				t.Errorf("EmptySlots.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptySlots_Remove(t *testing.T) {
	tests := []struct {
		n int
	}{
		{1},
		{2},
		{3},
	}
	for _, tt := range tests {
		t.Run("EmptySlots.Remove", func(t *testing.T) {
			set := *NewEmptySlots()
			set.set[tt.n] = true
			set.Remove(tt.n)
			if _, got := set.set[tt.n]; got == true {
				t.Errorf("EmptySlots.Remove() failed to remove element")
			}
		})
	}
}

func TestEmptySlots_Size(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{3, 3},
	}
	set := *NewEmptySlots()
	for _, tt := range tests {
		t.Run("EmptySlots.Size()", func(t *testing.T) {
				set.set[tt.n] = true
			if got := set.Size(); got != tt.want {
				t.Errorf("EmptySlots.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
