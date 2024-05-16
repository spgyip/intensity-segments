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

func TestIntensitySegmentsAdd(t *testing.T) {
	s := NewIntensitySegments()
	for i, tc := range []struct {
		from      int
		end       int
		intensity int
		expected  string
	}{
		// Overlap range
		{100, 300, 1, "[(100,1)(300,0)]"},
		{100, 300, 1, "[(100,2)(300,0)]"},                             // Repeatly add, no split, but increase the indensity
		{100, 200, 1, "[(100,3)(200,2)(300,0)]"},                      // tc.from==from, range in one segment
		{250, 300, 1, "[(100,3)(200,2)(250,3)(300,0)]"},               // tc.end==end, range in one segment
		{150, 280, 1, "[(100,3)(150,4)(200,3)(250,4)(280,3)(300,0)]"}, // Range multiple segments

		// No overlap range
		{20, 60, 1, "[(20,1)(60,0)(100,3)(150,4)(200,3)(250,4)(280,3)(300,0)]"},                 // No overlap, left
		{350, 400, 1, "[(20,1)(60,0)(100,3)(150,4)(200,3)(250,4)(280,3)(300,0)(350,1)(400,0)]"}, // No overlap, right
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
		// Overlap range
		{100, 300, 2, "[(100,2)(300,0)]"},               // tc.from==from and tc.end==end, in one segment, no split, set new intensity
		{100, 200, 3, "[(100,3)(200,2)(300,0)]"},        // tc.from==from, range in one segment
		{150, 200, 4, "[(100,3)(150,4)(200,2)(300,0)]"}, // tc.end==end, range in one segment
		{150, 200, 4, "[(100,3)(150,4)(200,2)(300,0)]"}, // Repeatly set, unchanged
		{120, 250, 5, "[(100,3)(120,5)(250,2)(300,0)]"}, // Range in multiple segments

		// No overlap range
		{20, 60, 1, "[(20,1)(60,0)(100,3)(120,5)(250,2)(300,0)]"},                 // No overlap, left
		{350, 400, 1, "[(20,1)(60,0)(100,3)(120,5)(250,2)(300,0)(350,1)(400,0)]"}, // No overlap, right
	} {
		s.Set(tc.from, tc.end, tc.intensity)
		got := s.String()
		if got != tc.expected {
			t.Fatalf("Case[%v] fail: got(%v)!=expected(%v)\n", i, got, tc.expected)
		}
	}
}
