package main

import "fmt"

// combination

//type Permutation struct {
//	array     []int
//	n, r, cnt int // n개 중 r개를 뽑음
//	visited   []bool
//	order     []int
//}
//
//func NewPermutation(array []int, n, r int) Permutation {
//	return Permutation{
//		array:   array,
//		n:       n,
//		r:       r,
//		cnt:     0,
//		visited: make([]bool, n),
//		order:   make([]int, n),
//	}
//}
//
//func (p *Permutation) Permu(cur int) {
//	if cur == p.r {
//		for i := 0; i < p.r; i++ {
//			fmt.Print(p.array[p.order[i]], " ")
//		}
//		fmt.Println()
//		p.cnt++
//		return
//	}
//
//	for i := 0; i < p.n; i++ {
//		if p.visited[i] {
//			continue
//		}
//
//		p.visited[i] = true
//		p.order[cur] = i
//		p.Permu(cur + 1)
//		p.visited[i] = false
//	}
//}

// startNum	부터 시작하는 size 크기의 리스트 반환
func MakeNumArray(size, startNum int) []int {
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = startNum + i
	}

	return result
}

// array값 중 r 개를 중복없이 선택한 경우의 수
func Permutation[T any](array []T, r int) [][]T {
	n := len(array)
	visited := make([]bool, n)
	order := make([]int, n)
	var result [][]T

	var run func(int)

	run = func(cur int) {
		if cur == r {
			resultOne := make([]T, r)

			for i := 0; i < r; i++ {
				resultOne[i] = array[order[i]]
			}

			result = append(result, resultOne)
			return
		}

		for i := cur; i < n; i++ {
			if visited[i] {
				continue
			}

			visited[i] = true
			order[cur] = i
			run(cur + 1)
			visited[i] = false
		}
	}

	run(0) // 실행

	return result
}

// 2차 배열 깊은 복사
func DeepCopyTwoDimension[T any](arr [][]T) [][]T {
	// 새로운 슬라이스를 만듭니다.
	newArr := make([][]T, len(arr))

	// 각각의 슬라이스를 새로운 슬라이스로 복사합니다.
	for i := range arr {
		newArr[i] = make([]T, len(arr[i]))
		copy(newArr[i], arr[i])
	}

	return newArr
}

// 리스트의 중복 제거
func RemoveDuplication[T any](arr []T) []T {
	var result []T

	unique := make(map[string]bool)
	for _, a := range arr {
		key := fmt.Sprintf("%v", a)

		if _, ok := unique[key]; !ok {
			unique[key] = true
			result = append(result, a)
		}
	}

	return result
}

// [a, b]의 인접 위치 반환 (상하좌우)
func findAdjacency(r, c int) [][2]int {
	xi := []int{1, 0, -1, 0}
	yi := []int{0, 1, 0, -1}

	var result [][2]int

	for i := 0; i < len(xi); i++ {
		result = append(result, [2]int{r + xi[i], c + yi[i]})
	}

	return result
}
