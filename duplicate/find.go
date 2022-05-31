package duplicate

import (
	"protek/tree"
	"protek/util"
)

type Duplicate struct {
	row    int
	column int
}

func Find(arr [][]string) []Duplicate {
	var duplicates []string
	bnrTrees := map[int]*tree.BnrTree{}
	for _, row := range arr {
		for j, value := range row {
			_, ok := bnrTrees[j]
			if !ok {
				bnrTrees[j] = &tree.BnrTree{Value: value}
			} else {
				isExist := bnrTrees[j].IsExist(value)
				if isExist {
					duplicates = append(duplicates, value)
				} else {
					bnrTrees[j].Insert(value)
				}
			}
		}
	}

	var result []Duplicate
	for i, row := range arr {
		for j, value := range row {
			if util.InArray(value, duplicates) {
				result = append(result, Duplicate{row: i, column: j})
			}
		}
	}

	return result
}
