package leetcode

func numTilePossibilities1(tiles string) int {
	res, tmp, tMap, used := 0, []byte{}, make(map[string]bool, 0), make([]bool, len(tiles))
	findTile([]byte(tiles), tmp, &used, 0, &res, tMap)
	return res
}

func findTile(tiles, tmp []byte, used *[]bool, index int, res *int, tMap map[string]bool) {
	flag := true
	for _, v := range *used {
		if !v {
			flag = false
			break
		}
	}
	if flag {
		return
	}
	for i := 0; i < len(tiles); i++ {
		if (*used)[i] {
			continue
		}
		tmp = append(tmp, tiles[i])
		(*used)[i] = true
		if _, ok := tMap[string(tmp)]; !ok {
			//fmt.Printf("i = %v tiles = %v 找到了结果 = %v\n", i, string(tiles), string(tmp))
			*res++
		}
		tMap[string(tmp)] = true

		findTile([]byte(tiles), tmp, used, i+1, res, tMap)
		tmp = tmp[:len(tmp)-1]
		(*used)[i] = false
	}
}

func numTilePossibilities(tiles string) int {
	m := make(map[byte]int)
	for i := range tiles {
		m[tiles[i]]++
	}
	arr := make([]int, 0)
	for _, v := range m {
		arr = append(arr, v)
	}
	return numTileDFS(arr)
}

func numTileDFS(arr []int) (r int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		r++
		arr[i]--
		r += numTileDFS(arr)
		arr[i]++
	}
	return
}
