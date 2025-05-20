# SetupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceName** | **string** | Appliance Name. Choose a name for your Morpheus Appliance.  This is stored in the Appliance Settings. | 
**ApplianceUrl** | Pointer to **string** | Appliance URL. Specify the full URL for your Morpheus Appliance. This is stored in the Appliance Settings. | [optional] 
**AccountName** | **string** | Name of the Master Tenant account being created. | 
**FirstName** | Pointer to **string** | First Name for the System Admin user being created. | [optional] 
**LastName** | Pointer to **string** | Last Name for the System Admin user being created. | [optional] 
**Username** | **string** | Username for the System Admin user being created. | 
**Email** | **string** | Email for the System Admin user being created. | 
**Password** | **string** | Password for the System Admin user being created. | 
**Hubmode** | Pointer to **string** | Hub Mode. Determines if and how the appliance should connect with the Morpheus Hub. The default value (skip) means do not connect with the hub, and you will be installing your license manually. If you login or register with the hub then a Community Edition license will be installed automatically. | [optional] [default to "skip"]
**LicenseKey** | Pointer to **string** | License Key to install on setup. | [optional] 

## Methods

### NewSetupRequest

`func NewSetupRequest(applianceName string, accountName string, username string, email string, password string, ) *SetupRequest`

NewSetupRequest instantiates a new SetupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupRequestWithDefaults

`func NewSetupRequestWithDefaults() *SetupRequest`

NewSetupRequestWithDefaults instantiates a new SetupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceName

`func (o *SetupRequest) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *SetupRequest) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *SetupRequest) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.


### GetApplianceUrl

`func (o *SetupRequest) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *SetupRequest) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *SetupRequest) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *SetupRequest) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetAccountName

`func (o *SetupRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SetupRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SetupRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetFirstName

`func (o *SetupRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SetupRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SetupRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SetupRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SetupRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SetupRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SetupRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SetupRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *SetupRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SetupRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SetupRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *SetupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SetupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SetupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SetupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHubmode

`func (o *SetupRequest) GetHubmode() string`

GetHubmode returns the Hubmode field if non-nil, zero value otherwise.

### GetHubmodeOk

`func (o *SetupRequest) GetHubmodeOk() (*string, bool)`

GetHubmodeOk returns a tuple with the Hubmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubmode

`func (o *SetupRequest) SetHubmode(v string)`

SetHubmode sets Hubmode field to given value.

### HasHubmode

`func (o *SetupRequest) HasHubmode() bool`

HasHubmode returns a boolean if a field has been set.

### GetLicenseKey

`func (o *SetupRequest) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *SetupRequest) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *SetupRequest) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *SetupRequest) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


