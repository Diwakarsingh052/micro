package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 13, 6, 5}
	BubbleSort(els)

	assert.EqualValues(t, 5, len(els))

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 8, els[2])
	assert.EqualValues(t, 9, els[3])
	assert.EqualValues(t, 13, els[4])

}
func TestBubbleSortBestCase(t *testing.T) {
	els := []int{5, 6, 8, 9, 13}
	BubbleSort(els)

	assert.EqualValues(t, 5, len(els))

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 8, els[2])
	assert.EqualValues(t, 9, els[3])
	assert.EqualValues(t, 13, els[4])

}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)

}

func getElements(n int) []int {
	result := make([]int, n)
	i:=0
	for j:=n-1;j>=0;j-- {
		result[i]=j
		i++
	}
	return result
}
func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.EqualValues(t, 5, len(els))


	assert.NotNil(t, els)
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	// we are benchmarking for worst case here as whole array is not sorted
	//10,9,8,7,6,5,4,3,2,1
	// see video once for more example
	els:= getElements(10)
	for i:=0; i<b.N;i++ {
		BubbleSort(els)
	}
}
