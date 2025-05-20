# OptionTypeListCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Description** | Pointer to **string** | Description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;. | [optional] [default to "rest"]
**SourceUrl** | Pointer to **string** | Source URL, the http(s) URL to request data from. Required when type is rest. | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**SourceMethod** | Pointer to **string** | Source Method, the HTTP method. | [optional] [default to "GET"]
**ApiType** | Pointer to **string** | Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. | [optional] 
**IgnoreSSLErrors** | Pointer to **bool** | Ignore SSL Errors. | [optional] [default to false]
**RealTime** | Pointer to **bool** | Real Time. | [optional] [default to false]
**Credential** | Pointer to [**AddOptionListRequestOptionTypeListCredential**](AddOptionListRequestOptionTypeListCredential.md) |  | [optional] 
**ServiceUsername** | Pointer to **string** | Username for authenticating with Basic Auth when type is rest or ldap username. | [optional] 
**ServicePassword** | Pointer to **string** | Password for authenticating with Basic Auth when type is rest or ldap password. | [optional] 
**InitialDataset** | Pointer to **string** | Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual. | [optional] 
**TranslationScript** | Pointer to **string** | Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**RequestScript** | Pointer to **string** | Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**Config** | Pointer to [**AddOptionListRequestOptionTypeListConfig**](AddOptionListRequestOptionTypeListConfig.md) |  | [optional] 

## Methods

### NewOptionTypeListCreate

`func NewOptionTypeListCreate(name string, ) *OptionTypeListCreate`

NewOptionTypeListCreate instantiates a new OptionTypeListCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeListCreateWithDefaults

`func NewOptionTypeListCreateWithDefaults() *OptionTypeListCreate`

NewOptionTypeListCreateWithDefaults instantiates a new OptionTypeListCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptionTypeListCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeListCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeListCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OptionTypeListCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeListCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeListCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeListCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *OptionTypeListCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OptionTypeListCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OptionTypeListCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OptionTypeListCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *OptionTypeListCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionTypeListCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionTypeListCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionTypeListCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *OptionTypeListCreate) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *OptionTypeListCreate) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *OptionTypeListCreate) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *OptionTypeListCreate) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OptionTypeListCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OptionTypeListCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OptionTypeListCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OptionTypeListCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSourceMethod

`func (o *OptionTypeListCreate) GetSourceMethod() string`

GetSourceMethod returns the SourceMethod field if non-nil, zero value otherwise.

### GetSourceMethodOk

`func (o *OptionTypeListCreate) GetSourceMethodOk() (*string, bool)`

GetSourceMethodOk returns a tuple with the SourceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethod

`func (o *OptionTypeListCreate) SetSourceMethod(v string)`

SetSourceMethod sets SourceMethod field to given value.

### HasSourceMethod

`func (o *OptionTypeListCreate) HasSourceMethod() bool`

HasSourceMethod returns a boolean if a field has been set.

### GetApiType

`func (o *OptionTypeListCreate) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OptionTypeListCreate) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OptionTypeListCreate) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OptionTypeListCreate) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetIgnoreSSLErrors

`func (o *OptionTypeListCreate) GetIgnoreSSLErrors() bool`

GetIgnoreSSLErrors returns the IgnoreSSLErrors field if non-nil, zero value otherwise.

### GetIgnoreSSLErrorsOk

`func (o *OptionTypeListCreate) GetIgnoreSSLErrorsOk() (*bool, bool)`

GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSLErrors

`func (o *OptionTypeListCreate) SetIgnoreSSLErrors(v bool)`

SetIgnoreSSLErrors sets IgnoreSSLErrors field to given value.

### HasIgnoreSSLErrors

`func (o *OptionTypeListCreate) HasIgnoreSSLErrors() bool`

HasIgnoreSSLErrors returns a boolean if a field has been set.

### GetRealTime

`func (o *OptionTypeListCreate) GetRealTime() bool`

GetRealTime returns the RealTime field if non-nil, zero value otherwise.

### GetRealTimeOk

`func (o *OptionTypeListCreate) GetRealTimeOk() (*bool, bool)`

GetRealTimeOk returns a tuple with the RealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTime

`func (o *OptionTypeListCreate) SetRealTime(v bool)`

SetRealTime sets RealTime field to given value.

### HasRealTime

`func (o *OptionTypeListCreate) HasRealTime() bool`

HasRealTime returns a boolean if a field has been set.

### GetCredential

`func (o *OptionTypeListCreate) GetCredential() AddOptionListRequestOptionTypeListCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OptionTypeListCreate) GetCredentialOk() (*AddOptionListRequestOptionTypeListCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OptionTypeListCreate) SetCredential(v AddOptionListRequestOptionTypeListCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OptionTypeListCreate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetServiceUsername

`func (o *OptionTypeListCreate) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *OptionTypeListCreate) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *OptionTypeListCreate) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *OptionTypeListCreate) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *OptionTypeListCreate) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *OptionTypeListCreate) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *OptionTypeListCreate) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *OptionTypeListCreate) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetInitialDataset

`func (o *OptionTypeListCreate) GetInitialDataset() string`

GetInitialDataset returns the InitialDataset field if non-nil, zero value otherwise.

### GetInitialDatasetOk

`func (o *OptionTypeListCreate) GetInitialDatasetOk() (*string, bool)`

GetInitialDatasetOk returns a tuple with the InitialDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDataset

`func (o *OptionTypeListCreate) SetInitialDataset(v string)`

SetInitialDataset sets InitialDataset field to given value.

### HasInitialDataset

`func (o *OptionTypeListCreate) HasInitialDataset() bool`

HasInitialDataset returns a boolean if a field has been set.

### GetTranslationScript

`func (o *OptionTypeListCreate) GetTranslationScript() string`

GetTranslationScript returns the TranslationScript field if non-nil, zero value otherwise.

### GetTranslationScriptOk

`func (o *OptionTypeListCreate) GetTranslationScriptOk() (*string, bool)`

GetTranslationScriptOk returns a tuple with the TranslationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationScript

`func (o *OptionTypeListCreate) SetTranslationScript(v string)`

SetTranslationScript sets TranslationScript field to given value.

### HasTranslationScript

`func (o *OptionTypeListCreate) HasTranslationScript() bool`

HasTranslationScript returns a boolean if a field has been set.

### GetRequestScript

`func (o *OptionTypeListCreate) GetRequestScript() string`

GetRequestScript returns the RequestScript field if non-nil, zero value otherwise.

### GetRequestScriptOk

`func (o *OptionTypeListCreate) GetRequestScriptOk() (*string, bool)`

GetRequestScriptOk returns a tuple with the RequestScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestScript

`func (o *OptionTypeListCreate) SetRequestScript(v string)`

SetRequestScript sets RequestScript field to given value.

### HasRequestScript

`func (o *OptionTypeListCreate) HasRequestScript() bool`

HasRequestScript returns a boolean if a field has been set.

### GetConfig

`func (o *OptionTypeListCreate) GetConfig() AddOptionListRequestOptionTypeListConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OptionTypeListCreate) GetConfigOk() (*AddOptionListRequestOptionTypeListConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OptionTypeListCreate) SetConfig(v AddOptionListRequestOptionTypeListConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OptionTypeListCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


