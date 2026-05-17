// Package lo is a Lodash-style Go library based on Go 1.18+ Generics.
// It provides a set of useful helpers for working with slices, maps, channels,
// strings, and other common data structures.
package lo

import (
	"math/rand"
)

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Filter iterates over elements of collection, returning an array of all elements
// for which predicate returns true.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))
	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// Reduce reduces collection to a value which is the accumulated result of running
// each element in collection through accumulator, where each successive invocation
// is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}
	return initial
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// Uniq returns a duplicate-free version of an array, in which only the first
// occurrence of each element is kept. The order of result values is determined
// by the order they occur in the array.
func Uniq[T comparable](collection []T) []T {
	seen := make(map[T]struct{}, len(collection))
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Keys creates an array of the map keys.
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))
	for k := range in {
		result = append(result, k)
	}
	return result
}

// Values creates an array of the map values.
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))
	for _, v := range in {
		result = append(result, v)
	}
	return result
}

// Shuffle returns an array of shuffled values. Uses a copy of the input so
// the original slice is never modified.
func Shuffle[T any](collection []T) []T {
	result := make([]T, len(collection))
	copy(result, collection)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// Reverse reverses array so that the first element becomes the last,
// the second element becomes the second to last, and so on.
// Returns a new slice and does not modify the original.
func Reverse[T any](collection []T) []T {
	result := make([]T, len(collection))
	copy(result, collection)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
