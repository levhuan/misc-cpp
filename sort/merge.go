/*
 * Block-style comment is for the package documentation.
 * 
 * Huan Le, mr.huanle@gmail.com 
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_INPUT_LENGTH = 10
	MAX_INPUT_VALUE  = 2000
)

var (
	num_msort_recursion int = 0
	num_qsort_recursion int = 0
	num_count_recursion int = 0
)

// Merge performs merge-sort algorithm by Divide-Conquer-Combine step.
// (1) Divide: merge-sort the left and right set of input
// (2) Conquer: works on subset
// (3) Combine: merge the result of two sorted sublists
// 
// Worst-case performance complexity: O(nlog(n))
// Memory performance: O(n^2)

func Merge(input []int) []int {

	num_msort_recursion++

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

func MergeCountInversion(input []int) ([]int, int) {

	num_count_recursion++

	length := len(input)
	output := make([]int, length, length)

	if length == 1 {
		output[0] = input[0]
		return output, 0
	}

	sublen := length/2

	left_slice, left_inversion := MergeCountInversion(input[0:sublen])
	right_slice, right_inversion := MergeCountInversion(input[sublen:length])

	lindex := 0
	rindex := 0
	my_inversion := 0

	for i := 0; i < length; i++ {
		if lindex == len(left_slice) {
			output[i] = right_slice[rindex]
			rindex++
		} else if rindex == len(right_slice) {
			output[i] = left_slice[lindex]
			lindex++
		} else if left_slice[lindex] > right_slice[rindex] {
			output[i] = right_slice[rindex]
			my_inversion +=  len(left_slice) - lindex
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
	return output, my_inversion + left_inversion + right_inversion;
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
	input[from] = input[pivot + 1]
	input[pivot + 1] = input[pivot]
	input[pivot] = temp
	return pivot + 1
}

// moveToRight moves "from" to the right of "pivot"
func moveToRight(input []int, pivot, from int) int {
	temp := input[from]
	input[from] = input[pivot-1]
	input[pivot-1] = input[pivot]
	input[pivot] = temp
	return pivot - 1
}

// selectPivot returns the index and value of the pivot
//
func selectRandomPivot(input []int) int {
	return rand.Intn(len(input))
}

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
			//
			// for duplicated value, don't need to swap as it will
			// be moved to its rightful position in subsequent recursions
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

	num_qsort_recursion++

	length := len(input)
	if length <= 1 {
		return 
	}

	pindex := selectRandomPivot(input)

	pindex = partition(input, pindex)

	Quick(input[0:pindex])
	Quick(input[pindex + 1: length])
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := make([]int, rand.Intn(MAX_INPUT_LENGTH))
	for i,_ := range input {
		input[i] = rand.Intn(MAX_INPUT_VALUE)
	}

	fmt.Println("Input:        ", input)

	stamp := time.Now().UTC().UnixNano()
	output := Merge(input)
	merge_perf := time.Now().UTC().UnixNano() - stamp

	fmt.Println("Merge-sorted: ", output)
	fmt.Println("Merge recursion: ", num_msort_recursion, 
		"Perf: ", merge_perf)

	stamp = time.Now().UTC().UnixNano()
	output, inversions := MergeCountInversion(input)
	merge_perf = time.Now().UTC().UnixNano() - stamp

	fmt.Println("Count Inversions: ", output)
	fmt.Println("Count recursion: ", num_count_recursion, " inversions: ", inversions,
		"Perf: ", merge_perf)

	stamp = time.Now().UTC().UnixNano()
	Quick(input)
	merge_perf = time.Now().UTC().UnixNano() - stamp
	fmt.Println("Q-sorted:     ", input)
	fmt.Println("Quick recursion: ", num_qsort_recursion,
		"Perf: ", merge_perf)
}
