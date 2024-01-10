package leetcode

import (
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	nums1 := []int{1, 2, 4, 4, 5}
	k1 := 5
	expected1 := 4
	if res := maxFrequency(nums1, k1); res != expected1 {
		t.Errorf("maxFrequency() = %d, want %d", res, expected1)
	}

	nums2 := []int{1, 4, 8, 13}
	k2 := 5
	expected2 := 2
	if res := maxFrequency(nums2, k2); res != expected2 {
		t.Errorf("maxFrequency() = %d, want %d", res, expected2)
	}

	nums3 := []int{3, 9, 10, 20}
	k3 := 0
	expected3 := 1
	if res := maxFrequency(nums3, k3); res != expected3 {
		t.Errorf("maxFrequency() = %d, want %d", res, expected3)
	}

	nums4 := []int{1, 1, 1, 1, 1}
	k4 := 10
	expected4 := 5
	if res := maxFrequency(nums4, k4); res != expected4 {
		t.Errorf("maxFrequency() = %d, want %d", res, expected4)
	}

	nums5 := []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966}
	k5 := 3056
	expected5 := 73
	if res := maxFrequency(nums5, k5); res != expected5 {
		t.Errorf("maxFrequency() = %d, want %d", res, expected5)
	}
}

func TestMaxFrequency1(t *testing.T) {
	nums1 := []int{1, 2, 4, 4, 5}
	k1 := 5
	expected1 := 4
	if res := maxFrequency1(nums1, k1); res != expected1 {
		t.Errorf("maxFrequency1() = %d, want %d", res, expected1)
	}

	nums2 := []int{1, 4, 8, 13}
	k2 := 5
	expected2 := 2
	if res := maxFrequency1(nums2, k2); res != expected2 {
		t.Errorf("maxFrequency1() = %d, want %d", res, expected2)
	}

	nums3 := []int{3, 9, 10, 20}
	k3 := 0
	expected3 := 1
	if res := maxFrequency1(nums3, k3); res != expected3 {
		t.Errorf("maxFrequency1() = %d, want %d", res, expected3)
	}

	nums4 := []int{1, 1, 1, 1, 1}
	k4 := 10
	expected4 := 5
	if res := maxFrequency1(nums4, k4); res != expected4 {
		t.Errorf("maxFrequency1() = %d, want %d", res, expected4)
	}

	nums5 := []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966}
	k5 := 3056
	expected5 := 73
	if res := maxFrequency1(nums5, k5); res != expected5 {
		t.Errorf("maxFrequency1() = %d, want %d", res, expected5)
	}
}
