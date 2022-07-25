package others

import (
	"crypto/tls"
	"time"
)

// type Server struct {
// 	Addr     string
// 	Port     int
// 	Protocol string
// 	Timeout  time.Duration
// 	MaxConns int
// 	TLS      *tls.Config
// }

// 因为go不支持函数重载

// func NewDefaultServer(addr string, port int) (*Server, error) {
// 	return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil

// }

// func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
// 	return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil

// }

// func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
// 	return &Server{addr, port, "tcp", timeout, 100, nil}, nil

// }

// func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
// 	return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil

// }

// ****************

// type Config struct {
// 	Protocol string
// 	Timeout  time.Duration
// 	Maxconns int
// 	TLS      *tls.Config
// }

// type Server struct {
// 	Addr string
// 	Port int
// 	Conf *Config
// }

// func NewServer(addr string, port int, conf *Config) (*Server, error) {
// 	//...

// }

// func main() {
//     //Using the default configuratrion
//     srv1, _ := NewServer("localhost", 9000, nil)

//     conf := Config{Protocol:"tcp", Timeout: 60}
//     srv2, _ := NewServer("locahost", 9000, &conf)
// }

// ***********************

// type Server struct {
// 	Addr     string
// 	Port     int
// 	Protocol string
// 	Timeout  time.Duration
// 	MaxConns int
// 	TLS      *tls.Config
// }

//使用一个builder类来做包装
// type ServerBuilder struct {
// 	Server
//       }

//       func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
// 	sb.Server.Addr = addr
// 	sb.Server.Port = port
// 	//其它代码设置其它成员的默认值
// 	return sb
//       }

//       func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
// 	sb.Server.Protocol = protocol
// 	return sb
//       }

//       func (sb *ServerBuilder) WithMaxConn( maxconn int) *ServerBuilder {
// 	sb.Server.MaxConns = maxconn
// 	return sb
//       }

//       func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
// 	sb.Server.Timeout = timeout
// 	return sb
//       }

//       func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
// 	sb.Server.TLS = tls
// 	return sb
//       }

//       func (sb *ServerBuilder) Build() (Server) {
// 	return  sb.Server
//       }

// ***************************

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxcoons int) Option {
	return func(s *Server) {
		s.MaxConns = maxcoons
	}
}

func NewServer(add string, port int, options ...func(*Server)) (*Server, error) {

	srv := Server{
		Addr:     add,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 100,
		TLS:      nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

// s1, _ := NewServer("localhost", 1024)
// s2, _ := NewServer("localhost", 2048, Protocol("udp"))
// s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
