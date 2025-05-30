// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/klucherev/logto-go-client/models/components"
)

type GetOneTimeTokenRequest struct {
	// The unique identifier of the one time token.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetOneTimeTokenRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetOneTimeTokenContext struct {
	JitOrganizationIds []string `json:"jitOrganizationIds,omitempty"`
}

func (o *GetOneTimeTokenContext) GetJitOrganizationIds() []string {
	if o == nil {
		return nil
	}
	return o.JitOrganizationIds
}

type GetOneTimeTokenStatus string

const (
	GetOneTimeTokenStatusActive   GetOneTimeTokenStatus = "active"
	GetOneTimeTokenStatusConsumed GetOneTimeTokenStatus = "consumed"
	GetOneTimeTokenStatusRevoked  GetOneTimeTokenStatus = "revoked"
	GetOneTimeTokenStatusExpired  GetOneTimeTokenStatus = "expired"
)

func (e GetOneTimeTokenStatus) ToPointer() *GetOneTimeTokenStatus {
	return &e
}
func (e *GetOneTimeTokenStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "consumed":
		fallthrough
	case "revoked":
		fallthrough
	case "expired":
		*e = GetOneTimeTokenStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetOneTimeTokenStatus: %v", v)
	}
}

// GetOneTimeTokenResponseBody - The one-time token found by ID.
type GetOneTimeTokenResponseBody struct {
	TenantID  string                 `json:"tenantId"`
	ID        string                 `json:"id"`
	Email     string                 `json:"email"`
	Token     string                 `json:"token"`
	Context   GetOneTimeTokenContext `json:"context"`
	Status    GetOneTimeTokenStatus  `json:"status"`
	CreatedAt float64                `json:"createdAt"`
	ExpiresAt float64                `json:"expiresAt"`
}

func (o *GetOneTimeTokenResponseBody) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *GetOneTimeTokenResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetOneTimeTokenResponseBody) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *GetOneTimeTokenResponseBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *GetOneTimeTokenResponseBody) GetContext() GetOneTimeTokenContext {
	if o == nil {
		return GetOneTimeTokenContext{}
	}
	return o.Context
}

func (o *GetOneTimeTokenResponseBody) GetStatus() GetOneTimeTokenStatus {
	if o == nil {
		return GetOneTimeTokenStatus("")
	}
	return o.Status
}

func (o *GetOneTimeTokenResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *GetOneTimeTokenResponseBody) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

type GetOneTimeTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The one-time token found by ID.
	Object *GetOneTimeTokenResponseBody
}

func (o *GetOneTimeTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetOneTimeTokenResponse) GetObject() *GetOneTimeTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
