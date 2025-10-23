# AddInstanceRequestInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the instance to be created. | 
**Description** | Pointer to **string** | Optional description field. | [optional] 
**Site** | [**AddInstanceRequestInstanceSite**](AddInstanceRequestInstanceSite.md) |  | 
**InstanceType** | [**AddInstanceRequestInstanceInstanceType**](AddInstanceRequestInstanceInstanceType.md) |  | 
**Layout** | [**AddInstanceRequestInstanceLayout**](AddInstanceRequestInstanceLayout.md) |  | 
**Plan** | [**AddInstanceRequestInstancePlan**](AddInstanceRequestInstancePlan.md) |  | 
**InstanceContext** | Pointer to **string** | Environment | [optional] 
**HostName** | Pointer to **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 

## Methods

### NewAddInstanceRequestInstance

`func NewAddInstanceRequestInstance(name string, site AddInstanceRequestInstanceSite, instanceType AddInstanceRequestInstanceInstanceType, layout AddInstanceRequestInstanceLayout, plan AddInstanceRequestInstancePlan, ) *AddInstanceRequestInstance`

NewAddInstanceRequestInstance instantiates a new AddInstanceRequestInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstanceRequestInstanceWithDefaults

`func NewAddInstanceRequestInstanceWithDefaults() *AddInstanceRequestInstance`

NewAddInstanceRequestInstanceWithDefaults instantiates a new AddInstanceRequestInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddInstanceRequestInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddInstanceRequestInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddInstanceRequestInstance) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddInstanceRequestInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddInstanceRequestInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddInstanceRequestInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddInstanceRequestInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSite

`func (o *AddInstanceRequestInstance) GetSite() AddInstanceRequestInstanceSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddInstanceRequestInstance) GetSiteOk() (*AddInstanceRequestInstanceSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddInstanceRequestInstance) SetSite(v AddInstanceRequestInstanceSite)`

SetSite sets Site field to given value.


### GetInstanceType

`func (o *AddInstanceRequestInstance) GetInstanceType() AddInstanceRequestInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddInstanceRequestInstance) GetInstanceTypeOk() (*AddInstanceRequestInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddInstanceRequestInstance) SetInstanceType(v AddInstanceRequestInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetLayout

`func (o *AddInstanceRequestInstance) GetLayout() AddInstanceRequestInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddInstanceRequestInstance) GetLayoutOk() (*AddInstanceRequestInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddInstanceRequestInstance) SetLayout(v AddInstanceRequestInstanceLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *AddInstanceRequestInstance) GetPlan() AddInstanceRequestInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddInstanceRequestInstance) GetPlanOk() (*AddInstanceRequestInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddInstanceRequestInstance) SetPlan(v AddInstanceRequestInstancePlan)`

SetPlan sets Plan field to given value.


### GetInstanceContext

`func (o *AddInstanceRequestInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *AddInstanceRequestInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *AddInstanceRequestInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *AddInstanceRequestInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetHostName

`func (o *AddInstanceRequestInstance) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddInstanceRequestInstance) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddInstanceRequestInstance) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddInstanceRequestInstance) HasHostName() bool`

HasHostName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


