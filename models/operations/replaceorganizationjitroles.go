// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type ReplaceOrganizationJitRolesRequestBody struct {
	// An array of organization role IDs to replace existing organization roles.
	OrganizationRoleIds []string `json:"organizationRoleIds"`
}

func (o *ReplaceOrganizationJitRolesRequestBody) GetOrganizationRoleIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OrganizationRoleIds
}

type ReplaceOrganizationJitRolesRequest struct {
	// The unique identifier of the organization.
	ID          string                                 `pathParam:"style=simple,explode=false,name=id"`
	RequestBody ReplaceOrganizationJitRolesRequestBody `request:"mediaType=application/json"`
}

func (o *ReplaceOrganizationJitRolesRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ReplaceOrganizationJitRolesRequest) GetRequestBody() ReplaceOrganizationJitRolesRequestBody {
	if o == nil {
		return ReplaceOrganizationJitRolesRequestBody{}
	}
	return o.RequestBody
}

type ReplaceOrganizationJitRolesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ReplaceOrganizationJitRolesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
