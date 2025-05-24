# Experience
(*Experience*)

## Overview

The Experience endpoints allow end-users to interact with Logto for identity verification and profile completion.

### Available Operations

* [InitInteraction](#initinteraction) - Init new interaction
* [UpdateInteractionEvent](#updateinteractionevent) - Update interaction event
* [IdentifyUser](#identifyuser) - Identify user for the current interaction
* [SubmitInteraction](#submitinteraction) - Submit interaction
* [CreatePasswordVerification](#createpasswordverification) - Create password verification record
* [CreateAndSendVerificationCode](#createandsendverificationcode) - Create and send verification code
* [VerifyVerificationCodeVerification](#verifyverificationcodeverification) - Verify verification code
* [CreateSocialVerification](#createsocialverification) - Create social verification
* [VerifySocialVerification](#verifysocialverification) - Verify social verification
* [CreateEnterpriseSsoVerification](#createenterprisessoverification) - Create enterprise SSO verification
* [VerifyEnterpriseSsoVerification](#verifyenterprisessoverification) - Verify enterprise SSO verification
* [CreateTotpSecret](#createtotpsecret) - Create TOTP secret
* [VerifyTotpVerification](#verifytotpverification) - Verify TOTP verification
* [CreateWebAuthnRegistrationVerification](#createwebauthnregistrationverification) - Create WebAuthn registration verification
* [VerifyWebAuthnRegistrationVerification](#verifywebauthnregistrationverification) - Verify WebAuthn registration verification
* [CreateWebAuthnAuthenticationVerification](#createwebauthnauthenticationverification) - Create WebAuthn authentication verification
* [VerifyWebAuthnAuthenticationVerification](#verifywebauthnauthenticationverification) - Verify WebAuthn authentication verification
* [GenerateBackupCodes](#generatebackupcodes) - Generate backup codes
* [VerifyBackupCode](#verifybackupcode) - Verify backup code
* [CreateNewPasswordIdentityVerification](#createnewpasswordidentityverification) - Create new password identity verification
* [VerifyOneTimeTokenVerification](#verifyonetimetokenverification) - Verify one-time token
* [AddUserProfile](#adduserprofile) - Add user profile
* [ResetUserPassword](#resetuserpassword) - Reset user password
* [SkipMfaBindingFlow](#skipmfabindingflow) - Skip MFA binding flow
* [BindMfaVerification](#bindmfaverification) - Bind MFA verification by verificationId
* [GetEnabledSsoConnectors](#getenabledssoconnectors) - Get enabled SSO connectors by the given email's domain

## InitInteraction

Init a new experience interaction with the given interaction type. Any existing interaction data will be cleared.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.InitInteraction(ctx, operations.InitInteractionRequest{
        InteractionEvent: operations.InitInteractionInteractionEventForgotPassword,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.InitInteractionRequest](../../models/operations/initinteractionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.InitInteractionResponse](../../models/operations/initinteractionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateInteractionEvent

Update the current experience interaction event to the given event type. This API is used to switch the interaction event between `SignIn` and `Register`, while keeping all the verification records data.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.UpdateInteractionEvent(ctx, operations.UpdateInteractionEventRequest{
        InteractionEvent: operations.UpdateInteractionEventInteractionEventForgotPassword,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateInteractionEventRequest](../../models/operations/updateinteractioneventrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateInteractionEventResponse](../../models/operations/updateinteractioneventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## IdentifyUser

This API identifies the user based on the verificationId within the current experience interaction: <br/>- `SignIn` and `ForgotPassword` interactions: Verifies the user's identity using the provided `verificationId`. <br/>- `Register` interaction: Creates a new user account using the profile data from the current interaction. If a verificationId is provided, the profile data will first be updated with the verification record before creating the account. If not, the account is created directly from the stored profile data.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.IdentifyUser(ctx, operations.IdentifyUserRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.IdentifyUserRequest](../../models/operations/identifyuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.IdentifyUserResponse](../../models/operations/identifyuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SubmitInteraction

Submit the current interaction. <br/>- Submit the verified user identity to the OIDC provider for further authentication (SignIn and Register). <br/>- Update the user's profile data if any (SignIn and Register). <br/>- Reset the password and clear all the interaction records (ForgotPassword).

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.SubmitInteraction(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubmitInteractionResponse](../../models/operations/submitinteractionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreatePasswordVerification

Create and verify a new Password verification record. The verification record can only be created if the provided user credentials are correct.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreatePasswordVerification(ctx, operations.CreatePasswordVerificationRequest{
        Identifier: operations.CreatePasswordVerificationIdentifier{
            Type: operations.CreatePasswordVerificationTypePhone,
            Value: "<value>",
        },
        Password: "Fud1_1gTN5Y96b1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreatePasswordVerificationRequest](../../models/operations/createpasswordverificationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreatePasswordVerificationResponse](../../models/operations/createpasswordverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAndSendVerificationCode

Create a new `CodeVerification` record and sends the code to the specified identifier. The code verification can be used to verify the given identifier.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateAndSendVerificationCode(ctx, operations.CreateAndSendVerificationCodeRequest{
        Identifier: operations.CreateAndSendVerificationCodeIdentifier{
            Type: operations.CreateAndSendVerificationCodeTypeEmail,
            Value: "<value>",
        },
        InteractionEvent: operations.CreateAndSendVerificationCodeInteractionEventRegister,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateAndSendVerificationCodeRequest](../../models/operations/createandsendverificationcoderequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateAndSendVerificationCodeResponse](../../models/operations/createandsendverificationcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyVerificationCodeVerification

Verify the provided verification code against the user's identifier. If successful, the verification record will be marked as verified.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyVerificationCodeVerification(ctx, operations.VerifyVerificationCodeVerificationRequest{
        Identifier: operations.VerifyVerificationCodeVerificationIdentifier{
            Type: operations.VerifyVerificationCodeVerificationTypePhone,
            Value: "<value>",
        },
        VerificationID: "<id>",
        Code: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.VerifyVerificationCodeVerificationRequest](../../models/operations/verifyverificationcodeverificationrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.VerifyVerificationCodeVerificationResponse](../../models/operations/verifyverificationcodeverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSocialVerification

Create a new SocialVerification record and return the provider's authorization URI for the given connector.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateSocialVerification(ctx, "<id>", operations.CreateSocialVerificationRequestBody{
        State: "Pennsylvania",
        RedirectURI: "https://limping-zen.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `connectorID`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The unique identifier of the connector.                                                                          |
| `requestBody`                                                                                                    | [operations.CreateSocialVerificationRequestBody](../../models/operations/createsocialverificationrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateSocialVerificationResponse](../../models/operations/createsocialverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifySocialVerification

Verify the social authorization response data and get the user's identity data from the social provider.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifySocialVerification(ctx, "<id>", operations.VerifySocialVerificationRequestBody{
        ConnectorData: operations.VerifySocialVerificationConnectorData{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `connectorID`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The unique identifier of the connector.                                                                          |
| `requestBody`                                                                                                    | [operations.VerifySocialVerificationRequestBody](../../models/operations/verifysocialverificationrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.VerifySocialVerificationResponse](../../models/operations/verifysocialverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateEnterpriseSsoVerification

Create a new EnterpriseSSO verification record and return the provider's authorization URI for the given connector.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateEnterpriseSsoVerification(ctx, "<id>", operations.CreateEnterpriseSsoVerificationRequestBody{
        State: "Montana",
        RedirectURI: "https://uncommon-spirit.org",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `connectorID`                                                                                                                  | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The unique identifier of the connector.                                                                                        |
| `requestBody`                                                                                                                  | [operations.CreateEnterpriseSsoVerificationRequestBody](../../models/operations/createenterprisessoverificationrequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.CreateEnterpriseSsoVerificationResponse](../../models/operations/createenterprisessoverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyEnterpriseSsoVerification

Verify the SSO authorization response data and get the user's identity from the SSO provider.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyEnterpriseSsoVerification(ctx, "<id>", operations.VerifyEnterpriseSsoVerificationRequestBody{
        ConnectorData: operations.VerifyEnterpriseSsoVerificationConnectorData{},
        VerificationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `connectorID`                                                                                                                  | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The unique identifier of the connector.                                                                                        |
| `requestBody`                                                                                                                  | [operations.VerifyEnterpriseSsoVerificationRequestBody](../../models/operations/verifyenterprisessoverificationrequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.VerifyEnterpriseSsoVerificationResponse](../../models/operations/verifyenterprisessoverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateTotpSecret

Create a new TOTP verification record and generate a new TOTP secret for the user. This secret can be used to bind a new TOTP verification to the user's profile. The verification record must be verified before the secret can be used to bind a new TOTP verification to the user's profile.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateTotpSecret(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateTotpSecretResponse](../../models/operations/createtotpsecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyTotpVerification

Verifies the provided TOTP code against the new created TOTP secret or the existing TOTP secret. If a verificationId is provided, this API will verify the code against the TOTP secret that is associated with the verification record. Otherwise, a new TOTP verification record will be created and verified against the user's existing TOTP secret.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyTotpVerification(ctx, operations.VerifyTotpVerificationRequest{
        Code: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.VerifyTotpVerificationRequest](../../models/operations/verifytotpverificationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.VerifyTotpVerificationResponse](../../models/operations/verifytotpverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateWebAuthnRegistrationVerification

Create a new WebAuthn registration verification record. The verification record can be used to bind a new WebAuthn credential to the user's profile.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateWebAuthnRegistrationVerification(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateWebAuthnRegistrationVerificationResponse](../../models/operations/createwebauthnregistrationverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyWebAuthnRegistrationVerification

Verify the WebAuthn registration response against the user's WebAuthn registration challenge. If the response is valid, the WebAuthn registration record will be marked as verified.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyWebAuthnRegistrationVerification(ctx, operations.VerifyWebAuthnRegistrationVerificationRequest{
        VerificationID: "<id>",
        Payload: operations.VerifyWebAuthnRegistrationVerificationPayload{
            Type: "<value>",
            ID: "<id>",
            RawID: "<id>",
            Response: operations.ExperienceVerifyWebAuthnRegistrationVerificationResponse{
                ClientDataJSON: "<value>",
                AttestationObject: "<value>",
            },
            ClientExtensionResults: operations.VerifyWebAuthnRegistrationVerificationClientExtensionResults{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.VerifyWebAuthnRegistrationVerificationRequest](../../models/operations/verifywebauthnregistrationverificationrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.VerifyWebAuthnRegistrationVerificationResponseResponse](../../models/operations/verifywebauthnregistrationverificationresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateWebAuthnAuthenticationVerification

Create a new WebAuthn authentication verification record based on the user's existing WebAuthn credential. This verification record can be used to verify the user's WebAuthn credential.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateWebAuthnAuthenticationVerification(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateWebAuthnAuthenticationVerificationResponse](../../models/operations/createwebauthnauthenticationverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyWebAuthnAuthenticationVerification

Verifies the WebAuthn authentication response against the user's authentication challenge. Upon successful verification, the verification record will be marked as verified.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyWebAuthnAuthenticationVerification(ctx, operations.VerifyWebAuthnAuthenticationVerificationRequest{
        VerificationID: "<id>",
        Payload: operations.VerifyWebAuthnAuthenticationVerificationPayload{
            Type: "<value>",
            ID: "<id>",
            RawID: "<id>",
            ClientExtensionResults: operations.VerifyWebAuthnAuthenticationVerificationClientExtensionResults{},
            Response: operations.ExperienceVerifyWebAuthnAuthenticationVerificationResponse{
                ClientDataJSON: "<value>",
                AuthenticatorData: "<value>",
                Signature: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.VerifyWebAuthnAuthenticationVerificationRequest](../../models/operations/verifywebauthnauthenticationverificationrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.VerifyWebAuthnAuthenticationVerificationResponseResponse](../../models/operations/verifywebauthnauthenticationverificationresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GenerateBackupCodes

Create a new BackupCode verification record with new backup codes generated. This verification record will be used to bind the backup codes to the user's profile.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.GenerateBackupCodes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GenerateBackupCodesResponse](../../models/operations/generatebackupcodesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyBackupCode

Create a new BackupCode verification record and verify the provided backup code against the user's backup codes. The verification record will be marked as verified if the code is correct.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyBackupCode(ctx, operations.VerifyBackupCodeRequest{
        Code: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.VerifyBackupCodeRequest](../../models/operations/verifybackupcoderequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.VerifyBackupCodeResponse](../../models/operations/verifybackupcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateNewPasswordIdentityVerification

Create a NewPasswordIdentity verification record for the new user registration use. The verification record includes a unique user identifier and a password that can be used to create a new user account.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.CreateNewPasswordIdentityVerification(ctx, operations.CreateNewPasswordIdentityVerificationRequest{
        Identifier: operations.CreateNewPasswordIdentityVerificationIdentifier{
            Type: "<value>",
            Value: "<value>",
        },
        Password: "6YRrbwdzv3yqVT4",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreateNewPasswordIdentityVerificationRequest](../../models/operations/createnewpasswordidentityverificationrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateNewPasswordIdentityVerificationResponse](../../models/operations/createnewpasswordidentityverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyOneTimeTokenVerification

Verify the provided one-time token against the user's email. If successful, the verification record will be marked as verified.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.VerifyOneTimeTokenVerification(ctx, operations.VerifyOneTimeTokenVerificationRequest{
        Identifier: operations.VerifyOneTimeTokenVerificationIdentifier{
            Type: "<value>",
            Value: "<value>",
        },
        Token: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.VerifyOneTimeTokenVerificationRequest](../../models/operations/verifyonetimetokenverificationrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.VerifyOneTimeTokenVerificationResponse](../../models/operations/verifyonetimetokenverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddUserProfile

Adds user profile data to the current experience interaction. <br/>- For `Register`: The profile data provided before the identification request will be used to create a new user account. <br/>- For `SignIn` and `Register`: The profile data provided after the user is identified will be used to update the user's profile when the interaction is submitted. <br/>- `ForgotPassword`: Not supported.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.AddUserProfile(ctx, operations.CreateAddUserProfileRequestAddUserProfileRequestBody3(
        operations.AddUserProfileRequestBody3{
            Type: "<value>",
            VerificationID: "<id>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.AddUserProfileRequest](../../models/operations/adduserprofilerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AddUserProfileResponse](../../models/operations/adduserprofileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ResetUserPassword

Reset the user's password. (`ForgotPassword` interaction only)

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.ResetUserPassword(ctx, operations.ResetUserPasswordRequest{
        Password: "VdpWl9R0tzVaIQF",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ResetUserPasswordRequest](../../models/operations/resetuserpasswordrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ResetUserPasswordResponse](../../models/operations/resetuserpasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SkipMfaBindingFlow

Skip MFA verification binding flow. If the MFA is enabled in the sign-in experience settings and marked as `UserControlled`, the user can skip the MFA verification binding flow by calling this API.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.SkipMfaBindingFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SkipMfaBindingFlowResponse](../../models/operations/skipmfabindingflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## BindMfaVerification

Bind new MFA verification to the user profile using the verificationId.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.BindMfaVerification(ctx, operations.BindMfaVerificationRequest{
        Type: operations.BindMfaVerificationTypeWebAuthn,
        VerificationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.BindMfaVerificationRequest](../../models/operations/bindmfaverificationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.BindMfaVerificationResponse](../../models/operations/bindmfaverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetEnabledSsoConnectors

Extract the email domain from the provided email address. Returns all the enabled SSO connectors that match the email domain.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Experience.GetEnabledSsoConnectors(ctx, "Aida.Mertz43@gmail.com")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `email`                                                  | *string*                                                 | :heavy_check_mark:                                       | The email address to find the enabled SSO connectors.    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetEnabledSsoConnectorsResponse](../../models/operations/getenabledssoconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |