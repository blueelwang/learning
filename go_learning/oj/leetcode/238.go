package leetcode

func moveZeroes(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	l := len(nums)
	head := 0
	tail := l
	for head < tail {
		for ; head < tail; head++ {
			if nums[head] == 0 {
				break
			}
		}
		for tail--; head < tail; tail-- {
			if nums[tail] != 0 {
				break
			}
		}
		for i := head; i < tail; i++ {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}

}

func moveZeroes2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := len(nums)
	firstZero := 0
	firstNonZero := firstZero

	for firstNonZero < l - 1 {
		for ;firstZero < l - 1; firstZero++ {
			if nums[firstZero] == 0 {
				break
			}
		}
		if firstZero >= l {
			return
		}

		for firstNonZero = firstZero + 1; firstNonZero < l; firstNonZero++ {
			if nums[firstNonZero] != 0 {
				break
			}
		}
		if firstNonZero >= l {
			return
		}
		nums[firstNonZero], nums[firstZero] = nums[firstZero], nums[firstNonZero]
	}
}

// NB
func moveZeroes3(nums []int) {
    i, j := 0, 0
    for i < len(nums) {
        if nums[i] != 0 { // i遇到非0
            nums[i], nums[j] = nums[j], nums[i] // 和j交换，安排到j的位置
            i++                                 // i推进，找下一个非0
            j++                                 // 安排了一个，j后移
        } else { // i遇到0，继续后移，找非0。j 继续指向待安排的坑位
            i++
        }
    }
}