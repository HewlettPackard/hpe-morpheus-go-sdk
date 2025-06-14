/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SSLCertificatesAPIService SSLCertificatesAPI service
type SSLCertificatesAPIService service

type ApiAddCertificateRequest struct {
	ctx                   context.Context
	ApiService            *SSLCertificatesAPIService
	addCertificateRequest *AddCertificateRequest
}

func (r ApiAddCertificateRequest) AddCertificateRequest(addCertificateRequest AddCertificateRequest) ApiAddCertificateRequest {
	r.addCertificateRequest = &addCertificateRequest
	return r
}

func (r ApiAddCertificateRequest) Execute() (*AddCertificate200Response, *http.Response, error) {
	return r.ApiService.AddCertificateExecute(r)
}

/*
AddCertificate Create a Certificate

Create a Certificate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCertificateRequest
*/
func (a *SSLCertificatesAPIService) AddCertificate(ctx context.Context) ApiAddCertificateRequest {
	return ApiAddCertificateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddCertificate200Response
func (a *SSLCertificatesAPIService) AddCertificateExecute(r ApiAddCertificateRequest) (*AddCertificate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCertificate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SSLCertificatesAPIService.AddCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{err: err}
	}

	localVarPath := localBasePath + "/api/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addCertificateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ListActivity4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ListActivity5XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
			err:  err,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteCertificateRequest struct {
	ctx        context.Context
	ApiService *SSLCertificatesAPIService
	id         int64
}

func (r ApiDeleteCertificateRequest) Execute() (*DeleteAlerts200Response, *http.Response, error) {
	return r.ApiService.DeleteCertificateExecute(r)
}

/*
DeleteCertificate Delete a Certificate

Will delete a certificate from the system and make it no longer usable.

If a certificate is actively in use, a delete will fail.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Morpheus ID of the Object being referenced
	@return ApiDeleteCertificateRequest
*/
func (a *SSLCertificatesAPIService) DeleteCertificate(ctx context.Context, id int64) ApiDeleteCertificateRequest {
	return ApiDeleteCertificateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return DeleteAlerts200Response
func (a *SSLCertificatesAPIService) DeleteCertificateExecute(r ApiDeleteCertificateRequest) (*DeleteAlerts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteAlerts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SSLCertificatesAPIService.DeleteCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{err: err}
	}

	localVarPath := localBasePath + "/api/certificates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ListActivity4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ListActivity5XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
			err:  err,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCertificateRequest struct {
	ctx        context.Context
	ApiService *SSLCertificatesAPIService
	id         int64
}

func (r ApiGetCertificateRequest) Execute() (*GetCertificate200Response, *http.Response, error) {
	return r.ApiService.GetCertificateExecute(r)
}

/*
GetCertificate Get a Specific Certificate

This endpoint retrieves a specific certificate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Morpheus ID of the Object being referenced
	@return ApiGetCertificateRequest
*/
func (a *SSLCertificatesAPIService) GetCertificate(ctx context.Context, id int64) ApiGetCertificateRequest {
	return ApiGetCertificateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetCertificate200Response
func (a *SSLCertificatesAPIService) GetCertificateExecute(r ApiGetCertificateRequest) (*GetCertificate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCertificate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SSLCertificatesAPIService.GetCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{err: err}
	}

	localVarPath := localBasePath + "/api/certificates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ListActivity4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ListActivity5XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
			err:  err,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCertificatesRequest struct {
	ctx        context.Context
	ApiService *SSLCertificatesAPIService
	name       *string
}

// Filter by name
func (r ApiListCertificatesRequest) Name(name string) ApiListCertificatesRequest {
	r.name = &name
	return r
}

func (r ApiListCertificatesRequest) Execute() (*ListCertificates200Response, *http.Response, error) {
	return r.ApiService.ListCertificatesExecute(r)
}

/*
ListCertificates Get All SSL Certificates

This endpoint retrieves all SSL certificates associated with the account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCertificatesRequest
*/
func (a *SSLCertificatesAPIService) ListCertificates(ctx context.Context) ApiListCertificatesRequest {
	return ApiListCertificatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListCertificates200Response
func (a *SSLCertificatesAPIService) ListCertificatesExecute(r ApiListCertificatesRequest) (*ListCertificates200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListCertificates200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SSLCertificatesAPIService.ListCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{err: err}
	}

	localVarPath := localBasePath + "/api/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ListActivity4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ListActivity5XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
			err:  err,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCertificateRequest struct {
	ctx                   context.Context
	ApiService            *SSLCertificatesAPIService
	id                    int64
	addCertificateRequest *AddCertificateRequest
}

func (r ApiUpdateCertificateRequest) AddCertificateRequest(addCertificateRequest AddCertificateRequest) ApiUpdateCertificateRequest {
	r.addCertificateRequest = &addCertificateRequest
	return r
}

func (r ApiUpdateCertificateRequest) Execute() (*GetCertificate200Response, *http.Response, error) {
	return r.ApiService.UpdateCertificateExecute(r)
}

/*
UpdateCertificate Update a Certificate

Update a Certificate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Morpheus ID of the Object being referenced
	@return ApiUpdateCertificateRequest
*/
func (a *SSLCertificatesAPIService) UpdateCertificate(ctx context.Context, id int64) ApiUpdateCertificateRequest {
	return ApiUpdateCertificateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetCertificate200Response
func (a *SSLCertificatesAPIService) UpdateCertificateExecute(r ApiUpdateCertificateRequest) (*GetCertificate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCertificate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SSLCertificatesAPIService.UpdateCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{err: err}
	}

	localVarPath := localBasePath + "/api/certificates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addCertificateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ListActivity4XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ListActivity5XXResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.err = err
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.err = errors.New(formatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body: localVarBody,
			err:  err,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
