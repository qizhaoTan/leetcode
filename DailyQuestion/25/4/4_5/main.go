package __5

func dg(nums []int) int {

}

func subsetXORSum(nums []int) int {
	var a, b int
	for _, num := range nums {
		b ^= num
	}

	var total int
	total += b

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			num := nums[i+j%len(nums)]
			a ^= num
			b ^= num
			total += b
			total += a
		}
	}

	return total / 2
}
