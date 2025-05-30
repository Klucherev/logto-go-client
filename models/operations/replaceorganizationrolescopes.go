// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type ReplaceOrganizationRoleScopesRequestBody struct {
	// An array of organization scope IDs to replace existing scopes.
	OrganizationScopeIds []string `json:"organizationScopeIds"`
}

func (o *ReplaceOrganizationRoleScopesRequestBody) GetOrganizationScopeIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OrganizationScopeIds
}

type ReplaceOrganizationRoleScopesRequest struct {
	// The unique identifier of the organization role.
	ID          string                                   `pathParam:"style=simple,explode=false,name=id"`
	RequestBody ReplaceOrganizationRoleScopesRequestBody `request:"mediaType=application/json"`
}

func (o *ReplaceOrganizationRoleScopesRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ReplaceOrganizationRoleScopesRequest) GetRequestBody() ReplaceOrganizationRoleScopesRequestBody {
	if o == nil {
		return ReplaceOrganizationRoleScopesRequestBody{}
	}
	return o.RequestBody
}

type ReplaceOrganizationRoleScopesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ReplaceOrganizationRoleScopesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
