package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers[:])
		want := 15
		assert.Equal(t, want, got, "numbers:", numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assert.Equal(t, want, got)

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v", got, want)
	// }
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 9, 2}, []int{})
		want := []int{2, 11, 0}

		assert.Equal(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		assert.Equal(t, want, got)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		assert.Equal(t, 6, Reduce([]int{1, 2, 3}, multiply, 1))
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		assert.Equal(t, "abc", Reduce([]string{"a", "b", "c"}, concatenate, ""))
	})
}

func TestBadBank(t *testing.T) {

	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	assert.EqualValues(t, 200, newBalanceFor(riya))
	assert.EqualValues(t, 0, newBalanceFor(chris))
	assert.EqualValues(t, 175, newBalanceFor(adil))
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		assert.True(t, found)
		assert.Equal(t, 2, firstEvenNumber)

	})

	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		assert.True(t, found)
		assert.Equal(t, Person{Name: "Chris James"}, king)
	})
}
