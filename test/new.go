package main

import ("fmt")

func main() {
    arr := []int{1, 2, 3, 4, 5}


    sum := sumArray(arr)


    fmt.Printf("Sum of array elements: %d\n", sum)
}

func sumArray(arr []int) int {
	sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
