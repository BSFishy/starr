package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/BSFishy/starr"
)

const bpApplications = APIver + "/applications"

type ApplicationInput struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	Fields             []*starr.FieldInput `json:"fields"`
	ImplementationName string              `json:"implementationName"`
	Implementation     string              `json:"implementation"`
	ConfigContract     string              `json:"configContract"`
	InfoLink           string              `json:"infoLink"`
	Tags               []int               `json:"tags"`
}

type ApplicationOutput struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	Fields             []*starr.FieldInput `json:"fields"`
	ImplementationName string              `json:"implementationName"`
	Implementation     string              `json:"implementation"`
	ConfigContract     string              `json:"configContract"`
	InfoLink           string              `json:"infoLink"`
	Tags               []int               `json:"tags"`
}

func (p *Prowlarr) GetApplications() ([]*ApplicationOutput, error) {
	return p.GetApplicationsContext(context.Background())
}

func (p *Prowlarr) GetApplicationsContext(ctx context.Context) ([]*ApplicationOutput, error) {
	var output []*ApplicationOutput

	req := starr.Request{URI: bpApplications}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

func (p *Prowlarr) GetApplication(applicationID int64) (*ApplicationOutput, error) {
	return p.GetApplicationContext(context.Background(), applicationID)
}

func (p *Prowlarr) GetApplicationContext(ctx context.Context, clientID int64) (*ApplicationOutput, error) {
	var output ApplicationOutput

	req := starr.Request{URI: path.Join(bpApplications, starr.Str(clientID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDownloadClient creates a download client without testing it.
func (p *Prowlarr) AddApplication(application *ApplicationInput) (*ApplicationOutput, error) {
	return p.AddApplicationContext(context.Background(), application)
}

// AddDownloadClientContext creates a download client without testing it.
func (p *Prowlarr) AddApplicationContext(ctx context.Context,
	client *ApplicationInput,
) (*ApplicationOutput, error) {
	var (
		output ApplicationOutput
		body   bytes.Buffer
	)

	client.ID = 0
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpApplications, err)
	}

	req := starr.Request{URI: bpApplications, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

func (p *Prowlarr) TestApplication(client *ApplicationInput) error {
	return p.TestApplicationContext(context.Background(), client)
}

func (p *Prowlarr) TestApplicationContext(ctx context.Context, client *ApplicationInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpApplications, err)
	}

	req := starr.Request{URI: path.Join(bpApplications, "test"), Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

func (p *Prowlarr) UpdateApplication(client *ApplicationInput, force bool) (*ApplicationOutput, error) {
	return p.UpdateApplicationContext(context.Background(), client, force)
}

func (p *Prowlarr) UpdateApplicationContext(ctx context.Context,
	client *ApplicationInput,
	force bool,
) (*ApplicationOutput, error) {
	var output ApplicationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpApplications, err)
	}

	req := starr.Request{
		URI:   path.Join(bpApplications, starr.Str(client.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(force)}},
	}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

func (p *Prowlarr) DeleteApplication(applicationID int64) error {
	return p.DeleteApplicationContext(context.Background(), applicationID)
}

func (p *Prowlarr) DeleteApplicationContext(ctx context.Context, applicationID int64) error {
	req := starr.Request{URI: path.Join(bpApplications, starr.Str(applicationID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
