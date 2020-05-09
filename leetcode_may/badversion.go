package main

import "fmt"

// You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.
// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
// You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
// Example:

// Given n = 5, and version = 4 is the first bad version.

// call isBadVersion(3) -> false
// call isBadVersion(5) -> true
// call isBadVersion(4) -> true

// Then 4 is the first bad version.

func isBadVersion(n int) bool {
	var badversion int = 1702766719
	var res bool
	if n >= badversion {
		res = true
	} else {
		res = false
	}
	return res
}

// 2126753390
// 1702766719

func firstBadVersion(n int) int {
	var left int = 1
	var right int = n
	for left < right {
		mid := left + (right-left)/2
		fmt.Println("Mid set to ", mid)
		if !isBadVersion(mid) {
			fmt.Println(mid, " is false")
			left = mid + 1
			fmt.Println("Left set to ", left)
		} else {
			right = mid
			fmt.Println("Mid is true, set mid as right ")
		}
		fmt.Println("==============================")
	}
	return left
}

func main() {
	var N int = 2126753390
	fmt.Println("First Bad Int is :- ", firstBadVersion(N))

}
