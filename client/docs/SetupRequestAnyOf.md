# SetupRequestAnyOf

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

### NewSetupRequestAnyOf

`func NewSetupRequestAnyOf(applianceName string, accountName string, username string, email string, password string, ) *SetupRequestAnyOf`

NewSetupRequestAnyOf instantiates a new SetupRequestAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupRequestAnyOfWithDefaults

`func NewSetupRequestAnyOfWithDefaults() *SetupRequestAnyOf`

NewSetupRequestAnyOfWithDefaults instantiates a new SetupRequestAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceName

`func (o *SetupRequestAnyOf) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *SetupRequestAnyOf) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *SetupRequestAnyOf) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.


### GetApplianceUrl

`func (o *SetupRequestAnyOf) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *SetupRequestAnyOf) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *SetupRequestAnyOf) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *SetupRequestAnyOf) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetAccountName

`func (o *SetupRequestAnyOf) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SetupRequestAnyOf) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SetupRequestAnyOf) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetFirstName

`func (o *SetupRequestAnyOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SetupRequestAnyOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SetupRequestAnyOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SetupRequestAnyOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SetupRequestAnyOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SetupRequestAnyOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SetupRequestAnyOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SetupRequestAnyOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *SetupRequestAnyOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SetupRequestAnyOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SetupRequestAnyOf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *SetupRequestAnyOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SetupRequestAnyOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SetupRequestAnyOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SetupRequestAnyOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetupRequestAnyOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetupRequestAnyOf) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHubmode

`func (o *SetupRequestAnyOf) GetHubmode() string`

GetHubmode returns the Hubmode field if non-nil, zero value otherwise.

### GetHubmodeOk

`func (o *SetupRequestAnyOf) GetHubmodeOk() (*string, bool)`

GetHubmodeOk returns a tuple with the Hubmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubmode

`func (o *SetupRequestAnyOf) SetHubmode(v string)`

SetHubmode sets Hubmode field to given value.

### HasHubmode

`func (o *SetupRequestAnyOf) HasHubmode() bool`

HasHubmode returns a boolean if a field has been set.

### GetLicenseKey

`func (o *SetupRequestAnyOf) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *SetupRequestAnyOf) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *SetupRequestAnyOf) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *SetupRequestAnyOf) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


