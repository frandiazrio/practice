func rotate(nums []int, k int) {
    k %= len(nums)
    if k == 0 {
        return
    }
    
    idx := 0
    for i := 0; i + k < len(nums); i++ {
		// swap
        nums[idx], nums[i+k] = nums[i+k], nums[idx]
        idx++
        idx %= k
    }
    
    rotate(nums[:k], k-idx)
}