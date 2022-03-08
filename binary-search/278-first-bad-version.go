package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) == false {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func isBadVersion(version int) bool {

}
