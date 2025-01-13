package main

import (
    "github.com/container-storage-interface/spec/lib/go/csi"
    "google.golang.org/grpc"
    "log"
    "net"
    "context"
	"github.com/hasirciogli/multicloud-csi-driver/utils"
)

func main() {
	utils.LogError(nil)
}

// type Driver struct {
//     csi.UnimplementedControllerServer
//     csi.UnimplementedNodeServer
//     // Other fields...
// }

// func (d *Driver) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
// 	var nodeId string := req.GetNodeID()
//     return &csi.NodeGetInfoResponse{
//         NodeId: nodeId,
//     }, nil
// }

// func main() {
//     listener, err := net.Listen("tcp", ":50051")
//     if err != nil {
//         log.Fatalf("Failed to listen: %v", err)
//     }

//     s := grpc.NewServer()
//     csi.RegisterControllerServer(s, &Driver{})
//     csi.RegisterNodeServer(s, &Driver{})

//     log.Println("Starting CSI driver server...")
//     if err := s.Serve(listener); err != nil {
//         log.Fatalf("Failed to serve: %v", err)
//     }
// } 