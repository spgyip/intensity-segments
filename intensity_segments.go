package main

// Segments holds array of segments. Each segment represent an interval `[from,end)`.
// Segments are always sorted as ascent, with adjacent intervals.
type IntensitySegments struct {
	segs []*Segment
}

// NewIntensitySegments creates a new `IntensitySegments`.
// The infinity interval `[MinInf, MaxInf)` is added by default, so there are 2 segments '(MinInf,0)(MaxInf,0)'.
//
// Usage:
//  s := NewIntensitySegments()
//	s.Add(10, 30, 1)   ==> '[(10,1)(30,0)]'
//	s.Add(20, 40, 1)   ==> '[(10,1)(20,2)(30,1)(40,0)]'
//	s.Add(10, 40, -1)  ==> '[(20,1)(30,0)]'
//	s.Add(10, 40, -1)  ==> '[(10,-1)(20,0)(30,-1)(40,0)]'
//
func NewIntensitySegments() *IntensitySegments {
	s := &IntensitySegments{}
	s.segs = append(s.segs, NewSegment(MinInf, MaxInf, 0))
	s.segs = append(s.segs, NewSegment(MaxInf, MaxInf, 0))
	return s
}

// Add intensity for interval `[from, end)`.
func (s *IntensitySegments) Add(from int, end int, intensity int) {
	s.splitInterval(from, end, intensity, true)
}

// Set intensity for interval `[from, end)`.
func (s *IntensitySegments) Set(from int, end int, intensity int) {
	s.splitInterval(from, end, intensity, false)
}

// Split segment for interval `[from, end)`, iterate from `from` until `end` and split on each segment.
// If `deltaIntensity` is true, the `intensity` is a delta intensity, which means the new intensity should be computed by segment intensity.
func (s *IntensitySegments) splitInterval(from int, end int, intensity int, deltaIntensity bool) {
	// Iterate until the first segment's interval includes `from`
	idx := 0
	for {
		if from < s.segs[idx].End() {
			break
		}
		idx++
	}

	// Iterate from `from` until `end`, split segments orderly.
	for from < end {
		seg := s.segs[idx]
		splitEnd := min(end, seg.End())
		setIntensity := intensity
		if deltaIntensity {
			setIntensity += seg.Intensity()
		}
		n := s.split(idx, from, splitEnd, setIntensity)

		// Next
		idx += (n + 1)
		from = splitEnd
	}
	s.compact()
}

// Split interval at the i(th) segment, with range `[from, end)`.
// `[from, end] ` must be included by `[segs[i].from, segs[i].end)`.
// If split with the same range, set the new intensity.
// Return the number of newly splitted segments.
//
// Show cases:
//         --------             <------ new intensity
//  --------------------------  <------ origin intensity
//  |________________________|
//  x     x1      y1         y
//
//  [x, y) = (seg[i].from, seg[i].end)
//  [x1, y1) = [from, end)
//
// Condition 1 - Not included:
//   Condition: (x1 < x) OR  (y1>y)
//   Action: Invalid input range, do nothing.
//   Output: "(x, )(y, )"
//   Number of new segments: 0
//
//  _______|________________|_______
//  x1     x1               y      y1
//
// Condition 2 - Equal
//    Condition: (x1==x) AND (y1==y)
//    Action: No split, set new intensity
//    Output: "(x, <new intensity>)(y, )"
//    Number of new segments: 0
//  --------------------------  <------ new intensity
//
//  |________________________|
//  x                        y
//  x1                       y1
//
// Condition 3 - Left split
//    Condition: (x1==x) AND (y1<y)
//    Action: Split new segment left to i(th)
//    Output: "(x, <new intensity>)(y1,)(y, )"
//    Number of new segments: 1
//  ----------                  <------ new intensity
//            ----------------  <------ origin intensity
//  |________________________|
//  x         y1             y
//  x1
//
// Condition 4 - Right split
//    Condition: (x1>x) AND (y1==y)
//    Action: Split new segment right to i(th)
//    Output: "(x, )(x1,<new intensity>)(y, )"
//    Number of new segments: 1
//                  ---------- <------ new intensity
//  ----------------           <------ origin intensity
//  |________________________|
//  x              x1        y
//                           y1
//
// Condition 5 - Inter-split
//    Condition: (x1>x) AND (y1<y)
//    Action: Split new segment inter the segment
//    Output: "(x, )(x1,<new intensity>)(y1,)(y, )"
//    Number of new segments: 2
//          ----------         <------ new intensity
//  --------           ------- <------ origin intensity
//  |________________________|
//  x      x1        y1      y
//
func (s *IntensitySegments) split(i int, from int, end int, intensity int) int {
	seg := s.segs[i]
	if !(from >= seg.From() && end <= seg.End()) {
		// Must be included or don't need to split
		return 0
	}

	// Make a copy of current segment,
	cpySeg := *seg

	var count = 0
	if from > seg.From() {
		seg.SetEnd(from)
		s.insertAfter(
			i,
			NewSegment(from, end, intensity),
		)
		i++
		count++
	} else { // from==seg.From()
		seg.SetIntensity(intensity)
		seg.SetEnd(end)
	}

	if end < cpySeg.End() {
		s.insertAfter(
			i,
			NewSegment(end, cpySeg.End(), cpySeg.Intensity()),
		)
		count++
	}
	return count
}

// Compact adjacent segments with the same intensity.
// For example '(10,1)(20,1)(30,0)' will be compacted to '(10,1)(30,0)'.
func (s *IntensitySegments) compact() {
	// Don't compact the last `(MaxInf, 0)` segment,
	// so the loop end at `len(s.segs)-2`.
	for i := 0; i < len(s.segs)-2; {
		cur := s.segs[i]
		next := s.segs[i+1]
		if cur.Intensity() != next.Intensity() {
			i++
			continue
		}
		// Compact segments, by removing the next segment
		cur.SetEnd(next.End())
		s.remove(i + 1)
	}
}

// Insert new segment after idx, at idx+1.
func (s *IntensitySegments) insertAfter(idx int, seg *Segment) {
	s.segs = append(s.segs, nil) // Append a nil segment
	for i := len(s.segs) - 1; i > idx+1; i-- {
		s.segs[i] = s.segs[i-1]
	}
	s.segs[idx+1] = seg
}

// Remove segment at index idx.
func (s *IntensitySegments) remove(idx int) {
	for i := idx; i < len(s.segs)-1; i++ {
		s.segs[i] = s.segs[i+1]
	}
	s.segs = s.segs[:len(s.segs)-1]
}

// Return printable string.
func (s *IntensitySegments) String() string {
	return s.stringWithOpt(false)
}

// Return printable string.
// If `verbose` is true, verbose details will be printed out.
func (s *IntensitySegments) stringWithOpt(verbose bool) string {
	i, j := 1, len(s.segs)-2
	if verbose {
		i--
		j++
	}

	ss := "["
	for ; i <= j; i++ {
		ss += s.segs[i].stringWithOpt(verbose)
	}
	ss += "]"
	return ss
}
