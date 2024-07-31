package leetcodego

import "testing"

func Test_minHeightShelves(t *testing.T) {
	expected := 6
	books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1},{1, 1},{1, 1},{1, 2}}
	shelfWidth := 4
	got := minHeightShelves(books, shelfWidth)

	if got != expected{
		t.Errorf("expected %d, got %d", expected, got)
	}
}
