# ManageHostPlacementRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Placement Strategy | [optional] 
**PreferredServer** | Pointer to [**ManageHostPlacementRequestServerPreferredServer**](ManageHostPlacementRequestServerPreferredServer.md) |  | [optional] 

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

### GetName

`func (o *ManageHostPlacementRequestServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManageHostPlacementRequestServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManageHostPlacementRequestServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManageHostPlacementRequestServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreferredServer

`func (o *ManageHostPlacementRequestServer) GetPreferredServer() ManageHostPlacementRequestServerPreferredServer`

GetPreferredServer returns the PreferredServer field if non-nil, zero value otherwise.

### GetPreferredServerOk

`func (o *ManageHostPlacementRequestServer) GetPreferredServerOk() (*ManageHostPlacementRequestServerPreferredServer, bool)`

GetPreferredServerOk returns a tuple with the PreferredServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredServer

`func (o *ManageHostPlacementRequestServer) SetPreferredServer(v ManageHostPlacementRequestServerPreferredServer)`

SetPreferredServer sets PreferredServer field to given value.

### HasPreferredServer

`func (o *ManageHostPlacementRequestServer) HasPreferredServer() bool`

HasPreferredServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


