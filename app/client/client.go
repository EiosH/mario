package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "app/proto"
)

const PORT = "9001"

func Client() {
	grpcConn, err := grpc.Dial(`127.0.0.1:`+PORT, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Dial err", err)

		return
	}

	defer grpcConn.Close()

	fmt.Println("Dial !")

	client := pb.NewSearchServiceClient(grpcConn)

	t, err := client.Search(context.TODO(), &pb.SearchRequest{Request: "你好"})

	if err != nil {
		fmt.Println("Dial err", err)
		return

	}

	fmt.Println(t.Response)

}
