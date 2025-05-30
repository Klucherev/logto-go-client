// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/internal/utils"
	"github.com/klucherev/logto-go-client/models/components"
)

type ListOrganizationRoleScopesRequest struct {
	// The unique identifier of the organization role.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Page number (starts from 1).
	Page *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Entries per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page_size"`
}

func (l ListOrganizationRoleScopesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListOrganizationRoleScopesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListOrganizationRoleScopesRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListOrganizationRoleScopesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListOrganizationRoleScopesRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListOrganizationRoleScopesResponseBody struct {
	TenantID    string  `json:"tenantId"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (o *ListOrganizationRoleScopesResponseBody) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *ListOrganizationRoleScopesResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListOrganizationRoleScopesResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListOrganizationRoleScopesResponseBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type ListOrganizationRoleScopesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A list of organization scopes.
	ResponseBodies []ListOrganizationRoleScopesResponseBody
}

func (o *ListOrganizationRoleScopesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListOrganizationRoleScopesResponse) GetResponseBodies() []ListOrganizationRoleScopesResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
