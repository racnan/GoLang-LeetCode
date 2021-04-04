func twoSum(nums []int, target int) []int {
    
    for i,_ := range nums {
        for j,_ := range nums {
            if i !=j {
                if nums[i] + nums[j] == target{
                    return []int{i,j}
                }
            }
        }
    }
    return []int{0,1}
}
