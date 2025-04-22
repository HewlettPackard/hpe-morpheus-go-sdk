# AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation

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

### NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation

`func NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation() *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation`

NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation instantiates a new AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocationWithDefaults

`func NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocationWithDefaults() *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation`

NewAddSecurityGroupLocations200ResponseAllOfSecurityGroupLocationWithDefaults instantiates a new AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetZone

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetZonePool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetZonePoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetZonePool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetStatus

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


