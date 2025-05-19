# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**IsEC2** | Pointer to **string** | Amazon Cloud Type | [optional] [default to "false"]
**AvailabilityId** | Pointer to **string** | Amazon Zone | [optional] 
**SecurityId** | Pointer to **string** | Security Group | [optional] 
**PublicIpType** | Pointer to **string** | Public IP | [optional] 
**InstanceProfile** | Pointer to **string** | IAM Profile | [optional] 
**KmsKeyId** | Pointer to **string** | KMS Key ID | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3 instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3WithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3WithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3WithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetIsEC2

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf3) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


