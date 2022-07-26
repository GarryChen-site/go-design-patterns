package singleton

import (
	"fmt"
	"testing"
)

// a creational design pattern, which ensures that only one object
// of its kind exists and provides a single point of access to it for any other code.
func Test_single(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
