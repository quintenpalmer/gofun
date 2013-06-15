package util

type PTree []Row

type Row []int

func Pascal(size int) PTree {
	tree := make(PTree,size)
	for i := 0; i < size; i++ {
		tree[i] = make(Row,i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				tree[i][j] = 1
			} else {
				tree[i][j] = tree[i-1][j-1] + tree[i-1][j]
			}
		}
	}
	return tree
}
