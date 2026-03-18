# UpdateRoleReportTypeAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportTypeId** | **int32** | &#x60;id&#x60; of the report type | 
**Access** | **string** | The new access level. | 
**AllReportTypes** | **bool** | Apply to all report types | 

## Methods

### NewUpdateRoleReportTypeAccessRequest

`func NewUpdateRoleReportTypeAccessRequest(reportTypeId int32, access string, allReportTypes bool, ) *UpdateRoleReportTypeAccessRequest`

NewUpdateRoleReportTypeAccessRequest instantiates a new UpdateRoleReportTypeAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleReportTypeAccessRequestWithDefaults

`func NewUpdateRoleReportTypeAccessRequestWithDefaults() *UpdateRoleReportTypeAccessRequest`

NewUpdateRoleReportTypeAccessRequestWithDefaults instantiates a new UpdateRoleReportTypeAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportTypeId

`func (o *UpdateRoleReportTypeAccessRequest) GetReportTypeId() int32`

GetReportTypeId returns the ReportTypeId field if non-nil, zero value otherwise.

### GetReportTypeIdOk

`func (o *UpdateRoleReportTypeAccessRequest) GetReportTypeIdOk() (*int32, bool)`

GetReportTypeIdOk returns a tuple with the ReportTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypeId

`func (o *UpdateRoleReportTypeAccessRequest) SetReportTypeId(v int32)`

SetReportTypeId sets ReportTypeId field to given value.


### GetAccess

`func (o *UpdateRoleReportTypeAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleReportTypeAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleReportTypeAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllReportTypes

`func (o *UpdateRoleReportTypeAccessRequest) GetAllReportTypes() bool`

GetAllReportTypes returns the AllReportTypes field if non-nil, zero value otherwise.

### GetAllReportTypesOk

`func (o *UpdateRoleReportTypeAccessRequest) GetAllReportTypesOk() (*bool, bool)`

GetAllReportTypesOk returns a tuple with the AllReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReportTypes

`func (o *UpdateRoleReportTypeAccessRequest) SetAllReportTypes(v bool)`

SetAllReportTypes sets AllReportTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


