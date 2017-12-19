func moveZeroes(nums []int)  { 
    n := 0
    for i := 0; i < len(nums); i ++ {
        if nums[i] != 0 {
            nums[i], nums[n] = nums[n], nums[i]
            n += 1
        }
    }
    // for i, v := range nums {
    //     if v != 0 {
    //         nums[i], nums[n] = nums[n], nums[i]
    //         n += 1
    //     }
    // }
}