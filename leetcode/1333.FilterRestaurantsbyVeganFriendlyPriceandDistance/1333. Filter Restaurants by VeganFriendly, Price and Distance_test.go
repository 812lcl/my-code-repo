package leetcode

import (
	"reflect"
	"testing"
)

func Test_filterRestaurants(t *testing.T) {
	restaurants := [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	veganFriendly := 1
	maxPrice := 50
	maxDistance := 10
	want := []int{3, 1, 5}
	if got := filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance); !reflect.DeepEqual(got, want) {
		t.Errorf("filterRestaurants() = %v, want %v", got, want)
	}

	restaurants = [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	veganFriendly = 0
	maxPrice = 50
	maxDistance = 10
	want = []int{4, 3, 2, 1, 5}
	if got := filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance); !reflect.DeepEqual(got, want) {
		t.Errorf("filterRestaurants() = %v, want %v", got, want)
	}

	restaurants = [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	veganFriendly = 0
	maxPrice = 30
	maxDistance = 3
	want = []int{4, 5}
	if got := filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance); !reflect.DeepEqual(got, want) {
		t.Errorf("filterRestaurants() = %v, want %v", got, want)
	}
}
