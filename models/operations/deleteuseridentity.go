// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type DeleteUserIdentityRequest struct {
	// The unique identifier of the user.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	Target string `pathParam:"style=simple,explode=false,name=target"`
}

func (o *DeleteUserIdentityRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *DeleteUserIdentityRequest) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

// DeleteUserIdentityCustomData - arbitrary
type DeleteUserIdentityCustomData struct {
}

// DeleteUserIdentityDetails - arbitrary
type DeleteUserIdentityDetails struct {
}

type DeleteUserIdentityIdentities struct {
	UserID string `json:"userId"`
	// arbitrary
	Details *DeleteUserIdentityDetails `json:"details,omitempty"`
}

func (o *DeleteUserIdentityIdentities) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *DeleteUserIdentityIdentities) GetDetails() *DeleteUserIdentityDetails {
	if o == nil {
		return nil
	}
	return o.Details
}

type DeleteUserIdentityAddress struct {
	Formatted     *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postalCode,omitempty"`
	Country       *string `json:"country,omitempty"`
}

func (o *DeleteUserIdentityAddress) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *DeleteUserIdentityAddress) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *DeleteUserIdentityAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *DeleteUserIdentityAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *DeleteUserIdentityAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *DeleteUserIdentityAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

type DeleteUserIdentityProfile struct {
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
	Address           *DeleteUserIdentityAddress `json:"address,omitempty"`
}

func (o *DeleteUserIdentityProfile) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *DeleteUserIdentityProfile) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *DeleteUserIdentityProfile) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *DeleteUserIdentityProfile) GetNickname() *string {
	if o == nil {
		return nil
	}
	return o.Nickname
}

func (o *DeleteUserIdentityProfile) GetPreferredUsername() *string {
	if o == nil {
		return nil
	}
	return o.PreferredUsername
}

func (o *DeleteUserIdentityProfile) GetProfile() *string {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *DeleteUserIdentityProfile) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *DeleteUserIdentityProfile) GetGender() *string {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *DeleteUserIdentityProfile) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *DeleteUserIdentityProfile) GetZoneinfo() *string {
	if o == nil {
		return nil
	}
	return o.Zoneinfo
}

func (o *DeleteUserIdentityProfile) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *DeleteUserIdentityProfile) GetAddress() *DeleteUserIdentityAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

// DeleteUserIdentityDetail - arbitrary
type DeleteUserIdentityDetail struct {
}

type DeleteUserIdentitySsoIdentity struct {
	TenantID   string `json:"tenantId"`
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	Issuer     string `json:"issuer"`
	IdentityID string `json:"identityId"`
	// arbitrary
	Detail         DeleteUserIdentityDetail `json:"detail"`
	CreatedAt      float64                  `json:"createdAt"`
	SsoConnectorID string                   `json:"ssoConnectorId"`
}

func (o *DeleteUserIdentitySsoIdentity) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *DeleteUserIdentitySsoIdentity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteUserIdentitySsoIdentity) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *DeleteUserIdentitySsoIdentity) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *DeleteUserIdentitySsoIdentity) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *DeleteUserIdentitySsoIdentity) GetDetail() DeleteUserIdentityDetail {
	if o == nil {
		return DeleteUserIdentityDetail{}
	}
	return o.Detail
}

func (o *DeleteUserIdentitySsoIdentity) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *DeleteUserIdentitySsoIdentity) GetSsoConnectorID() string {
	if o == nil {
		return ""
	}
	return o.SsoConnectorID
}

// DeleteUserIdentityResponseBody - The identity is deleted from the user.
type DeleteUserIdentityResponseBody struct {
	ID           string  `json:"id"`
	Username     *string `json:"username"`
	PrimaryEmail *string `json:"primaryEmail"`
	PrimaryPhone *string `json:"primaryPhone"`
	Name         *string `json:"name"`
	Avatar       *string `json:"avatar"`
	// arbitrary
	CustomData    DeleteUserIdentityCustomData            `json:"customData"`
	Identities    map[string]DeleteUserIdentityIdentities `json:"identities"`
	LastSignInAt  *float64                                `json:"lastSignInAt"`
	CreatedAt     float64                                 `json:"createdAt"`
	UpdatedAt     float64                                 `json:"updatedAt"`
	Profile       DeleteUserIdentityProfile               `json:"profile"`
	ApplicationID *string                                 `json:"applicationId"`
	IsSuspended   bool                                    `json:"isSuspended"`
	HasPassword   *bool                                   `json:"hasPassword,omitempty"`
	SsoIdentities []DeleteUserIdentitySsoIdentity         `json:"ssoIdentities,omitempty"`
}

func (o *DeleteUserIdentityResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteUserIdentityResponseBody) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *DeleteUserIdentityResponseBody) GetPrimaryEmail() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryEmail
}

func (o *DeleteUserIdentityResponseBody) GetPrimaryPhone() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryPhone
}

func (o *DeleteUserIdentityResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DeleteUserIdentityResponseBody) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *DeleteUserIdentityResponseBody) GetCustomData() DeleteUserIdentityCustomData {
	if o == nil {
		return DeleteUserIdentityCustomData{}
	}
	return o.CustomData
}

func (o *DeleteUserIdentityResponseBody) GetIdentities() map[string]DeleteUserIdentityIdentities {
	if o == nil {
		return map[string]DeleteUserIdentityIdentities{}
	}
	return o.Identities
}

func (o *DeleteUserIdentityResponseBody) GetLastSignInAt() *float64 {
	if o == nil {
		return nil
	}
	return o.LastSignInAt
}

func (o *DeleteUserIdentityResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *DeleteUserIdentityResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *DeleteUserIdentityResponseBody) GetProfile() DeleteUserIdentityProfile {
	if o == nil {
		return DeleteUserIdentityProfile{}
	}
	return o.Profile
}

func (o *DeleteUserIdentityResponseBody) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *DeleteUserIdentityResponseBody) GetIsSuspended() bool {
	if o == nil {
		return false
	}
	return o.IsSuspended
}

func (o *DeleteUserIdentityResponseBody) GetHasPassword() *bool {
	if o == nil {
		return nil
	}
	return o.HasPassword
}

func (o *DeleteUserIdentityResponseBody) GetSsoIdentities() []DeleteUserIdentitySsoIdentity {
	if o == nil {
		return nil
	}
	return o.SsoIdentities
}

type DeleteUserIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The identity is deleted from the user.
	Object *DeleteUserIdentityResponseBody
}

func (o *DeleteUserIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteUserIdentityResponse) GetObject() *DeleteUserIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
