// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type DeleteRoleUserRequest struct {
	// The unique identifier of the role.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique identifier of the user.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *DeleteRoleUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteRoleUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DeleteRoleUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteRoleUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
