package main



func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum 
}

// func SumAll(arr1 []int, arr2 [] int) []int {
// 	sumArr := [] int {}
// 	tot1 := 0
// 	tot2 := 0 

// 	for _, num := range arr1 {
// 		tot1 += num
// 	}
// 	for _, num := range arr2 {
// 		tot2 += num
// 	}

// 	sumArr = append(sumArr,tot1, tot2)

// 	return sumArr
// }   


// chapter example
// func SumAll(numbersToSum ...[]int) []int {
// 	// lengthOfNumbers is len of arguments
// 	lengthOfNumbers := len(numbersToSum)
// 	// sums will ba a slice that will have a size of lengthOfNumbers
// 	sums := make([]int, lengthOfNumbers)
//   // looping over numbers to sum wich are slices and calling Sum to get the sum 
// 	//  of each nested slice, then assigning that sum to sums at the index equal to the index of the numbersToSum slice
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums

// Refactor Example

// func SumAll(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		sums = append(sums, Sum(numbers))
// 	}
// 	return sums
// }


func SumAllTails(numbersToSum ...[]int) []int { 
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		}else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}