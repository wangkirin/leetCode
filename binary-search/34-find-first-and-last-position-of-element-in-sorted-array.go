package main

func main() {

}

//二分法找到元素，然后再往左右扩展
func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	fist := -1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			fist = mid
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if fist == -1 {
		return []int{-1, -1}
	}
	ls, rs := fist, fist
	for ls >= 1 {
		if nums[ls-1] == target {
			ls -= 1
			continue
		}
		break
	}
	for rs <= len(nums)-2 {
		if nums[rs+1] == target {
			rs += 1
			continue
		}
		break
	}
	return []int{ls, rs}
}

//考虑 target 开始和结束位置，其实我们要找的就是数组中「第一个等于target 的位置」和「第一个大于 target 的位置减一」
// 即，target向左收缩一次，target向右收缩一次
// 用两个边界方法
func searchRange(nums []int, target int) []int {
	// 目标值开始位置：为 target 的左侧边界
	start := findLeftBound(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 目标值结束位置：为 target 的右侧边界
	end := findRightBound(nums, target)
	return []int{start, end}
}

// 寻找左侧边界的二分查找
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // note: [left, right]
	for left <= right {           // note: 因为 right 是闭区间，所以可以取 =
		mid := left + ((right - left) >> 1) // mid = (left + right) / 2 的优化形式，防止溢出！
		if nums[mid] == target {
			right = mid - 1 // note: 收紧右侧边界以锁定左侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 返回左侧边界
	return left // note
}

// 寻找右侧边界的二分查找
func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 返回右侧边界
	return right
}

//调用系统接口解法
//go 语言是有自己实现的二分查找方法 sort.Search()

//func searchRange(nums []int, target int) []int {
//	leftmost := sort.SearchInts(nums, target)
//	if leftmost == len(nums) || nums[leftmost] != target {
//		return []int{-1, -1}
//	}
//	rightmost := sort.SearchInts(nums, target + 1) - 1
//	return []int{leftmost, rightmost}
//}
