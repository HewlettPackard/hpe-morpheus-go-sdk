# GetIntegrations200ResponseAllOfIntegrationOneOf15Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentAccess** | Pointer to **bool** |  | [optional] 
**RequestAccess** | Pointer to **bool** |  | [optional] 
**ServiceNowCMDBBusinessObject** | Pointer to **string** |  | [optional] 
**ServiceNowCustomCmdbMapping** | Pointer to **string** |  | [optional] 
**ServiceNowCmdbClassMapping** | Pointer to [**[]GetIntegrations200ResponseAllOfIntegrationOneOf15ConfigServiceNowCmdbClassMappingInner**](GetIntegrations200ResponseAllOfIntegrationOneOf15ConfigServiceNowCmdbClassMappingInner.md) |  | [optional] 
**WebServiceImportUrl** | Pointer to **string** |  | [optional] 
**WebServiceImportSysId** | Pointer to **string** |  | [optional] 
**WebServiceOperationUrl** | Pointer to **string** |  | [optional] 
**CmdbMode** | Pointer to **string** |  | [optional] [default to "TABLE"]
**PreparedForSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf15Config

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf15Config() *GetIntegrations200ResponseAllOfIntegrationOneOf15Config`

NewGetIntegrations200ResponseAllOfIntegrationOneOf15Config instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf15Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIncidentAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetIncidentAccess() bool`

GetIncidentAccess returns the IncidentAccess field if non-nil, zero value otherwise.

### GetIncidentAccessOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetIncidentAccessOk() (*bool, bool)`

GetIncidentAccessOk returns a tuple with the IncidentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetIncidentAccess(v bool)`

SetIncidentAccess sets IncidentAccess field to given value.

### HasIncidentAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasIncidentAccess() bool`

HasIncidentAccess returns a boolean if a field has been set.

### GetRequestAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetRequestAccess() bool`

GetRequestAccess returns the RequestAccess field if non-nil, zero value otherwise.

### GetRequestAccessOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetRequestAccessOk() (*bool, bool)`

GetRequestAccessOk returns a tuple with the RequestAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetRequestAccess(v bool)`

SetRequestAccess sets RequestAccess field to given value.

### HasRequestAccess

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasRequestAccess() bool`

HasRequestAccess returns a boolean if a field has been set.

### GetServiceNowCMDBBusinessObject

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCMDBBusinessObject() string`

GetServiceNowCMDBBusinessObject returns the ServiceNowCMDBBusinessObject field if non-nil, zero value otherwise.

### GetServiceNowCMDBBusinessObjectOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCMDBBusinessObjectOk() (*string, bool)`

GetServiceNowCMDBBusinessObjectOk returns a tuple with the ServiceNowCMDBBusinessObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCMDBBusinessObject

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetServiceNowCMDBBusinessObject(v string)`

SetServiceNowCMDBBusinessObject sets ServiceNowCMDBBusinessObject field to given value.

### HasServiceNowCMDBBusinessObject

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasServiceNowCMDBBusinessObject() bool`

HasServiceNowCMDBBusinessObject returns a boolean if a field has been set.

### GetServiceNowCustomCmdbMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCustomCmdbMapping() string`

GetServiceNowCustomCmdbMapping returns the ServiceNowCustomCmdbMapping field if non-nil, zero value otherwise.

### GetServiceNowCustomCmdbMappingOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCustomCmdbMappingOk() (*string, bool)`

GetServiceNowCustomCmdbMappingOk returns a tuple with the ServiceNowCustomCmdbMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCustomCmdbMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetServiceNowCustomCmdbMapping(v string)`

SetServiceNowCustomCmdbMapping sets ServiceNowCustomCmdbMapping field to given value.

### HasServiceNowCustomCmdbMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasServiceNowCustomCmdbMapping() bool`

HasServiceNowCustomCmdbMapping returns a boolean if a field has been set.

### GetServiceNowCmdbClassMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCmdbClassMapping() []GetIntegrations200ResponseAllOfIntegrationOneOf15ConfigServiceNowCmdbClassMappingInner`

GetServiceNowCmdbClassMapping returns the ServiceNowCmdbClassMapping field if non-nil, zero value otherwise.

### GetServiceNowCmdbClassMappingOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetServiceNowCmdbClassMappingOk() (*[]GetIntegrations200ResponseAllOfIntegrationOneOf15ConfigServiceNowCmdbClassMappingInner, bool)`

GetServiceNowCmdbClassMappingOk returns a tuple with the ServiceNowCmdbClassMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCmdbClassMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetServiceNowCmdbClassMapping(v []GetIntegrations200ResponseAllOfIntegrationOneOf15ConfigServiceNowCmdbClassMappingInner)`

SetServiceNowCmdbClassMapping sets ServiceNowCmdbClassMapping field to given value.

### HasServiceNowCmdbClassMapping

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasServiceNowCmdbClassMapping() bool`

HasServiceNowCmdbClassMapping returns a boolean if a field has been set.

### GetWebServiceImportUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceImportUrl() string`

GetWebServiceImportUrl returns the WebServiceImportUrl field if non-nil, zero value otherwise.

### GetWebServiceImportUrlOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceImportUrlOk() (*string, bool)`

GetWebServiceImportUrlOk returns a tuple with the WebServiceImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetWebServiceImportUrl(v string)`

SetWebServiceImportUrl sets WebServiceImportUrl field to given value.

### HasWebServiceImportUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasWebServiceImportUrl() bool`

HasWebServiceImportUrl returns a boolean if a field has been set.

### GetWebServiceImportSysId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceImportSysId() string`

GetWebServiceImportSysId returns the WebServiceImportSysId field if non-nil, zero value otherwise.

### GetWebServiceImportSysIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceImportSysIdOk() (*string, bool)`

GetWebServiceImportSysIdOk returns a tuple with the WebServiceImportSysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportSysId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetWebServiceImportSysId(v string)`

SetWebServiceImportSysId sets WebServiceImportSysId field to given value.

### HasWebServiceImportSysId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasWebServiceImportSysId() bool`

HasWebServiceImportSysId returns a boolean if a field has been set.

### GetWebServiceOperationUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceOperationUrl() string`

GetWebServiceOperationUrl returns the WebServiceOperationUrl field if non-nil, zero value otherwise.

### GetWebServiceOperationUrlOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetWebServiceOperationUrlOk() (*string, bool)`

GetWebServiceOperationUrlOk returns a tuple with the WebServiceOperationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceOperationUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetWebServiceOperationUrl(v string)`

SetWebServiceOperationUrl sets WebServiceOperationUrl field to given value.

### HasWebServiceOperationUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasWebServiceOperationUrl() bool`

HasWebServiceOperationUrl returns a boolean if a field has been set.

### GetCmdbMode

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetCmdbMode() string`

GetCmdbMode returns the CmdbMode field if non-nil, zero value otherwise.

### GetCmdbModeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetCmdbModeOk() (*string, bool)`

GetCmdbModeOk returns a tuple with the CmdbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbMode

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetCmdbMode(v string)`

SetCmdbMode sets CmdbMode field to given value.

### HasCmdbMode

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasCmdbMode() bool`

HasCmdbMode returns a boolean if a field has been set.

### GetPreparedForSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetPreparedForSync() bool`

GetPreparedForSync returns the PreparedForSync field if non-nil, zero value otherwise.

### GetPreparedForSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) GetPreparedForSyncOk() (*bool, bool)`

GetPreparedForSyncOk returns a tuple with the PreparedForSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedForSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) SetPreparedForSync(v bool)`

SetPreparedForSync sets PreparedForSync field to given value.

### HasPreparedForSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf15Config) HasPreparedForSync() bool`

HasPreparedForSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


