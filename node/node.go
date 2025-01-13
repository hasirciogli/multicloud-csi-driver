package node

import (
	"context"
	"log"
	"net"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type NodeDriver struct {
	csi.UnimplementedNodeServer
}

// NodeStageVolume prepares the volume for mounting by the node
func (d *NodeDriver) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	// Implementation to stage the volume
	return &csi.NodeStageVolumeResponse{}, nil
}

// NodeUnstageVolume unstage the volume from the node
func (d *NodeDriver) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	// Implementation to unstage the volume
	return &csi.NodeUnstageVolumeResponse{}, nil
}

// NodePublishVolume mounts the volume to the pod's mount point
func (d *NodeDriver) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// Implementation to mount the volume to the pod
	return &csi.NodePublishVolumeResponse{}, nil
}

// NodeUnpublishVolume unmounts the volume from the pod's mount point
func (d *NodeDriver) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	// Implementation to unmount the volume from the pod
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// NodeGetVolumeStats gets statistics for a volume
func (d *NodeDriver) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	// Implementation to get volume statistics
	return &csi.NodeGetVolumeStatsResponse{}, nil
}

// NodeExpandVolume expands the size of the volume
func (d *NodeDriver) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	// Implementation to expand the volume
	return &csi.NodeExpandVolumeResponse{}, nil
}

// NodeGetInfo gets the node info
func (d *NodeDriver) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	// Implementation to get the node info
	return &csi.NodeGetInfoResponse{
		NodeId:            "node-id-1-test",
		MaxVolumesPerNode: 10,
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				"region": "us-east-1",
				"zone":   "us-east-1a",
			},
		},
	}, nil
}

// get capabilities function
func (d *NodeDriver) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	// Implementation to get the capabilities
	return &csi.NodeGetCapabilitiesResponse{}, nil
}

// StartNodeServer Node Server'ı başlatır
func StartNodeServer() error {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	csi.RegisterNodeServer(s, &NodeDriver{})

	log.Println("Starting CSI node server on port 50052...")
	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}
