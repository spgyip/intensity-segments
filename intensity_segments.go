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
	curSegIdx := 0
	for from < end {
		// Iterate until the first segment's interval includes `from`
		curSeg := s.segs[curSegIdx]
		if from >= curSeg.End() {
			curSegIdx++
			continue
		}

		if from == curSeg.From() {
			if end >= curSeg.End() {
				// ------------------------  <------- Current segment
				//
				// |______________________|
				// from                   end
				//
				curSeg.AddIntensity(intensity)
				curSegIdx++
				from = curSeg.End()
			} else {
				// ------                     <------- Current segment
				//       -----------------    <------- New split segment
				// |______________________|
				// from                   end
				//
				s.insertAfter(
					curSegIdx,
					NewSegment(end, curSeg.End(), curSeg.Intensity()),
				)
				curSeg.SetEnd(end)
				curSeg.AddIntensity(intensity)
				curSegIdx++
				from = end // This will break the loop
			}
			continue
		}

		// Must be `from > curSeg.From()` now.
		// Shrink current segment's interval to `[curSeg.From(), from)`.
		// `oldCurEnd` is keeped for future used.
		oldCurEnd := curSeg.End()
		curSeg.SetEnd(from)

		//         ----------------  <------- New split segment
		// --------                  <------- Current segment
		// |______________________|
		// from                   end
		//
		if end >= oldCurEnd {
			s.insertAfter(
				curSegIdx,
				NewSegment(from, oldCurEnd, curSeg.Intensity()+intensity),
			)
			curSegIdx += 2
			from = oldCurEnd
			continue
		}

		// Must be `end < oldCurEnd`.
		// Split the interval in inter-segment.
		//
		//         ---------         <------- New split segment
		// --------         ------   <------- Current segment/New split segment
		// |______________________|
		// from                   end
		//
		// Break the loop here
		s.insertAfter(
			curSegIdx,
			NewSegment(from, end, curSeg.Intensity()+intensity),
		)
		s.insertAfter(
			curSegIdx+1,
			NewSegment(end, oldCurEnd, curSeg.Intensity()),
		)
		curSegIdx += 3
		from = end // This will break the loop
	}
	s.strip()
}

// Set intensity for interval `[from, end)`.
//
//	---------------         <------- New split segment
//
// --------               ------   <------- Current segment
// |__________|_______|_________|
func (s *IntensitySegments) Set(from int, end int, intensity int) {
	// Shrink the first segment by `from`
	firstSegIdx := 0
	for {
		if from >= s.segs[firstSegIdx].End() {
			firstSegIdx++
			continue
		}
		break
	}
	if from > s.segs[firstSegIdx].From() {
		s.segs[firstSegIdx].SetEnd(from)
	} else {
		firstSegIdx--
	}

	// Remove segments included by `[from, end)`
	curSegIdx := firstSegIdx + 1
	for {
		if end < s.segs[curSegIdx].End() {
			break
		}
		s.remove(curSegIdx)
	}
	// Shrink the last segment by `end`, or insert a new segment
	if end > s.segs[curSegIdx].From() {
		s.segs[curSegIdx].SetFrom(end)
	}

	// Keep for later use
	savedIntensity := s.segs[curSegIdx-1].Intensity()

	// Insert a new segment with interval '[from, to)' after `firstSegIdx`
	s.insertAfter(
		firstSegIdx,
		NewSegment(from, end, intensity),
	)
	curSegIdx++ // Move advance because a new segment is inserted before current
	if end < s.segs[curSegIdx].From() {
		s.insertAfter(
			firstSegIdx+1,
			NewSegment(end, s.segs[curSegIdx].From(), savedIntensity),
		)
	}
	s.strip()
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
	// Attention: The last segment should always be 0 intensity.
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
	return s.stringWithOpts(false)
}

// Return printable string.
// If `withInf` is true, the MinInf/MaxInf info will not be hidden.
func (s *IntensitySegments) stringWithOpts(withInf bool) string {
	i, j := 0, len(s.segs)-1
	if !withInf {
		i++
		j--
	}

	ss := "["
	for ; i <= j; i++ {
		ss += s.segs[i].String()
	}
	ss += "]"
	return ss
}
