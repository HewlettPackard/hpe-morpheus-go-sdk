# ManageHostPlacementRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementStrategy** | Pointer to **string** | Placement Strategy | [optional] 
**PreferredParentServer** | Pointer to [**ManageHostPlacementRequestServerPreferredParentServer**](ManageHostPlacementRequestServerPreferredParentServer.md) |  | [optional] 

## Methods

### NewManageHostPlacementRequestServer

`func NewManageHostPlacementRequestServer() *ManageHostPlacementRequestServer`

NewManageHostPlacementRequestServer instantiates a new ManageHostPlacementRequestServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageHostPlacementRequestServerWithDefaults

`func NewManageHostPlacementRequestServerWithDefaults() *ManageHostPlacementRequestServer`

NewManageHostPlacementRequestServerWithDefaults instantiates a new ManageHostPlacementRequestServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementStrategy

`func (o *ManageHostPlacementRequestServer) GetPlacementStrategy() string`

GetPlacementStrategy returns the PlacementStrategy field if non-nil, zero value otherwise.

### GetPlacementStrategyOk

`func (o *ManageHostPlacementRequestServer) GetPlacementStrategyOk() (*string, bool)`

GetPlacementStrategyOk returns a tuple with the PlacementStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementStrategy

`func (o *ManageHostPlacementRequestServer) SetPlacementStrategy(v string)`

SetPlacementStrategy sets PlacementStrategy field to given value.

### HasPlacementStrategy

`func (o *ManageHostPlacementRequestServer) HasPlacementStrategy() bool`

HasPlacementStrategy returns a boolean if a field has been set.

### GetPreferredParentServer

`func (o *ManageHostPlacementRequestServer) GetPreferredParentServer() ManageHostPlacementRequestServerPreferredParentServer`

GetPreferredParentServer returns the PreferredParentServer field if non-nil, zero value otherwise.

### GetPreferredParentServerOk

`func (o *ManageHostPlacementRequestServer) GetPreferredParentServerOk() (*ManageHostPlacementRequestServerPreferredParentServer, bool)`

GetPreferredParentServerOk returns a tuple with the PreferredParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredParentServer

`func (o *ManageHostPlacementRequestServer) SetPreferredParentServer(v ManageHostPlacementRequestServerPreferredParentServer)`

SetPreferredParentServer sets PreferredParentServer field to given value.

### HasPreferredParentServer

`func (o *ManageHostPlacementRequestServer) HasPreferredParentServer() bool`

HasPreferredParentServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


