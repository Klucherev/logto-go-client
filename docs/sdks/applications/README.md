# Applications
(*Applications*)

## Overview

Application represents your registered software program or service that has been authorized to access user information and perform actions on behalf of users within the system. Currently, Logto supports four types of applications:

- Traditional web

- Single-page app
- Native app
- Machine-to-machine app.

Depending on the application type, it may have different authentication flows and access to the system. See [🔗 Integrate Logto in your application](https://docs.logto.io/docs/recipes/integrate-logto/) to learn more about how to integrate Logto into your application.

Role-based access control (RBAC) is supported for machine-to-machine applications. See [🔐 Role-based access control (RBAC)](https://docs.logto.io/docs/recipes/rbac/) to get started with role-based access control.

### Available Operations

* [ListApplications](#listapplications) - Get applications
* [CreateApplication](#createapplication) - Create an application
* [GetApplication](#getapplication) - Get application
* [UpdateApplication](#updateapplication) - Update application
* [DeleteApplication](#deleteapplication) - Delete application
* [UpdateApplicationCustomData](#updateapplicationcustomdata) - Update application custom data
* [ListApplicationRoles](#listapplicationroles) - Get application API resource roles
* [AssignApplicationRoles](#assignapplicationroles) - Assign API resource roles to application
* [ReplaceApplicationRoles](#replaceapplicationroles) - Update API resource roles for application
* [DeleteApplicationRole](#deleteapplicationrole) - Remove a API resource role from application
* [ListApplicationProtectedAppMetadataCustomDomains](#listapplicationprotectedappmetadatacustomdomains) - Get application custom domains.
* [CreateApplicationProtectedAppMetadataCustomDomain](#createapplicationprotectedappmetadatacustomdomain) - Add a custom domain to the application.
* [DeleteApplicationProtectedAppMetadataCustomDomain](#deleteapplicationprotectedappmetadatacustomdomain) - Remove custom domain.
* [ListApplicationOrganizations](#listapplicationorganizations) - Get application organizations
* [DeleteApplicationLegacySecret](#deleteapplicationlegacysecret) - Delete application legacy secret
* [ListApplicationSecrets](#listapplicationsecrets) - Get application secrets
* [CreateApplicationSecret](#createapplicationsecret) - Add application secret
* [DeleteApplicationSecret](#deleteapplicationsecret) - Delete application secret
* [UpdateApplicationSecret](#updateapplicationsecret) - Update application secret
* [CreateApplicationUserConsentScope](#createapplicationuserconsentscope) - Assign user consent scopes to application.
* [ListApplicationUserConsentScopes](#listapplicationuserconsentscopes) - List all the user consent scopes of an application.
* [DeleteApplicationUserConsentScope](#deleteapplicationuserconsentscope) - Remove user consent scope from application.
* [ReplaceApplicationSignInExperience](#replaceapplicationsigninexperience) - Update application level sign-in experience
* [GetApplicationSignInExperience](#getapplicationsigninexperience) - Get the application level sign-in experience
* [ListApplicationUserConsentOrganizations](#listapplicationuserconsentorganizations) - List all the user consented organizations of a application.
* [ReplaceApplicationUserConsentOrganizations](#replaceapplicationuserconsentorganizations) - Grant a list of organization access of a user for a application.
* [CreateApplicationUserConsentOrganization](#createapplicationuserconsentorganization) - Grant a list of organization access of a user for a application.
* [DeleteApplicationUserConsentOrganization](#deleteapplicationuserconsentorganization) - Revoke a user's access to an organization for a application.

## ListApplications

Get applications that match the given query with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListApplicationsRequest](../../models/operations/listapplicationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListApplicationsResponse](../../models/operations/listapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateApplication

Create a new application with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.CreateApplication(ctx, operations.CreateApplicationRequest{
        Name: "<value>",
        Type: operations.CreateApplicationTypeRequestSaml,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateApplicationRequest](../../models/operations/createapplicationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateApplicationResponse](../../models/operations/createapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetApplication

Get application details by ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.GetApplication(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetApplicationResponse](../../models/operations/getapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateApplication

Update application details by ID with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.UpdateApplication(ctx, "<id>", operations.UpdateApplicationRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | The unique identifier of the application.                                                          |
| `requestBody`                                                                                      | [operations.UpdateApplicationRequestBody](../../models/operations/updateapplicationrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateApplicationResponse](../../models/operations/updateapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplication

Delete application by ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplication(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationResponse](../../models/operations/deleteapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateApplicationCustomData

Update the custom data of an application.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.UpdateApplicationCustomData(ctx, "<id>", operations.UpdateApplicationCustomDataRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `applicationID`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The unique identifier of the application.                                                                              |
| `requestBody`                                                                                                          | [operations.UpdateApplicationCustomDataRequestBody](../../models/operations/updateapplicationcustomdatarequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.UpdateApplicationCustomDataResponseResponse](../../models/operations/updateapplicationcustomdataresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationRoles

Get API resource roles assigned to the specified application with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationRoles(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `searchParams`                                           | map[string]*string*                                      | :heavy_minus_sign:                                       | Search query parameters.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationRolesResponse](../../models/operations/listapplicationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignApplicationRoles

Assign API resource roles to the specified application. The API resource roles will be added to the existing API resource roles.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.AssignApplicationRoles(ctx, "<id>", operations.AssignApplicationRolesRequestBody{
        RoleIds: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `applicationID`                                                                                              | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the application.                                                                    |
| `requestBody`                                                                                                | [operations.AssignApplicationRolesRequestBody](../../models/operations/assignapplicationrolesrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AssignApplicationRolesResponse](../../models/operations/assignapplicationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceApplicationRoles

Update API resource roles assigned to the specified application. This will replace the existing API resource roles.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ReplaceApplicationRoles(ctx, "<id>", operations.ReplaceApplicationRolesRequestBody{
        RoleIds: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `applicationID`                                                                                                | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The unique identifier of the application.                                                                      |
| `requestBody`                                                                                                  | [operations.ReplaceApplicationRolesRequestBody](../../models/operations/replaceapplicationrolesrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ReplaceApplicationRolesResponse](../../models/operations/replaceapplicationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationRole

Remove a API resource role from the specified application.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationRole(ctx, "<id>", "<id>")
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
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the role.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationRoleResponse](../../models/operations/deleteapplicationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationProtectedAppMetadataCustomDomains

Get custom domains of the specified application, the application type should be protected app.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationProtectedAppMetadataCustomDomains(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationProtectedAppMetadataCustomDomainsResponse](../../models/operations/listapplicationprotectedappmetadatacustomdomainsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateApplicationProtectedAppMetadataCustomDomain

Add a custom domain to the application. You'll need to setup DNS record later.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.CreateApplicationProtectedAppMetadataCustomDomain(ctx, "<id>", operations.CreateApplicationProtectedAppMetadataCustomDomainRequestBody{
        Domain: "simple-dead.com",
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

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `id`                                                                                                                                                               | *string*                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                 | The unique identifier of the application.                                                                                                                          |
| `requestBody`                                                                                                                                                      | [operations.CreateApplicationProtectedAppMetadataCustomDomainRequestBody](../../models/operations/createapplicationprotectedappmetadatacustomdomainrequestbody.md) | :heavy_check_mark:                                                                                                                                                 | N/A                                                                                                                                                                |
| `opts`                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.CreateApplicationProtectedAppMetadataCustomDomainResponse](../../models/operations/createapplicationprotectedappmetadatacustomdomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationProtectedAppMetadataCustomDomain

Remove custom domain from the specified application.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationProtectedAppMetadataCustomDomain(ctx, "<id>", "clueless-armchair.biz")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `domain`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationProtectedAppMetadataCustomDomainResponse](../../models/operations/deleteapplicationprotectedappmetadatacustomdomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationOrganizations

Get the list of organizations that an application is associated with.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationOrganizations(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationOrganizationsResponse](../../models/operations/listapplicationorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationLegacySecret

Delete the legacy secret for the application and replace it with a new internal secret.

Note: This operation does not "really" delete the legacy secret because it is still needed for internal validation. We may remove the display of the legacy secret (the `secret` field in the application response) in the future.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationLegacySecret(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationLegacySecretResponse](../../models/operations/deleteapplicationlegacysecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationSecrets

Get all the secrets for the application.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationSecrets(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationSecretsResponse](../../models/operations/listapplicationsecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateApplicationSecret

Add a new secret for the application.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.CreateApplicationSecret(ctx, "<id>", operations.CreateApplicationSecretRequestBody{
        Name: "<value>",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `id`                                                                                                           | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The unique identifier of the application.                                                                      |
| `requestBody`                                                                                                  | [operations.CreateApplicationSecretRequestBody](../../models/operations/createapplicationsecretrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateApplicationSecretResponse](../../models/operations/createapplicationsecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationSecret

Delete a secret for the application by name.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationSecret(ctx, "<id>", "<value>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `name`                                                   | *string*                                                 | :heavy_check_mark:                                       | The name of the secret.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationSecretResponse](../../models/operations/deleteapplicationsecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateApplicationSecret

Update a secret for the application by name.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.UpdateApplicationSecret(ctx, "<id>", "<value>", operations.UpdateApplicationSecretRequestBody{
        Name: "<value>",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `id`                                                                                                           | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The unique identifier of the application.                                                                      |
| `name`                                                                                                         | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The name of the secret.                                                                                        |
| `requestBody`                                                                                                  | [operations.UpdateApplicationSecretRequestBody](../../models/operations/updateapplicationsecretrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateApplicationSecretResponse](../../models/operations/updateapplicationsecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateApplicationUserConsentScope

Assign the user consent scopes to an application by application id

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.CreateApplicationUserConsentScope(ctx, "<id>", operations.CreateApplicationUserConsentScopeRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `applicationID`                                                                                                                    | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | The unique identifier of the application.                                                                                          |
| `requestBody`                                                                                                                      | [operations.CreateApplicationUserConsentScopeRequestBody](../../models/operations/createapplicationuserconsentscoperequestbody.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateApplicationUserConsentScopeResponse](../../models/operations/createapplicationuserconsentscoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationUserConsentScopes

List all the user consent scopes of an application by application id

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationUserConsentScopes(ctx, "<id>")
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
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationUserConsentScopesResponse](../../models/operations/listapplicationuserconsentscopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationUserConsentScope

Remove the user consent scope from an application by application id, scope type and scope id

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationUserConsentScope(ctx, "<id>", operations.ScopeTypeOrganizationScopes, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `applicationID`                                              | *string*                                                     | :heavy_check_mark:                                           | The unique identifier of the application.                    |
| `scopeType`                                                  | [operations.ScopeType](../../models/operations/scopetype.md) | :heavy_check_mark:                                           | N/A                                                          |
| `scopeID`                                                    | *string*                                                     | :heavy_check_mark:                                           | The unique identifier of the scope.                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.DeleteApplicationUserConsentScopeResponse](../../models/operations/deleteapplicationuserconsentscoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceApplicationSignInExperience

Update application level sign-in experience for the specified application. Create a new sign-in experience if it does not exist. 
 - Only branding properties and terms links customization is supported for now. 

 - Only third-party applications can be customized for now. 

 - Application level sign-in experience customization is optional, if provided, it will override the default branding and terms links.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ReplaceApplicationSignInExperience(ctx, "<id>", operations.ReplaceApplicationSignInExperienceRequestBody{
        TermsOfUseURL: logto.String("https://failing-pigsty.net"),
        PrivacyPolicyURL: logto.String("https://minor-markup.name/"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `applicationID`                                                                                                                      | *string*                                                                                                                             | :heavy_check_mark:                                                                                                                   | The unique identifier of the application.                                                                                            |
| `requestBody`                                                                                                                        | [operations.ReplaceApplicationSignInExperienceRequestBody](../../models/operations/replaceapplicationsigninexperiencerequestbody.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ReplaceApplicationSignInExperienceResponse](../../models/operations/replaceapplicationsigninexperienceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetApplicationSignInExperience

Get application level sign-in experience for a given application. 
 - Only branding properties and terms links customization is supported for now. 

 - Only third-party applications can have the sign-in experience customization for now.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.GetApplicationSignInExperience(ctx, "<id>")
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
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetApplicationSignInExperienceResponse](../../models/operations/getapplicationsigninexperienceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListApplicationUserConsentOrganizations

List all the user consented organizations for a application by application id and user id.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ListApplicationUserConsentOrganizations(ctx, "<id>", "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListApplicationUserConsentOrganizationsResponse](../../models/operations/listapplicationuserconsentorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceApplicationUserConsentOrganizations

Grant a list of organization access of a user for a application by application id and user id. <br/> The user must be a member of all the organizations. <br/> Only third-party application needs to be granted access to organizations, all the other applications can request for all the organizations' access by default.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.ReplaceApplicationUserConsentOrganizations(ctx, "<id>", "<id>", operations.ReplaceApplicationUserConsentOrganizationsRequestBody{
        OrganizationIds: []string{
            "<value 1>",
        },
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `id`                                                                                                                                                 | *string*                                                                                                                                             | :heavy_check_mark:                                                                                                                                   | The unique identifier of the application.                                                                                                            |
| `userID`                                                                                                                                             | *string*                                                                                                                                             | :heavy_check_mark:                                                                                                                                   | The unique identifier of the user.                                                                                                                   |
| `requestBody`                                                                                                                                        | [operations.ReplaceApplicationUserConsentOrganizationsRequestBody](../../models/operations/replaceapplicationuserconsentorganizationsrequestbody.md) | :heavy_check_mark:                                                                                                                                   | N/A                                                                                                                                                  |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.ReplaceApplicationUserConsentOrganizationsResponse](../../models/operations/replaceapplicationuserconsentorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateApplicationUserConsentOrganization

Grant a list of organization access of a user for a application by application id and user id. <br/> The user must be a member of all the organizations. <br/> Only third-party application needs to be granted access to organizations, all the other applications can request for all the organizations' access by default.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.CreateApplicationUserConsentOrganization(ctx, "<id>", "<id>", operations.CreateApplicationUserConsentOrganizationRequestBody{
        OrganizationIds: []string{
            "<value 1>",
        },
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `id`                                                                                                                                             | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | The unique identifier of the application.                                                                                                        |
| `userID`                                                                                                                                         | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | The unique identifier of the user.                                                                                                               |
| `requestBody`                                                                                                                                    | [operations.CreateApplicationUserConsentOrganizationRequestBody](../../models/operations/createapplicationuserconsentorganizationrequestbody.md) | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.CreateApplicationUserConsentOrganizationResponse](../../models/operations/createapplicationuserconsentorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteApplicationUserConsentOrganization

Revoke a user's access to an organization for a application by application id, user id and organization id.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Applications.DeleteApplicationUserConsentOrganization(ctx, "<id>", "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteApplicationUserConsentOrganizationResponse](../../models/operations/deleteapplicationuserconsentorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |