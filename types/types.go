package types

// An Ad represents a rental ad.
// It's distance must be calculated based on location.
type Ad struct {
	Price       int
	Distance    int
	Location    int
	Description string
}

// An Edge represents the street between two intersections.
// It holds two values: the destination vertex and the distance to it.
type Edge struct {
	V, D int
}
