/*
 * Block-style comment is for the package documentation.
 *
 * Quick sort algorithm, implemented in GO. GO "slice" makes in-place sorting
 * implementation very simple and elegant. 
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
	MAX_INPUT_LENGTH = 50
	MAX_INPUT_VALUE  = 2000
)

var (
	num_qsort_recursion int = 0
	num_rselect_recursion int = 0
)

/*
 * Quick sort functions.
 * Quick sort includes:
 * - selectPivot
 * - partitions:
 * - recurse the left partition
 * - recurse the right partition
 */

// moveToLeft moves "from" to the immediate left of "pivot" and return
// the new pivot index
func moveToLeft(input []int, pivot, from int) int {
	temp := input[from]
	input[from] = input[pivot + 1]
	input[pivot + 1] = input[pivot]
	input[pivot] = temp
	return pivot + 1
}

// moveToRight moves "from" to the immediate right of "pivot" and 
// return the new pivot index
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

// RSelect returns the i'th order value in the input array
// where i: 0-based value
//
func RSelect(input []int, order int) int {

	num_rselect_recursion++

	length := len(input)
	if length == 1 {
		if order != 0 {
			fmt.Println("Error: input ", input, " length ", length, " order ", order)
		}
		return input[0]
	}

	pindex := partition(input, selectRandomPivot(input));

	switch {
		case pindex > order:
			return RSelect(input[0:pindex], order)
		case pindex < order:
			return RSelect(input[pindex+1:len(input)], order - pindex - 1)
		default:
			return input[pindex]
	}

}

/*
 * Main test Quick and RSelect functions 
 */
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := make([]int, rand.Intn(MAX_INPUT_LENGTH))
	input_clone := make([]int, len(input))
	for i,_ := range input {
		input[i] = rand.Intn(MAX_INPUT_VALUE)
		input_clone[i] = input[i]
	}

	fmt.Println("Input:        ", input)

	stamp := time.Now().UTC().UnixNano()
	Quick(input)
	merge_perf := time.Now().UTC().UnixNano() - stamp
	fmt.Println("Q-sorted:     ", input)
	fmt.Println("Quick recursion: ", num_qsort_recursion,
		"Perf: ", merge_perf)
	
	stamp = time.Now().UTC().UnixNano()
	data := RSelect(input_clone, 3)
	merge_perf = time.Now().UTC().UnixNano() - stamp

	fmt.Println("RSelect:     ", input_clone, " 3rd order: ", data)
	fmt.Println("RSelect recursion: ", num_rselect_recursion,
		"Perf: ", merge_perf)

	num_rselect_recursion = 0
	stamp = time.Now().UTC().UnixNano()
	data = RSelect(input_clone, len(input_clone) - 1)
	merge_perf = time.Now().UTC().UnixNano() - stamp

	fmt.Println("RSelect:     ", input_clone, " Last order: ", data)
	fmt.Println("RSelect recursion: ", num_rselect_recursion,
		"Perf: ", merge_perf)
}
