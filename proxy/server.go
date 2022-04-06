package proxy

// Subject


type server interface {

	handleRequest(string, string) (int, string)
}