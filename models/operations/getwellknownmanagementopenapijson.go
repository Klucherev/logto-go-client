// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type GetWellKnownManagementOpenapiJSONResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetWellKnownManagementOpenapiJSONResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
