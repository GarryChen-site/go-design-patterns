package others

import (
	"net/http"
	"testing"
)

func Test_decorator(t *testing.T) {

	// decorator(Hello)("Hello, World!")

	// sum1 := timedSumFunc(Sum1)
	// sum2 := timedSumFunc(Sum2)

	// fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))

	http.HandleFunc("/v1/hello", WithServerHeader(hello))
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServer: ", err)
	// }

	handler := Handler(hello, WithBasicAuth, WithAuthCookie, WithDebugLog)
	http.HandleFunc("/v4/hello", handler)
}
