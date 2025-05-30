// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type ReplaceApplicationUserConsentOrganizationsRequestBody struct {
	// A list of organization ids to be granted. <br/> All the existing organizations' access will be revoked if not in the list. <br/> If the list is empty, all the organizations' access will be revoked.
	OrganizationIds []string `json:"organizationIds"`
}

func (o *ReplaceApplicationUserConsentOrganizationsRequestBody) GetOrganizationIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OrganizationIds
}

type ReplaceApplicationUserConsentOrganizationsRequest struct {
	// The unique identifier of the application.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique identifier of the user.
	UserID      string                                                `pathParam:"style=simple,explode=false,name=userId"`
	RequestBody ReplaceApplicationUserConsentOrganizationsRequestBody `request:"mediaType=application/json"`
}

func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetRequestBody() ReplaceApplicationUserConsentOrganizationsRequestBody {
	if o == nil {
		return ReplaceApplicationUserConsentOrganizationsRequestBody{}
	}
	return o.RequestBody
}

type ReplaceApplicationUserConsentOrganizationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ReplaceApplicationUserConsentOrganizationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
