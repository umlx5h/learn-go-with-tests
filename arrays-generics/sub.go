package main

func SumAll(numbersToSum ...[]int) []int {
	// sums := make([]int, 0, len(numbersToSum))
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	// var sum int
	// for _, number := range numbers {
	// 	sum += number
	// }
	// return sum

	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	if len(numbers) == 0 {
	// 		sums = append(sums, 0)
	// 	} else {
	// 		tail := numbers[1:]
	// 		sums = append(sums, Sum(tail))
	// 	}
	// }

	// return sums

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x[1:]))
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

// func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
func Reduce[A any, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

// Transaction{
// 	{
// 		From: "Chris",
// 		To:   "Riya",
// 		Sum:  100,
// 	},

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from Account, to Account, balance float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: balance}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

// func BalanceFor(transactions []Transaction, name string) float64 {
// 	var balance float64
// 	for _, t := range transactions {
// 		if t.From == name {
// 			balance -= t.Sum
// 		}

// 		if t.To == name {
// 			balance += t.Sum
// 		}
// 	}

// 	return balance
// }

// func BalanceFor(transactions []Transaction, name string) float64 {
// 	return Reduce(transactions, func(acc float64, t Transaction) float64 {
// 		if t.From == name {
// 			return acc - t.Sum
// 		}
// 		if t.To == name {
// 			return acc + t.Sum
// 		}

// 		return acc
// 	}, 0)
// }

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Find[T any](collections []T, searchFn func(e T) bool) (T, bool) {
	var zero T

	for _, e := range collections {
		if searchFn(e) {
			return e, true
		}
	}

	return zero, false
}
