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
}
