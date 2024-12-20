package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/BSFishy/starr"
)

/* Custom Formats do not exist in Sonarr v3; this is v4 only. */

const bpCustomFormat = APIver + "/customFormat"

// CustomFormatInput is the input for a new or updated CustomFormat.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
type CustomFormatInput struct {
	ID                    int64                    `json:"id,omitempty"`
	Name                  string                   `json:"name"`
	IncludeCFWhenRenaming bool                     `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatInputSpec `json:"specifications"`
}

// CustomFormatInputSpec is part of a CustomFormatInput.
type CustomFormatInputSpec struct {
	Name           string              `json:"name"`
	Implementation string              `json:"implementation"`
	Negate         bool                `json:"negate"`
	Required       bool                `json:"required"`
	Fields         []*starr.FieldInput `json:"fields"`
}

// CustomFormatOutput is the output from the CustomFormat methods.
type CustomFormatOutput struct {
	ID                    int64                     `json:"id"`
	Name                  string                    `json:"name"`
	IncludeCFWhenRenaming bool                      `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatOutputSpec `json:"specifications"`
}

// CustomFormatOutputSpec is part of a CustomFormatOutput.
type CustomFormatOutputSpec struct {
	Name               string               `json:"name"`
	Implementation     string               `json:"implementation"`
	ImplementationName string               `json:"implementationName"`
	InfoLink           string               `json:"infoLink"`
	Negate             bool                 `json:"negate"`
	Required           bool                 `json:"required"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// GetCustomFormats returns all configured Custom Formats.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) GetCustomFormats() ([]*CustomFormatOutput, error) {
	return s.GetCustomFormatsContext(context.Background())
}

// GetCustomFormatsContext returns all configured Custom Formats.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) GetCustomFormatsContext(ctx context.Context) ([]*CustomFormatOutput, error) {
	var output []*CustomFormatOutput

	req := starr.Request{URI: bpCustomFormat}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFormat returns a single customformat.
func (s *Sonarr) GetCustomFormat(customformatID int64) (*CustomFormatOutput, error) {
	return s.GetCustomFormatContext(context.Background(), customformatID)
}

// GetCustomFormatContext returns a single customformat.
func (s *Sonarr) GetCustomFormatContext(ctx context.Context, customformatID int64) (*CustomFormatOutput, error) {
	var output CustomFormatOutput

	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(customformatID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) AddCustomFormat(format *CustomFormatInput) (*CustomFormatOutput, error) {
	return s.AddCustomFormatContext(context.Background(), format)
}

// AddCustomFormatContext creates a new custom format and returns the response (with ID).
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) AddCustomFormatContext(ctx context.Context, format *CustomFormatInput) (*CustomFormatOutput, error) {
	var output CustomFormatOutput

	if format == nil {
		return &output, nil
	}

	format.ID = 0 // ID must be zero when adding.

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	req := starr.Request{URI: bpCustomFormat, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) UpdateCustomFormat(format *CustomFormatInput) (*CustomFormatOutput, error) {
	return s.UpdateCustomFormatContext(context.Background(), format)
}

// UpdateCustomFormatContext updates an existing custom format and returns the response.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) UpdateCustomFormatContext(
	ctx context.Context,
	format *CustomFormatInput,
) (*CustomFormatOutput, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	var output CustomFormatOutput

	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(format.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFormat deletes a custom format.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) DeleteCustomFormat(formatID int64) error {
	return s.DeleteCustomFormatContext(context.Background(), formatID)
}

// DeleteCustomFormatContext deletes a custom format.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) DeleteCustomFormatContext(ctx context.Context, formatID int64) error {
	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(formatID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
