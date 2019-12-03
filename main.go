package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nametake/grpc-with-http/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	// lis, err := net.Listen("tcp", ":9998")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v\n", err)
	// }
	// defer lis.Close()
	//
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to server: %v\n", err)
	// }

	// serverPair, err := tls.LoadX509KeyPair("./server.crt", "./server.key")
	// if err != nil {
	// 	log.Fatalf("failed server pair")
	// }
	//
	// serverCert := credentials.NewTLS(&tls.Config{
	// 	Certificates:       []tls.Certificate{serverPair},
	// 	InsecureSkipVerify: true,
	// })

	s := grpc.NewServer()
	// s := grpc.NewServer(grpc.Creds(serverCert))

	pb.RegisterPingAPIServer(s, &PingAPIServer{})

	b, err := ioutil.ReadFile("./server.crt")
	if err != nil {
		log.Fatalf("faltal load ioutil: %v", err)
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatalf("credentials: failed to append certificates")
	}

	// clientCert := credentials.NewTLS(&tls.Config{
	// 	RootCAs:            cp,
	// 	InsecureSkipVerify: true,
	// })

	conn, err := grpc.DialContext(ctx, "127.0.0.1:9998", grpc.WithInsecure())
	// conn, err := grpc.DialContext(ctx, "127.0.0.1:9998", grpc.WithTransportCredentials(clientCert))
	if err != nil {
		log.Fatalf("failed to create conn")
	}

	mux := runtime.NewServeMux()

	if err := pb.RegisterPingAPIHandler(ctx, mux, conn); err != nil {
		log.Fatalf("failed to register ping api hadler")
	}

	httpServer := &http.Server{
		Addr:    ":9998",
		Handler: httpGrpcRouter(s, mux),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// if err := httpServer.ListenAndServe(); err != nil {
	// 	log.Fatalf("failed to listen and serve")
	// }

	if err := httpServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("failed to listen and serve")
	}

	// if err := http.ListenAndServeTLS(":9998", "server.crt", "server.key", httpGrpcRouter(s, mux)); err != nil {
	// 	log.Fatalf("failed to listen and serve")
	// }
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "pong")
	// })
	// return mux
	return grpcServer
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
	// 		grpcServer.ServeHTTP(w, r)
	// 	} else {
	// 		httpHandler.ServeHTTP(w, r)
	// 	}
	// })
}

type PingAPIServer struct{}

func (p *PingAPIServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Msg: "pong",
	}, nil
}
