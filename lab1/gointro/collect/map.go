package collect

import (
	"maps"
	"slices"
)

func keys(x map[string]int) []string {
	return slices.Sorted(maps.Keys(x))
}

func values(x map[string]int) []int {
	return slices.Sorted(maps.Values(x))
}
