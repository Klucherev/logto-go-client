# github.com/klucherev/logto

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/klucherev/logto* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/klucherev/logto&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/bizapp/rentavita). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Logto API references: API references for Logto services.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/klucherev/logto](#githubcomklucherevlogto)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/Klucherev/logto-go-client
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                          | Type   | Scheme                         | Environment Variable                                              |
| ----------------------------- | ------ | ------------------------------ | ----------------------------------------------------------------- |
| `ClientID`<br/>`ClientSecret` | oauth2 | OAuth2 Client Credentials Flow | `LOGTO_CLIENT_ID`<br/>`LOGTO_CLIENT_SECRET`<br/>`LOGTO_TOKEN_URL` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountCenter](docs/sdks/accountcenter/README.md)

* [GetAccountCenterSettings](docs/sdks/accountcenter/README.md#getaccountcentersettings) - Get account center settings
* [UpdateAccountCenterSettings](docs/sdks/accountcenter/README.md#updateaccountcentersettings) - Update account center settings

### [Applications](docs/sdks/applications/README.md)

* [ListApplications](docs/sdks/applications/README.md#listapplications) - Get applications
* [CreateApplication](docs/sdks/applications/README.md#createapplication) - Create an application
* [GetApplication](docs/sdks/applications/README.md#getapplication) - Get application
* [UpdateApplication](docs/sdks/applications/README.md#updateapplication) - Update application
* [DeleteApplication](docs/sdks/applications/README.md#deleteapplication) - Delete application
* [UpdateApplicationCustomData](docs/sdks/applications/README.md#updateapplicationcustomdata) - Update application custom data
* [ListApplicationRoles](docs/sdks/applications/README.md#listapplicationroles) - Get application API resource roles
* [AssignApplicationRoles](docs/sdks/applications/README.md#assignapplicationroles) - Assign API resource roles to application
* [ReplaceApplicationRoles](docs/sdks/applications/README.md#replaceapplicationroles) - Update API resource roles for application
* [DeleteApplicationRole](docs/sdks/applications/README.md#deleteapplicationrole) - Remove a API resource role from application
* [ListApplicationProtectedAppMetadataCustomDomains](docs/sdks/applications/README.md#listapplicationprotectedappmetadatacustomdomains) - Get application custom domains.
* [CreateApplicationProtectedAppMetadataCustomDomain](docs/sdks/applications/README.md#createapplicationprotectedappmetadatacustomdomain) - Add a custom domain to the application.
* [DeleteApplicationProtectedAppMetadataCustomDomain](docs/sdks/applications/README.md#deleteapplicationprotectedappmetadatacustomdomain) - Remove custom domain.
* [ListApplicationOrganizations](docs/sdks/applications/README.md#listapplicationorganizations) - Get application organizations
* [DeleteApplicationLegacySecret](docs/sdks/applications/README.md#deleteapplicationlegacysecret) - Delete application legacy secret
* [ListApplicationSecrets](docs/sdks/applications/README.md#listapplicationsecrets) - Get application secrets
* [CreateApplicationSecret](docs/sdks/applications/README.md#createapplicationsecret) - Add application secret
* [DeleteApplicationSecret](docs/sdks/applications/README.md#deleteapplicationsecret) - Delete application secret
* [UpdateApplicationSecret](docs/sdks/applications/README.md#updateapplicationsecret) - Update application secret
* [CreateApplicationUserConsentScope](docs/sdks/applications/README.md#createapplicationuserconsentscope) - Assign user consent scopes to application.
* [ListApplicationUserConsentScopes](docs/sdks/applications/README.md#listapplicationuserconsentscopes) - List all the user consent scopes of an application.
* [DeleteApplicationUserConsentScope](docs/sdks/applications/README.md#deleteapplicationuserconsentscope) - Remove user consent scope from application.
* [ReplaceApplicationSignInExperience](docs/sdks/applications/README.md#replaceapplicationsigninexperience) - Update application level sign-in experience
* [GetApplicationSignInExperience](docs/sdks/applications/README.md#getapplicationsigninexperience) - Get the application level sign-in experience
* [ListApplicationUserConsentOrganizations](docs/sdks/applications/README.md#listapplicationuserconsentorganizations) - List all the user consented organizations of a application.
* [ReplaceApplicationUserConsentOrganizations](docs/sdks/applications/README.md#replaceapplicationuserconsentorganizations) - Grant a list of organization access of a user for a application.
* [CreateApplicationUserConsentOrganization](docs/sdks/applications/README.md#createapplicationuserconsentorganization) - Grant a list of organization access of a user for a application.
* [DeleteApplicationUserConsentOrganization](docs/sdks/applications/README.md#deleteapplicationuserconsentorganization) - Revoke a user's access to an organization for a application.

### [AuditLogs](docs/sdks/auditlogs/README.md)

* [ListLogs](docs/sdks/auditlogs/README.md#listlogs) - Get logs
* [GetLog](docs/sdks/auditlogs/README.md#getlog) - Get log

### [Authn](docs/sdks/authn/README.md)

* [GetHasuraAuth](docs/sdks/authn/README.md#gethasuraauth) - Hasura auth hook endpoint
* [~~AssertSaml~~](docs/sdks/authn/README.md#assertsaml) - SAML ACS endpoint (social) :warning: **Deprecated**
* [AssertSingleSignOnSaml](docs/sdks/authn/README.md#assertsinglesignonsaml) - SAML ACS endpoint (SSO)

### [CaptchaProvider](docs/sdks/captchaprovider/README.md)

* [GetCaptchaProvider](docs/sdks/captchaprovider/README.md#getcaptchaprovider) - Get captcha provider
* [UpdateCaptchaProvider](docs/sdks/captchaprovider/README.md#updatecaptchaprovider) - Update captcha provider
* [DeleteCaptchaProvider](docs/sdks/captchaprovider/README.md#deletecaptchaprovider) - Delete captcha provider

### [Configs](docs/sdks/configs/README.md)

* [GetAdminConsoleConfig](docs/sdks/configs/README.md#getadminconsoleconfig) - Get admin console config
* [UpdateAdminConsoleConfig](docs/sdks/configs/README.md#updateadminconsoleconfig) - Update admin console config
* [GetOidcKeys](docs/sdks/configs/README.md#getoidckeys) - Get OIDC keys
* [DeleteOidcKey](docs/sdks/configs/README.md#deleteoidckey) - Delete OIDC key
* [RotateOidcKeys](docs/sdks/configs/README.md#rotateoidckeys) - Rotate OIDC keys
* [UpsertJwtCustomizer](docs/sdks/configs/README.md#upsertjwtcustomizer) - Create or update JWT customizer
* [UpdateJwtCustomizer](docs/sdks/configs/README.md#updatejwtcustomizer) - Update JWT customizer
* [GetJwtCustomizer](docs/sdks/configs/README.md#getjwtcustomizer) - Get JWT customizer
* [DeleteJwtCustomizer](docs/sdks/configs/README.md#deletejwtcustomizer) - Delete JWT customizer
* [ListJwtCustomizers](docs/sdks/configs/README.md#listjwtcustomizers) - Get all JWT customizers
* [TestJwtCustomizer](docs/sdks/configs/README.md#testjwtcustomizer) - Test JWT customizer

### [ConnectorFactories](docs/sdks/connectorfactories/README.md)

* [ListConnectorFactories](docs/sdks/connectorfactories/README.md#listconnectorfactories) - Get connector factories
* [GetConnectorFactory](docs/sdks/connectorfactories/README.md#getconnectorfactory) - Get connector factory

### [Connectors](docs/sdks/connectors/README.md)

* [CreateConnector](docs/sdks/connectors/README.md#createconnector) - Create connector
* [ListConnectors](docs/sdks/connectors/README.md#listconnectors) - Get connectors
* [GetConnector](docs/sdks/connectors/README.md#getconnector) - Get connector
* [UpdateConnector](docs/sdks/connectors/README.md#updateconnector) - Update connector
* [DeleteConnector](docs/sdks/connectors/README.md#deleteconnector) - Delete connector
* [CreateConnectorTest](docs/sdks/connectors/README.md#createconnectortest) - Test passwordless connector
* [CreateConnectorAuthorizationURI](docs/sdks/connectors/README.md#createconnectorauthorizationuri) - Get connector's authorization URI

### [CustomPhrases](docs/sdks/customphrases/README.md)

* [ListCustomPhrases](docs/sdks/customphrases/README.md#listcustomphrases) - Get all custom phrases
* [GetCustomPhrase](docs/sdks/customphrases/README.md#getcustomphrase) - Get custom phrases
* [ReplaceCustomPhrase](docs/sdks/customphrases/README.md#replacecustomphrase) - Upsert custom phrases
* [DeleteCustomPhrase](docs/sdks/customphrases/README.md#deletecustomphrase) - Delete custom phrase

### [Dashboard](docs/sdks/dashboard/README.md)

* [GetTotalUserCount](docs/sdks/dashboard/README.md#gettotalusercount) - Get total user count
* [GetNewUserCounts](docs/sdks/dashboard/README.md#getnewusercounts) - Get new user count
* [GetActiveUserCounts](docs/sdks/dashboard/README.md#getactiveusercounts) - Get active user data

### [Domains](docs/sdks/domains/README.md)

* [ListDomains](docs/sdks/domains/README.md#listdomains) - Get domains
* [CreateDomain](docs/sdks/domains/README.md#createdomain) - Create domain
* [GetDomain](docs/sdks/domains/README.md#getdomain) - Get domain
* [DeleteDomain](docs/sdks/domains/README.md#deletedomain) - Delete domain

### [EmailTemplates](docs/sdks/emailtemplates/README.md)

* [ReplaceEmailTemplates](docs/sdks/emailtemplates/README.md#replaceemailtemplates) - Replace email templates
* [ListEmailTemplates](docs/sdks/emailtemplates/README.md#listemailtemplates) - Get email templates
* [DeleteEmailTemplates](docs/sdks/emailtemplates/README.md#deleteemailtemplates) - Delete email templates
* [GetEmailTemplate](docs/sdks/emailtemplates/README.md#getemailtemplate) - Get email template by ID
* [DeleteEmailTemplate](docs/sdks/emailtemplates/README.md#deleteemailtemplate) - Delete an email template
* [UpdateEmailTemplateDetails](docs/sdks/emailtemplates/README.md#updateemailtemplatedetails) - Update email template details

### [Experience](docs/sdks/experience/README.md)

* [InitInteraction](docs/sdks/experience/README.md#initinteraction) - Init new interaction
* [UpdateInteractionEvent](docs/sdks/experience/README.md#updateinteractionevent) - Update interaction event
* [IdentifyUser](docs/sdks/experience/README.md#identifyuser) - Identify user for the current interaction
* [SubmitInteraction](docs/sdks/experience/README.md#submitinteraction) - Submit interaction
* [CreatePasswordVerification](docs/sdks/experience/README.md#createpasswordverification) - Create password verification record
* [CreateAndSendVerificationCode](docs/sdks/experience/README.md#createandsendverificationcode) - Create and send verification code
* [VerifyVerificationCodeVerification](docs/sdks/experience/README.md#verifyverificationcodeverification) - Verify verification code
* [CreateSocialVerification](docs/sdks/experience/README.md#createsocialverification) - Create social verification
* [VerifySocialVerification](docs/sdks/experience/README.md#verifysocialverification) - Verify social verification
* [CreateEnterpriseSsoVerification](docs/sdks/experience/README.md#createenterprisessoverification) - Create enterprise SSO verification
* [VerifyEnterpriseSsoVerification](docs/sdks/experience/README.md#verifyenterprisessoverification) - Verify enterprise SSO verification
* [CreateTotpSecret](docs/sdks/experience/README.md#createtotpsecret) - Create TOTP secret
* [VerifyTotpVerification](docs/sdks/experience/README.md#verifytotpverification) - Verify TOTP verification
* [CreateWebAuthnRegistrationVerification](docs/sdks/experience/README.md#createwebauthnregistrationverification) - Create WebAuthn registration verification
* [VerifyWebAuthnRegistrationVerification](docs/sdks/experience/README.md#verifywebauthnregistrationverification) - Verify WebAuthn registration verification
* [CreateWebAuthnAuthenticationVerification](docs/sdks/experience/README.md#createwebauthnauthenticationverification) - Create WebAuthn authentication verification
* [VerifyWebAuthnAuthenticationVerification](docs/sdks/experience/README.md#verifywebauthnauthenticationverification) - Verify WebAuthn authentication verification
* [GenerateBackupCodes](docs/sdks/experience/README.md#generatebackupcodes) - Generate backup codes
* [VerifyBackupCode](docs/sdks/experience/README.md#verifybackupcode) - Verify backup code
* [CreateNewPasswordIdentityVerification](docs/sdks/experience/README.md#createnewpasswordidentityverification) - Create new password identity verification
* [VerifyOneTimeTokenVerification](docs/sdks/experience/README.md#verifyonetimetokenverification) - Verify one-time token
* [AddUserProfile](docs/sdks/experience/README.md#adduserprofile) - Add user profile
* [ResetUserPassword](docs/sdks/experience/README.md#resetuserpassword) - Reset user password
* [SkipMfaBindingFlow](docs/sdks/experience/README.md#skipmfabindingflow) - Skip MFA binding flow
* [BindMfaVerification](docs/sdks/experience/README.md#bindmfaverification) - Bind MFA verification by verificationId
* [GetEnabledSsoConnectors](docs/sdks/experience/README.md#getenabledssoconnectors) - Get enabled SSO connectors by the given email's domain

### [Hooks](docs/sdks/hooks/README.md)

* [ListHooks](docs/sdks/hooks/README.md#listhooks) - Get hooks
* [CreateHook](docs/sdks/hooks/README.md#createhook) - Create a hook
* [GetHook](docs/sdks/hooks/README.md#gethook) - Get hook
* [UpdateHook](docs/sdks/hooks/README.md#updatehook) - Update hook
* [DeleteHook](docs/sdks/hooks/README.md#deletehook) - Delete hook
* [ListHookRecentLogs](docs/sdks/hooks/README.md#listhookrecentlogs) - Get recent logs for a hook
* [CreateHookTest](docs/sdks/hooks/README.md#createhooktest) - Test hook
* [UpdateHookSigningKey](docs/sdks/hooks/README.md#updatehooksigningkey) - Update signing key for a hook

### [Interaction](docs/sdks/interaction/README.md)

* [PutAPIInteraction](docs/sdks/interaction/README.md#putapiinteraction)
* [DeleteAPIInteraction](docs/sdks/interaction/README.md#deleteapiinteraction)
* [PutAPIInteractionEvent](docs/sdks/interaction/README.md#putapiinteractionevent)
* [PatchAPIInteractionIdentifiers](docs/sdks/interaction/README.md#patchapiinteractionidentifiers)
* [PutAPIInteractionProfile](docs/sdks/interaction/README.md#putapiinteractionprofile)
* [PatchAPIInteractionProfile](docs/sdks/interaction/README.md#patchapiinteractionprofile)
* [DeleteAPIInteractionProfile](docs/sdks/interaction/README.md#deleteapiinteractionprofile)
* [PostAPIInteractionSubmit](docs/sdks/interaction/README.md#postapiinteractionsubmit)
* [PostAPIInteractionConsent](docs/sdks/interaction/README.md#postapiinteractionconsent)
* [GetAPIInteractionConsent](docs/sdks/interaction/README.md#getapiinteractionconsent)
* [PostAPIInteractionVerificationSocialAuthorizationURI](docs/sdks/interaction/README.md#postapiinteractionverificationsocialauthorizationuri)
* [PostAPIInteractionVerificationVerificationCode](docs/sdks/interaction/README.md#postapiinteractionverificationverificationcode)
* [PostAPIInteractionVerificationTotp](docs/sdks/interaction/README.md#postapiinteractionverificationtotp)
* [PostAPIInteractionVerificationWebauthnRegistration](docs/sdks/interaction/README.md#postapiinteractionverificationwebauthnregistration)
* [PostAPIInteractionVerificationWebauthnAuthentication](docs/sdks/interaction/README.md#postapiinteractionverificationwebauthnauthentication)
* [PostAPIInteractionBindMfa](docs/sdks/interaction/README.md#postapiinteractionbindmfa)
* [PutAPIInteractionMfa](docs/sdks/interaction/README.md#putapiinteractionmfa)
* [PutAPIInteractionMfaSkipped](docs/sdks/interaction/README.md#putapiinteractionmfaskipped)
* [PostAPIInteractionSingleSignOnConnectorIDAuthorizationURL](docs/sdks/interaction/README.md#postapiinteractionsinglesignonconnectoridauthorizationurl)
* [PostAPIInteractionSingleSignOnConnectorIDAuthentication](docs/sdks/interaction/README.md#postapiinteractionsinglesignonconnectoridauthentication)
* [PostAPIInteractionSingleSignOnConnectorIDRegistration](docs/sdks/interaction/README.md#postapiinteractionsinglesignonconnectoridregistration)
* [GetAPIInteractionSingleSignOnConnectors](docs/sdks/interaction/README.md#getapiinteractionsinglesignonconnectors)


### [MyAccount](docs/sdks/myaccount/README.md)

* [GetProfile](docs/sdks/myaccount/README.md#getprofile) - Get profile
* [UpdateProfile](docs/sdks/myaccount/README.md#updateprofile) - Update profile
* [UpdateOtherProfile](docs/sdks/myaccount/README.md#updateotherprofile) - Update other profile
* [UpdatePassword](docs/sdks/myaccount/README.md#updatepassword) - Update password
* [UpdatePrimaryEmail](docs/sdks/myaccount/README.md#updateprimaryemail) - Update primary email
* [DeletePrimaryEmail](docs/sdks/myaccount/README.md#deleteprimaryemail) - Delete primary email
* [UpdatePrimaryPhone](docs/sdks/myaccount/README.md#updateprimaryphone) - Update primary phone
* [DeletePrimaryPhone](docs/sdks/myaccount/README.md#deleteprimaryphone) - Delete primary phone
* [AddUserIdentities](docs/sdks/myaccount/README.md#adduseridentities) - Add a user identity
* [DeleteIdentity](docs/sdks/myaccount/README.md#deleteidentity) - Delete a user identity

### [OneTimeTokens](docs/sdks/onetimetokens/README.md)

* [ListOneTimeTokens](docs/sdks/onetimetokens/README.md#listonetimetokens) - Get one-time tokens
* [AddOneTimeTokens](docs/sdks/onetimetokens/README.md#addonetimetokens) - Create one-time token
* [GetOneTimeToken](docs/sdks/onetimetokens/README.md#getonetimetoken) - Get one-time token by ID
* [DeleteOneTimeToken](docs/sdks/onetimetokens/README.md#deleteonetimetoken) - Delete one-time token by ID
* [VerifyOneTimeToken](docs/sdks/onetimetokens/README.md#verifyonetimetoken) - Verify one-time token
* [ReplaceOneTimeTokenStatus](docs/sdks/onetimetokens/README.md#replaceonetimetokenstatus) - Update one-time token status

### [OrganizationInvitations](docs/sdks/organizationinvitations/README.md)

* [GetOrganizationInvitation](docs/sdks/organizationinvitations/README.md#getorganizationinvitation) - Get organization invitation
* [DeleteOrganizationInvitation](docs/sdks/organizationinvitations/README.md#deleteorganizationinvitation) - Delete organization invitation
* [ListOrganizationInvitations](docs/sdks/organizationinvitations/README.md#listorganizationinvitations) - Get organization invitations
* [CreateOrganizationInvitation](docs/sdks/organizationinvitations/README.md#createorganizationinvitation) - Create organization invitation
* [CreateOrganizationInvitationMessage](docs/sdks/organizationinvitations/README.md#createorganizationinvitationmessage) - Resend invitation message
* [ReplaceOrganizationInvitationStatus](docs/sdks/organizationinvitations/README.md#replaceorganizationinvitationstatus) - Update organization invitation status

### [OrganizationRoles](docs/sdks/organizationroles/README.md)

* [GetOrganizationRole](docs/sdks/organizationroles/README.md#getorganizationrole) - Get organization role
* [UpdateOrganizationRole](docs/sdks/organizationroles/README.md#updateorganizationrole) - Update organization role
* [DeleteOrganizationRole](docs/sdks/organizationroles/README.md#deleteorganizationrole) - Delete organization role
* [ListOrganizationRoles](docs/sdks/organizationroles/README.md#listorganizationroles) - Get organization roles
* [CreateOrganizationRole](docs/sdks/organizationroles/README.md#createorganizationrole) - Create an organization role
* [ListOrganizationRoleScopes](docs/sdks/organizationroles/README.md#listorganizationrolescopes) - Get organization role scopes
* [CreateOrganizationRoleScope](docs/sdks/organizationroles/README.md#createorganizationrolescope) - Assign organization scopes to organization role
* [ReplaceOrganizationRoleScopes](docs/sdks/organizationroles/README.md#replaceorganizationrolescopes) - Replace organization scopes for organization role
* [DeleteOrganizationRoleScope](docs/sdks/organizationroles/README.md#deleteorganizationrolescope) - Remove organization scope
* [ListOrganizationRoleResourceScopes](docs/sdks/organizationroles/README.md#listorganizationroleresourcescopes) - Get organization role resource scopes
* [CreateOrganizationRoleResourceScope](docs/sdks/organizationroles/README.md#createorganizationroleresourcescope) - Assign resource scopes to organization role
* [ReplaceOrganizationRoleResourceScopes](docs/sdks/organizationroles/README.md#replaceorganizationroleresourcescopes) - Replace resource scopes for organization role
* [DeleteOrganizationRoleResourceScope](docs/sdks/organizationroles/README.md#deleteorganizationroleresourcescope) - Remove resource scope

### [Organizations](docs/sdks/organizations/README.md)

* [CreateOrganization](docs/sdks/organizations/README.md#createorganization) - Create an organization
* [ListOrganizations](docs/sdks/organizations/README.md#listorganizations) - Get organizations
* [GetOrganization](docs/sdks/organizations/README.md#getorganization) - Get organization
* [UpdateOrganization](docs/sdks/organizations/README.md#updateorganization) - Update organization
* [DeleteOrganization](docs/sdks/organizations/README.md#deleteorganization) - Delete organization
* [AddOrganizationUsers](docs/sdks/organizations/README.md#addorganizationusers) - Add user members to organization
* [ReplaceOrganizationUsers](docs/sdks/organizations/README.md#replaceorganizationusers) - Replace organization user members
* [ListOrganizationUsers](docs/sdks/organizations/README.md#listorganizationusers) - Get organization user members
* [DeleteOrganizationUser](docs/sdks/organizations/README.md#deleteorganizationuser) - Remove user member from organization
* [AssignOrganizationRolesToUsers](docs/sdks/organizations/README.md#assignorganizationrolestousers) - Assign roles to organization user members
* [ListOrganizationUserRoles](docs/sdks/organizations/README.md#listorganizationuserroles) - Get roles for a user in an organization
* [AssignOrganizationRolesToUser](docs/sdks/organizations/README.md#assignorganizationrolestouser) - Assign roles to a user in an organization
* [ReplaceOrganizationUserRoles](docs/sdks/organizations/README.md#replaceorganizationuserroles) - Update roles for a user in an organization
* [DeleteOrganizationUserRole](docs/sdks/organizations/README.md#deleteorganizationuserrole) - Remove a role from a user in an organization
* [ListOrganizationUserScopes](docs/sdks/organizations/README.md#listorganizationuserscopes) - Get scopes for a user in an organization tailored by the organization roles
* [AddOrganizationApplications](docs/sdks/organizations/README.md#addorganizationapplications) - Add organization application
* [ReplaceOrganizationApplications](docs/sdks/organizations/README.md#replaceorganizationapplications) - Replace organization applications
* [ListOrganizationApplications](docs/sdks/organizations/README.md#listorganizationapplications) - Get organization applications
* [DeleteOrganizationApplication](docs/sdks/organizations/README.md#deleteorganizationapplication) - Remove organization application
* [AssignOrganizationRolesToApplications](docs/sdks/organizations/README.md#assignorganizationrolestoapplications) - Assign roles to applications in an organization
* [ListOrganizationApplicationRoles](docs/sdks/organizations/README.md#listorganizationapplicationroles) - Get organization application roles
* [AssignOrganizationRolesToApplication](docs/sdks/organizations/README.md#assignorganizationrolestoapplication) - Add organization application role
* [ReplaceOrganizationApplicationRoles](docs/sdks/organizations/README.md#replaceorganizationapplicationroles) - Replace organization application roles
* [DeleteOrganizationApplicationRole](docs/sdks/organizations/README.md#deleteorganizationapplicationrole) - Remove organization application role
* [ListOrganizationJitEmailDomains](docs/sdks/organizations/README.md#listorganizationjitemaildomains) - Get organization JIT email domains
* [CreateOrganizationJitEmailDomain](docs/sdks/organizations/README.md#createorganizationjitemaildomain) - Add organization JIT email domain
* [ReplaceOrganizationJitEmailDomains](docs/sdks/organizations/README.md#replaceorganizationjitemaildomains) - Replace organization JIT email domains
* [DeleteOrganizationJitEmailDomain](docs/sdks/organizations/README.md#deleteorganizationjitemaildomain) - Remove organization JIT email domain
* [ListOrganizationJitRoles](docs/sdks/organizations/README.md#listorganizationjitroles) - Get organization JIT default roles
* [CreateOrganizationJitRole](docs/sdks/organizations/README.md#createorganizationjitrole) - Add organization JIT default roles
* [ReplaceOrganizationJitRoles](docs/sdks/organizations/README.md#replaceorganizationjitroles) - Replace organization JIT default roles
* [DeleteOrganizationJitRole](docs/sdks/organizations/README.md#deleteorganizationjitrole) - Remove organization JIT default role
* [ListOrganizationJitSsoConnectors](docs/sdks/organizations/README.md#listorganizationjitssoconnectors) - Get organization JIT SSO connectors
* [CreateOrganizationJitSsoConnector](docs/sdks/organizations/README.md#createorganizationjitssoconnector) - Add organization JIT SSO connectors
* [ReplaceOrganizationJitSsoConnectors](docs/sdks/organizations/README.md#replaceorganizationjitssoconnectors) - Replace organization JIT SSO connectors
* [DeleteOrganizationJitSsoConnector](docs/sdks/organizations/README.md#deleteorganizationjitssoconnector) - Remove organization JIT SSO connector

### [OrganizationScopes](docs/sdks/organizationscopes/README.md)

* [ListOrganizationScopes](docs/sdks/organizationscopes/README.md#listorganizationscopes) - Get organization scopes
* [CreateOrganizationScope](docs/sdks/organizationscopes/README.md#createorganizationscope) - Create an organization scope
* [GetOrganizationScope](docs/sdks/organizationscopes/README.md#getorganizationscope) - Get organization scope
* [UpdateOrganizationScope](docs/sdks/organizationscopes/README.md#updateorganizationscope) - Update organization scope
* [DeleteOrganizationScope](docs/sdks/organizationscopes/README.md#deleteorganizationscope) - Delete organization scope

### [Resources](docs/sdks/resources/README.md)

* [ListResources](docs/sdks/resources/README.md#listresources) - Get API resources
* [CreateResource](docs/sdks/resources/README.md#createresource) - Create an API resource
* [GetResource](docs/sdks/resources/README.md#getresource) - Get API resource
* [UpdateResource](docs/sdks/resources/README.md#updateresource) - Update API resource
* [DeleteResource](docs/sdks/resources/README.md#deleteresource) - Delete API resource
* [UpdateResourceIsDefault](docs/sdks/resources/README.md#updateresourceisdefault) - Set API resource as default
* [ListResourceScopes](docs/sdks/resources/README.md#listresourcescopes) - Get API resource scopes
* [CreateResourceScope](docs/sdks/resources/README.md#createresourcescope) - Create API resource scope
* [UpdateResourceScope](docs/sdks/resources/README.md#updateresourcescope) - Update API resource scope
* [DeleteResourceScope](docs/sdks/resources/README.md#deleteresourcescope) - Delete API resource scope

### [Roles](docs/sdks/roles/README.md)

* [ListRoles](docs/sdks/roles/README.md#listroles) - Get roles
* [CreateRole](docs/sdks/roles/README.md#createrole) - Create a role
* [GetRole](docs/sdks/roles/README.md#getrole) - Get role
* [UpdateRole](docs/sdks/roles/README.md#updaterole) - Update role
* [DeleteRole](docs/sdks/roles/README.md#deleterole) - Delete role
* [ListRoleUsers](docs/sdks/roles/README.md#listroleusers) - Get role users
* [CreateRoleUser](docs/sdks/roles/README.md#createroleuser) - Assign role to users
* [DeleteRoleUser](docs/sdks/roles/README.md#deleteroleuser) - Remove role from user
* [ListRoleApplications](docs/sdks/roles/README.md#listroleapplications) - Get role applications
* [CreateRoleApplication](docs/sdks/roles/README.md#createroleapplication) - Assign role to applications
* [DeleteRoleApplication](docs/sdks/roles/README.md#deleteroleapplication) - Remove role from application
* [ListRoleScopes](docs/sdks/roles/README.md#listrolescopes) - Get role scopes
* [CreateRoleScope](docs/sdks/roles/README.md#createrolescope) - Link scopes to role
* [DeleteRoleScope](docs/sdks/roles/README.md#deleterolescope) - Unlink scope from role

### [SAMLApplications](docs/sdks/samlapplications/README.md)

* [CreateSamlApplication](docs/sdks/samlapplications/README.md#createsamlapplication) - Create SAML application
* [GetSamlApplication](docs/sdks/samlapplications/README.md#getsamlapplication) - Get SAML application
* [UpdateSamlApplication](docs/sdks/samlapplications/README.md#updatesamlapplication) - Update SAML application
* [DeleteSamlApplication](docs/sdks/samlapplications/README.md#deletesamlapplication) - Delete SAML application
* [CreateSamlApplicationSecret](docs/sdks/samlapplications/README.md#createsamlapplicationsecret) - Create SAML application secret
* [ListSamlApplicationSecrets](docs/sdks/samlapplications/README.md#listsamlapplicationsecrets) - List SAML application secrets
* [DeleteSamlApplicationSecret](docs/sdks/samlapplications/README.md#deletesamlapplicationsecret) - Delete SAML application secret
* [UpdateSamlApplicationSecret](docs/sdks/samlapplications/README.md#updatesamlapplicationsecret) - Update SAML application secret
* [ListSamlApplicationMetadata](docs/sdks/samlapplications/README.md#listsamlapplicationmetadata) - Get SAML application metadata
* [GetSamlApplicationCallback](docs/sdks/samlapplications/README.md#getsamlapplicationcallback) - SAML application callback

### [SAMLApplicationsAuthFlow](docs/sdks/samlapplicationsauthflow/README.md)

* [GetSamlAuthn](docs/sdks/samlapplicationsauthflow/README.md#getsamlauthn) - Handle SAML authentication request (Redirect binding)
* [CreateSamlAuthn](docs/sdks/samlapplicationsauthflow/README.md#createsamlauthn) - Handle SAML authentication request (POST binding)

### [SentinelActivities](docs/sdks/sentinelactivities/README.md)

* [DeleteSentinelActivities](docs/sdks/sentinelactivities/README.md#deletesentinelactivities) - Bulk delete sentinel activities

### [SignInExperience](docs/sdks/signinexperience/README.md)

* [GetSignInExp](docs/sdks/signinexperience/README.md#getsigninexp) - Get default sign-in experience settings
* [UpdateSignInExp](docs/sdks/signinexperience/README.md#updatesigninexp) - Update default sign-in experience settings
* [CheckPasswordWithDefaultSignInExperience](docs/sdks/signinexperience/README.md#checkpasswordwithdefaultsigninexperience) - Check if a password meets the password policy
* [UploadCustomUIAssets](docs/sdks/signinexperience/README.md#uploadcustomuiassets) - Upload custom UI assets

### [SSOConnectorProviders](docs/sdks/ssoconnectorproviders/README.md)

* [ListSsoConnectorProviders](docs/sdks/ssoconnectorproviders/README.md#listssoconnectorproviders) - List all the supported SSO connector provider details

### [SSOConnectors](docs/sdks/ssoconnectors/README.md)

* [CreateSsoConnector](docs/sdks/ssoconnectors/README.md#createssoconnector) - Create SSO connector
* [ListSsoConnectors](docs/sdks/ssoconnectors/README.md#listssoconnectors) - List SSO connectors
* [GetSsoConnector](docs/sdks/ssoconnectors/README.md#getssoconnector) - Get SSO connector
* [DeleteSsoConnector](docs/sdks/ssoconnectors/README.md#deletessoconnector) - Delete SSO connector
* [UpdateSsoConnector](docs/sdks/ssoconnectors/README.md#updatessoconnector) - Update SSO connector

### [Status](docs/sdks/status/README.md)

* [GetStatus](docs/sdks/status/README.md#getstatus) - Health check

### [SubjectTokens](docs/sdks/subjecttokens/README.md)

* [CreateSubjectToken](docs/sdks/subjecttokens/README.md#createsubjecttoken) - Create a new subject token.

### [Swagger](docs/sdks/swagger/README.md)


#### [Swagger.JSON](docs/sdks/json/README.md)

* [GetSwaggerJSON](docs/sdks/json/README.md#getswaggerjson) - Get Swagger JSON

### [Systems](docs/sdks/systems/README.md)

* [GetSystemApplicationConfig](docs/sdks/systems/README.md#getsystemapplicationconfig) - Get the application constants.

### [UserAssets](docs/sdks/userassets/README.md)

* [GetUserAssetServiceStatus](docs/sdks/userassets/README.md#getuserassetservicestatus) - Get service status
* [CreateUserAsset](docs/sdks/userassets/README.md#createuserasset) - Upload asset

### [Users](docs/sdks/users/README.md)

* [GetUser](docs/sdks/users/README.md#getuser) - Get user
* [UpdateUser](docs/sdks/users/README.md#updateuser) - Update user
* [DeleteUser](docs/sdks/users/README.md#deleteuser) - Delete user
* [ListUserCustomData](docs/sdks/users/README.md#listusercustomdata) - Get user custom data
* [UpdateUserCustomData](docs/sdks/users/README.md#updateusercustomdata) - Update user custom data
* [UpdateUserProfile](docs/sdks/users/README.md#updateuserprofile) - Update user profile
* [CreateUser](docs/sdks/users/README.md#createuser) - Create user
* [ListUsers](docs/sdks/users/README.md#listusers) - Get users
* [UpdateUserPassword](docs/sdks/users/README.md#updateuserpassword) - Update user password
* [VerifyUserPassword](docs/sdks/users/README.md#verifyuserpassword) - Verify user password
* [GetUserHasPassword](docs/sdks/users/README.md#getuserhaspassword) - Check if user has password
* [UpdateUserIsSuspended](docs/sdks/users/README.md#updateuserissuspended) - Update user suspension status
* [ListUserRoles](docs/sdks/users/README.md#listuserroles) - Get roles for user
* [AssignUserRoles](docs/sdks/users/README.md#assignuserroles) - Assign roles to user
* [ReplaceUserRoles](docs/sdks/users/README.md#replaceuserroles) - Update roles for user
* [DeleteUserRole](docs/sdks/users/README.md#deleteuserrole) - Remove role from user
* [ReplaceUserIdentity](docs/sdks/users/README.md#replaceuseridentity) - Update social identity of user
* [DeleteUserIdentity](docs/sdks/users/README.md#deleteuseridentity) - Delete social identity from user
* [CreateUserIdentity](docs/sdks/users/README.md#createuseridentity) - Link social identity to user
* [ListUserOrganizations](docs/sdks/users/README.md#listuserorganizations) - Get organizations for a user
* [ListUserMfaVerifications](docs/sdks/users/README.md#listusermfaverifications) - Get user's MFA verifications
* [CreateUserMfaVerification](docs/sdks/users/README.md#createusermfaverification) - Create an MFA verification for a user
* [DeleteUserMfaVerification](docs/sdks/users/README.md#deleteusermfaverification) - Delete an MFA verification for a user
* [ListUserPersonalAccessTokens](docs/sdks/users/README.md#listuserpersonalaccesstokens) - Get personal access tokens
* [CreateUserPersonalAccessToken](docs/sdks/users/README.md#createuserpersonalaccesstoken) - Add personal access token
* [DeleteUserPersonalAccessToken](docs/sdks/users/README.md#deleteuserpersonalaccesstoken) - Delete personal access token
* [UpdateUserPersonalAccessToken](docs/sdks/users/README.md#updateuserpersonalaccesstoken) - Update personal access token

### [VerificationCodes](docs/sdks/verificationcodes/README.md)

* [CreateVerificationCode](docs/sdks/verificationcodes/README.md#createverificationcode) - Request and send a verification code
* [VerifyVerificationCode](docs/sdks/verificationcodes/README.md#verifyverificationcode) - Verify a verification code

### [Verifications](docs/sdks/verifications/README.md)

* [CreateVerificationByPassword](docs/sdks/verifications/README.md#createverificationbypassword) - Create a record by password
* [CreateVerificationByVerificationCode](docs/sdks/verifications/README.md#createverificationbyverificationcode) - Create a record by verification code
* [VerifyVerificationByVerificationCode](docs/sdks/verifications/README.md#verifyverificationbyverificationcode) - Verify verification code
* [CreateVerificationBySocial](docs/sdks/verifications/README.md#createverificationbysocial) - Create a social verification record
* [VerifyVerificationBySocial](docs/sdks/verifications/README.md#verifyverificationbysocial) - Verify a social verification record

### [WellKnown](docs/sdks/wellknown/README.md)

* [GetWellKnownEndpoint](docs/sdks/wellknown/README.md#getwellknownendpoint)
* [~~GetSignInExperienceConfig~~](docs/sdks/wellknown/README.md#getsigninexperienceconfig) - Get full sign-in experience :warning: **Deprecated**
* [GetSignInExperiencePhrases](docs/sdks/wellknown/README.md#getsigninexperiencephrases) - Get localized phrases
* [GetWellKnownExperience](docs/sdks/wellknown/README.md#getwellknownexperience) - Get full sign-in experience
* [GetWellKnownManagementOpenapiJSON](docs/sdks/wellknown/README.md#getwellknownmanagementopenapijson) - Get Management API swagger JSON
* [GetWellKnownExperienceOpenapiJSON](docs/sdks/wellknown/README.md#getwellknownexperienceopenapijson) - Get Experience API swagger JSON
* [GetWellKnownUserOpenapiJSON](docs/sdks/wellknown/README.md#getwellknownuseropenapijson) - Get User API swagger JSON

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"github.com/klucherev/logto-go-client/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"github.com/klucherev/logto-go-client/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListApplications` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| apierrors.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/apierrors"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithServerURL("https://logto.rentavita.com"),
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/klucherev/logto&utm_campaign=go)
