func findUnsortedSubarray(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    ori_nums := make([]int, len(nums))
    copy(ori_nums, nums)
    quick_sort(nums, 0 , len(nums) - 1)
    var indexes []int
    for i := 0; i < len(nums); i ++ {
        if nums[i] != ori_nums[i] {
            indexes = append(indexes, i)
        }
    }
    if len(indexes) == 0 {
        return 0
    }
    return indexes[len(indexes) - 1] - indexes[0] + 1
}

func quick_sort(arr []int, start, end int) {
    if start < end {
        i, j := start, end
        key := arr[(start+end)/2]
        for i <= j {
            for arr[i] < key {
                i++
            }
            for arr[j] > key {
                j--
            }
            if i <= j {
                arr[i], arr[j] = arr[j], arr[i]
                i++
                j--
            }
        }

        if start < j {
            quick_sort(arr, start, j)
        }
        if end > i {
            quick_sort(arr, i, end)
        }
    }
}