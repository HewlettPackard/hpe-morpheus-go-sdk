# IntegrationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**HasCMDB** | Pointer to **bool** |  | [optional] 
**HasCMDBDiscovery** | Pointer to **bool** |  | [optional] 
**HasCM** | Pointer to **bool** |  | [optional] 
**HasDNS** | Pointer to **bool** |  | [optional] 
**HasApprovals** | Pointer to **bool** |  | [optional] 
**HasDeleteApprovals** | Pointer to **bool** |  | [optional] 
**HasReconfigureApprovals** | Pointer to **bool** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 

## Methods

### NewIntegrationType

`func NewIntegrationType() *IntegrationType`

NewIntegrationType instantiates a new IntegrationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationTypeWithDefaults

`func NewIntegrationTypeWithDefaults() *IntegrationType`

NewIntegrationTypeWithDefaults instantiates a new IntegrationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *IntegrationType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IntegrationType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IntegrationType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IntegrationType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *IntegrationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *IntegrationType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntegrationType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntegrationType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IntegrationType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *IntegrationType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *IntegrationType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *IntegrationType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *IntegrationType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasCMDB

`func (o *IntegrationType) GetHasCMDB() bool`

GetHasCMDB returns the HasCMDB field if non-nil, zero value otherwise.

### GetHasCMDBOk

`func (o *IntegrationType) GetHasCMDBOk() (*bool, bool)`

GetHasCMDBOk returns a tuple with the HasCMDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCMDB

`func (o *IntegrationType) SetHasCMDB(v bool)`

SetHasCMDB sets HasCMDB field to given value.

### HasHasCMDB

`func (o *IntegrationType) HasHasCMDB() bool`

HasHasCMDB returns a boolean if a field has been set.

### GetHasCMDBDiscovery

`func (o *IntegrationType) GetHasCMDBDiscovery() bool`

GetHasCMDBDiscovery returns the HasCMDBDiscovery field if non-nil, zero value otherwise.

### GetHasCMDBDiscoveryOk

`func (o *IntegrationType) GetHasCMDBDiscoveryOk() (*bool, bool)`

GetHasCMDBDiscoveryOk returns a tuple with the HasCMDBDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCMDBDiscovery

`func (o *IntegrationType) SetHasCMDBDiscovery(v bool)`

SetHasCMDBDiscovery sets HasCMDBDiscovery field to given value.

### HasHasCMDBDiscovery

`func (o *IntegrationType) HasHasCMDBDiscovery() bool`

HasHasCMDBDiscovery returns a boolean if a field has been set.

### GetHasCM

`func (o *IntegrationType) GetHasCM() bool`

GetHasCM returns the HasCM field if non-nil, zero value otherwise.

### GetHasCMOk

`func (o *IntegrationType) GetHasCMOk() (*bool, bool)`

GetHasCMOk returns a tuple with the HasCM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCM

`func (o *IntegrationType) SetHasCM(v bool)`

SetHasCM sets HasCM field to given value.

### HasHasCM

`func (o *IntegrationType) HasHasCM() bool`

HasHasCM returns a boolean if a field has been set.

### GetHasDNS

`func (o *IntegrationType) GetHasDNS() bool`

GetHasDNS returns the HasDNS field if non-nil, zero value otherwise.

### GetHasDNSOk

`func (o *IntegrationType) GetHasDNSOk() (*bool, bool)`

GetHasDNSOk returns a tuple with the HasDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDNS

`func (o *IntegrationType) SetHasDNS(v bool)`

SetHasDNS sets HasDNS field to given value.

### HasHasDNS

`func (o *IntegrationType) HasHasDNS() bool`

HasHasDNS returns a boolean if a field has been set.

### GetHasApprovals

`func (o *IntegrationType) GetHasApprovals() bool`

GetHasApprovals returns the HasApprovals field if non-nil, zero value otherwise.

### GetHasApprovalsOk

`func (o *IntegrationType) GetHasApprovalsOk() (*bool, bool)`

GetHasApprovalsOk returns a tuple with the HasApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApprovals

`func (o *IntegrationType) SetHasApprovals(v bool)`

SetHasApprovals sets HasApprovals field to given value.

### HasHasApprovals

`func (o *IntegrationType) HasHasApprovals() bool`

HasHasApprovals returns a boolean if a field has been set.

### GetHasDeleteApprovals

`func (o *IntegrationType) GetHasDeleteApprovals() bool`

GetHasDeleteApprovals returns the HasDeleteApprovals field if non-nil, zero value otherwise.

### GetHasDeleteApprovalsOk

`func (o *IntegrationType) GetHasDeleteApprovalsOk() (*bool, bool)`

GetHasDeleteApprovalsOk returns a tuple with the HasDeleteApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeleteApprovals

`func (o *IntegrationType) SetHasDeleteApprovals(v bool)`

SetHasDeleteApprovals sets HasDeleteApprovals field to given value.

### HasHasDeleteApprovals

`func (o *IntegrationType) HasHasDeleteApprovals() bool`

HasHasDeleteApprovals returns a boolean if a field has been set.

### GetHasReconfigureApprovals

`func (o *IntegrationType) GetHasReconfigureApprovals() bool`

GetHasReconfigureApprovals returns the HasReconfigureApprovals field if non-nil, zero value otherwise.

### GetHasReconfigureApprovalsOk

`func (o *IntegrationType) GetHasReconfigureApprovalsOk() (*bool, bool)`

GetHasReconfigureApprovalsOk returns a tuple with the HasReconfigureApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReconfigureApprovals

`func (o *IntegrationType) SetHasReconfigureApprovals(v bool)`

SetHasReconfigureApprovals sets HasReconfigureApprovals field to given value.

### HasHasReconfigureApprovals

`func (o *IntegrationType) HasHasReconfigureApprovals() bool`

HasHasReconfigureApprovals returns a boolean if a field has been set.

### GetIsPlugin

`func (o *IntegrationType) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *IntegrationType) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *IntegrationType) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *IntegrationType) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


