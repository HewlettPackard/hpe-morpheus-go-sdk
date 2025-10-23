# UpdateVDIAppsRequestVdiApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | VDI App name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**IconPath** | Pointer to ***os.File** | Icon Path. A relative location of an icon image | [optional] 
**LaunchPrefix** | Pointer to **string** | The RDS App Name Prefix.  Must start with &#39;||&#39; | [optional] 

## Methods

### NewUpdateVDIAppsRequestVdiApp

`func NewUpdateVDIAppsRequestVdiApp() *UpdateVDIAppsRequestVdiApp`

NewUpdateVDIAppsRequestVdiApp instantiates a new UpdateVDIAppsRequestVdiApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVDIAppsRequestVdiAppWithDefaults

`func NewUpdateVDIAppsRequestVdiAppWithDefaults() *UpdateVDIAppsRequestVdiApp`

NewUpdateVDIAppsRequestVdiAppWithDefaults instantiates a new UpdateVDIAppsRequestVdiApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVDIAppsRequestVdiApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDIAppsRequestVdiApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDIAppsRequestVdiApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDIAppsRequestVdiApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateVDIAppsRequestVdiApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVDIAppsRequestVdiApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVDIAppsRequestVdiApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVDIAppsRequestVdiApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconPath

`func (o *UpdateVDIAppsRequestVdiApp) GetIconPath() *os.File`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateVDIAppsRequestVdiApp) GetIconPathOk() (**os.File, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateVDIAppsRequestVdiApp) SetIconPath(v *os.File)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateVDIAppsRequestVdiApp) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLaunchPrefix

`func (o *UpdateVDIAppsRequestVdiApp) GetLaunchPrefix() string`

GetLaunchPrefix returns the LaunchPrefix field if non-nil, zero value otherwise.

### GetLaunchPrefixOk

`func (o *UpdateVDIAppsRequestVdiApp) GetLaunchPrefixOk() (*string, bool)`

GetLaunchPrefixOk returns a tuple with the LaunchPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchPrefix

`func (o *UpdateVDIAppsRequestVdiApp) SetLaunchPrefix(v string)`

SetLaunchPrefix sets LaunchPrefix field to given value.

### HasLaunchPrefix

`func (o *UpdateVDIAppsRequestVdiApp) HasLaunchPrefix() bool`

HasLaunchPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


