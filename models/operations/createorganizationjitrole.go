// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type CreateOrganizationJitRoleRequestBody struct {
	// The organization role IDs to add.
	OrganizationRoleIds []string `json:"organizationRoleIds"`
}

func (o *CreateOrganizationJitRoleRequestBody) GetOrganizationRoleIds() []string {
	if o == nil {
		return []string{}
	}
	return o.OrganizationRoleIds
}

type CreateOrganizationJitRoleRequest struct {
	// The unique identifier of the organization.
	ID          string                               `pathParam:"style=simple,explode=false,name=id"`
	RequestBody CreateOrganizationJitRoleRequestBody `request:"mediaType=application/json"`
}

func (o *CreateOrganizationJitRoleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateOrganizationJitRoleRequest) GetRequestBody() CreateOrganizationJitRoleRequestBody {
	if o == nil {
		return CreateOrganizationJitRoleRequestBody{}
	}
	return o.RequestBody
}

type CreateOrganizationJitRoleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateOrganizationJitRoleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
