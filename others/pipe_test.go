package others

import (
	"fmt"
	"testing"
)

func Test_pipe(t *testing.T) {

	//var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for n := range sum(sq(odd(echo(nums)))) {
	//	fmt.Println(n)
	//}

	//var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for n := range pipeline(nums, echo, odd, sq, sum) {
	//	fmt.Println(n)
	//}

	nums := makeRange(1, 10000)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int

	for i := range chans {
		chans[i] = sum(prime(in)) // return quickly
	}

	for n := range sum(merge(chans[:])) { // will block
		fmt.Println(n)
	}
}
