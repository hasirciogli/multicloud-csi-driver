package controller

import (
	"context"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"log"
	"net"
	
)

type ControllerDriver struct {
	csi.UnimplementedControllerServer
}

// ControllerGetCapabilities handles requests for the capabilities of the controller
func (d *ControllerDriver) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	capabilities := []*csi.ControllerServiceCapability{
		{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
				},
			},
		},
	}

	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: capabilities,
	}, nil
}

// CreateVolume handles requests to create a volume
func (d *ControllerDriver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	fmt.Printf("Creating volume with name %s\n", req.GetName())
	
	volumeID := "unique-volume-id"
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId: volumeID,
			CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
		},
	}, nil
}

// DeleteVolume handles requests to delete a volume
func (d *ControllerDriver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	fmt.Printf("Deleting volume with ID %s\n", req.GetVolumeId())
	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerGetVolume handles requests to get details about a volume
func (d *ControllerDriver) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	// Example of getting volume details
	// You will integrate with your storage backend here to get the volume information

	volumeID := req.GetVolumeId()

	// Here you would query the volume details and return the response
	return &csi.ControllerGetVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      volumeID,
			CapacityBytes: 1024 * 1024 * 1024, // Example: 1GB volume
		},
	}, nil
}

// ControllerPublishVolume handles requests to publish a volume to a node
func (d *ControllerDriver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	// Example of publishing a volume to a node (i.e., mounting the volume)
	fmt.Printf("Publishing volume %s to node %s\n", req.GetVolumeId(), req.GetNodeId())

	// In a real implementation, you would instruct the node to mount the volume
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// ControllerUnpublishVolume handles requests to unpublish a volume from a node
func (d *ControllerDriver) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	// Example of unpublishing a volume (i.e., unmounting the volume)
	fmt.Printf("Unpublishing volume %s from node %s\n", req.GetVolumeId(), req.GetNodeId())

	// In a real implementation, you would instruct the node to unmount the volume
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// ControllerExpandVolume handles requests to expand a volume
func (d *ControllerDriver) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	// Example of expanding a volume
	// You would integrate with your storage backend here to expand the volume

	fmt.Printf("Expanding volume %s to new size %d bytes\n", req.GetVolumeId(), req.GetCapacityRange().GetRequiredBytes())

	// Returning the updated volume information
	return &csi.ControllerExpandVolumeResponse{
		CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
	}, nil
}

// CreateSnapshot handles requests to create a snapshot of a volume
func (d *ControllerDriver) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	fmt.Printf("Creating snapshot for volume %s\n", req.GetSourceVolumeId())
	
	snapshotID := "unique-snapshot-id"
	return &csi.CreateSnapshotResponse{
		Snapshot: &csi.Snapshot{
			SnapshotId: snapshotID,
			SourceVolumeId: req.GetSourceVolumeId(),
		},
	}, nil
}

// ListSnapshots handles requests to list snapshots
func (d *ControllerDriver) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	entries := []*csi.ListSnapshotsResponse_Entry{
		{
			Snapshot: &csi.Snapshot{
				SnapshotId:     "snapshot-id-1",
				SourceVolumeId: "volume-id-1",
			},
		},
	}
	
	return &csi.ListSnapshotsResponse{
		Entries: entries,
	}, nil
}

// GetPluginInfo handles requests to get plugin information
func (d *ControllerDriver) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name: "custom-csi-driver",
		VendorVersion: "1.0.0",
	}, nil
}

// StartControllerServer Controller Server'ı başlatır
func StartControllerServer() error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	csi.RegisterControllerServer(s, &ControllerDriver{})

	log.Println("Starting CSI controller server on port 50051...")
	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}