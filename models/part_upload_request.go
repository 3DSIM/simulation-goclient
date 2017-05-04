package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PartUploadRequest part upload request
// swagger:model PartUploadRequest
type PartUploadRequest struct {

	// s3 key
	S3Key string `json:"s3Key,omitempty"`

	// signed Url
	SignedURL string `json:"signedUrl,omitempty"`
}

// Validate validates this part upload request
func (m *PartUploadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
