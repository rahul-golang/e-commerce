package main

import (
	"context"
	"fmt"
	"gokit/ecommerse/users/pkg/grpc/pb"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"

	"google.golang.org/grpc"
)

var zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"
var hostPort = "0.0.0.0:0"

func main() {

	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, false, hostPort, " Users client")

	// Create our tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)

	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	// Explicitly set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)

	conn, err := grpc.Dial("localhost:8801", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error in Dial(): gRPC connection : ", err)
	}

	client := pb.NewUsersClient(conn)

	span := opentracing.StartSpan("Run")

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	resp, err := client.GetUser(ctx, &pb.GetUserRequest{
		ID: "1",
	})

	if err != nil {
		fmt.Println("Error in Calling server Method : ", err)
	}
	fmt.Println("Responce from server ", resp)

	resp1, err := client.GetAllUser(ctx, &pb.GetAllUserRequest{})

	if err != nil {
		fmt.Println("Error in Calling server Method : ", err)
	}
	fmt.Println("Responce from server ", resp1)

	// Finish our CLI span
	span.Finish()

	// Close collector to ensure spans are sent before exiting.
	collector.Close()

}
