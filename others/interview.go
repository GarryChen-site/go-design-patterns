package others

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// only edit foo function
// foo must call slow function
// foo must return when the ctx is timeout
// if slow finished before ctx timeout, foo must return too

func foo(ctx context.Context) {
	// slow();

	doneCh := make(chan struct{})

	go func() {
		slow()
		close(doneCh)
	}()

	select {
	case <-ctx.Done():
		log.Println("foo: parent ctx done, reason:", ctx.Err())
		return
	case <-doneCh:
		log.Println("slow finished before ctx timeout")
		return
	}
}

func slow() {
	n := rand.Intn(7)
	log.Printf("sleeping for %d seconds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
