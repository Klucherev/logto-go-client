// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klucherev/logto-go-client/internal/utils"
	"github.com/klucherev/logto-go-client/models/components"
)

// UpdateJwtCustomizerTokenTypePath - The token type to update a JWT customizer for.
type UpdateJwtCustomizerTokenTypePath string

const (
	UpdateJwtCustomizerTokenTypePathAccessToken       UpdateJwtCustomizerTokenTypePath = "access-token"
	UpdateJwtCustomizerTokenTypePathClientCredentials UpdateJwtCustomizerTokenTypePath = "client-credentials"
)

func (e UpdateJwtCustomizerTokenTypePath) ToPointer() *UpdateJwtCustomizerTokenTypePath {
	return &e
}
func (e *UpdateJwtCustomizerTokenTypePath) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access-token":
		fallthrough
	case "client-credentials":
		*e = UpdateJwtCustomizerTokenTypePath(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateJwtCustomizerTokenTypePath: %v", v)
	}
}

type UpdateJwtCustomizerRequestBody struct {
	// The script of the JWT customizer.
	Script any `json:"script,omitempty"`
	// The environment variables for the JWT customizer.
	EnvironmentVariables any `json:"environmentVariables,omitempty"`
	// The sample context for the JWT customizer script testing purpose.
	ContextSample any `json:"contextSample,omitempty"`
	// The sample raw token payload for the JWT customizer script testing purpose.
	TokenSample any `json:"tokenSample,omitempty"`
}

func (o *UpdateJwtCustomizerRequestBody) GetScript() any {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *UpdateJwtCustomizerRequestBody) GetEnvironmentVariables() any {
	if o == nil {
		return nil
	}
	return o.EnvironmentVariables
}

func (o *UpdateJwtCustomizerRequestBody) GetContextSample() any {
	if o == nil {
		return nil
	}
	return o.ContextSample
}

func (o *UpdateJwtCustomizerRequestBody) GetTokenSample() any {
	if o == nil {
		return nil
	}
	return o.TokenSample
}

type UpdateJwtCustomizerRequest struct {
	// The token type to update a JWT customizer for.
	TokenTypePath UpdateJwtCustomizerTokenTypePath `pathParam:"style=simple,explode=false,name=tokenTypePath"`
	RequestBody   UpdateJwtCustomizerRequestBody   `request:"mediaType=application/json"`
}

func (o *UpdateJwtCustomizerRequest) GetTokenTypePath() UpdateJwtCustomizerTokenTypePath {
	if o == nil {
		return UpdateJwtCustomizerTokenTypePath("")
	}
	return o.TokenTypePath
}

func (o *UpdateJwtCustomizerRequest) GetRequestBody() UpdateJwtCustomizerRequestBody {
	if o == nil {
		return UpdateJwtCustomizerRequestBody{}
	}
	return o.RequestBody
}

// UpdateJwtCustomizerContextSample2 - arbitrary
type UpdateJwtCustomizerContextSample2 struct {
}

type UpdateJwtCustomizerAud2Type string

const (
	UpdateJwtCustomizerAud2TypeStr        UpdateJwtCustomizerAud2Type = "str"
	UpdateJwtCustomizerAud2TypeArrayOfStr UpdateJwtCustomizerAud2Type = "arrayOfStr"
)

type UpdateJwtCustomizerAud2 struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type UpdateJwtCustomizerAud2Type
}

func CreateUpdateJwtCustomizerAud2Str(str string) UpdateJwtCustomizerAud2 {
	typ := UpdateJwtCustomizerAud2TypeStr

	return UpdateJwtCustomizerAud2{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpdateJwtCustomizerAud2ArrayOfStr(arrayOfStr []string) UpdateJwtCustomizerAud2 {
	typ := UpdateJwtCustomizerAud2TypeArrayOfStr

	return UpdateJwtCustomizerAud2{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpdateJwtCustomizerAud2) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpdateJwtCustomizerAud2TypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpdateJwtCustomizerAud2TypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateJwtCustomizerAud2", string(data))
}

func (u UpdateJwtCustomizerAud2) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateJwtCustomizerAud2: all fields are null")
}

type UpdateJwtCustomizerTokenSample2 struct {
	Jti      *string                  `json:"jti,omitempty"`
	Aud      *UpdateJwtCustomizerAud2 `json:"aud,omitempty"`
	Scope    *string                  `json:"scope,omitempty"`
	ClientID *string                  `json:"clientId,omitempty"`
	Kind     *string                  `json:"kind,omitempty"`
}

func (o *UpdateJwtCustomizerTokenSample2) GetJti() *string {
	if o == nil {
		return nil
	}
	return o.Jti
}

func (o *UpdateJwtCustomizerTokenSample2) GetAud() *UpdateJwtCustomizerAud2 {
	if o == nil {
		return nil
	}
	return o.Aud
}

func (o *UpdateJwtCustomizerTokenSample2) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *UpdateJwtCustomizerTokenSample2) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *UpdateJwtCustomizerTokenSample2) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

type UpdateJwtCustomizerResponseBody2 struct {
	Script               string            `json:"script"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	// arbitrary
	ContextSample *UpdateJwtCustomizerContextSample2 `json:"contextSample,omitempty"`
	TokenSample   *UpdateJwtCustomizerTokenSample2   `json:"tokenSample,omitempty"`
}

func (o *UpdateJwtCustomizerResponseBody2) GetScript() string {
	if o == nil {
		return ""
	}
	return o.Script
}

func (o *UpdateJwtCustomizerResponseBody2) GetEnvironmentVariables() map[string]string {
	if o == nil {
		return nil
	}
	return o.EnvironmentVariables
}

func (o *UpdateJwtCustomizerResponseBody2) GetContextSample() *UpdateJwtCustomizerContextSample2 {
	if o == nil {
		return nil
	}
	return o.ContextSample
}

func (o *UpdateJwtCustomizerResponseBody2) GetTokenSample() *UpdateJwtCustomizerTokenSample2 {
	if o == nil {
		return nil
	}
	return o.TokenSample
}

// UpdateJwtCustomizerCustomData - arbitrary
type UpdateJwtCustomizerCustomData struct {
}

// UpdateJwtCustomizerDetails - arbitrary
type UpdateJwtCustomizerDetails struct {
}

type UpdateJwtCustomizerIdentities struct {
	UserID string `json:"userId"`
	// arbitrary
	Details *UpdateJwtCustomizerDetails `json:"details,omitempty"`
}

func (o *UpdateJwtCustomizerIdentities) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateJwtCustomizerIdentities) GetDetails() *UpdateJwtCustomizerDetails {
	if o == nil {
		return nil
	}
	return o.Details
}

type UpdateJwtCustomizerAddress struct {
	Formatted     *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postalCode,omitempty"`
	Country       *string `json:"country,omitempty"`
}

func (o *UpdateJwtCustomizerAddress) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *UpdateJwtCustomizerAddress) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *UpdateJwtCustomizerAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *UpdateJwtCustomizerAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *UpdateJwtCustomizerAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *UpdateJwtCustomizerAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type UpdateJwtCustomizerProfile struct {
	FamilyName        *string                     `json:"familyName,omitempty"`
	GivenName         *string                     `json:"givenName,omitempty"`
	MiddleName        *string                     `json:"middleName,omitempty"`
	Nickname          *string                     `json:"nickname,omitempty"`
	PreferredUsername *string                     `json:"preferredUsername,omitempty"`
	Profile           *string                     `json:"profile,omitempty"`
	Website           *string                     `json:"website,omitempty"`
	Gender            *string                     `json:"gender,omitempty"`
	Birthdate         *string                     `json:"birthdate,omitempty"`
	Zoneinfo          *string                     `json:"zoneinfo,omitempty"`
	Locale            *string                     `json:"locale,omitempty"`
	Address           *UpdateJwtCustomizerAddress `json:"address,omitempty"`
}

func (o *UpdateJwtCustomizerProfile) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *UpdateJwtCustomizerProfile) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *UpdateJwtCustomizerProfile) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *UpdateJwtCustomizerProfile) GetNickname() *string {
	if o == nil {
		return nil
	}
	return o.Nickname
}

func (o *UpdateJwtCustomizerProfile) GetPreferredUsername() *string {
	if o == nil {
		return nil
	}
	return o.PreferredUsername
}

func (o *UpdateJwtCustomizerProfile) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UpdateJwtCustomizerProfile) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *UpdateJwtCustomizerProfile) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UpdateJwtCustomizerProfile) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *UpdateJwtCustomizerProfile) GetZoneinfo() *string {
	if o == nil {
		return nil
	}
	return o.Zoneinfo
}

func (o *UpdateJwtCustomizerProfile) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *UpdateJwtCustomizerProfile) GetAddress() *UpdateJwtCustomizerAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

// UpdateJwtCustomizerDetail - arbitrary
type UpdateJwtCustomizerDetail struct {
}

type UpdateJwtCustomizerSsoIdentity struct {
	Issuer     string `json:"issuer"`
	IdentityID string `json:"identityId"`
	// arbitrary
	Detail UpdateJwtCustomizerDetail `json:"detail"`
}

func (o *UpdateJwtCustomizerSsoIdentity) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *UpdateJwtCustomizerSsoIdentity) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *UpdateJwtCustomizerSsoIdentity) GetDetail() UpdateJwtCustomizerDetail {
	if o == nil {
		return UpdateJwtCustomizerDetail{}
	}
	return o.Detail
}

type UpdateJwtCustomizerMfaVerificationFactor string

const (
	UpdateJwtCustomizerMfaVerificationFactorTotp       UpdateJwtCustomizerMfaVerificationFactor = "Totp"
	UpdateJwtCustomizerMfaVerificationFactorWebAuthn   UpdateJwtCustomizerMfaVerificationFactor = "WebAuthn"
	UpdateJwtCustomizerMfaVerificationFactorBackupCode UpdateJwtCustomizerMfaVerificationFactor = "BackupCode"
)

func (e UpdateJwtCustomizerMfaVerificationFactor) ToPointer() *UpdateJwtCustomizerMfaVerificationFactor {
	return &e
}
func (e *UpdateJwtCustomizerMfaVerificationFactor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Totp":
		fallthrough
	case "WebAuthn":
		fallthrough
	case "BackupCode":
		*e = UpdateJwtCustomizerMfaVerificationFactor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateJwtCustomizerMfaVerificationFactor: %v", v)
	}
}

type UpdateJwtCustomizerResource struct {
	TenantID       string  `json:"tenantId"`
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Indicator      string  `json:"indicator"`
	IsDefault      bool    `json:"isDefault"`
	AccessTokenTTL float64 `json:"accessTokenTtl"`
}

func (o *UpdateJwtCustomizerResource) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *UpdateJwtCustomizerResource) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateJwtCustomizerResource) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateJwtCustomizerResource) GetIndicator() string {
	if o == nil {
		return ""
	}
	return o.Indicator
}

func (o *UpdateJwtCustomizerResource) GetIsDefault() bool {
	if o == nil {
		return false
	}
	return o.IsDefault
}

func (o *UpdateJwtCustomizerResource) GetAccessTokenTTL() float64 {
	if o == nil {
		return 0.0
	}
	return o.AccessTokenTTL
}

type UpdateJwtCustomizerScope struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description *string                     `json:"description"`
	ResourceID  string                      `json:"resourceId"`
	Resource    UpdateJwtCustomizerResource `json:"resource"`
}

func (o *UpdateJwtCustomizerScope) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateJwtCustomizerScope) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateJwtCustomizerScope) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateJwtCustomizerScope) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *UpdateJwtCustomizerScope) GetResource() UpdateJwtCustomizerResource {
	if o == nil {
		return UpdateJwtCustomizerResource{}
	}
	return o.Resource
}

type UpdateJwtCustomizerRole struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Scopes      []UpdateJwtCustomizerScope `json:"scopes"`
}

func (o *UpdateJwtCustomizerRole) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateJwtCustomizerRole) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateJwtCustomizerRole) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *UpdateJwtCustomizerRole) GetScopes() []UpdateJwtCustomizerScope {
	if o == nil {
		return []UpdateJwtCustomizerScope{}
	}
	return o.Scopes
}

type UpdateJwtCustomizerOrganization struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (o *UpdateJwtCustomizerOrganization) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateJwtCustomizerOrganization) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateJwtCustomizerOrganization) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type UpdateJwtCustomizerOrganizationRole struct {
	OrganizationID string `json:"organizationId"`
	RoleID         string `json:"roleId"`
	RoleName       string `json:"roleName"`
}

func (o *UpdateJwtCustomizerOrganizationRole) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *UpdateJwtCustomizerOrganizationRole) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

func (o *UpdateJwtCustomizerOrganizationRole) GetRoleName() string {
	if o == nil {
		return ""
	}
	return o.RoleName
}

type UpdateJwtCustomizerUser struct {
	ID           *string `json:"id,omitempty"`
	Username     *string `json:"username,omitempty"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	Name         *string `json:"name,omitempty"`
	Avatar       *string `json:"avatar,omitempty"`
	// arbitrary
	CustomData             *UpdateJwtCustomizerCustomData             `json:"customData,omitempty"`
	Identities             map[string]UpdateJwtCustomizerIdentities   `json:"identities,omitempty"`
	LastSignInAt           *float64                                   `json:"lastSignInAt,omitempty"`
	CreatedAt              *float64                                   `json:"createdAt,omitempty"`
	UpdatedAt              *float64                                   `json:"updatedAt,omitempty"`
	Profile                *UpdateJwtCustomizerProfile                `json:"profile,omitempty"`
	ApplicationID          *string                                    `json:"applicationId,omitempty"`
	IsSuspended            *bool                                      `json:"isSuspended,omitempty"`
	HasPassword            *bool                                      `json:"hasPassword,omitempty"`
	SsoIdentities          []UpdateJwtCustomizerSsoIdentity           `json:"ssoIdentities,omitempty"`
	MfaVerificationFactors []UpdateJwtCustomizerMfaVerificationFactor `json:"mfaVerificationFactors,omitempty"`
	Roles                  []UpdateJwtCustomizerRole                  `json:"roles,omitempty"`
	Organizations          []UpdateJwtCustomizerOrganization          `json:"organizations,omitempty"`
	OrganizationRoles      []UpdateJwtCustomizerOrganizationRole      `json:"organizationRoles,omitempty"`
}

func (o *UpdateJwtCustomizerUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateJwtCustomizerUser) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UpdateJwtCustomizerUser) GetPrimaryEmail() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryEmail
}

func (o *UpdateJwtCustomizerUser) GetPrimaryPhone() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryPhone
}

func (o *UpdateJwtCustomizerUser) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateJwtCustomizerUser) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *UpdateJwtCustomizerUser) GetCustomData() *UpdateJwtCustomizerCustomData {
	if o == nil {
		return nil
	}
	return o.CustomData
}

func (o *UpdateJwtCustomizerUser) GetIdentities() map[string]UpdateJwtCustomizerIdentities {
	if o == nil {
		return nil
	}
	return o.Identities
}

func (o *UpdateJwtCustomizerUser) GetLastSignInAt() *float64 {
	if o == nil {
		return nil
	}
	return o.LastSignInAt
}

func (o *UpdateJwtCustomizerUser) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpdateJwtCustomizerUser) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *UpdateJwtCustomizerUser) GetProfile() *UpdateJwtCustomizerProfile {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UpdateJwtCustomizerUser) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *UpdateJwtCustomizerUser) GetIsSuspended() *bool {
	if o == nil {
		return nil
	}
	return o.IsSuspended
}

func (o *UpdateJwtCustomizerUser) GetHasPassword() *bool {
	if o == nil {
		return nil
	}
	return o.HasPassword
}

func (o *UpdateJwtCustomizerUser) GetSsoIdentities() []UpdateJwtCustomizerSsoIdentity {
	if o == nil {
		return nil
	}
	return o.SsoIdentities
}

func (o *UpdateJwtCustomizerUser) GetMfaVerificationFactors() []UpdateJwtCustomizerMfaVerificationFactor {
	if o == nil {
		return nil
	}
	return o.MfaVerificationFactors
}

func (o *UpdateJwtCustomizerUser) GetRoles() []UpdateJwtCustomizerRole {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *UpdateJwtCustomizerUser) GetOrganizations() []UpdateJwtCustomizerOrganization {
	if o == nil {
		return nil
	}
	return o.Organizations
}

func (o *UpdateJwtCustomizerUser) GetOrganizationRoles() []UpdateJwtCustomizerOrganizationRole {
	if o == nil {
		return nil
	}
	return o.OrganizationRoles
}

// UpdateJwtCustomizerSubjectTokenContext - arbitrary
type UpdateJwtCustomizerSubjectTokenContext struct {
}

type UpdateJwtCustomizerGrant struct {
	Type *string `json:"type,omitempty"`
	// arbitrary
	SubjectTokenContext *UpdateJwtCustomizerSubjectTokenContext `json:"subjectTokenContext,omitempty"`
}

func (o *UpdateJwtCustomizerGrant) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateJwtCustomizerGrant) GetSubjectTokenContext() *UpdateJwtCustomizerSubjectTokenContext {
	if o == nil {
		return nil
	}
	return o.SubjectTokenContext
}

type UpdateJwtCustomizerContextSample1 struct {
	User  UpdateJwtCustomizerUser   `json:"user"`
	Grant *UpdateJwtCustomizerGrant `json:"grant,omitempty"`
}

func (o *UpdateJwtCustomizerContextSample1) GetUser() UpdateJwtCustomizerUser {
	if o == nil {
		return UpdateJwtCustomizerUser{}
	}
	return o.User
}

func (o *UpdateJwtCustomizerContextSample1) GetGrant() *UpdateJwtCustomizerGrant {
	if o == nil {
		return nil
	}
	return o.Grant
}

type UpdateJwtCustomizerAud1Type string

const (
	UpdateJwtCustomizerAud1TypeStr        UpdateJwtCustomizerAud1Type = "str"
	UpdateJwtCustomizerAud1TypeArrayOfStr UpdateJwtCustomizerAud1Type = "arrayOfStr"
)

type UpdateJwtCustomizerAud1 struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type UpdateJwtCustomizerAud1Type
}

func CreateUpdateJwtCustomizerAud1Str(str string) UpdateJwtCustomizerAud1 {
	typ := UpdateJwtCustomizerAud1TypeStr

	return UpdateJwtCustomizerAud1{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpdateJwtCustomizerAud1ArrayOfStr(arrayOfStr []string) UpdateJwtCustomizerAud1 {
	typ := UpdateJwtCustomizerAud1TypeArrayOfStr

	return UpdateJwtCustomizerAud1{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpdateJwtCustomizerAud1) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpdateJwtCustomizerAud1TypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpdateJwtCustomizerAud1TypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateJwtCustomizerAud1", string(data))
}

func (u UpdateJwtCustomizerAud1) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateJwtCustomizerAud1: all fields are null")
}

type UpdateJwtCustomizerTokenSample1 struct {
	Jti                *string                  `json:"jti,omitempty"`
	Aud                *UpdateJwtCustomizerAud1 `json:"aud,omitempty"`
	Scope              *string                  `json:"scope,omitempty"`
	ClientID           *string                  `json:"clientId,omitempty"`
	AccountID          *string                  `json:"accountId,omitempty"`
	ExpiresWithSession *bool                    `json:"expiresWithSession,omitempty"`
	GrantID            *string                  `json:"grantId,omitempty"`
	Gty                *string                  `json:"gty,omitempty"`
	SessionUID         *string                  `json:"sessionUid,omitempty"`
	Sid                *string                  `json:"sid,omitempty"`
	Kind               *string                  `json:"kind,omitempty"`
}

func (o *UpdateJwtCustomizerTokenSample1) GetJti() *string {
	if o == nil {
		return nil
	}
	return o.Jti
}

func (o *UpdateJwtCustomizerTokenSample1) GetAud() *UpdateJwtCustomizerAud1 {
	if o == nil {
		return nil
	}
	return o.Aud
}

func (o *UpdateJwtCustomizerTokenSample1) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *UpdateJwtCustomizerTokenSample1) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *UpdateJwtCustomizerTokenSample1) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UpdateJwtCustomizerTokenSample1) GetExpiresWithSession() *bool {
	if o == nil {
		return nil
	}
	return o.ExpiresWithSession
}

func (o *UpdateJwtCustomizerTokenSample1) GetGrantID() *string {
	if o == nil {
		return nil
	}
	return o.GrantID
}

func (o *UpdateJwtCustomizerTokenSample1) GetGty() *string {
	if o == nil {
		return nil
	}
	return o.Gty
}

func (o *UpdateJwtCustomizerTokenSample1) GetSessionUID() *string {
	if o == nil {
		return nil
	}
	return o.SessionUID
}

func (o *UpdateJwtCustomizerTokenSample1) GetSid() *string {
	if o == nil {
		return nil
	}
	return o.Sid
}

func (o *UpdateJwtCustomizerTokenSample1) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

type UpdateJwtCustomizerResponseBody1 struct {
	Script               string                             `json:"script"`
	EnvironmentVariables map[string]string                  `json:"environmentVariables,omitempty"`
	ContextSample        *UpdateJwtCustomizerContextSample1 `json:"contextSample,omitempty"`
	TokenSample          *UpdateJwtCustomizerTokenSample1   `json:"tokenSample,omitempty"`
}

func (o *UpdateJwtCustomizerResponseBody1) GetScript() string {
	if o == nil {
		return ""
	}
	return o.Script
}

func (o *UpdateJwtCustomizerResponseBody1) GetEnvironmentVariables() map[string]string {
	if o == nil {
		return nil
	}
	return o.EnvironmentVariables
}

func (o *UpdateJwtCustomizerResponseBody1) GetContextSample() *UpdateJwtCustomizerContextSample1 {
	if o == nil {
		return nil
	}
	return o.ContextSample
}

func (o *UpdateJwtCustomizerResponseBody1) GetTokenSample() *UpdateJwtCustomizerTokenSample1 {
	if o == nil {
		return nil
	}
	return o.TokenSample
}

type UpdateJwtCustomizerResponseBodyType string

const (
	UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody1 UpdateJwtCustomizerResponseBodyType = "UpdateJwtCustomizer_ResponseBody_1"
	UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody2 UpdateJwtCustomizerResponseBodyType = "UpdateJwtCustomizer_ResponseBody_2"
)

// UpdateJwtCustomizerResponseBody - The updated JWT customizer.
type UpdateJwtCustomizerResponseBody struct {
	UpdateJwtCustomizerResponseBody1 *UpdateJwtCustomizerResponseBody1 `queryParam:"inline"`
	UpdateJwtCustomizerResponseBody2 *UpdateJwtCustomizerResponseBody2 `queryParam:"inline"`

	Type UpdateJwtCustomizerResponseBodyType
}

func CreateUpdateJwtCustomizerResponseBodyUpdateJwtCustomizerResponseBody1(updateJwtCustomizerResponseBody1 UpdateJwtCustomizerResponseBody1) UpdateJwtCustomizerResponseBody {
	typ := UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody1

	return UpdateJwtCustomizerResponseBody{
		UpdateJwtCustomizerResponseBody1: &updateJwtCustomizerResponseBody1,
		Type:                             typ,
	}
}

func CreateUpdateJwtCustomizerResponseBodyUpdateJwtCustomizerResponseBody2(updateJwtCustomizerResponseBody2 UpdateJwtCustomizerResponseBody2) UpdateJwtCustomizerResponseBody {
	typ := UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody2

	return UpdateJwtCustomizerResponseBody{
		UpdateJwtCustomizerResponseBody2: &updateJwtCustomizerResponseBody2,
		Type:                             typ,
	}
}

func (u *UpdateJwtCustomizerResponseBody) UnmarshalJSON(data []byte) error {

	var updateJwtCustomizerResponseBody1 UpdateJwtCustomizerResponseBody1 = UpdateJwtCustomizerResponseBody1{}
	if err := utils.UnmarshalJSON(data, &updateJwtCustomizerResponseBody1, "", true, true); err == nil {
		u.UpdateJwtCustomizerResponseBody1 = &updateJwtCustomizerResponseBody1
		u.Type = UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody1
		return nil
	}

	var updateJwtCustomizerResponseBody2 UpdateJwtCustomizerResponseBody2 = UpdateJwtCustomizerResponseBody2{}
	if err := utils.UnmarshalJSON(data, &updateJwtCustomizerResponseBody2, "", true, true); err == nil {
		u.UpdateJwtCustomizerResponseBody2 = &updateJwtCustomizerResponseBody2
		u.Type = UpdateJwtCustomizerResponseBodyTypeUpdateJwtCustomizerResponseBody2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateJwtCustomizerResponseBody", string(data))
}

func (u UpdateJwtCustomizerResponseBody) MarshalJSON() ([]byte, error) {
	if u.UpdateJwtCustomizerResponseBody1 != nil {
		return utils.MarshalJSON(u.UpdateJwtCustomizerResponseBody1, "", true)
	}

	if u.UpdateJwtCustomizerResponseBody2 != nil {
		return utils.MarshalJSON(u.UpdateJwtCustomizerResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateJwtCustomizerResponseBody: all fields are null")
}

type UpdateJwtCustomizerResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The updated JWT customizer.
	OneOf *UpdateJwtCustomizerResponseBody
}

func (o *UpdateJwtCustomizerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateJwtCustomizerResponse) GetOneOf() *UpdateJwtCustomizerResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
