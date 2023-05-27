package logicprogramming

func removeDuplicates(nums []int) int {
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }

    return i+1
}