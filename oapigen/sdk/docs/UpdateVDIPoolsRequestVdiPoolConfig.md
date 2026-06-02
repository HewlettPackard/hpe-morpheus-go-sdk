# UpdateVDIPoolsRequestVdiPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of instance | 
**Group** | [**AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup**](AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup.md) |  | 
**Cloud** | [**AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud**](AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud.md) |  | 
**Type** | [**AddVDIPoolsRequestVdiPoolOneOf1ConfigType**](AddVDIPoolsRequestVdiPoolOneOf1ConfigType.md) |  | 
**Layout** | [**AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout**](AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout.md) |  | 
**Plan** | [**AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan**](AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan.md) |  | 

## Methods

### NewUpdateVDIPoolsRequestVdiPoolConfig

`func NewUpdateVDIPoolsRequestVdiPoolConfig(name string, group AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup, cloud AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud, type_ AddVDIPoolsRequestVdiPoolOneOf1ConfigType, layout AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout, plan AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan, ) *UpdateVDIPoolsRequestVdiPoolConfig`

NewUpdateVDIPoolsRequestVdiPoolConfig instantiates a new UpdateVDIPoolsRequestVdiPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetGroup() AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetGroupOk() (*AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetGroup(v AddVDIPoolsRequestVdiPoolOneOf1ConfigGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetCloud() AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetCloudOk() (*AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetCloud(v AddVDIPoolsRequestVdiPoolOneOf1ConfigCloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetType() AddVDIPoolsRequestVdiPoolOneOf1ConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetTypeOk() (*AddVDIPoolsRequestVdiPoolOneOf1ConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetType(v AddVDIPoolsRequestVdiPoolOneOf1ConfigType)`

SetType sets Type field to given value.


### GetLayout

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetLayout() AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetLayoutOk() (*AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetLayout(v AddVDIPoolsRequestVdiPoolOneOf1ConfigLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetPlan() AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) GetPlanOk() (*AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateVDIPoolsRequestVdiPoolConfig) SetPlan(v AddVDIPoolsRequestVdiPoolOneOf1ConfigPlan)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


