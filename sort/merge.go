/*
 * Block-style comment is for the package documentation.
 * 
 * Huan Le, mr.huanle@gmail.com 
 */
package main

import (
	"fmt"
	"math/rand"
)

// Merge performs merge-sort algorithm by Divide-Conquer-Combine step.
// (1) Divide: merge-sort the left and right set of input
// (2) Conquer: works on subset
// (3) Combine: merge the result of two sorted sublists
// 
// Worst-case performance complexity: O(nlog(n))
// Memory performance: O(n^2)

func Merge(input []int) []int {

	length := len(input)
	output := make([]int, length, length)

	if length == 1 {
		output[0] = input[0]
		return output
	}

    sublen := length/2

    left_slice := Merge(input[0:sublen])
	right_slice := Merge(input[sublen:length])

	lindex := 0
	rindex := 0

	for i := 0; i < length; i++ {
		if lindex == len(left_slice) {
			output[i] = right_slice[rindex]
			rindex++
		} else if rindex == len(right_slice) {
			output[i] = left_slice[lindex]
			lindex++
		} else if left_slice[lindex] > right_slice[rindex] {
			output[i] = right_slice[rindex]
			rindex++
		} else if left_slice[lindex] == right_slice[rindex] {
			output[i] = left_slice[lindex]
			lindex++
			i++
			output[i] = right_slice[rindex]
			rindex++
		} else { // if left_slice[lindex] < right_slice[rindex] 
			output[i] = left_slice[lindex]
			lindex++
		}
	}
	return output;
}

/*
 * Quick merge functions.
 * Quick merge includes:
 * - selectPivot
 * - partitions:
 * - recurse the left partition
 * - recurse the right partition
 */

// moveToLeft moves "from" to the left of "pivot"
func moveToLeft(input []int, pivot, from int) int {
	temp := input[from]
	for i := from - 1; i >= pivot; i-- {
		input[i+1] = input[i]
	}
	input[pivot] = temp
	return pivot + 1
}

// moveToRight moves "from" to the right of "pivot"
func moveToRight(input []int, pivot, from int) int {
	temp := input[from]
	for i := from + 1; i <= pivot; i++ {
		input[i-1] = input[i]
	}
	input[pivot] = temp
	return pivot - 1
}

// selectPivot returns the index and value of the pivot
//
func selectPivot(input []int) int {
	length := len(input)
	first := input[0]
	mid := input[length/2]
	last := input[length - 1]

	if first > last  &&  first < mid {
		return 0
	}
	if last > first  &&  last < mid {
		return length - 1
	}
	return length/2
}

func partition(input []int, pindex int) int {
	pivot := input[pindex]
	for i := 0; i < len(input); {
		value := input[i]
		if pivot < value  &&  pindex > i {
			pindex = moveToRight(input, pindex, i)
		} else if pivot > value  &&  pindex < i {
			pindex = moveToLeft(input, pindex, i)
		} else {
			// advance to next element only when not swapping as
			// swapping can move new value to the current index
			i++;
		}
	}

	return pindex
}

// Quick sorts the input in-place by choosing a pivot and move the pivot to
// it location/index in the slice. It recursively works on the left- and right-
// sub-array of the pivot index until the length of the sub-arrays are zero.
//
// Average performance complexity: O(nlog(n))
// Memory performance: O(1) - in-place
//
func Quick(input []int) {
	length := len(input)
	if length <= 1 {
		return 
	}

	pindex := selectPivot(input)

	pindex = partition(input, pindex)

	Quick(input[0:pindex])
	Quick(input[pindex + 1: length])
}

func main() {
	input := make([]int, 15, 15)
	random := rand.New(rand.NewSource(99))
	for i,_ := range input {
		input[i] = random.Int() % 20000
	}

	fmt.Println("Input:        ", input)

	output := Merge(input)
	fmt.Println("Merge-sorted: ", output)

	Quick(input)
	fmt.Println("Q-sorted:     ", input)
}
