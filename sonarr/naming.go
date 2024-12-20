package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/BSFishy/starr"
)

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

// CRF is ColonReplacementFormat, for naming config.
type CRF int

// These are all of the possible Colon Replacement Formats (for naming config) in Sonarr.
const (
	ColonDelete CRF = iota
	ColonReplaceWithDash
	ColonReplaceWithSpaceDash
	ColonReplaceWithSpaceDashSpace
	ColonSmartReplace
	Custom
)

// Naming represents the config/naming endpoint in Sonarr.
type Naming struct {
	RenameEpisodes               bool   `json:"renameEpisodes"`
	ReplaceIllegalCharacters     bool   `json:"replaceIllegalCharacters"`
	ColonReplacementFormat       CRF    `json:"colonReplacementFormat"`
	ID                           int64  `json:"id"`
	MultiEpisodeStyle            int64  `json:"multiEpisodeStyle"`
	DailyEpisodeFormat           string `json:"dailyEpisodeFormat"`
	AnimeEpisodeFormat           string `json:"animeEpisodeFormat"`
	SeriesFolderFormat           string `json:"seriesFolderFormat"`
	SeasonFolderFormat           string `json:"seasonFolderFormat"`
	SpecialsFolderFormat         string `json:"specialsFolderFormat"`
	StandardEpisodeFormat        string `json:"standardEpisodeFormat"`
	CustomColonReplacementFormat string `json:"customColonReplacementFormat"`
}

// GetNaming returns the naming.
func (s *Sonarr) GetNaming() (*Naming, error) {
	return s.GetNamingContext(context.Background())
}

// GetNamingContext returns the naming.
func (s *Sonarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output Naming

	req := starr.Request{URI: bpNaming}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNaming updates the naming.
func (s *Sonarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return s.UpdateNamingContext(context.Background(), naming)
}

// UpdateNamingContext updates the naming.
func (s *Sonarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
	var (
		output Naming
		body   bytes.Buffer
	)

	naming.ID = 1
	if err := json.NewEncoder(&body).Encode(naming); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNaming, err)
	}

	req := starr.Request{URI: bpNaming, Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
