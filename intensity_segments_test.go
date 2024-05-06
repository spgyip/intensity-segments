package main

import (
	"testing"
)

func TestIntensitySegmentsEmpty(t *testing.T) {
	s := NewIntensitySegments()
	var expected = "[]"
	var got = s.String()
	if got != expected {
		t.Errorf("got(%v)!=expectd(%v)\n", got, expected)
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
			t.Fatalf("Case[%v] fail: got(%v)!=expectd(%v)\n", i, got, tc.expected)
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
			t.Fatalf("Case[%v] fail: got(%v)!=expectd(%v)\n", i, got, tc.expected)
		}
	}
}
