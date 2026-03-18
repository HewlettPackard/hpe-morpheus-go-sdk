# ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentAccess** | Pointer to **bool** |  | [optional] 
**RequestAccess** | Pointer to **bool** |  | [optional] 
**ServiceNowCMDBBusinessObject** | Pointer to **string** |  | [optional] 
**ServiceNowCustomCmdbMapping** | Pointer to **string** |  | [optional] 
**ServiceNowCmdbClassMapping** | Pointer to [**[]ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner.md) |  | [optional] 
**WebServiceImportUrl** | Pointer to **string** |  | [optional] 
**WebServiceImportSysId** | Pointer to **string** |  | [optional] 
**WebServiceOperationUrl** | Pointer to **string** |  | [optional] 
**CmdbMode** | Pointer to **string** |  | [optional] [default to "TABLE"]
**PreparedForSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigWithDefaults

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetIncidentAccess() bool`

GetIncidentAccess returns the IncidentAccess field if non-nil, zero value otherwise.

### GetIncidentAccessOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetIncidentAccessOk() (*bool, bool)`

GetIncidentAccessOk returns a tuple with the IncidentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetIncidentAccess(v bool)`

SetIncidentAccess sets IncidentAccess field to given value.

### HasIncidentAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasIncidentAccess() bool`

HasIncidentAccess returns a boolean if a field has been set.

### GetRequestAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetRequestAccess() bool`

GetRequestAccess returns the RequestAccess field if non-nil, zero value otherwise.

### GetRequestAccessOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetRequestAccessOk() (*bool, bool)`

GetRequestAccessOk returns a tuple with the RequestAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetRequestAccess(v bool)`

SetRequestAccess sets RequestAccess field to given value.

### HasRequestAccess

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasRequestAccess() bool`

HasRequestAccess returns a boolean if a field has been set.

### GetServiceNowCMDBBusinessObject

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCMDBBusinessObject() string`

GetServiceNowCMDBBusinessObject returns the ServiceNowCMDBBusinessObject field if non-nil, zero value otherwise.

### GetServiceNowCMDBBusinessObjectOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCMDBBusinessObjectOk() (*string, bool)`

GetServiceNowCMDBBusinessObjectOk returns a tuple with the ServiceNowCMDBBusinessObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCMDBBusinessObject

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCMDBBusinessObject(v string)`

SetServiceNowCMDBBusinessObject sets ServiceNowCMDBBusinessObject field to given value.

### HasServiceNowCMDBBusinessObject

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasServiceNowCMDBBusinessObject() bool`

HasServiceNowCMDBBusinessObject returns a boolean if a field has been set.

### GetServiceNowCustomCmdbMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCustomCmdbMapping() string`

GetServiceNowCustomCmdbMapping returns the ServiceNowCustomCmdbMapping field if non-nil, zero value otherwise.

### GetServiceNowCustomCmdbMappingOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCustomCmdbMappingOk() (*string, bool)`

GetServiceNowCustomCmdbMappingOk returns a tuple with the ServiceNowCustomCmdbMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCustomCmdbMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCustomCmdbMapping(v string)`

SetServiceNowCustomCmdbMapping sets ServiceNowCustomCmdbMapping field to given value.

### HasServiceNowCustomCmdbMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasServiceNowCustomCmdbMapping() bool`

HasServiceNowCustomCmdbMapping returns a boolean if a field has been set.

### GetServiceNowCmdbClassMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCmdbClassMapping() []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner`

GetServiceNowCmdbClassMapping returns the ServiceNowCmdbClassMapping field if non-nil, zero value otherwise.

### GetServiceNowCmdbClassMappingOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCmdbClassMappingOk() (*[]ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner, bool)`

GetServiceNowCmdbClassMappingOk returns a tuple with the ServiceNowCmdbClassMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCmdbClassMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCmdbClassMapping(v []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner)`

SetServiceNowCmdbClassMapping sets ServiceNowCmdbClassMapping field to given value.

### HasServiceNowCmdbClassMapping

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasServiceNowCmdbClassMapping() bool`

HasServiceNowCmdbClassMapping returns a boolean if a field has been set.

### GetWebServiceImportUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportUrl() string`

GetWebServiceImportUrl returns the WebServiceImportUrl field if non-nil, zero value otherwise.

### GetWebServiceImportUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportUrlOk() (*string, bool)`

GetWebServiceImportUrlOk returns a tuple with the WebServiceImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceImportUrl(v string)`

SetWebServiceImportUrl sets WebServiceImportUrl field to given value.

### HasWebServiceImportUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasWebServiceImportUrl() bool`

HasWebServiceImportUrl returns a boolean if a field has been set.

### GetWebServiceImportSysId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportSysId() string`

GetWebServiceImportSysId returns the WebServiceImportSysId field if non-nil, zero value otherwise.

### GetWebServiceImportSysIdOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportSysIdOk() (*string, bool)`

GetWebServiceImportSysIdOk returns a tuple with the WebServiceImportSysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportSysId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceImportSysId(v string)`

SetWebServiceImportSysId sets WebServiceImportSysId field to given value.

### HasWebServiceImportSysId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasWebServiceImportSysId() bool`

HasWebServiceImportSysId returns a boolean if a field has been set.

### GetWebServiceOperationUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceOperationUrl() string`

GetWebServiceOperationUrl returns the WebServiceOperationUrl field if non-nil, zero value otherwise.

### GetWebServiceOperationUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceOperationUrlOk() (*string, bool)`

GetWebServiceOperationUrlOk returns a tuple with the WebServiceOperationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceOperationUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceOperationUrl(v string)`

SetWebServiceOperationUrl sets WebServiceOperationUrl field to given value.

### HasWebServiceOperationUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasWebServiceOperationUrl() bool`

HasWebServiceOperationUrl returns a boolean if a field has been set.

### GetCmdbMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetCmdbMode() string`

GetCmdbMode returns the CmdbMode field if non-nil, zero value otherwise.

### GetCmdbModeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetCmdbModeOk() (*string, bool)`

GetCmdbModeOk returns a tuple with the CmdbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetCmdbMode(v string)`

SetCmdbMode sets CmdbMode field to given value.

### HasCmdbMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasCmdbMode() bool`

HasCmdbMode returns a boolean if a field has been set.

### GetPreparedForSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetPreparedForSync() bool`

GetPreparedForSync returns the PreparedForSync field if non-nil, zero value otherwise.

### GetPreparedForSyncOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetPreparedForSyncOk() (*bool, bool)`

GetPreparedForSyncOk returns a tuple with the PreparedForSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedForSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetPreparedForSync(v bool)`

SetPreparedForSync sets PreparedForSync field to given value.

### HasPreparedForSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) HasPreparedForSync() bool`

HasPreparedForSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


