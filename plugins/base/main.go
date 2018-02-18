package main

import (
	t "github.com/marianogappa/rungopherrun/types"
)

// Name is the name of the plugin
var Name = "base"

// CalculateDistance returns the distance from vertices i and j (intersections)
func CalculateDistance(i, j int, graphEdges [][]t.Edge) int {
	return floydWarshall(i, j, graphEdges)
}

// SortBasedOnPriceAndDistance sorts in-place; doesn't matter if it's stable
func SortBasedOnPriceAndDistance(ads []t.Ad, councilApproved func(t.Ad) bool) {
	bubbleSort(ads, councilApproved)
}

func main() {
}

func bubbleSort(ads []t.Ad, councilApproved func(t.Ad) bool) {
	for i := 0; i < len(ads); i++ {
		for j := 0; j < len(ads); j++ {
			if i < j {
				if councilApproved(ads[i]) && !councilApproved(ads[j]) {
					continue
				}
				if councilApproved(ads[j]) && !councilApproved(ads[i]) {
					ads[i], ads[j] = ads[j], ads[i]
					continue
				}
				var pi, pj = ads[i].Price / 100, ads[j].Price / 100
				if pi > pj || (pi == pj && ads[i].Distance > ads[j].Distance) {
					ads[i], ads[j] = ads[j], ads[i]
				}
			}
		}
	}
}

func floydWarshall(source, target int, g [][]t.Edge) int {
	var d = make([][]int, len(g))
	for i := range g {
		d[i] = make([]int, len(g))
		for j := range g {
			if i == j {
				d[i][j] = 0
				continue
			}
			d[i][j] = 1<<31 - 1
			for _, e := range g[i] {
				if e.V == j {
					d[i][j] = e.D
				}
			}
		}
	}
	for k := range g {
		for i := range g {
			for j := range g {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return d[source][target]
}
