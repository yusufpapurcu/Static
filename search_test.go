package main

import (
	"testing"
)

// TestSearch function
func TestFileSearch(t *testing.T) {
	result := Search("./example", "example.docx", SEARCH_FILE)
	expected := "example/page/test/justtest/as/example.docx"
	if result[0] != expected {
		t.Errorf("Expected : %v  Got : %v", expected, result[0])
	}
}
func TestKeywordSearch(t *testing.T) {
	result := Search("./example", "GOPHER IS HERE", SEARCH_KEYWORD)
	expected := "Found in line 41 in example/page/test/justtest/as/example.docx : Vivamus ut ex quis justo egestas vehicula ut GOPHER IS HERE feugiat quam. Proin quis sapien in sem sodales placerat et vel ligula. Mauris malesuada sapien eu felis egestas ultricies. Nam quis nulla lorem. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nulla rhoncus, turpis quis ultrices dictum, elit ligula accumsan risus, quis condimentum arcu nisi sit amet nulla. Proin sed dui congue, interdum sapien non, convallis augue."
	if result[0] != expected {
		t.Errorf("Expected : %v  Got : %v", expected, result[0])
	}
}
