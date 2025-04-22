# UpdateServicePlansRequestServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Service plan name | [optional] 
**Code** | Pointer to **string** | Service plan code, must be unique | [optional] 
**Description** | Pointer to **string** | Service plan description | [optional] 
**Editable** | Pointer to **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**MaxStorage** | Pointer to **float32** | Max storage size in bytes | [optional] 
**MaxMemory** | Pointer to **float32** | Max memory size in bytes | [optional] 
**MaxCores** | Pointer to **float32** | Max cores | [optional] 
**MaxDisks** | Pointer to **float32** | Max disks allowed | [optional] 
**ProvisionType** | Pointer to [**AddClusterLayoutsRequestLayoutProvisionType**](AddClusterLayoutsRequestLayoutProvisionType.md) |  | [optional] 
**CustomCores** | Pointer to **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**CustomMaxStorage** | Pointer to **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**CustomMaxDataStorage** | Pointer to **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**CustomMaxMemory** | Pointer to **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**AddVolumes** | Pointer to **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**SortOrder** | Pointer to **float32** | Sort order | [optional] 
**PriceSets** | Pointer to [**[]AddServicePlansRequestServicePlanPriceSetsInner**](AddServicePlansRequestServicePlanPriceSetsInner.md) | List of price sets to include in service plan | [optional] 
**Config** | Pointer to [**UpdateServicePlansRequestServicePlanConfig**](UpdateServicePlansRequestServicePlanConfig.md) |  | [optional] 

## Methods

### NewUpdateServicePlansRequestServicePlan

`func NewUpdateServicePlansRequestServicePlan() *UpdateServicePlansRequestServicePlan`

NewUpdateServicePlansRequestServicePlan instantiates a new UpdateServicePlansRequestServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServicePlansRequestServicePlanWithDefaults

`func NewUpdateServicePlansRequestServicePlanWithDefaults() *UpdateServicePlansRequestServicePlan`

NewUpdateServicePlansRequestServicePlanWithDefaults instantiates a new UpdateServicePlansRequestServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServicePlansRequestServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServicePlansRequestServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServicePlansRequestServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServicePlansRequestServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateServicePlansRequestServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateServicePlansRequestServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateServicePlansRequestServicePlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateServicePlansRequestServicePlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServicePlansRequestServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServicePlansRequestServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServicePlansRequestServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServicePlansRequestServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *UpdateServicePlansRequestServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateServicePlansRequestServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateServicePlansRequestServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateServicePlansRequestServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateServicePlansRequestServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateServicePlansRequestServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *UpdateServicePlansRequestServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdateServicePlansRequestServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdateServicePlansRequestServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *UpdateServicePlansRequestServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *UpdateServicePlansRequestServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *UpdateServicePlansRequestServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *UpdateServicePlansRequestServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *UpdateServicePlansRequestServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetProvisionType

`func (o *UpdateServicePlansRequestServicePlan) GetProvisionType() AddClusterLayoutsRequestLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *UpdateServicePlansRequestServicePlan) GetProvisionTypeOk() (*AddClusterLayoutsRequestLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *UpdateServicePlansRequestServicePlan) SetProvisionType(v AddClusterLayoutsRequestLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *UpdateServicePlansRequestServicePlan) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetCustomCores

`func (o *UpdateServicePlansRequestServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *UpdateServicePlansRequestServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *UpdateServicePlansRequestServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *UpdateServicePlansRequestServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *UpdateServicePlansRequestServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *UpdateServicePlansRequestServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *UpdateServicePlansRequestServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *UpdateServicePlansRequestServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *UpdateServicePlansRequestServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *UpdateServicePlansRequestServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *UpdateServicePlansRequestServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *UpdateServicePlansRequestServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *UpdateServicePlansRequestServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateServicePlansRequestServicePlan) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateServicePlansRequestServicePlan) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateServicePlansRequestServicePlan) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateServicePlansRequestServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPriceSets

`func (o *UpdateServicePlansRequestServicePlan) GetPriceSets() []AddServicePlansRequestServicePlanPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *UpdateServicePlansRequestServicePlan) GetPriceSetsOk() (*[]AddServicePlansRequestServicePlanPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *UpdateServicePlansRequestServicePlan) SetPriceSets(v []AddServicePlansRequestServicePlanPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *UpdateServicePlansRequestServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateServicePlansRequestServicePlan) GetConfig() UpdateServicePlansRequestServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateServicePlansRequestServicePlan) GetConfigOk() (*UpdateServicePlansRequestServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateServicePlansRequestServicePlan) SetConfig(v UpdateServicePlansRequestServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateServicePlansRequestServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


