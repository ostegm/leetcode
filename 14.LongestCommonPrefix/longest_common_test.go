package main

import "testing"

func TestLongestCommonPrefixWithCommon(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	got := longestCommonPrefix(strs)
	want := "fl"
	if got != want {
		t.Errorf("Got %s, want %s.", got, want)
	}
}

func TestLongestCommonPrefixWhenAllSame(t *testing.T) {
	strs := []string{"flower", "flower", "flower"}
	got := longestCommonPrefix(strs)
	want := "flower"
	if got != want {
		t.Errorf("Got %s, want %s.", got, want)
	}
}

func TestLongestCommonPrefixWithNoCommon(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	got := longestCommonPrefix(strs)
	want := ""
	if got != want {
		t.Errorf("Got %s, want %s.", got, want)
	}
}

func TestLongestCommonPrefixWhenAllEmpty(t *testing.T) {
	strs := []string{"", "", ""}
	got := longestCommonPrefix(strs)
	want := ""
	if got != want {
		t.Errorf("Got %s, want %s.", got, want)
	}
}