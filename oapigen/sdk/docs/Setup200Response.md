# Setup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**BuildVersion** | Pointer to **string** | Morpheus build version that the server is running. | [optional] 
**Uuid** | Pointer to **string** | The Appliance Unique ID that is auto generated. | [optional] 
**ApplianceUrl** | Pointer to **string** | The Appliance Server URL as defined under Appliance Settings. | [optional] 
**SetupNeeded** | Pointer to **bool** | Flag to determine if the appliance has been setup, only true when appliance is a fresh install and has not been initialized. | [optional] 

## Methods

### NewSetup200Response

`func NewSetup200Response() *Setup200Response`

NewSetup200Response instantiates a new Setup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetup200ResponseWithDefaults

`func NewSetup200ResponseWithDefaults() *Setup200Response`

NewSetup200ResponseWithDefaults instantiates a new Setup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Setup200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Setup200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Setup200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Setup200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBuildVersion

`func (o *Setup200Response) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Setup200Response) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Setup200Response) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Setup200Response) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetUuid

`func (o *Setup200Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Setup200Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Setup200Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Setup200Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *Setup200Response) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *Setup200Response) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *Setup200Response) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *Setup200Response) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetSetupNeeded

`func (o *Setup200Response) GetSetupNeeded() bool`

GetSetupNeeded returns the SetupNeeded field if non-nil, zero value otherwise.

### GetSetupNeededOk

`func (o *Setup200Response) GetSetupNeededOk() (*bool, bool)`

GetSetupNeededOk returns a tuple with the SetupNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupNeeded

`func (o *Setup200Response) SetSetupNeeded(v bool)`

SetSetupNeeded sets SetupNeeded field to given value.

### HasSetupNeeded

`func (o *Setup200Response) HasSetupNeeded() bool`

HasSetupNeeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


