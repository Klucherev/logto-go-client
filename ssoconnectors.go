// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package logtogoclient

import (
	"bytes"
	"context"
	"fmt"
	"github.com/klucherev/logto-go-client/internal/hooks"
	"github.com/klucherev/logto-go-client/internal/utils"
	"github.com/klucherev/logto-go-client/models/apierrors"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"github.com/klucherev/logto-go-client/retry"
	"net/http"
	"net/url"
)

// SSOConnectors - Endpoints for managing single sign-on (SSO) connectors. Your sign-in experience can use these well-configured SSO connectors to authenticate users and sync user attributes from external identity providers (IdPs).
//
// SSO connectors are created by SSO connector provider factories.
type SSOConnectors struct {
	sdkConfiguration sdkConfiguration
}

func newSSOConnectors(sdkConfig sdkConfiguration) *SSOConnectors {
	return &SSOConnectors{
		sdkConfiguration: sdkConfig,
	}
}

// CreateSsoConnector - Create SSO connector
// Create an new SSO connector instance for a given provider.
func (s *SSOConnectors) CreateSsoConnector(ctx context.Context, request operations.CreateSsoConnectorRequest, opts ...operations.Option) (*operations.CreateSsoConnectorResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := url.JoinPath(baseURL, "/api/sso-connectors")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "CreateSsoConnector",
		OAuth2Scopes:   []string{"all"},
		SecuritySource: s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"400", "401", "403", "409", "422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.CreateSsoConnectorResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out operations.CreateSsoConnectorResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// ListSsoConnectors - List SSO connectors
// Get SSO connectors with pagination. In addition to the raw SSO connector data, a copy of fetched or parsed IdP configs and a copy of connector provider's data will be attached.
func (s *SSOConnectors) ListSsoConnectors(ctx context.Context, page *int64, pageSize *int64, opts ...operations.Option) (*operations.ListSsoConnectorsResponse, error) {
	request := operations.ListSsoConnectorsRequest{
		Page:     page,
		PageSize: pageSize,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := url.JoinPath(baseURL, "/api/sso-connectors")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "ListSsoConnectors",
		OAuth2Scopes:   []string{"all"},
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "403", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListSsoConnectorsResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out []operations.ListSsoConnectorsResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ResponseBodies = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// GetSsoConnector - Get SSO connector
// Get SSO connector data by ID. In addition to the raw SSO connector data, a copy of fetched or parsed IdP configs and a copy of connector provider's data will be attached.
func (s *SSOConnectors) GetSsoConnector(ctx context.Context, id string, opts ...operations.Option) (*operations.GetSsoConnectorResponse, error) {
	request := operations.GetSsoConnectorRequest{
		ID: id,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/sso-connectors/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "GetSsoConnector",
		OAuth2Scopes:   []string{"all"},
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.GetSsoConnectorResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out operations.GetSsoConnectorResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// DeleteSsoConnector - Delete SSO connector
// Delete an SSO connector by ID.
func (s *SSOConnectors) DeleteSsoConnector(ctx context.Context, id string, opts ...operations.Option) (*operations.DeleteSsoConnectorResponse, error) {
	request := operations.DeleteSsoConnectorRequest{
		ID: id,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/sso-connectors/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "DeleteSsoConnector",
		OAuth2Scopes:   []string{"all"},
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.DeleteSsoConnectorResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 204:
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// UpdateSsoConnector - Update SSO connector
// Update an SSO connector by ID. This method performs a partial update.
func (s *SSOConnectors) UpdateSsoConnector(ctx context.Context, id string, requestBody operations.UpdateSsoConnectorRequestBody, opts ...operations.Option) (*operations.UpdateSsoConnectorResponse, error) {
	request := operations.UpdateSsoConnectorRequest{
		ID:          id,
		RequestBody: requestBody,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/sso-connectors/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "UpdateSsoConnector",
		OAuth2Scopes:   []string{"all"},
		SecuritySource: s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"400", "401", "403", "404", "409", "422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.UpdateSsoConnectorResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out operations.UpdateSsoConnectorResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
