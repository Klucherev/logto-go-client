// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/klucherev/logto-go-client/models/components"
)

type ListSsoConnectorProvidersProviderName string

const (
	ListSsoConnectorProvidersProviderNameOidc            ListSsoConnectorProvidersProviderName = "OIDC"
	ListSsoConnectorProvidersProviderNameSaml            ListSsoConnectorProvidersProviderName = "SAML"
	ListSsoConnectorProvidersProviderNameAzureAd         ListSsoConnectorProvidersProviderName = "AzureAD"
	ListSsoConnectorProvidersProviderNameGoogleWorkspace ListSsoConnectorProvidersProviderName = "GoogleWorkspace"
	ListSsoConnectorProvidersProviderNameOkta            ListSsoConnectorProvidersProviderName = "Okta"
	ListSsoConnectorProvidersProviderNameAzureAdOidc     ListSsoConnectorProvidersProviderName = "AzureAdOidc"
)

func (e ListSsoConnectorProvidersProviderName) ToPointer() *ListSsoConnectorProvidersProviderName {
	return &e
}
func (e *ListSsoConnectorProvidersProviderName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OIDC":
		fallthrough
	case "SAML":
		fallthrough
	case "AzureAD":
		fallthrough
	case "GoogleWorkspace":
		fallthrough
	case "Okta":
		fallthrough
	case "AzureAdOidc":
		*e = ListSsoConnectorProvidersProviderName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSsoConnectorProvidersProviderName: %v", v)
	}
}

type ListSsoConnectorProvidersProviderType string

const (
	ListSsoConnectorProvidersProviderTypeOidc ListSsoConnectorProvidersProviderType = "oidc"
	ListSsoConnectorProvidersProviderTypeSaml ListSsoConnectorProvidersProviderType = "saml"
)

func (e ListSsoConnectorProvidersProviderType) ToPointer() *ListSsoConnectorProvidersProviderType {
	return &e
}
func (e *ListSsoConnectorProvidersProviderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oidc":
		fallthrough
	case "saml":
		*e = ListSsoConnectorProvidersProviderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSsoConnectorProvidersProviderType: %v", v)
	}
}

type ListSsoConnectorProvidersResponseBody struct {
	ProviderName ListSsoConnectorProvidersProviderName `json:"providerName"`
	ProviderType ListSsoConnectorProvidersProviderType `json:"providerType"`
	Logo         string                                `json:"logo"`
	LogoDark     string                                `json:"logoDark"`
	Description  string                                `json:"description"`
	Name         string                                `json:"name"`
}

func (o *ListSsoConnectorProvidersResponseBody) GetProviderName() ListSsoConnectorProvidersProviderName {
	if o == nil {
		return ListSsoConnectorProvidersProviderName("")
	}
	return o.ProviderName
}

func (o *ListSsoConnectorProvidersResponseBody) GetProviderType() ListSsoConnectorProvidersProviderType {
	if o == nil {
		return ListSsoConnectorProvidersProviderType("")
	}
	return o.ProviderType
}

func (o *ListSsoConnectorProvidersResponseBody) GetLogo() string {
	if o == nil {
		return ""
	}
	return o.Logo
}

func (o *ListSsoConnectorProvidersResponseBody) GetLogoDark() string {
	if o == nil {
		return ""
	}
	return o.LogoDark
}

func (o *ListSsoConnectorProvidersResponseBody) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *ListSsoConnectorProvidersResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ListSsoConnectorProvidersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A list of SSO provider data.
	ResponseBodies []ListSsoConnectorProvidersResponseBody
}

func (o *ListSsoConnectorProvidersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListSsoConnectorProvidersResponse) GetResponseBodies() []ListSsoConnectorProvidersResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
