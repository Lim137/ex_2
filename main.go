package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func normalizeString(str string) string {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, ",", "")
	return str
}

func getIntArrFromStrArr(strings []string) []int {
	var nums []int
	for _, numStr := range strings {
		if numStr == "" || numStr == "0" || numStr == " " {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func getMinNum(nums []int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func getCommonDivisorsForNums(numsString string) []int {
	var commonDivisors []int
	numsString = normalizeString(numsString)
	numsArrString := strings.Split(numsString, " ")
	nums := getIntArrFromStrArr(numsArrString)
	minNum := getMinNum(nums)
	for divisor := 2; divisor <= minNum/2; divisor++ {
		for _, num := range nums {
			if num%divisor != 0 {
				break
			}
			if num == nums[len(nums)-1] {
				commonDivisors = append(commonDivisors, divisor)
			}
		}
	}
	if minNum > 1 {
		for _, num := range nums {
			if num%minNum != 0 {
				break
			}
			if num == nums[len(nums)-1] {
				commonDivisors = append(commonDivisors, minNum)
			}
		}
	}
	return commonDivisors
}

func main() {
	var numsString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numsString = scanner.Text()
	commonDivisors := getCommonDivisorsForNums(numsString)
	fmt.Println(commonDivisors)
}
