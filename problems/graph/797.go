package graph

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	curPath := []int{}
	traverse(graph, &curPath, 0, &res)
	return res
}


func traverse(graph [][]int, curPath *[]int, curNode int, res *[][]int){

	*curPath = append(*curPath, curNode)
	if curNode == len(graph) - 1{
		tmp := make([]int, len(*curPath))
		copy(tmp, *curPath)
		*res = append(*res, tmp)
	}else{
		curNeighbor := graph[curNode]
		for  i := 0; i < len(curNeighbor); i ++{
			traverse(graph, curPath, curNeighbor[i], res)
		}
	}
	*curPath = (*curPath)[:len(*curPath) - 1]
}

