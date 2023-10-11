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
	return Reduce(slices, addTails, []int{})

}

func Reduce[A, B any](collection []B, fn func(A, B) A, initialValue A) A {
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
    fn := func(currentBalance float64, t Transaction) float64 {
        if t.From == name {
            return currentBalance - t.Sum
        }
        if t.To == name {
            return currentBalance + t.Sum
        }
        return currentBalance
    }
    return Reduce(trs, fn, 0.0)
}
