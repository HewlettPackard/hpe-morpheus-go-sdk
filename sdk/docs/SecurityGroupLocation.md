# SecurityGroupLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ZonePool** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**GroupLayer** | Pointer to **string** |  | [optional] 

## Methods

### NewSecurityGroupLocation

`func NewSecurityGroupLocation() *SecurityGroupLocation`

NewSecurityGroupLocation instantiates a new SecurityGroupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupLocationWithDefaults

`func NewSecurityGroupLocationWithDefaults() *SecurityGroupLocation`

NewSecurityGroupLocationWithDefaults instantiates a new SecurityGroupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupLocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupLocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupLocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityGroupLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupLocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupLocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *SecurityGroupLocation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityGroupLocation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityGroupLocation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityGroupLocation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *SecurityGroupLocation) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *SecurityGroupLocation) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *SecurityGroupLocation) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *SecurityGroupLocation) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetZone

`func (o *SecurityGroupLocation) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SecurityGroupLocation) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SecurityGroupLocation) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SecurityGroupLocation) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *SecurityGroupLocation) GetZonePool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *SecurityGroupLocation) GetZonePoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *SecurityGroupLocation) SetZonePool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *SecurityGroupLocation) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityGroupLocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityGroupLocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityGroupLocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityGroupLocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *SecurityGroupLocation) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SecurityGroupLocation) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SecurityGroupLocation) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SecurityGroupLocation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *SecurityGroupLocation) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *SecurityGroupLocation) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *SecurityGroupLocation) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *SecurityGroupLocation) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


