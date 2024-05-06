package main

import (
	"testing"
)

func TestIntensitySegmentsEmpty(t *testing.T) {
	s := NewIntensitySegments()
	var expected = "[]"
	var got = s.String()
	if got != expected {
		t.Errorf("got(%v)!=expected(%v)\n", got, expected)
	}
}

func TestIntensitySegmentsSplit(t *testing.T) {
	s := NewIntensitySegments()
	for i, tc := range []struct {
		index     int
		from      int
		end       int
		intensity int
		expectedN int
		expected  string
	}{
		{0, 50, 200, 1, 2, "[(50,1)(200,0)]"},
		{1, 50, 200, 2, 0, "[(50,2)(200,0)]"},                              // tc.from==from and tc.end==end, no split, set new intensity
		{1, 20, 210, 3, 0, "[(50,2)(200,0)]"},                              // tc.from<from and tc.end>end, no split
		{1, 50, 100, 3, 1, "[(50,3)(100,2)(200,0)]"},                       // tc.from==from
		{2, 120, 200, 4, 1, "[(50,3)(100,2)(120,4)(200,0)]"},               // tc.end==end
		{3, 130, 150, 5, 2, "[(50,3)(100,2)(120,4)(130,5)(150,4)(200,0)]"}, // tc.from>from && tc.end<end
	} {
		gotN := s.split(tc.index, tc.from, tc.end, tc.intensity)
		got := s.String()
		if gotN != tc.expectedN {
			t.Fatalf("Case[%v] fail: gotN(%v)!=expectedN(%v)\n", i, gotN, tc.expectedN)
		}
		if got != tc.expected {
			t.Fatalf("Case[%v] fail: got(%v)!=expected(%v)\n", i, got, tc.expected)
		}
	}
}

func TestIntensitySegmentsAdd(t *testing.T) {
	s := NewIntensitySegments()
	for i, tc := range []struct {
		from      int
		end       int
		intensity int
		expected  string
	}{
		{10, 30, 1, "[(10,1)(30,0)]"},
		{20, 40, 1, "[(10,1)(20,2)(30,1)(40,0)]"},
		{10, 40, -1, "[(20,1)(30,0)]"},
		{10, 40, -1, "[(10,-1)(20,0)(30,-1)(40,0)]"},
	} {
		s.Add(tc.from, tc.end, tc.intensity)
		got := s.String()
		if got != tc.expected {
			t.Fatalf("Case[%v] fail: got(%v)!=expected(%v)\n", i, got, tc.expected)
		}
	}
}

func TestIntensitySegmentsSet(t *testing.T) {
	s := NewIntensitySegments()
	s.Add(100, 300, 1)
	for i, tc := range []struct {
		from      int
		end       int
		intensity int
		expected  string
	}{
		{100, 300, 2, "[(100,2)(300,0)]"},               // tc.from==from and tc.end==end, in one segment, no split, set new intensity
		{100, 200, 3, "[(100,3)(200,2)(300,0)]"},        // tc.from==from, range in one segment
		{150, 200, 4, "[(100,3)(150,4)(200,2)(300,0)]"}, // tc.end==end, range in one segment
		{150, 200, 4, "[(100,3)(150,4)(200,2)(300,0)]"}, // Repeatly set, unchanged
		{120, 250, 5, "[(100,3)(120,5)(250,2)(300,0)]"}, // Range in multiple segments
	} {
		s.Set(tc.from, tc.end, tc.intensity)
		got := s.String()
		if got != tc.expected {
			t.Fatalf("Case[%v] fail: got(%v)!=expected(%v)\n", i, got, tc.expected)
		}
	}
}
