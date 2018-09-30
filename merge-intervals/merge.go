package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//result := merge([]Interval{
	//	{1, 2},
	//	{2, 6},
	//	{8, 10},
	//	{15, 18},
	//})

	result := merge([]Interval{
		{1, 4},
		{0, 0},
	})

	fmt.Println(result)
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	var merged []Interval

	if len(intervals) == 0{
		return merged
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	open := intervals[0]
	for _, v := range intervals[1:] {
		if v.Start <= open.End {
			open.End = int(math.Max(float64(open.End), float64(v.End)))
			open.Start = int(math.Min(float64(open.Start), float64(v.Start)))

		} else {
			merged = append(merged, open)
			open = v
		}

	}
	merged = append(merged, open)

	return merged
}
