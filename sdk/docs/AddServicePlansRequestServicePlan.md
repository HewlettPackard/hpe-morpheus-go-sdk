# AddServicePlansRequestServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service plan name | 
**Code** | **string** | Service plan code, must be unique | 
**Description** | Pointer to **string** | Service plan description | [optional] 
**Editable** | Pointer to **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**MaxStorage** | **float32** | Max storage size in bytes | 
**MaxMemory** | **float32** | Max memory size in bytes | 
**MaxCores** | Pointer to **float32** | Max cores | [optional] 
**MaxDisks** | Pointer to **float32** | Max disks allowed | [optional] 
**ProvisionType** | [**AddClusterLayoutsRequestLayoutProvisionType**](AddClusterLayoutsRequestLayoutProvisionType.md) |  | 
**CustomCores** | Pointer to **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**CustomMaxStorage** | Pointer to **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**CustomMaxDataStorage** | Pointer to **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**CustomMaxMemory** | Pointer to **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**AddVolumes** | Pointer to **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**SortOrder** | Pointer to **float32** | Sort order | [optional] 
**PriceSets** | Pointer to [**[]AddServicePlansRequestServicePlanPriceSetsInner**](AddServicePlansRequestServicePlanPriceSetsInner.md) | List of price sets to include in service plan | [optional] 
**Config** | Pointer to [**AddServicePlansRequestServicePlanConfig**](AddServicePlansRequestServicePlanConfig.md) |  | [optional] 

## Methods

### NewAddServicePlansRequestServicePlan

`func NewAddServicePlansRequestServicePlan(name string, code string, maxStorage float32, maxMemory float32, provisionType AddClusterLayoutsRequestLayoutProvisionType, ) *AddServicePlansRequestServicePlan`

NewAddServicePlansRequestServicePlan instantiates a new AddServicePlansRequestServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlansRequestServicePlanWithDefaults

`func NewAddServicePlansRequestServicePlanWithDefaults() *AddServicePlansRequestServicePlan`

NewAddServicePlansRequestServicePlanWithDefaults instantiates a new AddServicePlansRequestServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddServicePlansRequestServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddServicePlansRequestServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddServicePlansRequestServicePlan) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddServicePlansRequestServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddServicePlansRequestServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddServicePlansRequestServicePlan) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *AddServicePlansRequestServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddServicePlansRequestServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddServicePlansRequestServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddServicePlansRequestServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *AddServicePlansRequestServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddServicePlansRequestServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddServicePlansRequestServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddServicePlansRequestServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddServicePlansRequestServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddServicePlansRequestServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddServicePlansRequestServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxMemory

`func (o *AddServicePlansRequestServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddServicePlansRequestServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddServicePlansRequestServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxCores

`func (o *AddServicePlansRequestServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddServicePlansRequestServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddServicePlansRequestServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddServicePlansRequestServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *AddServicePlansRequestServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *AddServicePlansRequestServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *AddServicePlansRequestServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *AddServicePlansRequestServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetProvisionType

`func (o *AddServicePlansRequestServicePlan) GetProvisionType() AddClusterLayoutsRequestLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddServicePlansRequestServicePlan) GetProvisionTypeOk() (*AddClusterLayoutsRequestLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddServicePlansRequestServicePlan) SetProvisionType(v AddClusterLayoutsRequestLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.


### GetCustomCores

`func (o *AddServicePlansRequestServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *AddServicePlansRequestServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *AddServicePlansRequestServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *AddServicePlansRequestServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *AddServicePlansRequestServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *AddServicePlansRequestServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *AddServicePlansRequestServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *AddServicePlansRequestServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *AddServicePlansRequestServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *AddServicePlansRequestServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddServicePlansRequestServicePlan) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddServicePlansRequestServicePlan) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddServicePlansRequestServicePlan) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddServicePlansRequestServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPriceSets

`func (o *AddServicePlansRequestServicePlan) GetPriceSets() []AddServicePlansRequestServicePlanPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *AddServicePlansRequestServicePlan) GetPriceSetsOk() (*[]AddServicePlansRequestServicePlanPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *AddServicePlansRequestServicePlan) SetPriceSets(v []AddServicePlansRequestServicePlanPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *AddServicePlansRequestServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *AddServicePlansRequestServicePlan) GetConfig() AddServicePlansRequestServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddServicePlansRequestServicePlan) GetConfigOk() (*AddServicePlansRequestServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddServicePlansRequestServicePlan) SetConfig(v AddServicePlansRequestServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddServicePlansRequestServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


