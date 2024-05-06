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
		{0, 50, 200, 1, 0, "[(50,1)(200,0)]"},                              // tc.from==from and tc.end==end, no split
		{0, 20, 210, 1, 0, "[(50,1)(200,0)]"},                              // tc.from<from and tc.end>end, no split
		{1, 50, 100, 2, 1, "[(50,2)(100,1)(200,0)]"},                       // tc.from==from
		{2, 120, 200, 3, 1, "[(50,2)(100,1)(120,3)(200,0)]"},               // tc.end==end
		{3, 130, 150, 4, 2, "[(50,2)(100,1)(120,3)(130,4)(150,3)(200,0)]"}, // tc.from>from && tc.end<end
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
	s.Add(10, 30, 1)
	s.Add(20, 40, 1)
	for i, tc := range []struct {
		from      int
		end       int
		intensity int
		expected  string
	}{
		{15, 25, 3, "[(10,1)(15,3)(25,2)(30,1)(40,0)]"},
		{15, 35, 3, "[(10,1)(15,3)(35,1)(40,0)]"},
	} {
		s.Set(tc.from, tc.end, tc.intensity)
		got := s.String()
		if got != tc.expected {
			t.Fatalf("Case[%v] fail: got(%v)!=expected(%v)\n", i, got, tc.expected)
		}
	}
}
