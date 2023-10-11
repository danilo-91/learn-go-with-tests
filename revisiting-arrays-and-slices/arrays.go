package main

func Sum(nums []int) int {
	add := func(acc, el int) int { return acc + el }
	return Reduce[int](nums, add, 0)
}

func SumAllTails(slices ...[]int) []int {
	addTails := func(acc, el []int) []int {
		if len(el) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(el[1:]))
		}
	}
	return Reduce[[]int](slices, addTails, []int{})

}

func Reduce[A any](collection []A, acc func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, el := range collection {
		result = acc(result, el)
	}
	return result
}
