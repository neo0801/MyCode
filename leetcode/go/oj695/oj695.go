package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

var directs = []inode{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func printGrid(grid [][]int, msg string) {
	fmt.Printf("++++++++++++++   %s   +++++++++++++++\n", msg)
	for _, line := range grid {
		fmt.Print("[")
		for _, num := range line {
			if num == -1 {
				fmt.Print("# ")
			} else {
				fmt.Print(num, " ")
			}
		}
		fmt.Println("]")
	}
	fmt.Println("++++++++++++++   end     +++++++++++++++")
}

func maxAreaOfIsland(grid [][]int) int {
	m, n, mm := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ni := bfs(grid, m, n, i, j)
				if ni > mm {
					mm = ni
				}
			}
		}
	}
	return mm
}

type inode struct {
	x, y int
}

func isInside(cur inode, m, n int) bool {
	if cur.x < 0 || cur.x >= m || cur.y < 0 || cur.y >= n {
		return false
	}
	return true
}

func bfs(grid [][]int, m, n, x, y int) int {
	printGrid(grid, "Before BFS")
	que := []inode{}
	que = append(que, inode{x: x, y: y})
	grid[x][y] = -1 // must marked as visited when append to queue, otherwise may be append twice
	var maxSize int = 0
	for len(que) > 0 {
		head := que[0]
		fmt.Println("Current queue is: ", que)
		maxSize++
		fmt.Println("MaxSize is", maxSize)
		que = que[1:] // pop left
		for _, direct := range directs {
			var curNode inode
			curNode.x = head.x + direct.x
			curNode.y = head.y + direct.y
			if isInside(curNode, m, n) && grid[curNode.x][curNode.y] == 1 {
				que = append(que, curNode)
				grid[curNode.x][curNode.y] = -1 // must marked as visited when append to queue, otherwise may be append twice
				fmt.Printf("Append node: x=%d, y=%d to queue with value: %d\n", curNode.x, curNode.y, grid[curNode.x][curNode.y])
			}
		}
	}
	printGrid(grid, "After BFS")
	return maxSize
}
