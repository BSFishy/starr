package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/BSFishy/starr"
)

// Define Base Path for download client calls.
const bpDownloadClient = APIver + "/downloadClient"

// DownloadClientInput is the input for a new or updated download client.
type DownloadClientInput struct {
	Enable             bool                `json:"enable"`
	Priority           int                 `json:"priority"`
	ID                 int64               `json:"id,omitempty"`
	ConfigContract     string              `json:"configContract"`
	Implementation     string              `json:"implementation"`
	ImplementationName string              `json:"implementationName"`
	Name               string              `json:"name"`
	Protocol           starr.Protocol      `json:"protocol"`
	Tags               []int               `json:"tags"`
	Fields             []*starr.FieldInput `json:"fields"`
}

// DownloadClientOutput is the output from the download client methods.
type DownloadClientOutput struct {
	Enable             bool                 `json:"enable"`
	Priority           int                  `json:"priority"`
	ID                 int64                `json:"id,omitempty"`
	ConfigContract     string               `json:"configContract"`
	Implementation     string               `json:"implementation"`
	ImplementationName string               `json:"implementationName"`
	InfoLink           string               `json:"infoLink"`
	Name               string               `json:"name"`
	Protocol           starr.Protocol       `json:"protocol"`
	Tags               []int                `json:"tags"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// GetDownloadClients returns all configured download clients.
func (r *Readarr) GetDownloadClients() ([]*DownloadClientOutput, error) {
	return r.GetDownloadClientsContext(context.Background())
}

// GetDownloadClientsContext returns all configured download clients.
func (r *Readarr) GetDownloadClientsContext(ctx context.Context) ([]*DownloadClientOutput, error) {
	var output []*DownloadClientOutput

	req := starr.Request{URI: bpDownloadClient}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetDownloadClient returns a single download client.
func (r *Readarr) GetDownloadClient(downloadclientID int64) (*DownloadClientOutput, error) {
	return r.GetDownloadClientContext(context.Background(), downloadclientID)
}

// GetDownloadClientContext returns a single download client.
func (r *Readarr) GetDownloadClientContext(ctx context.Context, downloadclientID int64) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	req := starr.Request{URI: path.Join(bpDownloadClient, starr.Str(downloadclientID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDownloadClient creates a download client without testing it.
func (r *Readarr) AddDownloadClient(downloadclient *DownloadClientInput) (*DownloadClientOutput, error) {
	return r.AddDownloadClientContext(context.Background(), downloadclient)
}

// AddDownloadClientContext creates a download client without testing it.
func (r *Readarr) AddDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
) (*DownloadClientOutput, error) {
	var (
		output DownloadClientOutput
		body   bytes.Buffer
	)

	client.ID = 0
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: bpDownloadClient, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// TestDownloadClient tests a download client.
func (r *Readarr) TestDownloadClient(client *DownloadClientInput) error {
	return r.TestDownloadClientContext(context.Background(), client)
}

// TestDownloadClientContext tests a download client.
func (r *Readarr) TestDownloadClientContext(ctx context.Context, client *DownloadClientInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: path.Join(bpDownloadClient, "test"), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// UpdateDownloadClient updates the download client.
func (r *Readarr) UpdateDownloadClient(downloadclient *DownloadClientInput, force bool) (*DownloadClientOutput, error) {
	return r.UpdateDownloadClientContext(context.Background(), downloadclient, force)
}

// UpdateDownloadClientContext updates the download client.
func (r *Readarr) UpdateDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
	force bool,
) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{
		URI:   path.Join(bpDownloadClient, starr.Str(client.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(force)}},
	}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDownloadClient removes a single download client.
func (r *Readarr) DeleteDownloadClient(downloadclientID int64) error {
	return r.DeleteDownloadClientContext(context.Background(), downloadclientID)
}

// DeleteDownloadClientContext removes a single download client.
func (r *Readarr) DeleteDownloadClientContext(ctx context.Context, downloadclientID int64) error {
	req := starr.Request{URI: path.Join(bpDownloadClient, starr.Str(downloadclientID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
