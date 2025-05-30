// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

// UsersUpdateUserCustomDataRequest - Custom data object to update for the given user ID. Note this will replace the entire custom data object.
//
// If you want to perform a partial update, use the `PATCH /api/users/{userId}/custom-data` endpoint instead.
type UsersUpdateUserCustomDataRequest struct {
}

type UpdateUserAddressRequest struct {
	Formatted     *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postalCode,omitempty"`
	Country       *string `json:"country,omitempty"`
}

func (o *UpdateUserAddressRequest) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *UpdateUserAddressRequest) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *UpdateUserAddressRequest) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *UpdateUserAddressRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *UpdateUserAddressRequest) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *UpdateUserAddressRequest) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type UsersUpdateUserProfileRequest struct {
	FamilyName        *string                   `json:"familyName,omitempty"`
	GivenName         *string                   `json:"givenName,omitempty"`
	MiddleName        *string                   `json:"middleName,omitempty"`
	Nickname          *string                   `json:"nickname,omitempty"`
	PreferredUsername *string                   `json:"preferredUsername,omitempty"`
	Profile           *string                   `json:"profile,omitempty"`
	Website           *string                   `json:"website,omitempty"`
	Gender            *string                   `json:"gender,omitempty"`
	Birthdate         *string                   `json:"birthdate,omitempty"`
	Zoneinfo          *string                   `json:"zoneinfo,omitempty"`
	Locale            *string                   `json:"locale,omitempty"`
	Address           *UpdateUserAddressRequest `json:"address,omitempty"`
}

func (o *UsersUpdateUserProfileRequest) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *UsersUpdateUserProfileRequest) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *UsersUpdateUserProfileRequest) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *UsersUpdateUserProfileRequest) GetNickname() *string {
	if o == nil {
		return nil
	}
	return o.Nickname
}

func (o *UsersUpdateUserProfileRequest) GetPreferredUsername() *string {
	if o == nil {
		return nil
	}
	return o.PreferredUsername
}

func (o *UsersUpdateUserProfileRequest) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UsersUpdateUserProfileRequest) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *UsersUpdateUserProfileRequest) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UsersUpdateUserProfileRequest) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *UsersUpdateUserProfileRequest) GetZoneinfo() *string {
	if o == nil {
		return nil
	}
	return o.Zoneinfo
}

func (o *UsersUpdateUserProfileRequest) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *UsersUpdateUserProfileRequest) GetAddress() *UpdateUserAddressRequest {
	if o == nil {
		return nil
	}
	return o.Address
}

type UpdateUserRequestBody struct {
	// Updated username for the user. It should be unique across all users.
	Username *string `json:"username,omitempty"`
	// Updated primary email address for the user. It should be unique across all users.
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	// Updated primary phone number for the user. It should be unique across all users.
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	Name         *string `json:"name,omitempty"`
	Avatar       *string `json:"avatar,omitempty"`
	// Custom data object to update for the given user ID. Note this will replace the entire custom data object.
	//
	// If you want to perform a partial update, use the `PATCH /api/users/{userId}/custom-data` endpoint instead.
	CustomData *UsersUpdateUserCustomDataRequest `json:"customData,omitempty"`
	Profile    *UsersUpdateUserProfileRequest    `json:"profile,omitempty"`
}

func (o *UpdateUserRequestBody) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UpdateUserRequestBody) GetPrimaryEmail() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryEmail
}

func (o *UpdateUserRequestBody) GetPrimaryPhone() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryPhone
}

func (o *UpdateUserRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateUserRequestBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *UpdateUserRequestBody) GetCustomData() *UsersUpdateUserCustomDataRequest {
	if o == nil {
		return nil
	}
	return o.CustomData
}

func (o *UpdateUserRequestBody) GetProfile() *UsersUpdateUserProfileRequest {
	if o == nil {
		return nil
	}
	return o.Profile
}

type UpdateUserRequest struct {
	// The unique identifier of the user.
	UserID      string                `pathParam:"style=simple,explode=false,name=userId"`
	RequestBody UpdateUserRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateUserRequest) GetRequestBody() UpdateUserRequestBody {
	if o == nil {
		return UpdateUserRequestBody{}
	}
	return o.RequestBody
}

// UsersUpdateUserCustomDataResponse - arbitrary
type UsersUpdateUserCustomDataResponse struct {
}

// UpdateUserDetails - arbitrary
type UpdateUserDetails struct {
}

type UpdateUserIdentities struct {
	UserID string `json:"userId"`
	// arbitrary
	Details *UpdateUserDetails `json:"details,omitempty"`
}

func (o *UpdateUserIdentities) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateUserIdentities) GetDetails() *UpdateUserDetails {
	if o == nil {
		return nil
	}
	return o.Details
}

type UpdateUserAddressResponse struct {
	Formatted     *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postalCode,omitempty"`
	Country       *string `json:"country,omitempty"`
}

func (o *UpdateUserAddressResponse) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *UpdateUserAddressResponse) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *UpdateUserAddressResponse) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *UpdateUserAddressResponse) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *UpdateUserAddressResponse) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *UpdateUserAddressResponse) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type UsersUpdateUserProfileResponse struct {
	FamilyName        *string                    `json:"familyName,omitempty"`
	GivenName         *string                    `json:"givenName,omitempty"`
	MiddleName        *string                    `json:"middleName,omitempty"`
	Nickname          *string                    `json:"nickname,omitempty"`
	PreferredUsername *string                    `json:"preferredUsername,omitempty"`
	Profile           *string                    `json:"profile,omitempty"`
	Website           *string                    `json:"website,omitempty"`
	Gender            *string                    `json:"gender,omitempty"`
	Birthdate         *string                    `json:"birthdate,omitempty"`
	Zoneinfo          *string                    `json:"zoneinfo,omitempty"`
	Locale            *string                    `json:"locale,omitempty"`
	Address           *UpdateUserAddressResponse `json:"address,omitempty"`
}

func (o *UsersUpdateUserProfileResponse) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *UsersUpdateUserProfileResponse) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *UsersUpdateUserProfileResponse) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *UsersUpdateUserProfileResponse) GetNickname() *string {
	if o == nil {
		return nil
	}
	return o.Nickname
}

func (o *UsersUpdateUserProfileResponse) GetPreferredUsername() *string {
	if o == nil {
		return nil
	}
	return o.PreferredUsername
}

func (o *UsersUpdateUserProfileResponse) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *UsersUpdateUserProfileResponse) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *UsersUpdateUserProfileResponse) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *UsersUpdateUserProfileResponse) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *UsersUpdateUserProfileResponse) GetZoneinfo() *string {
	if o == nil {
		return nil
	}
	return o.Zoneinfo
}

func (o *UsersUpdateUserProfileResponse) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *UsersUpdateUserProfileResponse) GetAddress() *UpdateUserAddressResponse {
	if o == nil {
		return nil
	}
	return o.Address
}

// UpdateUserDetail - arbitrary
type UpdateUserDetail struct {
}

type UpdateUserSsoIdentity struct {
	TenantID   string `json:"tenantId"`
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	Issuer     string `json:"issuer"`
	IdentityID string `json:"identityId"`
	// arbitrary
	Detail         UpdateUserDetail `json:"detail"`
	CreatedAt      float64          `json:"createdAt"`
	SsoConnectorID string           `json:"ssoConnectorId"`
}

func (o *UpdateUserSsoIdentity) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *UpdateUserSsoIdentity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateUserSsoIdentity) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateUserSsoIdentity) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *UpdateUserSsoIdentity) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *UpdateUserSsoIdentity) GetDetail() UpdateUserDetail {
	if o == nil {
		return UpdateUserDetail{}
	}
	return o.Detail
}

func (o *UpdateUserSsoIdentity) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *UpdateUserSsoIdentity) GetSsoConnectorID() string {
	if o == nil {
		return ""
	}
	return o.SsoConnectorID
}

// UpdateUserResponseBody - Updated user data for the given ID.
type UpdateUserResponseBody struct {
	ID           string  `json:"id"`
	Username     *string `json:"username"`
	PrimaryEmail *string `json:"primaryEmail"`
	PrimaryPhone *string `json:"primaryPhone"`
	Name         *string `json:"name"`
	Avatar       *string `json:"avatar"`
	// arbitrary
	CustomData    UsersUpdateUserCustomDataResponse `json:"customData"`
	Identities    map[string]UpdateUserIdentities   `json:"identities"`
	LastSignInAt  *float64                          `json:"lastSignInAt"`
	CreatedAt     float64                           `json:"createdAt"`
	UpdatedAt     float64                           `json:"updatedAt"`
	Profile       UsersUpdateUserProfileResponse    `json:"profile"`
	ApplicationID *string                           `json:"applicationId"`
	IsSuspended   bool                              `json:"isSuspended"`
	HasPassword   *bool                             `json:"hasPassword,omitempty"`
	SsoIdentities []UpdateUserSsoIdentity           `json:"ssoIdentities,omitempty"`
}

func (o *UpdateUserResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateUserResponseBody) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UpdateUserResponseBody) GetPrimaryEmail() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryEmail
}

func (o *UpdateUserResponseBody) GetPrimaryPhone() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryPhone
}

func (o *UpdateUserResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateUserResponseBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *UpdateUserResponseBody) GetCustomData() UsersUpdateUserCustomDataResponse {
	if o == nil {
		return UsersUpdateUserCustomDataResponse{}
	}
	return o.CustomData
}

func (o *UpdateUserResponseBody) GetIdentities() map[string]UpdateUserIdentities {
	if o == nil {
		return map[string]UpdateUserIdentities{}
	}
	return o.Identities
}

func (o *UpdateUserResponseBody) GetLastSignInAt() *float64 {
	if o == nil {
		return nil
	}
	return o.LastSignInAt
}

func (o *UpdateUserResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *UpdateUserResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *UpdateUserResponseBody) GetProfile() UsersUpdateUserProfileResponse {
	if o == nil {
		return UsersUpdateUserProfileResponse{}
	}
	return o.Profile
}

func (o *UpdateUserResponseBody) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *UpdateUserResponseBody) GetIsSuspended() bool {
	if o == nil {
		return false
	}
	return o.IsSuspended
}

func (o *UpdateUserResponseBody) GetHasPassword() *bool {
	if o == nil {
		return nil
	}
	return o.HasPassword
}

func (o *UpdateUserResponseBody) GetSsoIdentities() []UpdateUserSsoIdentity {
	if o == nil {
		return nil
	}
	return o.SsoIdentities
}

type UpdateUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updated user data for the given ID.
	Object *UpdateUserResponseBody
}

func (o *UpdateUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateUserResponse) GetObject() *UpdateUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
