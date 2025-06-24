# ListClusterServices200ResponseAllOfServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalPort** | Pointer to **NullableString** |  | [optional] 
**InternalPort** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListClusterServices200ResponseAllOfServicesInner

`func NewListClusterServices200ResponseAllOfServicesInner() *ListClusterServices200ResponseAllOfServicesInner`

NewListClusterServices200ResponseAllOfServicesInner instantiates a new ListClusterServices200ResponseAllOfServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterServices200ResponseAllOfServicesInnerWithDefaults

`func NewListClusterServices200ResponseAllOfServicesInnerWithDefaults() *ListClusterServices200ResponseAllOfServicesInner`

NewListClusterServices200ResponseAllOfServicesInnerWithDefaults instantiates a new ListClusterServices200ResponseAllOfServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetExternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetExternalPort() string`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetExternalPortOk() (*string, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetExternalPort(v string)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### SetExternalPortNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetExternalPortNil(b bool)`

 SetExternalPortNil sets the value for ExternalPort to be an explicit nil

### UnsetExternalPort
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetExternalPort()`

UnsetExternalPort ensures that no value is present for ExternalPort, not even an explicit nil
### GetInternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetInternalPort() string`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetInternalPortOk() (*string, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetInternalPort(v string)`

SetInternalPort sets InternalPort field to given value.

### HasInternalPort

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasInternalPort() bool`

HasInternalPort returns a boolean if a field has been set.

### SetInternalPortNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetInternalPortNil(b bool)`

 SetInternalPortNil sets the value for InternalPort to be an explicit nil

### UnsetInternalPort
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetInternalPort()`

UnsetInternalPort ensures that no value is present for InternalPort, not even an explicit nil
### GetStatus

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDateCreated

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetLastUpdated

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterServices200ResponseAllOfServicesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterServices200ResponseAllOfServicesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ListClusterServices200ResponseAllOfServicesInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ListClusterServices200ResponseAllOfServicesInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


