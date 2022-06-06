package sol

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		// break condition
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}
		// prune not possible path
		if i >= len(candidates) || total > target {
			return
		}

		// choose i
		cur = append(cur, candidates[i])
		dfs(i, cur, total+candidates[i])

		// not choose i, next choose
		cur = cur[:len(cur)-1]
		dfs(i+1, cur, total)
	}
	dfs(0, []int{}, 0)
	return result
}
