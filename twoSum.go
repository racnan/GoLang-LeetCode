//Runtime: 4ms; Runtime Percentile: 92.13%
//Memory: 3.2MB; Memory Percentile: 99.76%


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
