// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package widget

import (
	"time"
)

type GetTokenRequest struct {
	GrantType *string `json:"grant_type,omitempty"`
}

type GetTokenResult struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

type Tags = map[string]string

// Widget a widget.
type Widget struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`

	// Parent the widget this one was cloned from. Omitted when it was created
	// directly.
	Parent      *Widget     `json:"parent,omitempty"`
	PrimaryIPv6 string      `json:"primary_ipv6,omitempty"`
	Tags        Tags        `json:"tags,omitempty"`
	WidgetState WidgetState `json:"widget_state,omitempty"`
}

type WidgetCreateRequest struct {
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`

	// Required.
	Name   string `json:"name"`
	SizeGB *int   `json:"size_gb,omitempty"`
	Tags   Tags   `json:"tags,omitempty"`
}

// WidgetState lifecycle state of a widget.
type WidgetState string

// Values WidgetState accepts.
const (
	WidgetStatePending WidgetState = "pending"
	WidgetStateActive  WidgetState = "active"
	WidgetStateRetired WidgetState = "retired"
)
