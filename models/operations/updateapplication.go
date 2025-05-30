// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klucherev/logto-go-client/internal/utils"
	"github.com/klucherev/logto-go-client/models/components"
)

// UpdateApplicationRedirectUrisRequest2 - Validator function
type UpdateApplicationRedirectUrisRequest2 struct {
}

// UpdateApplicationRedirectUrisRequest1 - Validator function
type UpdateApplicationRedirectUrisRequest1 struct {
}

type UpdateApplicationRedirectUrisRequestUnionType string

const (
	UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest1 UpdateApplicationRedirectUrisRequestUnionType = "UpdateApplication_redirectUris_request_1"
	UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest2 UpdateApplicationRedirectUrisRequestUnionType = "UpdateApplication_redirectUris_request_2"
)

type UpdateApplicationRedirectUrisRequestUnion struct {
	UpdateApplicationRedirectUrisRequest1 *UpdateApplicationRedirectUrisRequest1 `queryParam:"inline"`
	UpdateApplicationRedirectUrisRequest2 *UpdateApplicationRedirectUrisRequest2 `queryParam:"inline"`

	Type UpdateApplicationRedirectUrisRequestUnionType
}

func CreateUpdateApplicationRedirectUrisRequestUnionUpdateApplicationRedirectUrisRequest1(updateApplicationRedirectUrisRequest1 UpdateApplicationRedirectUrisRequest1) UpdateApplicationRedirectUrisRequestUnion {
	typ := UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest1

	return UpdateApplicationRedirectUrisRequestUnion{
		UpdateApplicationRedirectUrisRequest1: &updateApplicationRedirectUrisRequest1,
		Type:                                  typ,
	}
}

func CreateUpdateApplicationRedirectUrisRequestUnionUpdateApplicationRedirectUrisRequest2(updateApplicationRedirectUrisRequest2 UpdateApplicationRedirectUrisRequest2) UpdateApplicationRedirectUrisRequestUnion {
	typ := UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest2

	return UpdateApplicationRedirectUrisRequestUnion{
		UpdateApplicationRedirectUrisRequest2: &updateApplicationRedirectUrisRequest2,
		Type:                                  typ,
	}
}

func (u *UpdateApplicationRedirectUrisRequestUnion) UnmarshalJSON(data []byte) error {

	var updateApplicationRedirectUrisRequest1 UpdateApplicationRedirectUrisRequest1 = UpdateApplicationRedirectUrisRequest1{}
	if err := utils.UnmarshalJSON(data, &updateApplicationRedirectUrisRequest1, "", true, true); err == nil {
		u.UpdateApplicationRedirectUrisRequest1 = &updateApplicationRedirectUrisRequest1
		u.Type = UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest1
		return nil
	}

	var updateApplicationRedirectUrisRequest2 UpdateApplicationRedirectUrisRequest2 = UpdateApplicationRedirectUrisRequest2{}
	if err := utils.UnmarshalJSON(data, &updateApplicationRedirectUrisRequest2, "", true, true); err == nil {
		u.UpdateApplicationRedirectUrisRequest2 = &updateApplicationRedirectUrisRequest2
		u.Type = UpdateApplicationRedirectUrisRequestUnionTypeUpdateApplicationRedirectUrisRequest2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateApplicationRedirectUrisRequestUnion", string(data))
}

func (u UpdateApplicationRedirectUrisRequestUnion) MarshalJSON() ([]byte, error) {
	if u.UpdateApplicationRedirectUrisRequest1 != nil {
		return utils.MarshalJSON(u.UpdateApplicationRedirectUrisRequest1, "", true)
	}

	if u.UpdateApplicationRedirectUrisRequest2 != nil {
		return utils.MarshalJSON(u.UpdateApplicationRedirectUrisRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateApplicationRedirectUrisRequestUnion: all fields are null")
}

type UpdateApplicationOidcClientMetadataRequest struct {
	RedirectUris                     []UpdateApplicationRedirectUrisRequestUnion `json:"redirectUris"`
	PostLogoutRedirectUris           []string                                    `json:"postLogoutRedirectUris"`
	BackchannelLogoutURI             *string                                     `json:"backchannelLogoutUri,omitempty"`
	BackchannelLogoutSessionRequired *bool                                       `json:"backchannelLogoutSessionRequired,omitempty"`
	LogoURI                          *string                                     `json:"logoUri,omitempty"`
}

func (o *UpdateApplicationOidcClientMetadataRequest) GetRedirectUris() []UpdateApplicationRedirectUrisRequestUnion {
	if o == nil {
		return []UpdateApplicationRedirectUrisRequestUnion{}
	}
	return o.RedirectUris
}

func (o *UpdateApplicationOidcClientMetadataRequest) GetPostLogoutRedirectUris() []string {
	if o == nil {
		return []string{}
	}
	return o.PostLogoutRedirectUris
}

func (o *UpdateApplicationOidcClientMetadataRequest) GetBackchannelLogoutURI() *string {
	if o == nil {
		return nil
	}
	return o.BackchannelLogoutURI
}

func (o *UpdateApplicationOidcClientMetadataRequest) GetBackchannelLogoutSessionRequired() *bool {
	if o == nil {
		return nil
	}
	return o.BackchannelLogoutSessionRequired
}

func (o *UpdateApplicationOidcClientMetadataRequest) GetLogoURI() *string {
	if o == nil {
		return nil
	}
	return o.LogoURI
}

type UpdateApplicationCustomClientMetadataRequest struct {
	CorsAllowedOrigins      []string `json:"corsAllowedOrigins,omitempty"`
	IDTokenTTL              *float64 `json:"idTokenTtl,omitempty"`
	RefreshTokenTTL         *float64 `json:"refreshTokenTtl,omitempty"`
	RefreshTokenTTLInDays   *float64 `json:"refreshTokenTtlInDays,omitempty"`
	TenantID                *string  `json:"tenantId,omitempty"`
	AlwaysIssueRefreshToken *bool    `json:"alwaysIssueRefreshToken,omitempty"`
	RotateRefreshToken      *bool    `json:"rotateRefreshToken,omitempty"`
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetCorsAllowedOrigins() []string {
	if o == nil {
		return nil
	}
	return o.CorsAllowedOrigins
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetIDTokenTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.IDTokenTTL
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetRefreshTokenTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenTTL
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetRefreshTokenTTLInDays() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenTTLInDays
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetAlwaysIssueRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.AlwaysIssueRefreshToken
}

func (o *UpdateApplicationCustomClientMetadataRequest) GetRotateRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.RotateRefreshToken
}

// ApplicationsUpdateApplicationCustomDataRequest - arbitrary
type ApplicationsUpdateApplicationCustomDataRequest struct {
}

type PageRuleRequest struct {
	Path string `json:"path"`
}

func (o *PageRuleRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type UpdateApplicationProtectedAppMetadataRequest struct {
	Origin          *string           `json:"origin,omitempty"`
	SessionDuration *float64          `json:"sessionDuration,omitempty"`
	PageRules       []PageRuleRequest `json:"pageRules,omitempty"`
}

func (o *UpdateApplicationProtectedAppMetadataRequest) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *UpdateApplicationProtectedAppMetadataRequest) GetSessionDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.SessionDuration
}

func (o *UpdateApplicationProtectedAppMetadataRequest) GetPageRules() []PageRuleRequest {
	if o == nil {
		return nil
	}
	return o.PageRules
}

type UpdateApplicationRequestBody struct {
	Name                 *string                                       `json:"name,omitempty"`
	Description          *string                                       `json:"description,omitempty"`
	OidcClientMetadata   *UpdateApplicationOidcClientMetadataRequest   `json:"oidcClientMetadata,omitempty"`
	CustomClientMetadata *UpdateApplicationCustomClientMetadataRequest `json:"customClientMetadata,omitempty"`
	// arbitrary
	CustomData           *ApplicationsUpdateApplicationCustomDataRequest `json:"customData,omitempty"`
	ProtectedAppMetadata *UpdateApplicationProtectedAppMetadataRequest   `json:"protectedAppMetadata,omitempty"`
	// Whether the application has admin access. User can enable the admin access for Machine-to-Machine apps.
	IsAdmin *bool `json:"isAdmin,omitempty"`
}

func (o *UpdateApplicationRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateApplicationRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateApplicationRequestBody) GetOidcClientMetadata() *UpdateApplicationOidcClientMetadataRequest {
	if o == nil {
		return nil
	}
	return o.OidcClientMetadata
}

func (o *UpdateApplicationRequestBody) GetCustomClientMetadata() *UpdateApplicationCustomClientMetadataRequest {
	if o == nil {
		return nil
	}
	return o.CustomClientMetadata
}

func (o *UpdateApplicationRequestBody) GetCustomData() *ApplicationsUpdateApplicationCustomDataRequest {
	if o == nil {
		return nil
	}
	return o.CustomData
}

func (o *UpdateApplicationRequestBody) GetProtectedAppMetadata() *UpdateApplicationProtectedAppMetadataRequest {
	if o == nil {
		return nil
	}
	return o.ProtectedAppMetadata
}

func (o *UpdateApplicationRequestBody) GetIsAdmin() *bool {
	if o == nil {
		return nil
	}
	return o.IsAdmin
}

type UpdateApplicationRequest struct {
	// The unique identifier of the application.
	ID          string                       `pathParam:"style=simple,explode=false,name=id"`
	RequestBody UpdateApplicationRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateApplicationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateApplicationRequest) GetRequestBody() UpdateApplicationRequestBody {
	if o == nil {
		return UpdateApplicationRequestBody{}
	}
	return o.RequestBody
}

type UpdateApplicationType string

const (
	UpdateApplicationTypeNative           UpdateApplicationType = "Native"
	UpdateApplicationTypeSpa              UpdateApplicationType = "SPA"
	UpdateApplicationTypeTraditional      UpdateApplicationType = "Traditional"
	UpdateApplicationTypeMachineToMachine UpdateApplicationType = "MachineToMachine"
	UpdateApplicationTypeProtected        UpdateApplicationType = "Protected"
	UpdateApplicationTypeSaml             UpdateApplicationType = "SAML"
)

func (e UpdateApplicationType) ToPointer() *UpdateApplicationType {
	return &e
}
func (e *UpdateApplicationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Native":
		fallthrough
	case "SPA":
		fallthrough
	case "Traditional":
		fallthrough
	case "MachineToMachine":
		fallthrough
	case "Protected":
		fallthrough
	case "SAML":
		*e = UpdateApplicationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateApplicationType: %v", v)
	}
}

// UpdateApplicationRedirectUrisResponse2 - Validator function
type UpdateApplicationRedirectUrisResponse2 struct {
}

// UpdateApplicationRedirectUrisResponse1 - Validator function
type UpdateApplicationRedirectUrisResponse1 struct {
}

type UpdateApplicationRedirectUrisResponseUnionType string

const (
	UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse1 UpdateApplicationRedirectUrisResponseUnionType = "UpdateApplication_redirectUris_response_1"
	UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse2 UpdateApplicationRedirectUrisResponseUnionType = "UpdateApplication_redirectUris_response_2"
)

type UpdateApplicationRedirectUrisResponseUnion struct {
	UpdateApplicationRedirectUrisResponse1 *UpdateApplicationRedirectUrisResponse1 `queryParam:"inline"`
	UpdateApplicationRedirectUrisResponse2 *UpdateApplicationRedirectUrisResponse2 `queryParam:"inline"`

	Type UpdateApplicationRedirectUrisResponseUnionType
}

func CreateUpdateApplicationRedirectUrisResponseUnionUpdateApplicationRedirectUrisResponse1(updateApplicationRedirectUrisResponse1 UpdateApplicationRedirectUrisResponse1) UpdateApplicationRedirectUrisResponseUnion {
	typ := UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse1

	return UpdateApplicationRedirectUrisResponseUnion{
		UpdateApplicationRedirectUrisResponse1: &updateApplicationRedirectUrisResponse1,
		Type:                                   typ,
	}
}

func CreateUpdateApplicationRedirectUrisResponseUnionUpdateApplicationRedirectUrisResponse2(updateApplicationRedirectUrisResponse2 UpdateApplicationRedirectUrisResponse2) UpdateApplicationRedirectUrisResponseUnion {
	typ := UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse2

	return UpdateApplicationRedirectUrisResponseUnion{
		UpdateApplicationRedirectUrisResponse2: &updateApplicationRedirectUrisResponse2,
		Type:                                   typ,
	}
}

func (u *UpdateApplicationRedirectUrisResponseUnion) UnmarshalJSON(data []byte) error {

	var updateApplicationRedirectUrisResponse1 UpdateApplicationRedirectUrisResponse1 = UpdateApplicationRedirectUrisResponse1{}
	if err := utils.UnmarshalJSON(data, &updateApplicationRedirectUrisResponse1, "", true, true); err == nil {
		u.UpdateApplicationRedirectUrisResponse1 = &updateApplicationRedirectUrisResponse1
		u.Type = UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse1
		return nil
	}

	var updateApplicationRedirectUrisResponse2 UpdateApplicationRedirectUrisResponse2 = UpdateApplicationRedirectUrisResponse2{}
	if err := utils.UnmarshalJSON(data, &updateApplicationRedirectUrisResponse2, "", true, true); err == nil {
		u.UpdateApplicationRedirectUrisResponse2 = &updateApplicationRedirectUrisResponse2
		u.Type = UpdateApplicationRedirectUrisResponseUnionTypeUpdateApplicationRedirectUrisResponse2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateApplicationRedirectUrisResponseUnion", string(data))
}

func (u UpdateApplicationRedirectUrisResponseUnion) MarshalJSON() ([]byte, error) {
	if u.UpdateApplicationRedirectUrisResponse1 != nil {
		return utils.MarshalJSON(u.UpdateApplicationRedirectUrisResponse1, "", true)
	}

	if u.UpdateApplicationRedirectUrisResponse2 != nil {
		return utils.MarshalJSON(u.UpdateApplicationRedirectUrisResponse2, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateApplicationRedirectUrisResponseUnion: all fields are null")
}

type UpdateApplicationOidcClientMetadataResponse struct {
	RedirectUris                     []UpdateApplicationRedirectUrisResponseUnion `json:"redirectUris"`
	PostLogoutRedirectUris           []string                                     `json:"postLogoutRedirectUris"`
	BackchannelLogoutURI             *string                                      `json:"backchannelLogoutUri,omitempty"`
	BackchannelLogoutSessionRequired *bool                                        `json:"backchannelLogoutSessionRequired,omitempty"`
	LogoURI                          *string                                      `json:"logoUri,omitempty"`
}

func (o *UpdateApplicationOidcClientMetadataResponse) GetRedirectUris() []UpdateApplicationRedirectUrisResponseUnion {
	if o == nil {
		return []UpdateApplicationRedirectUrisResponseUnion{}
	}
	return o.RedirectUris
}

func (o *UpdateApplicationOidcClientMetadataResponse) GetPostLogoutRedirectUris() []string {
	if o == nil {
		return []string{}
	}
	return o.PostLogoutRedirectUris
}

func (o *UpdateApplicationOidcClientMetadataResponse) GetBackchannelLogoutURI() *string {
	if o == nil {
		return nil
	}
	return o.BackchannelLogoutURI
}

func (o *UpdateApplicationOidcClientMetadataResponse) GetBackchannelLogoutSessionRequired() *bool {
	if o == nil {
		return nil
	}
	return o.BackchannelLogoutSessionRequired
}

func (o *UpdateApplicationOidcClientMetadataResponse) GetLogoURI() *string {
	if o == nil {
		return nil
	}
	return o.LogoURI
}

type UpdateApplicationCustomClientMetadataResponse struct {
	CorsAllowedOrigins      []string `json:"corsAllowedOrigins,omitempty"`
	IDTokenTTL              *float64 `json:"idTokenTtl,omitempty"`
	RefreshTokenTTL         *float64 `json:"refreshTokenTtl,omitempty"`
	RefreshTokenTTLInDays   *float64 `json:"refreshTokenTtlInDays,omitempty"`
	TenantID                *string  `json:"tenantId,omitempty"`
	AlwaysIssueRefreshToken *bool    `json:"alwaysIssueRefreshToken,omitempty"`
	RotateRefreshToken      *bool    `json:"rotateRefreshToken,omitempty"`
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetCorsAllowedOrigins() []string {
	if o == nil {
		return nil
	}
	return o.CorsAllowedOrigins
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetIDTokenTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.IDTokenTTL
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetRefreshTokenTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenTTL
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetRefreshTokenTTLInDays() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenTTLInDays
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetAlwaysIssueRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.AlwaysIssueRefreshToken
}

func (o *UpdateApplicationCustomClientMetadataResponse) GetRotateRefreshToken() *bool {
	if o == nil {
		return nil
	}
	return o.RotateRefreshToken
}

type UpdateApplicationPageRuleResponse struct {
	Path string `json:"path"`
}

func (o *UpdateApplicationPageRuleResponse) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type UpdateApplicationStatus string

const (
	UpdateApplicationStatusPendingVerification UpdateApplicationStatus = "PendingVerification"
	UpdateApplicationStatusPendingSsl          UpdateApplicationStatus = "PendingSsl"
	UpdateApplicationStatusActive              UpdateApplicationStatus = "Active"
	UpdateApplicationStatusError               UpdateApplicationStatus = "Error"
)

func (e UpdateApplicationStatus) ToPointer() *UpdateApplicationStatus {
	return &e
}
func (e *UpdateApplicationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PendingVerification":
		fallthrough
	case "PendingSsl":
		fallthrough
	case "Active":
		fallthrough
	case "Error":
		*e = UpdateApplicationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateApplicationStatus: %v", v)
	}
}

type UpdateApplicationDNSRecord struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (o *UpdateApplicationDNSRecord) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateApplicationDNSRecord) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateApplicationDNSRecord) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type UpdateApplicationValidationError struct {
	Message string `json:"message"`
}

func (o *UpdateApplicationValidationError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type UpdateApplicationSsl struct {
	Status           string                             `json:"status"`
	ValidationErrors []UpdateApplicationValidationError `json:"validation_errors,omitempty"`
}

func (o *UpdateApplicationSsl) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateApplicationSsl) GetValidationErrors() []UpdateApplicationValidationError {
	if o == nil {
		return nil
	}
	return o.ValidationErrors
}

type UpdateApplicationCloudflareData struct {
	ID                 string               `json:"id"`
	Status             string               `json:"status"`
	Ssl                UpdateApplicationSsl `json:"ssl"`
	VerificationErrors []string             `json:"verification_errors,omitempty"`
}

func (o *UpdateApplicationCloudflareData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateApplicationCloudflareData) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateApplicationCloudflareData) GetSsl() UpdateApplicationSsl {
	if o == nil {
		return UpdateApplicationSsl{}
	}
	return o.Ssl
}

func (o *UpdateApplicationCloudflareData) GetVerificationErrors() []string {
	if o == nil {
		return nil
	}
	return o.VerificationErrors
}

type UpdateApplicationCustomDomain struct {
	Domain         string                           `json:"domain"`
	Status         UpdateApplicationStatus          `json:"status"`
	ErrorMessage   *string                          `json:"errorMessage"`
	DNSRecords     []UpdateApplicationDNSRecord     `json:"dnsRecords"`
	CloudflareData *UpdateApplicationCloudflareData `json:"cloudflareData"`
}

func (o *UpdateApplicationCustomDomain) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *UpdateApplicationCustomDomain) GetStatus() UpdateApplicationStatus {
	if o == nil {
		return UpdateApplicationStatus("")
	}
	return o.Status
}

func (o *UpdateApplicationCustomDomain) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateApplicationCustomDomain) GetDNSRecords() []UpdateApplicationDNSRecord {
	if o == nil {
		return []UpdateApplicationDNSRecord{}
	}
	return o.DNSRecords
}

func (o *UpdateApplicationCustomDomain) GetCloudflareData() *UpdateApplicationCloudflareData {
	if o == nil {
		return nil
	}
	return o.CloudflareData
}

type UpdateApplicationProtectedAppMetadataResponse struct {
	Host            string                              `json:"host"`
	Origin          string                              `json:"origin"`
	SessionDuration float64                             `json:"sessionDuration"`
	PageRules       []UpdateApplicationPageRuleResponse `json:"pageRules"`
	CustomDomains   []UpdateApplicationCustomDomain     `json:"customDomains,omitempty"`
}

func (o *UpdateApplicationProtectedAppMetadataResponse) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *UpdateApplicationProtectedAppMetadataResponse) GetOrigin() string {
	if o == nil {
		return ""
	}
	return o.Origin
}

func (o *UpdateApplicationProtectedAppMetadataResponse) GetSessionDuration() float64 {
	if o == nil {
		return 0.0
	}
	return o.SessionDuration
}

func (o *UpdateApplicationProtectedAppMetadataResponse) GetPageRules() []UpdateApplicationPageRuleResponse {
	if o == nil {
		return []UpdateApplicationPageRuleResponse{}
	}
	return o.PageRules
}

func (o *UpdateApplicationProtectedAppMetadataResponse) GetCustomDomains() []UpdateApplicationCustomDomain {
	if o == nil {
		return nil
	}
	return o.CustomDomains
}

// ApplicationsUpdateApplicationCustomDataResponse - arbitrary
type ApplicationsUpdateApplicationCustomDataResponse struct {
}

// UpdateApplicationResponseBody - The application was updated successfully.
type UpdateApplicationResponseBody struct {
	TenantID string `json:"tenantId"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	// The internal client secret. Note it is only used for internal validation, and the actual secrets should be retrieved from `/api/applications/{id}/secrets` endpoints.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Secret               string                                         `json:"secret"`
	Description          *string                                        `json:"description"`
	Type                 UpdateApplicationType                          `json:"type"`
	OidcClientMetadata   UpdateApplicationOidcClientMetadataResponse    `json:"oidcClientMetadata"`
	CustomClientMetadata UpdateApplicationCustomClientMetadataResponse  `json:"customClientMetadata"`
	ProtectedAppMetadata *UpdateApplicationProtectedAppMetadataResponse `json:"protectedAppMetadata"`
	// arbitrary
	CustomData   ApplicationsUpdateApplicationCustomDataResponse `json:"customData"`
	IsThirdParty bool                                            `json:"isThirdParty"`
	CreatedAt    float64                                         `json:"createdAt"`
}

func (o *UpdateApplicationResponseBody) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *UpdateApplicationResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateApplicationResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateApplicationResponseBody) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

func (o *UpdateApplicationResponseBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateApplicationResponseBody) GetType() UpdateApplicationType {
	if o == nil {
		return UpdateApplicationType("")
	}
	return o.Type
}

func (o *UpdateApplicationResponseBody) GetOidcClientMetadata() UpdateApplicationOidcClientMetadataResponse {
	if o == nil {
		return UpdateApplicationOidcClientMetadataResponse{}
	}
	return o.OidcClientMetadata
}

func (o *UpdateApplicationResponseBody) GetCustomClientMetadata() UpdateApplicationCustomClientMetadataResponse {
	if o == nil {
		return UpdateApplicationCustomClientMetadataResponse{}
	}
	return o.CustomClientMetadata
}

func (o *UpdateApplicationResponseBody) GetProtectedAppMetadata() *UpdateApplicationProtectedAppMetadataResponse {
	if o == nil {
		return nil
	}
	return o.ProtectedAppMetadata
}

func (o *UpdateApplicationResponseBody) GetCustomData() ApplicationsUpdateApplicationCustomDataResponse {
	if o == nil {
		return ApplicationsUpdateApplicationCustomDataResponse{}
	}
	return o.CustomData
}

func (o *UpdateApplicationResponseBody) GetIsThirdParty() bool {
	if o == nil {
		return false
	}
	return o.IsThirdParty
}

func (o *UpdateApplicationResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

type UpdateApplicationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The application was updated successfully.
	Object *UpdateApplicationResponseBody
}

func (o *UpdateApplicationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateApplicationResponse) GetObject() *UpdateApplicationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
