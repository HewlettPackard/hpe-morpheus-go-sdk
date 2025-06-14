/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config{}

// ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config struct for ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config
type ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config struct {
	Url                        *string                `json:"url,omitempty"`
	LogoutUrl                  *string                `json:"logoutUrl,omitempty"`
	DoNotIncludeSAMLRequest    *bool                  `json:"doNotIncludeSAMLRequest,omitempty"`
	SAMLSignatureMode          *string                `json:"SAMLSignatureMode,omitempty"`
	DoNotValidateSignature     *bool                  `json:"doNotValidateSignature,omitempty"`
	DoNotValidateStatusCode    *bool                  `json:"doNotValidateStatusCode,omitempty"`
	DoNotValidateDestination   *bool                  `json:"doNotValidateDestination,omitempty"`
	DoNotValidateIssueInstants *bool                  `json:"doNotValidateIssueInstants,omitempty"`
	DoNotValidateAssertions    *bool                  `json:"doNotValidateAssertions,omitempty"`
	GivenNameAttribute         *string                `json:"givenNameAttribute,omitempty"`
	SurnameAttribute           *string                `json:"surnameAttribute,omitempty"`
	EmailAttribute             *string                `json:"emailAttribute,omitempty"`
	RequiredAttributeValue     *string                `json:"requiredAttributeValue,omitempty"`
	RoleAttributeName          *string                `json:"roleAttributeName,omitempty"`
	AzureTenantId              *string                `json:"azureTenantId,omitempty"`
	AzureAppId                 *string                `json:"azureAppId,omitempty"`
	AzureAppSecret             *string                `json:"azureAppSecret,omitempty"`
	RoleLinkAttributeName      *string                `json:"roleLinkAttributeName,omitempty"`
	PublicKey                  *string                `json:"publicKey,omitempty"`
	AzureAppSecretHash         *string                `json:"azureAppSecretHash,omitempty"`
	AdditionalProperties       map[string]interface{} `json:",remain"`
}

type _ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config{}
	return &this
}

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6ConfigWithDefaults instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6ConfigWithDefaults() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// IsSetUrl returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetUrl(v string) {
	o.Url = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetLogoutUrl() string {
	if o == nil || IsNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutUrl) {
		return nil, false
	}
	return o.LogoutUrl, true
}

// IsSetLogoutUrl returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetLogoutUrl() bool {
	if o != nil && !IsNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotIncludeSAMLRequest() bool {
	if o == nil || IsNil(o.DoNotIncludeSAMLRequest) {
		var ret bool
		return ret
	}
	return *o.DoNotIncludeSAMLRequest
}

// GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotIncludeSAMLRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotIncludeSAMLRequest) {
		return nil, false
	}
	return o.DoNotIncludeSAMLRequest, true
}

// IsSetDoNotIncludeSAMLRequest returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotIncludeSAMLRequest() bool {
	if o != nil && !IsNil(o.DoNotIncludeSAMLRequest) {
		return true
	}

	return false
}

// SetDoNotIncludeSAMLRequest gets a reference to the given bool and assigns it to the DoNotIncludeSAMLRequest field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotIncludeSAMLRequest(v bool) {
	o.DoNotIncludeSAMLRequest = &v
}

// GetSAMLSignatureMode returns the SAMLSignatureMode field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSAMLSignatureMode() string {
	if o == nil || IsNil(o.SAMLSignatureMode) {
		var ret string
		return ret
	}
	return *o.SAMLSignatureMode
}

// GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSAMLSignatureModeOk() (*string, bool) {
	if o == nil || IsNil(o.SAMLSignatureMode) {
		return nil, false
	}
	return o.SAMLSignatureMode, true
}

// IsSetSAMLSignatureMode returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetSAMLSignatureMode() bool {
	if o != nil && !IsNil(o.SAMLSignatureMode) {
		return true
	}

	return false
}

// SetSAMLSignatureMode gets a reference to the given string and assigns it to the SAMLSignatureMode field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetSAMLSignatureMode(v string) {
	o.SAMLSignatureMode = &v
}

// GetDoNotValidateSignature returns the DoNotValidateSignature field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateSignature() bool {
	if o == nil || IsNil(o.DoNotValidateSignature) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateSignature
}

// GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotValidateSignature) {
		return nil, false
	}
	return o.DoNotValidateSignature, true
}

// IsSetDoNotValidateSignature returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotValidateSignature() bool {
	if o != nil && !IsNil(o.DoNotValidateSignature) {
		return true
	}

	return false
}

// SetDoNotValidateSignature gets a reference to the given bool and assigns it to the DoNotValidateSignature field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateSignature(v bool) {
	o.DoNotValidateSignature = &v
}

// GetDoNotValidateStatusCode returns the DoNotValidateStatusCode field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateStatusCode() bool {
	if o == nil || IsNil(o.DoNotValidateStatusCode) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateStatusCode
}

// GetDoNotValidateStatusCodeOk returns a tuple with the DoNotValidateStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateStatusCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotValidateStatusCode) {
		return nil, false
	}
	return o.DoNotValidateStatusCode, true
}

// IsSetDoNotValidateStatusCode returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotValidateStatusCode() bool {
	if o != nil && !IsNil(o.DoNotValidateStatusCode) {
		return true
	}

	return false
}

// SetDoNotValidateStatusCode gets a reference to the given bool and assigns it to the DoNotValidateStatusCode field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateStatusCode(v bool) {
	o.DoNotValidateStatusCode = &v
}

// GetDoNotValidateDestination returns the DoNotValidateDestination field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateDestination() bool {
	if o == nil || IsNil(o.DoNotValidateDestination) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateDestination
}

// GetDoNotValidateDestinationOk returns a tuple with the DoNotValidateDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateDestinationOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotValidateDestination) {
		return nil, false
	}
	return o.DoNotValidateDestination, true
}

// IsSetDoNotValidateDestination returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotValidateDestination() bool {
	if o != nil && !IsNil(o.DoNotValidateDestination) {
		return true
	}

	return false
}

// SetDoNotValidateDestination gets a reference to the given bool and assigns it to the DoNotValidateDestination field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateDestination(v bool) {
	o.DoNotValidateDestination = &v
}

// GetDoNotValidateIssueInstants returns the DoNotValidateIssueInstants field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateIssueInstants() bool {
	if o == nil || IsNil(o.DoNotValidateIssueInstants) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateIssueInstants
}

// GetDoNotValidateIssueInstantsOk returns a tuple with the DoNotValidateIssueInstants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateIssueInstantsOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotValidateIssueInstants) {
		return nil, false
	}
	return o.DoNotValidateIssueInstants, true
}

// IsSetDoNotValidateIssueInstants returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotValidateIssueInstants() bool {
	if o != nil && !IsNil(o.DoNotValidateIssueInstants) {
		return true
	}

	return false
}

// SetDoNotValidateIssueInstants gets a reference to the given bool and assigns it to the DoNotValidateIssueInstants field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateIssueInstants(v bool) {
	o.DoNotValidateIssueInstants = &v
}

// GetDoNotValidateAssertions returns the DoNotValidateAssertions field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateAssertions() bool {
	if o == nil || IsNil(o.DoNotValidateAssertions) {
		var ret bool
		return ret
	}
	return *o.DoNotValidateAssertions
}

// GetDoNotValidateAssertionsOk returns a tuple with the DoNotValidateAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotValidateAssertions) {
		return nil, false
	}
	return o.DoNotValidateAssertions, true
}

// IsSetDoNotValidateAssertions returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetDoNotValidateAssertions() bool {
	if o != nil && !IsNil(o.DoNotValidateAssertions) {
		return true
	}

	return false
}

// SetDoNotValidateAssertions gets a reference to the given bool and assigns it to the DoNotValidateAssertions field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateAssertions(v bool) {
	o.DoNotValidateAssertions = &v
}

// GetGivenNameAttribute returns the GivenNameAttribute field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetGivenNameAttribute() string {
	if o == nil || IsNil(o.GivenNameAttribute) {
		var ret string
		return ret
	}
	return *o.GivenNameAttribute
}

// GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetGivenNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GivenNameAttribute) {
		return nil, false
	}
	return o.GivenNameAttribute, true
}

// IsSetGivenNameAttribute returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetGivenNameAttribute() bool {
	if o != nil && !IsNil(o.GivenNameAttribute) {
		return true
	}

	return false
}

// SetGivenNameAttribute gets a reference to the given string and assigns it to the GivenNameAttribute field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetGivenNameAttribute(v string) {
	o.GivenNameAttribute = &v
}

// GetSurnameAttribute returns the SurnameAttribute field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSurnameAttribute() string {
	if o == nil || IsNil(o.SurnameAttribute) {
		var ret string
		return ret
	}
	return *o.SurnameAttribute
}

// GetSurnameAttributeOk returns a tuple with the SurnameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSurnameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.SurnameAttribute) {
		return nil, false
	}
	return o.SurnameAttribute, true
}

// IsSetSurnameAttribute returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetSurnameAttribute() bool {
	if o != nil && !IsNil(o.SurnameAttribute) {
		return true
	}

	return false
}

// SetSurnameAttribute gets a reference to the given string and assigns it to the SurnameAttribute field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetSurnameAttribute(v string) {
	o.SurnameAttribute = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetEmailAttribute() string {
	if o == nil || IsNil(o.EmailAttribute) {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetEmailAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttribute) {
		return nil, false
	}
	return o.EmailAttribute, true
}

// IsSetEmailAttribute returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetEmailAttribute() bool {
	if o != nil && !IsNil(o.EmailAttribute) {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetRequiredAttributeValue returns the RequiredAttributeValue field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRequiredAttributeValue() string {
	if o == nil || IsNil(o.RequiredAttributeValue) {
		var ret string
		return ret
	}
	return *o.RequiredAttributeValue
}

// GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRequiredAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredAttributeValue) {
		return nil, false
	}
	return o.RequiredAttributeValue, true
}

// IsSetRequiredAttributeValue returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetRequiredAttributeValue() bool {
	if o != nil && !IsNil(o.RequiredAttributeValue) {
		return true
	}

	return false
}

// SetRequiredAttributeValue gets a reference to the given string and assigns it to the RequiredAttributeValue field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRequiredAttributeValue(v string) {
	o.RequiredAttributeValue = &v
}

// GetRoleAttributeName returns the RoleAttributeName field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleAttributeName() string {
	if o == nil || IsNil(o.RoleAttributeName) {
		var ret string
		return ret
	}
	return *o.RoleAttributeName
}

// GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleAttributeName) {
		return nil, false
	}
	return o.RoleAttributeName, true
}

// IsSetRoleAttributeName returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetRoleAttributeName() bool {
	if o != nil && !IsNil(o.RoleAttributeName) {
		return true
	}

	return false
}

// SetRoleAttributeName gets a reference to the given string and assigns it to the RoleAttributeName field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRoleAttributeName(v string) {
	o.RoleAttributeName = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureTenantId() string {
	if o == nil || IsNil(o.AzureTenantId) {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.AzureTenantId) {
		return nil, false
	}
	return o.AzureTenantId, true
}

// IsSetAzureTenantId returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetAzureTenantId() bool {
	if o != nil && !IsNil(o.AzureTenantId) {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetAzureAppId returns the AzureAppId field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppId() string {
	if o == nil || IsNil(o.AzureAppId) {
		var ret string
		return ret
	}
	return *o.AzureAppId
}

// GetAzureAppIdOk returns a tuple with the AzureAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AzureAppId) {
		return nil, false
	}
	return o.AzureAppId, true
}

// IsSetAzureAppId returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetAzureAppId() bool {
	if o != nil && !IsNil(o.AzureAppId) {
		return true
	}

	return false
}

// SetAzureAppId gets a reference to the given string and assigns it to the AzureAppId field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppId(v string) {
	o.AzureAppId = &v
}

// GetAzureAppSecret returns the AzureAppSecret field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecret() string {
	if o == nil || IsNil(o.AzureAppSecret) {
		var ret string
		return ret
	}
	return *o.AzureAppSecret
}

// GetAzureAppSecretOk returns a tuple with the AzureAppSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretOk() (*string, bool) {
	if o == nil || IsNil(o.AzureAppSecret) {
		return nil, false
	}
	return o.AzureAppSecret, true
}

// IsSetAzureAppSecret returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetAzureAppSecret() bool {
	if o != nil && !IsNil(o.AzureAppSecret) {
		return true
	}

	return false
}

// SetAzureAppSecret gets a reference to the given string and assigns it to the AzureAppSecret field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppSecret(v string) {
	o.AzureAppSecret = &v
}

// GetRoleLinkAttributeName returns the RoleLinkAttributeName field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleLinkAttributeName() string {
	if o == nil || IsNil(o.RoleLinkAttributeName) {
		var ret string
		return ret
	}
	return *o.RoleLinkAttributeName
}

// GetRoleLinkAttributeNameOk returns a tuple with the RoleLinkAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleLinkAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleLinkAttributeName) {
		return nil, false
	}
	return o.RoleLinkAttributeName, true
}

// IsSetRoleLinkAttributeName returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetRoleLinkAttributeName() bool {
	if o != nil && !IsNil(o.RoleLinkAttributeName) {
		return true
	}

	return false
}

// SetRoleLinkAttributeName gets a reference to the given string and assigns it to the RoleLinkAttributeName field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRoleLinkAttributeName(v string) {
	o.RoleLinkAttributeName = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// IsSetPublicKey returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetAzureAppSecretHash returns the AzureAppSecretHash field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretHash() string {
	if o == nil || IsNil(o.AzureAppSecretHash) {
		var ret string
		return ret
	}
	return *o.AzureAppSecretHash
}

// GetAzureAppSecretHashOk returns a tuple with the AzureAppSecretHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretHashOk() (*string, bool) {
	if o == nil || IsNil(o.AzureAppSecretHash) {
		return nil, false
	}
	return o.AzureAppSecretHash, true
}

// IsSetAzureAppSecretHash returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) IsSetAzureAppSecretHash() bool {
	if o != nil && !IsNil(o.AzureAppSecretHash) {
		return true
	}

	return false
}

// SetAzureAppSecretHash gets a reference to the given string and assigns it to the AzureAppSecretHash field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppSecretHash(v string) {
	o.AzureAppSecretHash = &v
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.LogoutUrl) {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if !IsNil(o.DoNotIncludeSAMLRequest) {
		toSerialize["doNotIncludeSAMLRequest"] = o.DoNotIncludeSAMLRequest
	}
	if !IsNil(o.SAMLSignatureMode) {
		toSerialize["SAMLSignatureMode"] = o.SAMLSignatureMode
	}
	if !IsNil(o.DoNotValidateSignature) {
		toSerialize["doNotValidateSignature"] = o.DoNotValidateSignature
	}
	if !IsNil(o.DoNotValidateStatusCode) {
		toSerialize["doNotValidateStatusCode"] = o.DoNotValidateStatusCode
	}
	if !IsNil(o.DoNotValidateDestination) {
		toSerialize["doNotValidateDestination"] = o.DoNotValidateDestination
	}
	if !IsNil(o.DoNotValidateIssueInstants) {
		toSerialize["doNotValidateIssueInstants"] = o.DoNotValidateIssueInstants
	}
	if !IsNil(o.DoNotValidateAssertions) {
		toSerialize["doNotValidateAssertions"] = o.DoNotValidateAssertions
	}
	if !IsNil(o.GivenNameAttribute) {
		toSerialize["givenNameAttribute"] = o.GivenNameAttribute
	}
	if !IsNil(o.SurnameAttribute) {
		toSerialize["surnameAttribute"] = o.SurnameAttribute
	}
	if !IsNil(o.EmailAttribute) {
		toSerialize["emailAttribute"] = o.EmailAttribute
	}
	if !IsNil(o.RequiredAttributeValue) {
		toSerialize["requiredAttributeValue"] = o.RequiredAttributeValue
	}
	if !IsNil(o.RoleAttributeName) {
		toSerialize["roleAttributeName"] = o.RoleAttributeName
	}
	if !IsNil(o.AzureTenantId) {
		toSerialize["azureTenantId"] = o.AzureTenantId
	}
	if !IsNil(o.AzureAppId) {
		toSerialize["azureAppId"] = o.AzureAppId
	}
	if !IsNil(o.AzureAppSecret) {
		toSerialize["azureAppSecret"] = o.AzureAppSecret
	}
	if !IsNil(o.RoleLinkAttributeName) {
		toSerialize["roleLinkAttributeName"] = o.RoleLinkAttributeName
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !IsNil(o.AzureAppSecretHash) {
		toSerialize["azureAppSecretHash"] = o.AzureAppSecretHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
