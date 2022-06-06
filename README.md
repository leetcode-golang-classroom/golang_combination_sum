# golang_combination_sum

Given an array of **distinct** integers `candidates` and a target integer `target`, return *a list of all **unique combinations** of* `candidates` *where the chosen numbers sum to* `target`*.* You may return the combinations in **any order**.

The **same** number may be chosen from `candidates` an **unlimited number of times**. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is **guaranteed** that the number of unique combinations that sum up to `target` is less than `150` combinations for the given input.

## Examples

**Example 1:**

```
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

```

**Example 2:**

```
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

```

**Example 3:**

```
Input: candidates = [2], target = 1
Output: []

```

**Constraints:**

- `1 <= candidates.length <= 30`
- `1 <= candidates[i] <= 200`
- All elements of `candidates` are **distinct**.
- `1 <= target <= 500`

## 解析

給定一個整數陣列 nums 還有一個整數 target 

要求要寫出一個演算法來找出和為 target 的所有 nums 內元素不重複組合

可以利用類似 PowerSet 個概念

分為選擇該數值或是不選該數值的 decision tree

每次先檢查當選了該數值還可以繼續則及續選

否則就往下一個數值檢查

概念圖如下

![](https://i.imgur.com/dtvXlz9.png)

## 程式碼
```go
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

```
## 困難點

1. 理解如何做窮舉法
2. 理解如何窮舉出不重複的方式
3. 需要理解 golang slice 只是 array 的 reference, 要保存值只能透過 copy 方式

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity