package csiutils

import (
	"errors"

	"github.com/codenrhoden/csi-nfs-plugin/csi"
)

// ErrMissingCSIEndpoint occurs when the value for the environment
// variable CSI_ENDPOINT is not set.
var ErrMissingCSIEndpoint = errors.New("missing CSI_ENDPOINT")

// ErrInvalidCSIEndpoint occurs when the value for the environment
// variable CSI_ENDPOINT is an invalid network address.
var ErrInvalidCSIEndpoint = errors.New("invalid CSI_ENDPOINT")

// ErrCreateVolume returns a CreateVolumeResponse with a CreateVolumeError.
func ErrCreateVolume(
	code csi.Error_CreateVolumeError_CreateVolumeErrorCode,
	msg string) *csi.CreateVolumeResponse {

	return &csi.CreateVolumeResponse{
		Reply: &csi.CreateVolumeResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_CreateVolumeError_{
					CreateVolumeError: &csi.Error_CreateVolumeError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}

// ErrDeleteVolume returns a DeleteVolumeResponse with a DeleteVolumeError.
func ErrDeleteVolume(
	code csi.Error_DeleteVolumeError_DeleteVolumeErrorCode,
	msg string) *csi.DeleteVolumeResponse {

	return &csi.DeleteVolumeResponse{
		Reply: &csi.DeleteVolumeResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_DeleteVolumeError_{
					DeleteVolumeError: &csi.Error_DeleteVolumeError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}

// ErrControllerPublishVolume returns a
// ControllerPublishVolumeResponse with a
// ControllerPublishVolumeVolumeError.
func ErrControllerPublishVolume(
	code csi.Error_ControllerPublishVolumeError_ControllerPublishVolumeErrorCode,
	msg string) *csi.ControllerPublishVolumeResponse {

	return &csi.ControllerPublishVolumeResponse{
		Reply: &csi.ControllerPublishVolumeResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_ControllerPublishVolumeError_{
					ControllerPublishVolumeError: &csi.Error_ControllerPublishVolumeError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}

// ErrControllerUnpublishVolume returns a
// ControllerUnpublishVolumeResponse with a
// ControllerUnpublishVolumeVolumeError.
func ErrControllerUnpublishVolume(
	code csi.Error_ControllerUnpublishVolumeError_ControllerUnpublishVolumeErrorCode,
	msg string) *csi.ControllerUnpublishVolumeResponse {

	return &csi.ControllerUnpublishVolumeResponse{
		Reply: &csi.ControllerUnpublishVolumeResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_ControllerUnpublishVolumeError_{
					ControllerUnpublishVolumeError: &csi.Error_ControllerUnpublishVolumeError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}

// ErrGetCapacity returns a
// GetCapacityResponse with a
// GeneralError.
func ErrGetCapacity(
	code csi.Error_GeneralError_GeneralErrorCode,
	msg string) *csi.GetCapacityResponse {

	return &csi.GetCapacityResponse{
		Reply: &csi.GetCapacityResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_GeneralError_{
					GeneralError: &csi.Error_GeneralError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}

// ErrListVolumes returns a
// ListVolumesResponse with a
// GeneralError.
func ErrListVolumes(
	code csi.Error_GeneralError_GeneralErrorCode,
	msg string) *csi.ListVolumesResponse {

	return &csi.ListVolumesResponse{
		Reply: &csi.ListVolumesResponse_Error{
			Error: &csi.Error{
				Value: &csi.Error_GeneralError_{
					GeneralError: &csi.Error_GeneralError{
						ErrorCode:        code,
						ErrorDescription: msg,
					},
				},
			},
		},
	}
}