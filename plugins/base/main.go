package main

import (
	t "github.com/marianogappa/rungopherrun/types"
)

// Name is the name of the plugin
var Name = "base"

// FilterMandatoryRequirements returns all ads that fulfill all requirements
func FilterMandatoryRequirements(ads []t.Ad, requirements []string) []t.Ad {
	var filtered = make([]t.Ad, 0)
	for _, ad := range ads {
		var found = make(map[string]struct{}, 0)
		for _, r := range requirements {
			for i := 0; i < len(ad.Description)-len(r)+1; i++ {
				if ad.Description[i:i+len(r)] == r {
					found[r] = struct{}{}
				}
			}
		}
		if len(found) == len(requirements) {
			filtered = append(filtered, ad)
		}
	}
	return filtered
}

// CalculateDistance returns the distance from vertices i and j (intersections)
func CalculateDistance(i, j int, graphEdges [][]t.Edge) int {
	return floydWarshall(i, j, graphEdges)
}

// SortBasedOnPriceAndDistance sorts in-place; doesn't matter if it's stable
func SortBasedOnPriceAndDistance(ads []t.Ad) {
	bubbleSort(ads)
}

func main() {
}

func bubbleSort(ads []t.Ad) {
	for i := 0; i < len(ads); i++ {
		for j := 0; j < len(ads); j++ {
			if i < j {
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

func splitBySpace(s string) []string {
	var (
		r  = make([]string, 0)
		bs = make([]byte, 0)
	)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && len(bs) > 0 {
			r = append(r, string(bs))
			bs = make([]byte, 0)
		} else {
			bs = append(bs, s[i])
		}
	}
	if len(bs) > 0 {
		r = append(r, string(bs))
	}
	return r
}
