// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type ReplaceUserRolesRequestBody struct {
	// An array of API resource role IDs to assign.
	RoleIds []string `json:"roleIds"`
}

func (o *ReplaceUserRolesRequestBody) GetRoleIds() []string {
	if o == nil {
		return []string{}
	}
	return o.RoleIds
}

type ReplaceUserRolesRequest struct {
	// The unique identifier of the user.
	UserID      string                      `pathParam:"style=simple,explode=false,name=userId"`
	RequestBody ReplaceUserRolesRequestBody `request:"mediaType=application/json"`
}

func (o *ReplaceUserRolesRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ReplaceUserRolesRequest) GetRequestBody() ReplaceUserRolesRequestBody {
	if o == nil {
		return ReplaceUserRolesRequestBody{}
	}
	return o.RequestBody
}

type ReplaceUserRolesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ReplaceUserRolesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
