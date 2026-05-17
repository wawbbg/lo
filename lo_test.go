package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()

	result1 := Contains([]int{1, 2, 3, 4}, 3)
	result2 := Contains([]int{1, 2, 3, 4}, 5)
	result3 := Contains([]string{"foo", "bar"}, "bar")
	result4 := Contains([]string{"foo", "bar"}, "baz")

	assert.True(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
	assert.False(t, result4)
}

func TestMap(t *testing.T) {
	t.Parallel()

	result1 := Map([]int{1, 2, 3, 4}, func(x int, _ int) int {
		return x * 2
	})
	result2 := Map([]string{"foo", "bar"}, func(x string, _ int) string {
		return x + "!"
	})

	assert.Equal(t, []int{2, 4, 6, 8}, result1)
	assert.Equal(t, []string{"foo!", "bar!"}, result2)
}

func TestFilter(t *testing.T) {
	t.Parallel()

	result1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	result2 := Filter([]string{"foo", "bar", "baz"}, func(x string, _ int) bool {
		return len(x) == 3
	})

	assert.Equal(t, []int{2, 4}, result1)
	assert.Equal(t, []string{"foo", "bar", "baz"}, result2)
}

func TestReduce(t *testing.T) {
	t.Parallel()

	result1 := Reduce([]int{1, 2, 3, 4}, func(acc int, x int, _ int) int {
		return acc + x
	}, 0)
	result2 := Reduce([]string{"foo", "bar", "baz"}, func(acc string, x string, _ int) string {
		return acc + x
	}, "")

	assert.Equal(t, 10, result1)
	assert.Equal(t, "foobarbaz", result2)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	collected := []int{}
	ForEach([]int{1, 2, 3, 4}, func(x int, _ int) {
		collected = append(collected, x)
	})

	assert.Equal(t, []int{1, 2, 3, 4}, collected)
}

func TestUniq(t *testing.T) {
	t.Parallel()

	result1 := Uniq([]int{1, 2, 2, 3, 4, 4})
	result2 := Uniq([]string{"foo", "bar", "foo", "baz"})

	assert.Equal(t, []int{1, 2, 3, 4}, result1)
	assert.Equal(t, []string{"foo", "bar", "baz"}, result2)
}

func TestFlatten(t *testing.T) {
	t.Parallel()

	result1 := Flatten([][]int{{1, 2}, {3, 4}, {5}})
	result2 := Flatten([][]string{{"foo", "bar"}, {"baz"}})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, result1)
	assert.Equal(t, []string{"foo", "bar", "baz"}, result2)
}

func TestKeys(t *testing.T) {
	t.Parallel()

	result := Keys(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	assert.Len(t, result, 3)
	assert.Contains(t, result, "foo")
	assert.Contains(t, result, "bar")
	assert.Contains(t, result, "baz")
}

func TestValues(t *testing.T) {
	t.Parallel()

	result := Values(map[string]int{"foo": 1, "bar": 2, "baz": 3})

	assert.Len(t, result, 3)
	assert.Contains(t, result, 1)
	assert.Contains(t, result, 2)
	assert.Contains(t, result, 3)
}

func TestReverse(t *testing.T) {
	t.Parallel()

	result1 := Reverse([]int{1, 2, 3, 4, 5})
	result2 := Reverse([]string{"foo", "bar", "baz"})

	assert.Equal(t, []int{5, 4, 3, 2, 1}, result1)
	assert.Equal(t, []string{"baz", "bar", "foo"}, result2)
}
