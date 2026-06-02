# GetMigration200ResponseMigrationServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Migration Server Status. The possible status values are: &#39;pending&#39;, &#39;precheck&#39;, &#39;running&#39;, &#39;failed&#39;, &#39;completed&#39; | [optional] 
**StatusMessage** | Pointer to **NullableString** | Status Message | [optional] 
**SourceServer** | Pointer to [**GetMigration200ResponseMigrationServersInnerSourceServer**](GetMigration200ResponseMigrationServersInnerSourceServer.md) |  | [optional] 
**DestinationServer** | Pointer to [**GetMigration200ResponseMigrationServersInnerDestinationServer**](GetMigration200ResponseMigrationServersInnerDestinationServer.md) |  | [optional] 

## Methods

### NewGetMigration200ResponseMigrationServersInner

`func NewGetMigration200ResponseMigrationServersInner() *GetMigration200ResponseMigrationServersInner`

NewGetMigration200ResponseMigrationServersInner instantiates a new GetMigration200ResponseMigrationServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetStatus

`func (o *GetMigration200ResponseMigrationServersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMigration200ResponseMigrationServersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMigration200ResponseMigrationServersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMigration200ResponseMigrationServersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetMigration200ResponseMigrationServersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetMigration200ResponseMigrationServersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetMigration200ResponseMigrationServersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetMigration200ResponseMigrationServersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetMigration200ResponseMigrationServersInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetMigration200ResponseMigrationServersInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetSourceServer

`func (o *GetMigration200ResponseMigrationServersInner) GetSourceServer() GetMigration200ResponseMigrationServersInnerSourceServer`

GetSourceServer returns the SourceServer field if non-nil, zero value otherwise.

### GetSourceServerOk

`func (o *GetMigration200ResponseMigrationServersInner) GetSourceServerOk() (*GetMigration200ResponseMigrationServersInnerSourceServer, bool)`

GetSourceServerOk returns a tuple with the SourceServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceServer

`func (o *GetMigration200ResponseMigrationServersInner) SetSourceServer(v GetMigration200ResponseMigrationServersInnerSourceServer)`

SetSourceServer sets SourceServer field to given value.

### HasSourceServer

`func (o *GetMigration200ResponseMigrationServersInner) HasSourceServer() bool`

HasSourceServer returns a boolean if a field has been set.

### GetDestinationServer

`func (o *GetMigration200ResponseMigrationServersInner) GetDestinationServer() GetMigration200ResponseMigrationServersInnerDestinationServer`

GetDestinationServer returns the DestinationServer field if non-nil, zero value otherwise.

### GetDestinationServerOk

`func (o *GetMigration200ResponseMigrationServersInner) GetDestinationServerOk() (*GetMigration200ResponseMigrationServersInnerDestinationServer, bool)`

GetDestinationServerOk returns a tuple with the DestinationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationServer

`func (o *GetMigration200ResponseMigrationServersInner) SetDestinationServer(v GetMigration200ResponseMigrationServersInnerDestinationServer)`

SetDestinationServer sets DestinationServer field to given value.

### HasDestinationServer

`func (o *GetMigration200ResponseMigrationServersInner) HasDestinationServer() bool`

HasDestinationServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


