package main

func Sum(nums []int) int {
    var sum int
    for _, num := range nums {
        sum += num
    }
    return sum
}

func SumAll(slices ...[]int) []int {
    var sums []int
    for _, slice := range slices {
        sums = append(sums, Sum(slice))
    }
    return sums
}

func SumAllTails(slices ...[]int) []int {
    var sums []int
    for _, slice := range slices {
        if len(slice) < 2 {
            sums = append(sums, 0)
            continue
        }
        sums = append(sums, Sum(slice[1:]))
    }
    return sums
}

