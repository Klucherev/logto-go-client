// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type UpdateProfileRequest struct {
	// The new name for the user.
	Name *string `json:"name,omitempty"`
	// The new avatar for the user, must be a URL.
	Avatar *string `json:"avatar,omitempty"`
	// The new username for the user, must be a valid username and unique.
	Username *string `json:"username,omitempty"`
}

func (o *UpdateProfileRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateProfileRequest) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *UpdateProfileRequest) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// UpdateProfileCustomData - arbitrary
type UpdateProfileCustomData struct {
}

// UpdateProfileDetails - arbitrary
type UpdateProfileDetails struct {
}

type UpdateProfileIdentities struct {
	UserID string `json:"userId"`
	// arbitrary
	Details *UpdateProfileDetails `json:"details,omitempty"`
}

func (o *UpdateProfileIdentities) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateProfileIdentities) GetDetails() *UpdateProfileDetails {
	if o == nil {
		return nil
	}
	return o.Details
}

type UpdateProfileAddress struct {
	Formatted     *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postalCode,omitempty"`
	Country       *string `json:"country,omitempty"`
}

func (o *UpdateProfileAddress) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *UpdateProfileAddress) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *UpdateProfileAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *UpdateProfileAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *UpdateProfileAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *UpdateProfileAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type UpdateProfileProfile struct {
	FamilyName        *string               `json:"familyName,omitempty"`
	GivenName         *string               `json:"givenName,omitempty"`
	MiddleName        *string               `json:"middleName,omitempty"`
	Nickname          *string               `json:"nickname,omitempty"`
	PreferredUsername *string               `json:"preferredUsername,omitempty"`
	Profile           *string               `json:"profile,omitempty"`
	Website           *string               `json:"website,omitempty"`
	Gender            *string               `json:"gender,omitempty"`
	Birthdate         *string               `json:"birthdate,omitempty"`
	Zoneinfo          *string               `json:"zoneinfo,omitempty"`
	Locale            *string               `json:"locale,omitempty"`
	Address           *UpdateProfileAddress `json:"address,omitempty"`
}

func (o *UpdateProfileProfile) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *UpdateProfileProfile) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *UpdateProfileProfile) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *UpdateProfileProfile) GetNickname() *string {
	if o == nil {
		return nil
	}
	return o.Nickname
}

func (o *UpdateProfileProfile) GetPreferredUsername() *string {
	if o == nil {
		return nil
	}
	return o.PreferredUsername
}

func (o *UpdateProfileProfile) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UpdateProfileProfile) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *UpdateProfileProfile) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UpdateProfileProfile) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *UpdateProfileProfile) GetZoneinfo() *string {
	if o == nil {
		return nil
	}
	return o.Zoneinfo
}

func (o *UpdateProfileProfile) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *UpdateProfileProfile) GetAddress() *UpdateProfileAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

// UpdateProfileDetail - arbitrary
type UpdateProfileDetail struct {
}

type UpdateProfileSsoIdentity struct {
	TenantID   string `json:"tenantId"`
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	Issuer     string `json:"issuer"`
	IdentityID string `json:"identityId"`
	// arbitrary
	Detail         UpdateProfileDetail `json:"detail"`
	CreatedAt      float64             `json:"createdAt"`
	SsoConnectorID string              `json:"ssoConnectorId"`
}

func (o *UpdateProfileSsoIdentity) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *UpdateProfileSsoIdentity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateProfileSsoIdentity) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateProfileSsoIdentity) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *UpdateProfileSsoIdentity) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *UpdateProfileSsoIdentity) GetDetail() UpdateProfileDetail {
	if o == nil {
		return UpdateProfileDetail{}
	}
	return o.Detail
}

func (o *UpdateProfileSsoIdentity) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *UpdateProfileSsoIdentity) GetSsoConnectorID() string {
	if o == nil {
		return ""
	}
	return o.SsoConnectorID
}

// UpdateProfileResponseBody - The profile was updated successfully.
type UpdateProfileResponseBody struct {
	ID           *string `json:"id,omitempty"`
	Username     *string `json:"username,omitempty"`
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	Name         *string `json:"name,omitempty"`
	Avatar       *string `json:"avatar,omitempty"`
	// arbitrary
	CustomData    *UpdateProfileCustomData           `json:"customData,omitempty"`
	Identities    map[string]UpdateProfileIdentities `json:"identities,omitempty"`
	LastSignInAt  *float64                           `json:"lastSignInAt,omitempty"`
	CreatedAt     *float64                           `json:"createdAt,omitempty"`
	UpdatedAt     *float64                           `json:"updatedAt,omitempty"`
	Profile       *UpdateProfileProfile              `json:"profile,omitempty"`
	ApplicationID *string                            `json:"applicationId,omitempty"`
	IsSuspended   *bool                              `json:"isSuspended,omitempty"`
	HasPassword   *bool                              `json:"hasPassword,omitempty"`
	SsoIdentities []UpdateProfileSsoIdentity         `json:"ssoIdentities,omitempty"`
}

func (o *UpdateProfileResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateProfileResponseBody) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UpdateProfileResponseBody) GetPrimaryEmail() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryEmail
}

func (o *UpdateProfileResponseBody) GetPrimaryPhone() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryPhone
}

func (o *UpdateProfileResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateProfileResponseBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *UpdateProfileResponseBody) GetCustomData() *UpdateProfileCustomData {
	if o == nil {
		return nil
	}
	return o.CustomData
}

func (o *UpdateProfileResponseBody) GetIdentities() map[string]UpdateProfileIdentities {
	if o == nil {
		return nil
	}
	return o.Identities
}

func (o *UpdateProfileResponseBody) GetLastSignInAt() *float64 {
	if o == nil {
		return nil
	}
	return o.LastSignInAt
}

func (o *UpdateProfileResponseBody) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpdateProfileResponseBody) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *UpdateProfileResponseBody) GetProfile() *UpdateProfileProfile {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UpdateProfileResponseBody) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *UpdateProfileResponseBody) GetIsSuspended() *bool {
	if o == nil {
		return nil
	}
	return o.IsSuspended
}

func (o *UpdateProfileResponseBody) GetHasPassword() *bool {
	if o == nil {
		return nil
	}
	return o.HasPassword
}

func (o *UpdateProfileResponseBody) GetSsoIdentities() []UpdateProfileSsoIdentity {
	if o == nil {
		return nil
	}
	return o.SsoIdentities
}

type UpdateProfileResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The profile was updated successfully.
	Object *UpdateProfileResponseBody
}

func (o *UpdateProfileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateProfileResponse) GetObject() *UpdateProfileResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
