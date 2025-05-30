// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type ReplaceOrganizationApplicationRolesRequestBody struct {
	// An array of role IDs to replace existing roles.
	OrganizationRoleIds []string `json:"organizationRoleIds"`
}

func (o *ReplaceOrganizationApplicationRolesRequestBody) GetOrganizationRoleIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OrganizationRoleIds
}

type ReplaceOrganizationApplicationRolesRequest struct {
	// The unique identifier of the organization.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique identifier of the application.
	ApplicationID string                                         `pathParam:"style=simple,explode=false,name=applicationId"`
	RequestBody   ReplaceOrganizationApplicationRolesRequestBody `request:"mediaType=application/json"`
}

func (o *ReplaceOrganizationApplicationRolesRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ReplaceOrganizationApplicationRolesRequest) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *ReplaceOrganizationApplicationRolesRequest) GetRequestBody() ReplaceOrganizationApplicationRolesRequestBody {
	if o == nil {
		return ReplaceOrganizationApplicationRolesRequestBody{}
	}
	return o.RequestBody
}

type ReplaceOrganizationApplicationRolesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ReplaceOrganizationApplicationRolesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
