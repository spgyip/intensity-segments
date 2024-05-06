package main

import "fmt"

const (
	MaxInf = int(^uint(0) >> 1)
	MinInf = -MaxInf - 1
)

// Segment is a 3-elements-tupple represented by array of int, which represents an interval `[from, end)` of intensity.
// Array[0] represents the left closed interval `from`,
// Array[1] represents the right opened interval `end`,
// Array[2] represents the `intensity`.
type Segment [3]int

// Create a segment with interval and intensity
func NewSegment(from int, end int, intensity int) *Segment {
	seg := &Segment{}
	seg[0] = from
	seg[1] = end
	seg[2] = intensity
	return seg
}

// Get the left closed interval `from`
func (seg *Segment) From() int {
	return seg[0]
}

// Get the right opened interval `end`
func (seg *Segment) End() int {
	return seg[1]
}

// Set the right opened interval `end`
func (seg *Segment) SetEnd(end int) {
	seg[1] = end
}

// Get the intensity
func (seg *Segment) Intensity() int {
	return seg[2]
}

// Add the intensity
func (seg *Segment) AddIntensity(delta int) {
	seg[2] += delta
}

// Return printable string
func (seg *Segment) String() string {
	return fmt.Sprintf("(%d, %d, %d)", seg[0], seg[1], seg[2])
}

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
	for curSegIdx < len(s.segs) {
		// Iterate until the first segment's interval includes `from`
		curSeg := s.segs[curSegIdx]
		if from >= curSeg.End() {
			curSegIdx++
			continue
		}

		// `[from, end)` includes current segment's interval
		//
		// ------------------------  <------- New intensity
		//
		// |______________________|
		// from                   end
		//
		if from == curSeg.From() && end >= curSeg.End() {
			curSeg.AddIntensity(intensity)
			curSegIdx++
			from = curSeg.End()
			continue
		}

		// Shrink current segment's interval to `[curSeg.From(), from)`.
		// `oldCurEnd` is keeped for future used.
		oldCurEnd := curSeg.End()
		curSeg.SetEnd(from)

		// Must be `from > curSeg.From()`
		//
		//         ----------------  <------- New intensity
		// --------                  <------- Old intensity
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
		// Split the interval in inter-segment
		//
		//         ---------         <------- New intensity
		// --------         ------   <------- Old intensity
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
		break
	}

}

// Insert new segment after index i, which means at index i+1.
func (s *IntensitySegments) insertAfter(idx int, seg *Segment) {
	s.segs = append(s.segs, nil) // Append a nil segment
	for i := len(s.segs) - 1; i > idx+1; i-- {
		s.segs[i] = s.segs[i-1]
	}
	s.segs[idx+1] = seg
}

// Set intensity for interval `[from, end)`.
func (s *IntensitySegments) Set(from int, end int, intensity int) {

}

// Return printable string
func (s *IntensitySegments) String() string {
	ss := "["
	for _, seg := range s.segs {
		ss += seg.String()
	}
	ss += "]"
	return ss
}

func main() {
	s := NewIntensitySegments()
	fmt.Println(s)
	s.Add(10, 30, 1)
	fmt.Println(s)
	s.Add(20, 40, 1)
	fmt.Println(s)
}
