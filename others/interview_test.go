package others

import (
	"context"
	"testing"
	"time"

	"log"
)

func Test_foo(t *testing.T) {
	ctx := context.Background()
	log.Println("start")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	foo(ctx)
}
