package main

import "fmt"

// Segment represents `intensity` of an interval `[from, end)`.
type Segment struct {
	from      int
	end       int
	intensity int
}

// Create a segment with interval and intensity.
func NewSegment(from int, end int, intensity int) *Segment {
	return &Segment{
		from:      from,
		end:       end,
		intensity: intensity,
	}
}

// Get printable string
func (seg *Segment) String() string {
	return fmt.Sprintf("(%d,%d,%d)", seg.from, seg.end, seg.intensity)
}
