package A

import "testing"

func Test_main(t *testing.T) {
	a := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		a = append(a, i)
	}
	getDistances(a, 1000000)
}
