package main

// Segments holds array of segments representing multiple adjacent intervals.
// `segs` are always sorted as ascent.
type IntensitySegments struct {
	segs []*Segment
}

// NewIntensitySegments creates a new `IntensitySegments` with initiated interval `[MinInf, MaxInf)`.
func NewIntensitySegments() *IntensitySegments {
	s := &IntensitySegments{}
	s.segs = append(s.segs, NewSegment(MinInf, MaxInf, 0))
	s.segs = append(s.segs, NewSegment(MaxInf, MaxInf, 0))
	return s
}

// Add intensity for interval `[from, end)`.
func (s *IntensitySegments) Add(from int, end int, intensity int) {
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
		n := s.split(idx, from, splitEnd, seg.Intensity()+intensity)

		// Next
		idx += (n + 1)
		from = splitEnd
	}
	s.strip()
}

// Set intensity for interval `[from, end)`.
func (s *IntensitySegments) Set(from int, end int, intensity int) {
	// Iterate until the first segment's interval includes `from`
	idx := 0
	for {
		if from < s.segs[idx].End() {
			break
		}
		idx++
	}

	// Split first segment
	splitEnd := min(end, s.segs[idx].End())
	n := s.split(idx, from, splitEnd, intensity)
	idx += n

	// The new segment has beed splitted
	if s.segs[idx].From() >= end {
		return
	}

	// Remove segments included by `[from, end)`
	for {
		if end < s.segs[idx].End() {
			break
		}
		s.remove(idx)
	}

	// Set last segment's from to `end`
	s.segs[idx].SetFrom(end)

	// Insert a new segment with interval '[from, to)'
	s.insertAfter(
		idx-1,
		NewSegment(from, end, intensity),
	)
	//idx++ // Move advance

	// Strip
	s.strip()
}

// Split interval at the i(th) segment, with range `[from, end)`.
// `[from, end)` must be included by `[segs[i].from, segs[i].end)`.
// If `[from, end)` equals to `[segs[i].from, segs[i].end)`, no split will do but new intensity is set for `segs[i]`.
// Return the number of newly splitted segments.
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

// Strip segments with extra 0 intensity in head or tail.
// Don't iterate the MinInf and MaxInf segment.
func (s *IntensitySegments) strip() {
	var i int
	// Strip head
	for i = 1; i <= len(s.segs)-2; {
		cur := s.segs[i]
		if cur.Intensity() != 0 {
			break
		}
		s.segs[i-1].SetEnd(cur.End())
		s.remove(i)
	}

	// Strip tail
	// Iterate from tail until the first segment with no-zero intensity.
	// Causion: The last segment is always 0 intensity, the iteration must be from second-last segment.
	for i = len(s.segs) - 3; i >= 1; i-- {
		if s.segs[i].Intensity() != 0 {
			break
		}
	}
	i++ // Switch to next
	// Remove the next segment
	for i <= len(s.segs)-3 {
		s.segs[i].SetEnd(s.segs[i+1].End())
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
