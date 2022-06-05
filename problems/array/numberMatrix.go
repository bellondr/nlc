package main

type NumMatrix struct {
	innerMat [][]int
}


func ConstructorN(matrix [][]int) NumMatrix {
	innerMat := make([][]int, len(matrix)+1)
	for i := range innerMat {
		innerMat[i] = make([]int, len(matrix[0]) + 1)
	}

	for i, list := range matrix {
		for j, v := range list {
			innerMat[i+1][j+1] = innerMat[i][j+1] + innerMat[i+1][j] - innerMat[i][j] + v
		}
	}
	return NumMatrix{
		innerMat: innerMat,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.innerMat[row2+1][col2+1] + this.innerMat[row1][col1] - this.innerMat[row1][col2+1] - this.innerMat[row2+1][col1]
}

