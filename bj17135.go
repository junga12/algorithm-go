package main

import (
	"fmt"
	"sort"
)

func Q17135() {
	var n, m, d int
	fmt.Scan(&n, &m, &d)

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&board[i][j])
		}
	}

	q := q17135{
		n:     n,
		m:     m,
		d:     d,
		board: board,
	}

	q.setFinalRound()

	q.play()

	fmt.Println(q.result)
}

type q17135 struct {
	n     int
	m     int
	d     int
	board [][]int

	finalRound int
	result     int // 제거할 수 있는 최대의 적
}

func (q *q17135) setFinalRound() {
	for r := 0; r < q.n; r++ {
		for c := 0; c < q.m; c++ {

			if q.board[r][c] == 1 { // 적 존재
				q.finalRound = q.n - r
				return
			}
		}
	}
}

func (q *q17135) play() {
	archerPositions := Permutation(MakeNumArray(q.m, 0), 3) // 궁수의 위치

	for _, archerPosition := range archerPositions {
		q.playByArcherPosition(archerPosition)
	}
}

func (q *q17135) playByArcherPosition(archerPosition []int) {
	killedEnemy := 0

	board := DeepCopyTwoDimension(q.board)

	for round := 1; round <= q.finalRound; round++ {
		// 죽을 적 찾기
		target := q.findKillEnemy(board, archerPosition)
		fmt.Println(target)

		// 적 없애기 (1 -> 0)
		killedEnemy += len(target)
		for _, t := range target {
			board[t[0]][t[1]] = 0
		}

		// 적 이동
		board = append([][]int{make([]int, q.m)}, board[:len(board)-1]...)
	}

	q.setResult(killedEnemy)
}

func (q *q17135) findKillEnemy(board [][]int, archerPosition []int) [][2]int {
	var target [][2]int // 각 궁수가 공격하는 위치

	for _, archer := range archerPosition {

		canAttackPosition := [][2]int{{len(board), archer}} // 공격가능 위치

		for i := 0; i < q.d; i++ { // 거리 d
			var attacked [][2]int // i 거리에 위치한 적 리스트

			var nextAttackPosition [][2]int
			// i 거리에 위치한 적 구하기
			for _, beforeAttackPosition := range canAttackPosition {

				for _, cur := range findAdjacency(beforeAttackPosition[0], beforeAttackPosition[1]) {

					// board 안에 위치하지 않으면 pass
					if cur[0] < 0 || cur[1] < 0 || cur[0] >= q.n || cur[1] >= q.m {
						continue
					}

					switch board[cur[0]][cur[1]] {
					case 0: // 빈 칸
						nextAttackPosition = append(nextAttackPosition, cur)
					case 1: // 적 존재
						attacked = append(attacked, cur)
					}
				}
			}

			if len(attacked) == 0 {
				canAttackPosition = nextAttackPosition
				continue
			}

			// 가장 왼쪽에 존재하는 적 찾기
			sort.Slice(attacked, func(i, j int) bool {
				return attacked[i][1] < attacked[j][1]
			})

			target = append(target, attacked[0])
			break
		}
	}

	return RemoveDuplication(target)
}

func (q *q17135) findNearestEnemy() {

}

func (q *q17135) setResult(r int) {
	if r > q.result {
		q.result = r
	}
}
