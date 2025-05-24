# Organizations
(*Organizations*)

## Overview

Organization is a concept that brings together multiple identities (mostly users). Logto supports multiple organizations, and each organization can have multiple users.

Every organization shares the same set (organization template) of roles and permissions. Each user can have different roles in different organizations. See [🏢 Organizations (Multi-tenancy)](https://docs.logto.io/docs/recipes/organizations/) to get started with organizations and organization template.

### Available Operations

* [CreateOrganization](#createorganization) - Create an organization
* [ListOrganizations](#listorganizations) - Get organizations
* [GetOrganization](#getorganization) - Get organization
* [UpdateOrganization](#updateorganization) - Update organization
* [DeleteOrganization](#deleteorganization) - Delete organization
* [AddOrganizationUsers](#addorganizationusers) - Add user members to organization
* [ReplaceOrganizationUsers](#replaceorganizationusers) - Replace organization user members
* [ListOrganizationUsers](#listorganizationusers) - Get organization user members
* [DeleteOrganizationUser](#deleteorganizationuser) - Remove user member from organization
* [AssignOrganizationRolesToUsers](#assignorganizationrolestousers) - Assign roles to organization user members
* [ListOrganizationUserRoles](#listorganizationuserroles) - Get roles for a user in an organization
* [AssignOrganizationRolesToUser](#assignorganizationrolestouser) - Assign roles to a user in an organization
* [ReplaceOrganizationUserRoles](#replaceorganizationuserroles) - Update roles for a user in an organization
* [DeleteOrganizationUserRole](#deleteorganizationuserrole) - Remove a role from a user in an organization
* [ListOrganizationUserScopes](#listorganizationuserscopes) - Get scopes for a user in an organization tailored by the organization roles
* [AddOrganizationApplications](#addorganizationapplications) - Add organization application
* [ReplaceOrganizationApplications](#replaceorganizationapplications) - Replace organization applications
* [ListOrganizationApplications](#listorganizationapplications) - Get organization applications
* [DeleteOrganizationApplication](#deleteorganizationapplication) - Remove organization application
* [AssignOrganizationRolesToApplications](#assignorganizationrolestoapplications) - Assign roles to applications in an organization
* [ListOrganizationApplicationRoles](#listorganizationapplicationroles) - Get organization application roles
* [AssignOrganizationRolesToApplication](#assignorganizationrolestoapplication) - Add organization application role
* [ReplaceOrganizationApplicationRoles](#replaceorganizationapplicationroles) - Replace organization application roles
* [DeleteOrganizationApplicationRole](#deleteorganizationapplicationrole) - Remove organization application role
* [ListOrganizationJitEmailDomains](#listorganizationjitemaildomains) - Get organization JIT email domains
* [CreateOrganizationJitEmailDomain](#createorganizationjitemaildomain) - Add organization JIT email domain
* [ReplaceOrganizationJitEmailDomains](#replaceorganizationjitemaildomains) - Replace organization JIT email domains
* [DeleteOrganizationJitEmailDomain](#deleteorganizationjitemaildomain) - Remove organization JIT email domain
* [ListOrganizationJitRoles](#listorganizationjitroles) - Get organization JIT default roles
* [CreateOrganizationJitRole](#createorganizationjitrole) - Add organization JIT default roles
* [ReplaceOrganizationJitRoles](#replaceorganizationjitroles) - Replace organization JIT default roles
* [DeleteOrganizationJitRole](#deleteorganizationjitrole) - Remove organization JIT default role
* [ListOrganizationJitSsoConnectors](#listorganizationjitssoconnectors) - Get organization JIT SSO connectors
* [CreateOrganizationJitSsoConnector](#createorganizationjitssoconnector) - Add organization JIT SSO connectors
* [ReplaceOrganizationJitSsoConnectors](#replaceorganizationjitssoconnectors) - Replace organization JIT SSO connectors
* [DeleteOrganizationJitSsoConnector](#deleteorganizationjitssoconnector) - Remove organization JIT SSO connector

## CreateOrganization

Create a new organization with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.CreateOrganization(ctx, operations.CreateOrganizationRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateOrganizationRequest](../../models/operations/createorganizationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateOrganizationResponse](../../models/operations/createorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizations

Get organizations that match the given query with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizations(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                       | Type                                                                                                                                                                                            | Required                                                                                                                                                                                        | Description                                                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                              | The context to use for the request.                                                                                                                                                             |
| `q`                                                                                                                                                                                             | **string*                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                              | The query to filter organizations. It can be a partial ID or name.<br/><br/>If not provided, all organizations will be returned.                                                                |
| `showFeatured`                                                                                                                                                                                  | **string*                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                              | Whether to show featured users in the organization. Featured users are randomly selected from the organization members.<br/><br/>If not provided, `featuredUsers` will not be included in the response. |
| `page`                                                                                                                                                                                          | **int64*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                              | Page number (starts from 1).                                                                                                                                                                    |
| `pageSize`                                                                                                                                                                                      | **int64*                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                              | Entries per page.                                                                                                                                                                               |
| `opts`                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                              | The options for this request.                                                                                                                                                                   |

### Response

**[*operations.ListOrganizationsResponse](../../models/operations/listorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetOrganization

Get organization details by ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.GetOrganization(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOrganizationResponse](../../models/operations/getorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateOrganization

Update organization details by ID with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.UpdateOrganization(ctx, "<id>", operations.UpdateOrganizationRequestBody{})
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
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | The unique identifier of the organization.                                                           |
| `requestBody`                                                                                        | [operations.UpdateOrganizationRequestBody](../../models/operations/updateorganizationrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateOrganizationResponse](../../models/operations/updateorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganization

Delete organization by ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganization(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationResponse](../../models/operations/deleteorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddOrganizationUsers

Add users as members to the specified organization with the given user IDs.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AddOrganizationUsers(ctx, "<id>", operations.AddOrganizationUsersRequestBody{
        UserIds: []string{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The unique identifier of the organization.                                                               |
| `requestBody`                                                                                            | [operations.AddOrganizationUsersRequestBody](../../models/operations/addorganizationusersrequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AddOrganizationUsersResponse](../../models/operations/addorganizationusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationUsers

Replace all user members for the specified organization with the given users. This effectively removing all existing user memberships in the organization and adding the new users as members.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationUsers(ctx, "<id>", operations.ReplaceOrganizationUsersRequestBody{
        UserIds: []string{
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `id`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The unique identifier of the organization.                                                                       |
| `requestBody`                                                                                                    | [operations.ReplaceOrganizationUsersRequestBody](../../models/operations/replaceorganizationusersrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ReplaceOrganizationUsersResponse](../../models/operations/replaceorganizationusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationUsers

Get users that are members of the specified organization for the given query with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationUsers(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | The context to use for the request.                                                                                                                                     |
| `id`                                                                                                                                                                    | *string*                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                      | The unique identifier of the organization.                                                                                                                              |
| `q`                                                                                                                                                                     | **string*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                      | The query to filter users. It will match multiple fields of users, including ID, name, username, email, and phone number.<br/><br/>If not provided, all users will be returned. |
| `page`                                                                                                                                                                  | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Page number (starts from 1).                                                                                                                                            |
| `pageSize`                                                                                                                                                              | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Entries per page.                                                                                                                                                       |
| `opts`                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | The options for this request.                                                                                                                                           |

### Response

**[*operations.ListOrganizationUsersResponse](../../models/operations/listorganizationusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationUser

Remove a user's membership from the specified organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationUser(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationUserResponse](../../models/operations/deleteorganizationuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignOrganizationRolesToUsers

Assign roles to user members of the specified organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AssignOrganizationRolesToUsers(ctx, "<id>", operations.AssignOrganizationRolesToUsersRequestBody{
        UserIds: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        OrganizationRoleIds: []string{
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `id`                                                                                                                         | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The unique identifier of the organization.                                                                                   |
| `requestBody`                                                                                                                | [operations.AssignOrganizationRolesToUsersRequestBody](../../models/operations/assignorganizationrolestousersrequestbody.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.AssignOrganizationRolesToUsersResponse](../../models/operations/assignorganizationrolestousersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationUserRoles

Get roles assigned to a user in the specified organization with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationUserRoles(ctx, "<id>", "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationUserRolesResponse](../../models/operations/listorganizationuserrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignOrganizationRolesToUser

Assign roles to a user in the specified organization with the provided data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AssignOrganizationRolesToUser(ctx, "<id>", "<id>", operations.AssignOrganizationRolesToUserRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `id`                                                                                                                       | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the organization.                                                                                 |
| `userID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the user.                                                                                         |
| `requestBody`                                                                                                              | [operations.AssignOrganizationRolesToUserRequestBody](../../models/operations/assignorganizationrolestouserrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.AssignOrganizationRolesToUserResponse](../../models/operations/assignorganizationrolestouserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationUserRoles

Update roles assigned to a user in the specified organization with the provided data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationUserRoles(ctx, "<id>", "<id>", operations.ReplaceOrganizationUserRolesRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `id`                                                                                                                     | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The unique identifier of the organization.                                                                               |
| `userID`                                                                                                                 | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The unique identifier of the user.                                                                                       |
| `requestBody`                                                                                                            | [operations.ReplaceOrganizationUserRolesRequestBody](../../models/operations/replaceorganizationuserrolesrequestbody.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.ReplaceOrganizationUserRolesResponse](../../models/operations/replaceorganizationuserrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationUserRole

Remove a role assignment from a user in the specified organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationUserRole(ctx, "<id>", "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `organizationRoleID`                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationUserRoleResponse](../../models/operations/deleteorganizationuserroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationUserScopes

Get scopes assigned to a user in the specified organization tailored by the organization roles. The scopes are derived from the organization roles assigned to the user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationUserScopes(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationUserScopesResponse](../../models/operations/listorganizationuserscopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddOrganizationApplications

Add an application to the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AddOrganizationApplications(ctx, "<id>", operations.AddOrganizationApplicationsRequestBody{
        ApplicationIds: []string{
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The unique identifier of the organization.                                                                             |
| `requestBody`                                                                                                          | [operations.AddOrganizationApplicationsRequestBody](../../models/operations/addorganizationapplicationsrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AddOrganizationApplicationsResponse](../../models/operations/addorganizationapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationApplications

Replace all applications associated with the organization with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationApplications(ctx, "<id>", operations.ReplaceOrganizationApplicationsRequestBody{
        ApplicationIds: []string{},
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `id`                                                                                                                           | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The unique identifier of the organization.                                                                                     |
| `requestBody`                                                                                                                  | [operations.ReplaceOrganizationApplicationsRequestBody](../../models/operations/replaceorganizationapplicationsrequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ReplaceOrganizationApplicationsResponse](../../models/operations/replaceorganizationapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationApplications

Get applications associated with the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationApplications(ctx, "<id>", nil, nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `q`                                                      | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationApplicationsResponse](../../models/operations/listorganizationapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationApplication

Remove an application from the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationApplication(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationApplicationResponse](../../models/operations/deleteorganizationapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignOrganizationRolesToApplications

Assign roles to applications in the specified organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AssignOrganizationRolesToApplications(ctx, "<id>", operations.AssignOrganizationRolesToApplicationsRequestBody{
        ApplicationIds: []string{},
        OrganizationRoleIds: []string{},
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `id`                                                                                                                                       | *string*                                                                                                                                   | :heavy_check_mark:                                                                                                                         | The unique identifier of the organization.                                                                                                 |
| `requestBody`                                                                                                                              | [operations.AssignOrganizationRolesToApplicationsRequestBody](../../models/operations/assignorganizationrolestoapplicationsrequestbody.md) | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.AssignOrganizationRolesToApplicationsResponse](../../models/operations/assignorganizationrolestoapplicationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationApplicationRoles

Get roles associated with the application in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationApplicationRoles(ctx, "<id>", "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationApplicationRolesResponse](../../models/operations/listorganizationapplicationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignOrganizationRolesToApplication

Add a role to the application in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.AssignOrganizationRolesToApplication(ctx, "<id>", "<id>", operations.AssignOrganizationRolesToApplicationRequestBody{
        OrganizationRoleIds: []string{
            "<value 1>",
            "<value 2>",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `id`                                                                                                                                     | *string*                                                                                                                                 | :heavy_check_mark:                                                                                                                       | The unique identifier of the organization.                                                                                               |
| `applicationID`                                                                                                                          | *string*                                                                                                                                 | :heavy_check_mark:                                                                                                                       | The unique identifier of the application.                                                                                                |
| `requestBody`                                                                                                                            | [operations.AssignOrganizationRolesToApplicationRequestBody](../../models/operations/assignorganizationrolestoapplicationrequestbody.md) | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.AssignOrganizationRolesToApplicationResponse](../../models/operations/assignorganizationrolestoapplicationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationApplicationRoles

Replace all roles associated with the application in the organization with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationApplicationRoles(ctx, "<id>", "<id>", operations.ReplaceOrganizationApplicationRolesRequestBody{
        OrganizationRoleIds: []string{},
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The unique identifier of the organization.                                                                                             |
| `applicationID`                                                                                                                        | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The unique identifier of the application.                                                                                              |
| `requestBody`                                                                                                                          | [operations.ReplaceOrganizationApplicationRolesRequestBody](../../models/operations/replaceorganizationapplicationrolesrequestbody.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.ReplaceOrganizationApplicationRolesResponse](../../models/operations/replaceorganizationapplicationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationApplicationRole

Remove a role from the application in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationApplicationRole(ctx, "<id>", "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `applicationID`                                          | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the application.                |
| `organizationRoleID`                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationApplicationRoleResponse](../../models/operations/deleteorganizationapplicationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationJitEmailDomains

Get email domains for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationJitEmailDomains(ctx, "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationJitEmailDomainsResponse](../../models/operations/listorganizationjitemaildomainsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateOrganizationJitEmailDomain

Add a new email domain for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.CreateOrganizationJitEmailDomain(ctx, "<id>", operations.CreateOrganizationJitEmailDomainRequestBody{
        EmailDomain: "<value>",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `id`                                                                                                                             | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The unique identifier of the organization.                                                                                       |
| `requestBody`                                                                                                                    | [operations.CreateOrganizationJitEmailDomainRequestBody](../../models/operations/createorganizationjitemaildomainrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.CreateOrganizationJitEmailDomainResponse](../../models/operations/createorganizationjitemaildomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationJitEmailDomains

Replace all just-in-time provisioning email domains for the organization with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationJitEmailDomains(ctx, "<id>", operations.ReplaceOrganizationJitEmailDomainsRequestBody{
        EmailDomains: []string{
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `id`                                                                                                                                 | *string*                                                                                                                             | :heavy_check_mark:                                                                                                                   | The unique identifier of the organization.                                                                                           |
| `requestBody`                                                                                                                        | [operations.ReplaceOrganizationJitEmailDomainsRequestBody](../../models/operations/replaceorganizationjitemaildomainsrequestbody.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ReplaceOrganizationJitEmailDomainsResponse](../../models/operations/replaceorganizationjitemaildomainsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationJitEmailDomain

Remove an email domain for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationJitEmailDomain(ctx, "<id>", "<value>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `emailDomain`                                            | *string*                                                 | :heavy_check_mark:                                       | The email domain to remove.                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationJitEmailDomainResponse](../../models/operations/deleteorganizationjitemaildomainresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationJitRoles

Get organization roles that will be assigned to users during just-in-time provisioning.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationJitRoles(ctx, "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationJitRolesResponse](../../models/operations/listorganizationjitrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateOrganizationJitRole

Add new organization roles that will be assigned to users during just-in-time provisioning.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.CreateOrganizationJitRole(ctx, "<id>", operations.CreateOrganizationJitRoleRequestBody{
        OrganizationRoleIds: []string{},
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The unique identifier of the organization.                                                                         |
| `requestBody`                                                                                                      | [operations.CreateOrganizationJitRoleRequestBody](../../models/operations/createorganizationjitrolerequestbody.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateOrganizationJitRoleResponse](../../models/operations/createorganizationjitroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationJitRoles

Replace all organization roles that will be assigned to users during just-in-time provisioning with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationJitRoles(ctx, "<id>", operations.ReplaceOrganizationJitRolesRequestBody{
        OrganizationRoleIds: []string{
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The unique identifier of the organization.                                                                             |
| `requestBody`                                                                                                          | [operations.ReplaceOrganizationJitRolesRequestBody](../../models/operations/replaceorganizationjitrolesrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ReplaceOrganizationJitRolesResponse](../../models/operations/replaceorganizationjitrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationJitRole

Remove an organization role that will be assigned to users during just-in-time provisioning.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationJitRole(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `organizationRoleID`                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationJitRoleResponse](../../models/operations/deleteorganizationjitroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizationJitSsoConnectors

Get enterprise SSO connectors for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ListOrganizationJitSsoConnectors(ctx, "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationJitSsoConnectorsResponse](../../models/operations/listorganizationjitssoconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateOrganizationJitSsoConnector

Add new enterprise SSO connectors for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.CreateOrganizationJitSsoConnector(ctx, "<id>", operations.CreateOrganizationJitSsoConnectorRequestBody{
        SsoConnectorIds: []string{},
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `id`                                                                                                                               | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | The unique identifier of the organization.                                                                                         |
| `requestBody`                                                                                                                      | [operations.CreateOrganizationJitSsoConnectorRequestBody](../../models/operations/createorganizationjitssoconnectorrequestbody.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateOrganizationJitSsoConnectorResponse](../../models/operations/createorganizationjitssoconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceOrganizationJitSsoConnectors

Replace all enterprise SSO connectors for just-in-time provisioning of users in the organization with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.ReplaceOrganizationJitSsoConnectors(ctx, "<id>", operations.ReplaceOrganizationJitSsoConnectorsRequestBody{
        SsoConnectorIds: []string{
            "<value 1>",
            "<value 2>",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The unique identifier of the organization.                                                                                             |
| `requestBody`                                                                                                                          | [operations.ReplaceOrganizationJitSsoConnectorsRequestBody](../../models/operations/replaceorganizationjitssoconnectorsrequestbody.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.ReplaceOrganizationJitSsoConnectorsResponse](../../models/operations/replaceorganizationjitssoconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOrganizationJitSsoConnector

Remove an enterprise SSO connector for just-in-time provisioning of users in the organization.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Organizations.DeleteOrganizationJitSsoConnector(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization.               |
| `ssoConnectorID`                                         | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the sso connector.              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationJitSsoConnectorResponse](../../models/operations/deleteorganizationjitssoconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |