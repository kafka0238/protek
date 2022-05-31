package tree

type BnrTree struct {
	Value string
	Left  *BnrTree
	Right *BnrTree
}

func (tree *BnrTree) Insert(value string) {
	if tree.Value == value {
		return
	}

	if tree.Value > value {
		if tree.Left == nil {
			tree.Left = &BnrTree{Value: value}
			return
		}
		tree.Left.Insert(value)
	} else if tree.Value < value {
		if tree.Right == nil {
			tree.Right = &BnrTree{Value: value}
			return
		}
		tree.Right.Insert(value)
	} else {
		return
	}
}

func (tree *BnrTree) IsExist(value string) bool {
	if tree == nil {
		return false
	}

	if tree.Value == value {
		return true
	} else if tree.Value > value {
		return tree.Left.IsExist(value)
	} else {
		return tree.Right.IsExist(value)
	}
}
