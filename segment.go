package main

import "fmt"

// Segment is a 3-elements-tupple represented by array of int, which represents an interval `[from, end)` of intensity.
// Array[0] represents the left closed interval `from`,
// Array[1] represents the right opened interval `end`,
// Array[2] represents the `intensity`.
type Segment [3]int

// Create a segment with interval and intensity.
func NewSegment(from int, end int, intensity int) *Segment {
	seg := &Segment{}
	seg[0] = from
	seg[1] = end
	seg[2] = intensity
	return seg
}

// Get the left closed interval `from`.
func (seg *Segment) From() int {
	return seg[0]
}

// Set the left closed interval `from`.
func (seg *Segment) SetFrom(from int) {
	seg[0] = from
}

// Get the right opened interval `end`.
func (seg *Segment) End() int {
	return seg[1]
}

// Set the right opened interval `end`.
func (seg *Segment) SetEnd(end int) {
	seg[1] = end
}

// Get the intensity.
func (seg *Segment) Intensity() int {
	return seg[2]
}

// Add the intensity.
func (seg *Segment) AddIntensity(delta int) {
	seg[2] += delta
}

// Return printable string.
func (seg *Segment) String() string {
	return seg.stringWithOpt(false)
}

// Return printable string.
// If `withEnd` is true, the end will be printed.
func (seg *Segment) stringWithOpt(withEnd bool) string {
	if withEnd {
		return fmt.Sprintf("(%d,%d,%d)", seg[0], seg[1], seg[2])
	}
	return fmt.Sprintf("(%d,%d)", seg[0], seg[2])
}
