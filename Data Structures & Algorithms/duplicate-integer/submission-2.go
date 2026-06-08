func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})

    for _, n := range(nums) {
        if _, exists := (set[n]); exists {
            return true
        }
        set[n] = struct{}{}
    }
    return false
}
