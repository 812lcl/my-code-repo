package leetcode

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var res [][]int
	for _, r := range restaurants {
		if r[2] < veganFriendly || r[3] > maxPrice || r[4] > maxDistance {
			continue
		}
		res = append(res, r)
	}
	sort.Sort(stores(res))
	var result []int
	for _, v := range res {
		result = append(result, v[0])
	}
	return result
}

type stores [][]int

// Len is the number of elements in the collection.
func (s stores) Len() int {
	return len(s)
}

// Less reports whether the element with index i
func (s stores) Less(i int, j int) bool {
	if s[j][1] < s[i][1] {
		return true
	} else if s[j][1] > s[i][1] {
		return false
	} else {
		if s[j][0] < s[i][0] {
			return true
		} else {
			return false
		}
	}
}

// Swap swaps the elements with indexes i and j.
func (s stores) Swap(i int, j int) {
	(s)[i], (s)[j] = (s)[j], (s)[i]
}
