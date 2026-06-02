# GetIntegrationTypes200ResponseIntegrationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
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

### NewGetIntegrationTypes200ResponseIntegrationType

`func NewGetIntegrationTypes200ResponseIntegrationType() *GetIntegrationTypes200ResponseIntegrationType`

NewGetIntegrationTypes200ResponseIntegrationType instantiates a new GetIntegrationTypes200ResponseIntegrationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetIntegrationTypes200ResponseIntegrationType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasCMDB

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCMDB() bool`

GetHasCMDB returns the HasCMDB field if non-nil, zero value otherwise.

### GetHasCMDBOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCMDBOk() (*bool, bool)`

GetHasCMDBOk returns a tuple with the HasCMDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCMDB

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasCMDB(v bool)`

SetHasCMDB sets HasCMDB field to given value.

### HasHasCMDB

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasCMDB() bool`

HasHasCMDB returns a boolean if a field has been set.

### GetHasCMDBDiscovery

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCMDBDiscovery() bool`

GetHasCMDBDiscovery returns the HasCMDBDiscovery field if non-nil, zero value otherwise.

### GetHasCMDBDiscoveryOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCMDBDiscoveryOk() (*bool, bool)`

GetHasCMDBDiscoveryOk returns a tuple with the HasCMDBDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCMDBDiscovery

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasCMDBDiscovery(v bool)`

SetHasCMDBDiscovery sets HasCMDBDiscovery field to given value.

### HasHasCMDBDiscovery

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasCMDBDiscovery() bool`

HasHasCMDBDiscovery returns a boolean if a field has been set.

### GetHasCM

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCM() bool`

GetHasCM returns the HasCM field if non-nil, zero value otherwise.

### GetHasCMOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasCMOk() (*bool, bool)`

GetHasCMOk returns a tuple with the HasCM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCM

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasCM(v bool)`

SetHasCM sets HasCM field to given value.

### HasHasCM

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasCM() bool`

HasHasCM returns a boolean if a field has been set.

### GetHasDNS

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasDNS() bool`

GetHasDNS returns the HasDNS field if non-nil, zero value otherwise.

### GetHasDNSOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasDNSOk() (*bool, bool)`

GetHasDNSOk returns a tuple with the HasDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDNS

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasDNS(v bool)`

SetHasDNS sets HasDNS field to given value.

### HasHasDNS

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasDNS() bool`

HasHasDNS returns a boolean if a field has been set.

### GetHasApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasApprovals() bool`

GetHasApprovals returns the HasApprovals field if non-nil, zero value otherwise.

### GetHasApprovalsOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasApprovalsOk() (*bool, bool)`

GetHasApprovalsOk returns a tuple with the HasApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasApprovals(v bool)`

SetHasApprovals sets HasApprovals field to given value.

### HasHasApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasApprovals() bool`

HasHasApprovals returns a boolean if a field has been set.

### GetHasDeleteApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasDeleteApprovals() bool`

GetHasDeleteApprovals returns the HasDeleteApprovals field if non-nil, zero value otherwise.

### GetHasDeleteApprovalsOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasDeleteApprovalsOk() (*bool, bool)`

GetHasDeleteApprovalsOk returns a tuple with the HasDeleteApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeleteApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasDeleteApprovals(v bool)`

SetHasDeleteApprovals sets HasDeleteApprovals field to given value.

### HasHasDeleteApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasDeleteApprovals() bool`

HasHasDeleteApprovals returns a boolean if a field has been set.

### GetHasReconfigureApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasReconfigureApprovals() bool`

GetHasReconfigureApprovals returns the HasReconfigureApprovals field if non-nil, zero value otherwise.

### GetHasReconfigureApprovalsOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetHasReconfigureApprovalsOk() (*bool, bool)`

GetHasReconfigureApprovalsOk returns a tuple with the HasReconfigureApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReconfigureApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetHasReconfigureApprovals(v bool)`

SetHasReconfigureApprovals sets HasReconfigureApprovals field to given value.

### HasHasReconfigureApprovals

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasHasReconfigureApprovals() bool`

HasHasReconfigureApprovals returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrationTypes200ResponseIntegrationType) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrationTypes200ResponseIntegrationType) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrationTypes200ResponseIntegrationType) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


