package main

import (
	"fmt"
)

func Q9465() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		sticker := make([][]int, 2)
		for i := 0; i < 2; i++ {
			sticker[i] = make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scan(&sticker[i][j])
			}
		}

		visited := make([][]bool, 2)
		for j := 0; j < n; j++ {
			visited[j] = make([]bool, n)
		}

		q := q9465{
			sticker:     sticker,
			row:         2,
			col:         n,
			maxScore:    0,
			curStickers: make([][]int, 0),
			visited:     visited,
		}
		q.play()

		fmt.Println(q.maxScore)
	}
}

type q9465 struct {
	sticker  [][]int
	row, col int

	maxScore int

	curStickers [][]int
	visited     [][]bool
}

func (q *q9465) play() {
	fmt.Println("sticker: ", q.sticker)

	q.pickOne(0, 0)

	fmt.Println("max score: ", q.maxScore)
}

func (q *q9465) pickOne(r, c int) {
	if q.isPossibleVisit(r, c) {
		q.visited[r][c] = true

		q.curStickers = append(q.curStickers)

		q.visited[r][c] = false
	}
}

func (q *q9465) isPossibleVisit(r, c int) bool {
	//for _, adj := range findAdjacency(r, c) {
	//}
	return true
}

// 범위 안에 존재 & 스티커 선택 가능
func (q *q9465) isLocatedIn(r, c int) bool {
	return r >= 0 && c >= 0 && r < q.row && c < q.col
}

// 다음 위치 (오른쪽으로 이동, 오른쪽 끝이면 아래로 이동)
func (q *q9465) next(r, c int) (int, int, bool) {
	c += 1

	if c == q.col {
		r += 1
		c = 0
	}

	if r == q.row {
		return 0, 0, false
	}

	return r, c, true
}
