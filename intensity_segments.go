package main

import (
	"container/list"
	"fmt"
)

// Segments holds list of segments, which are sorted as ascent by its interval.
type IntensitySegments struct {
	segs *list.List
}

// NewIntensitySegments creates a new `IntensitySegments` with a default segment with interval `[MinInf, MaxInf)`.
//
// Usage:
//
//	 s := NewIntensitySegments()
//		s.Add(10, 30, 1)   ==> '[(10,1)(30,0)]'
//		s.Add(20, 40, 1)   ==> '[(10,1)(20,2)(30,1)(40,0)]'
//		s.Add(10, 40, -1)  ==> '[(20,1)(30,0)]'
//		s.Add(10, 40, -1)  ==> '[(10,-1)(20,0)(30,-1)(40,0)]'
func NewIntensitySegments() *IntensitySegments {
	s := &IntensitySegments{
		segs: list.New(),
	}
	s.segs.PushBack(NewSegment(MinInf, MaxInf, 0))
	return s
}

// Add amount of intensity for range `[from, end)`
func (s *IntensitySegments) Add(from int, end int, amount int) {
	s.setWithRange(from, end, amount, true)
	s.compact()
}

// Set amount of intensity for range `[from, end)`
func (s *IntensitySegments) Set(from int, end int, amount int) {
	s.setWithRange(from, end, amount, false)
	s.compact()
}

// Set indensity for range `[from, end)` with `amount`.
// If `delta` is true, `amount` is treated as a delta value.
func (s *IntensitySegments) setWithRange(from int, end int, amount int, delta bool) {
	for e := s.segs.Front(); e != nil && from < end; e = e.Next() {
		seg, _ := e.Value.(*Segment)
		if from >= seg.end {
			continue
		}
		// Make a copy of current segment, because `seg` maybe moved later
		cpySeg := *seg
		curEnd := min(seg.end, end)

		// Set with `newIntensity` with segments included in `[from, end)`
		newIntensity := amount
		if delta {
			newIntensity = seg.intensity + amount
		}

		if from == seg.from {
			seg.intensity = newIntensity
			seg.end = curEnd
		} else {
			seg.end = from
			// Insert new segment and move `e` to next
			e = s.segs.InsertAfter(NewSegment(from, curEnd, newIntensity), e)
			seg, _ = e.Value.(*Segment)
		}

		if curEnd < cpySeg.end {
			// Insert new segment and move `e` to next,
			// use the current's intensity.
			e = s.segs.InsertAfter(NewSegment(curEnd, cpySeg.end, cpySeg.intensity), e)
			seg, _ = e.Value.(*Segment)
		}

		// Increase the `from` position
		from += (seg.end - from)
	}
	// The last segment
	if from < end {
		// Use intensity as `amount`
		s.segs.PushBack(NewSegment(from, end, amount))
	}
}

// Compact adjacent segments with the same intensity.
// For example '(10,1)(20,1)(30,0)' will be compacted to '(10,1)(30,0)'.
func (s *IntensitySegments) compact() {
	for e := s.segs.Front(); e != nil && e.Next() != nil; {
		nextE := e.Next()
		curSeg, _ := e.Value.(*Segment)
		nextSeg, _ := nextE.Value.(*Segment)
		if curSeg.intensity != nextSeg.intensity {
			// Move next
			e = e.Next()
			continue
		}
		// Set current segment's end to next segment's end,
		// and remove the next segment from list.
		curSeg.end = nextSeg.end
		s.segs.Remove(nextE)
	}
}

func (s *IntensitySegments) String() string {
	// Don't print the head element
	e := s.segs.Front().Next()
	ss := "["
	for ; e != nil; e = e.Next() {
		seg, _ := e.Value.(*Segment)
		ss += fmt.Sprintf("(%v,%v)", seg.from, seg.intensity)
	}
	ss += "]"
	return ss
}
