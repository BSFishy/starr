package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/BSFishy/starr"
)

const bpExclusions = APIver + "/exclusions"

// Exclusion is a Radarr excluded item.
type Exclusion struct {
	TMDBID int64  `json:"tmdbId"`
	Title  string `json:"movieTitle"`
	Year   int    `json:"movieYear"`
	ID     int64  `json:"id,omitempty"`
}

// GetExclusions returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusions() ([]*Exclusion, error) {
	return r.GetExclusionsContext(context.Background())
}

// GetExclusionsContext returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusionsContext(ctx context.Context) ([]*Exclusion, error) {
	var output []*Exclusion

	req := starr.Request{URI: bpExclusions}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateExclusion changes an exclusions in Radarr.
func (r *Radarr) UpdateExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return r.UpdateExclusionContext(context.Background(), exclusion)
}

// UpdateExclusionContext changes an exclusions in Radarr.
func (r *Radarr) UpdateExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions, starr.Str(exclusion.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteExclusions removes exclusions from Radarr.
func (r *Radarr) DeleteExclusions(ids []int64) error {
	return r.DeleteExclusionsContext(context.Background(), ids)
}

// DeleteExclusionsContext removes exclusions from Radarr.
func (r *Radarr) DeleteExclusionsContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		req := starr.Request{URI: path.Join(bpExclusions, starr.Str(id))}
		if err := r.DeleteAny(ctx, req); err != nil {
			errs += fmt.Sprintf("api.Post(%s): %v ", &req, err)
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// AddExclusions adds multiple exclusions to Radarr.
func (r *Radarr) AddExclusions(exclusions []*Exclusion) error {
	return r.AddExclusionsContext(context.Background(), exclusions)
}

// AddExclusionsContext adds exclusions to Radarr.
func (r *Radarr) AddExclusionsContext(ctx context.Context, exclusions []*Exclusion) error {
	for i := range exclusions {
		exclusions[i].ID = 0
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusions); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output interface{} // any ok

	req := starr.Request{URI: path.Join(bpExclusions, "bulk"), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// AddExclusion adds one exclusion to Radarr.
func (r *Radarr) AddExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return r.AddExclusionContext(context.Background(), exclusion)
}

// AddExclusionContext adds one exclusion to Radarr.
func (r *Radarr) AddExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	exclusion.ID = 0 // if this panics, fix your code

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}
