package leetcode

func reverseBits(num uint32) uint32 {
	var res, bit uint32
	bit = 1
	for i := 0; i < 32; i++ {
		b := bit & num
		b = b >> i
		res += b
		bit = bit << 1
		if i != 31 {
			res = res << 1
		}
	}
	return res
}
