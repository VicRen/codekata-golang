package main

import (
	"sort"
)

//56. 合并区间
//
//给出一个区间的集合，请合并所有重叠的区间。
//
//
//示例 1:
//
//输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
//
//提示：
//intervals[i][0] <= intervals[i][1]

func main() {

}

func mergeAndSortSlices(src [][]int) []int {
	var ret []int
	for _, s := range src {
		for _, n := range s {
			ret = append(ret, n)
		}
	}
	sort.Ints(ret)
	return ret
}

func pickInterval(src []int) [][]int {
	left := 0
	i := 1
	var ret [][]int
	for {
		if i+1 > len(src)-1 {
			if left < len(src) {
				ret = append(ret, [][]int{{src[left], src[i]}}...)
			}
			break
		}
		if src[i]+1 < src[i+1] {
			ret = append(ret, [][]int{{src[left], src[i+1]}}...)
			i += 2
			left = i
			continue
		}
		i++
	}
	return ret
}

func merge(src [][]int) [][]int {
	return pickInterval(mergeAndSortSlices(src))
}
