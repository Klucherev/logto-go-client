// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type UpdatePrimaryEmailRequest struct {
	// The new email for the user.
	Email string `json:"email"`
	// The identifier verification record ID for the new email ownership verification.
	NewIdentifierVerificationRecordID string `json:"newIdentifierVerificationRecordId"`
}

func (o *UpdatePrimaryEmailRequest) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *UpdatePrimaryEmailRequest) GetNewIdentifierVerificationRecordID() string {
	if o == nil {
		return ""
	}
	return o.NewIdentifierVerificationRecordID
}

type UpdatePrimaryEmailResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdatePrimaryEmailResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
