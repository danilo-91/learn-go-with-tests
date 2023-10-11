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

func Reduce[A any](collection []A, fn func(A, A) A, initialValue A) A {
	var acc = initialValue
	for _, el := range collection {
		acc = fn(acc, el)
	}
	return acc
}

type Transaction struct {
    From, To string
    Sum float64
}

func BalanceFor(trs []Transaction, name string) float64 {
    var balance float64
    for _, t := range trs {
        if t.From == name {
            balance -= t.Sum
        }
        if t.To == name {
            balance += t.Sum
        }
    }
    return balance
}
