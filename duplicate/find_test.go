package duplicate

import "testing"

func TestFind(t *testing.T) {
	arr := [][]string{
		{"1014065", "1", "4708", "4709"},
		{"1014071", "1", "4708", "4697"},
		{"1014130", "2", "9416", "8004"},
		{"1014130", "2", "9416", "1412"},
		{"1014238", "1", "4708", "4694"},
		{"1014240", "3", "14124", "14090"},
	}
	got := Find(arr)
	want := []Duplicate{
		{row: 0, column: 1},
		{row: 0, column: 2},
		{row: 1, column: 1},
		{row: 1, column: 2},
		{row: 2, column: 0},
		{row: 2, column: 1},
		{row: 2, column: 2},
		{row: 3, column: 0},
		{row: 3, column: 1},
		{row: 3, column: 2},
		{row: 4, column: 1},
		{row: 4, column: 2},
	}

	if len(got) != len(want) {
		t.Errorf("got count duplicate found %d, wanted count duplicate found %d", len(got), len(want))
	} else {
		for i := 0; i < len(got); i++ {
			if got[i].row != want[i].row {
				t.Errorf("got row %d, wanted row %d on index %d", got[i].row, want[i].row, i)
			}
			if got[i].column != want[i].column {
				t.Errorf("got column %d, wanted column %d on index %d", got[i].column, want[i].column, i)
			}
		}
	}
}
