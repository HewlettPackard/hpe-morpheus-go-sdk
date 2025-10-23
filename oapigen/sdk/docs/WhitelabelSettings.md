# WhitelabelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ApplianceName** | Pointer to **string** |  | [optional] 
**DisableSupportMenu** | Pointer to **bool** |  | [optional] 
**HeaderLogo** | Pointer to **string** |  | [optional] 
**FooterLogo** | Pointer to **string** |  | [optional] 
**LoginLogo** | Pointer to **string** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**HeaderBgColor** | Pointer to **string** |  | [optional] 
**HeaderFgColor** | Pointer to **string** |  | [optional] 
**NavBgColor** | Pointer to **string** |  | [optional] 
**NavFgColor** | Pointer to **string** |  | [optional] 
**NavHoverColor** | Pointer to **string** |  | [optional] 
**PrimaryButtonBgColor** | Pointer to **string** |  | [optional] 
**PrimaryButtonFgColor** | Pointer to **string** |  | [optional] 
**PrimaryButtonHoverBgColor** | Pointer to **string** |  | [optional] 
**PrimaryButtonHoverFgColor** | Pointer to **string** |  | [optional] 
**FooterBgColor** | Pointer to **string** |  | [optional] 
**FooterFgColor** | Pointer to **string** |  | [optional] 
**LoginBgColor** | Pointer to **string** |  | [optional] 
**OverrideCss** | Pointer to **string** |  | [optional] 
**CopyrightString** | Pointer to **string** |  | [optional] 
**TermsOfUse** | Pointer to **string** |  | [optional] 
**PrivacyPolicy** | Pointer to **string** |  | [optional] 
**SupportMenuLinks** | Pointer to [**[]ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner**](ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner.md) |  | [optional] 

## Methods

### NewWhitelabelSettings

`func NewWhitelabelSettings() *WhitelabelSettings`

NewWhitelabelSettings instantiates a new WhitelabelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhitelabelSettingsWithDefaults

`func NewWhitelabelSettingsWithDefaults() *WhitelabelSettings`

NewWhitelabelSettingsWithDefaults instantiates a new WhitelabelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WhitelabelSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WhitelabelSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WhitelabelSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WhitelabelSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetApplianceName

`func (o *WhitelabelSettings) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *WhitelabelSettings) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *WhitelabelSettings) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.

### HasApplianceName

`func (o *WhitelabelSettings) HasApplianceName() bool`

HasApplianceName returns a boolean if a field has been set.

### GetDisableSupportMenu

`func (o *WhitelabelSettings) GetDisableSupportMenu() bool`

GetDisableSupportMenu returns the DisableSupportMenu field if non-nil, zero value otherwise.

### GetDisableSupportMenuOk

`func (o *WhitelabelSettings) GetDisableSupportMenuOk() (*bool, bool)`

GetDisableSupportMenuOk returns a tuple with the DisableSupportMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSupportMenu

`func (o *WhitelabelSettings) SetDisableSupportMenu(v bool)`

SetDisableSupportMenu sets DisableSupportMenu field to given value.

### HasDisableSupportMenu

`func (o *WhitelabelSettings) HasDisableSupportMenu() bool`

HasDisableSupportMenu returns a boolean if a field has been set.

### GetHeaderLogo

`func (o *WhitelabelSettings) GetHeaderLogo() string`

GetHeaderLogo returns the HeaderLogo field if non-nil, zero value otherwise.

### GetHeaderLogoOk

`func (o *WhitelabelSettings) GetHeaderLogoOk() (*string, bool)`

GetHeaderLogoOk returns a tuple with the HeaderLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderLogo

`func (o *WhitelabelSettings) SetHeaderLogo(v string)`

SetHeaderLogo sets HeaderLogo field to given value.

### HasHeaderLogo

`func (o *WhitelabelSettings) HasHeaderLogo() bool`

HasHeaderLogo returns a boolean if a field has been set.

### GetFooterLogo

`func (o *WhitelabelSettings) GetFooterLogo() string`

GetFooterLogo returns the FooterLogo field if non-nil, zero value otherwise.

### GetFooterLogoOk

`func (o *WhitelabelSettings) GetFooterLogoOk() (*string, bool)`

GetFooterLogoOk returns a tuple with the FooterLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLogo

`func (o *WhitelabelSettings) SetFooterLogo(v string)`

SetFooterLogo sets FooterLogo field to given value.

### HasFooterLogo

`func (o *WhitelabelSettings) HasFooterLogo() bool`

HasFooterLogo returns a boolean if a field has been set.

### GetLoginLogo

`func (o *WhitelabelSettings) GetLoginLogo() string`

GetLoginLogo returns the LoginLogo field if non-nil, zero value otherwise.

### GetLoginLogoOk

`func (o *WhitelabelSettings) GetLoginLogoOk() (*string, bool)`

GetLoginLogoOk returns a tuple with the LoginLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginLogo

`func (o *WhitelabelSettings) SetLoginLogo(v string)`

SetLoginLogo sets LoginLogo field to given value.

### HasLoginLogo

`func (o *WhitelabelSettings) HasLoginLogo() bool`

HasLoginLogo returns a boolean if a field has been set.

### GetFavicon

`func (o *WhitelabelSettings) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *WhitelabelSettings) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *WhitelabelSettings) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *WhitelabelSettings) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetHeaderBgColor

`func (o *WhitelabelSettings) GetHeaderBgColor() string`

GetHeaderBgColor returns the HeaderBgColor field if non-nil, zero value otherwise.

### GetHeaderBgColorOk

`func (o *WhitelabelSettings) GetHeaderBgColorOk() (*string, bool)`

GetHeaderBgColorOk returns a tuple with the HeaderBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBgColor

`func (o *WhitelabelSettings) SetHeaderBgColor(v string)`

SetHeaderBgColor sets HeaderBgColor field to given value.

### HasHeaderBgColor

`func (o *WhitelabelSettings) HasHeaderBgColor() bool`

HasHeaderBgColor returns a boolean if a field has been set.

### GetHeaderFgColor

`func (o *WhitelabelSettings) GetHeaderFgColor() string`

GetHeaderFgColor returns the HeaderFgColor field if non-nil, zero value otherwise.

### GetHeaderFgColorOk

`func (o *WhitelabelSettings) GetHeaderFgColorOk() (*string, bool)`

GetHeaderFgColorOk returns a tuple with the HeaderFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFgColor

`func (o *WhitelabelSettings) SetHeaderFgColor(v string)`

SetHeaderFgColor sets HeaderFgColor field to given value.

### HasHeaderFgColor

`func (o *WhitelabelSettings) HasHeaderFgColor() bool`

HasHeaderFgColor returns a boolean if a field has been set.

### GetNavBgColor

`func (o *WhitelabelSettings) GetNavBgColor() string`

GetNavBgColor returns the NavBgColor field if non-nil, zero value otherwise.

### GetNavBgColorOk

`func (o *WhitelabelSettings) GetNavBgColorOk() (*string, bool)`

GetNavBgColorOk returns a tuple with the NavBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavBgColor

`func (o *WhitelabelSettings) SetNavBgColor(v string)`

SetNavBgColor sets NavBgColor field to given value.

### HasNavBgColor

`func (o *WhitelabelSettings) HasNavBgColor() bool`

HasNavBgColor returns a boolean if a field has been set.

### GetNavFgColor

`func (o *WhitelabelSettings) GetNavFgColor() string`

GetNavFgColor returns the NavFgColor field if non-nil, zero value otherwise.

### GetNavFgColorOk

`func (o *WhitelabelSettings) GetNavFgColorOk() (*string, bool)`

GetNavFgColorOk returns a tuple with the NavFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavFgColor

`func (o *WhitelabelSettings) SetNavFgColor(v string)`

SetNavFgColor sets NavFgColor field to given value.

### HasNavFgColor

`func (o *WhitelabelSettings) HasNavFgColor() bool`

HasNavFgColor returns a boolean if a field has been set.

### GetNavHoverColor

`func (o *WhitelabelSettings) GetNavHoverColor() string`

GetNavHoverColor returns the NavHoverColor field if non-nil, zero value otherwise.

### GetNavHoverColorOk

`func (o *WhitelabelSettings) GetNavHoverColorOk() (*string, bool)`

GetNavHoverColorOk returns a tuple with the NavHoverColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavHoverColor

`func (o *WhitelabelSettings) SetNavHoverColor(v string)`

SetNavHoverColor sets NavHoverColor field to given value.

### HasNavHoverColor

`func (o *WhitelabelSettings) HasNavHoverColor() bool`

HasNavHoverColor returns a boolean if a field has been set.

### GetPrimaryButtonBgColor

`func (o *WhitelabelSettings) GetPrimaryButtonBgColor() string`

GetPrimaryButtonBgColor returns the PrimaryButtonBgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonBgColorOk

`func (o *WhitelabelSettings) GetPrimaryButtonBgColorOk() (*string, bool)`

GetPrimaryButtonBgColorOk returns a tuple with the PrimaryButtonBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonBgColor

`func (o *WhitelabelSettings) SetPrimaryButtonBgColor(v string)`

SetPrimaryButtonBgColor sets PrimaryButtonBgColor field to given value.

### HasPrimaryButtonBgColor

`func (o *WhitelabelSettings) HasPrimaryButtonBgColor() bool`

HasPrimaryButtonBgColor returns a boolean if a field has been set.

### GetPrimaryButtonFgColor

`func (o *WhitelabelSettings) GetPrimaryButtonFgColor() string`

GetPrimaryButtonFgColor returns the PrimaryButtonFgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonFgColorOk

`func (o *WhitelabelSettings) GetPrimaryButtonFgColorOk() (*string, bool)`

GetPrimaryButtonFgColorOk returns a tuple with the PrimaryButtonFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonFgColor

`func (o *WhitelabelSettings) SetPrimaryButtonFgColor(v string)`

SetPrimaryButtonFgColor sets PrimaryButtonFgColor field to given value.

### HasPrimaryButtonFgColor

`func (o *WhitelabelSettings) HasPrimaryButtonFgColor() bool`

HasPrimaryButtonFgColor returns a boolean if a field has been set.

### GetPrimaryButtonHoverBgColor

`func (o *WhitelabelSettings) GetPrimaryButtonHoverBgColor() string`

GetPrimaryButtonHoverBgColor returns the PrimaryButtonHoverBgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonHoverBgColorOk

`func (o *WhitelabelSettings) GetPrimaryButtonHoverBgColorOk() (*string, bool)`

GetPrimaryButtonHoverBgColorOk returns a tuple with the PrimaryButtonHoverBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonHoverBgColor

`func (o *WhitelabelSettings) SetPrimaryButtonHoverBgColor(v string)`

SetPrimaryButtonHoverBgColor sets PrimaryButtonHoverBgColor field to given value.

### HasPrimaryButtonHoverBgColor

`func (o *WhitelabelSettings) HasPrimaryButtonHoverBgColor() bool`

HasPrimaryButtonHoverBgColor returns a boolean if a field has been set.

### GetPrimaryButtonHoverFgColor

`func (o *WhitelabelSettings) GetPrimaryButtonHoverFgColor() string`

GetPrimaryButtonHoverFgColor returns the PrimaryButtonHoverFgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonHoverFgColorOk

`func (o *WhitelabelSettings) GetPrimaryButtonHoverFgColorOk() (*string, bool)`

GetPrimaryButtonHoverFgColorOk returns a tuple with the PrimaryButtonHoverFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonHoverFgColor

`func (o *WhitelabelSettings) SetPrimaryButtonHoverFgColor(v string)`

SetPrimaryButtonHoverFgColor sets PrimaryButtonHoverFgColor field to given value.

### HasPrimaryButtonHoverFgColor

`func (o *WhitelabelSettings) HasPrimaryButtonHoverFgColor() bool`

HasPrimaryButtonHoverFgColor returns a boolean if a field has been set.

### GetFooterBgColor

`func (o *WhitelabelSettings) GetFooterBgColor() string`

GetFooterBgColor returns the FooterBgColor field if non-nil, zero value otherwise.

### GetFooterBgColorOk

`func (o *WhitelabelSettings) GetFooterBgColorOk() (*string, bool)`

GetFooterBgColorOk returns a tuple with the FooterBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterBgColor

`func (o *WhitelabelSettings) SetFooterBgColor(v string)`

SetFooterBgColor sets FooterBgColor field to given value.

### HasFooterBgColor

`func (o *WhitelabelSettings) HasFooterBgColor() bool`

HasFooterBgColor returns a boolean if a field has been set.

### GetFooterFgColor

`func (o *WhitelabelSettings) GetFooterFgColor() string`

GetFooterFgColor returns the FooterFgColor field if non-nil, zero value otherwise.

### GetFooterFgColorOk

`func (o *WhitelabelSettings) GetFooterFgColorOk() (*string, bool)`

GetFooterFgColorOk returns a tuple with the FooterFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterFgColor

`func (o *WhitelabelSettings) SetFooterFgColor(v string)`

SetFooterFgColor sets FooterFgColor field to given value.

### HasFooterFgColor

`func (o *WhitelabelSettings) HasFooterFgColor() bool`

HasFooterFgColor returns a boolean if a field has been set.

### GetLoginBgColor

`func (o *WhitelabelSettings) GetLoginBgColor() string`

GetLoginBgColor returns the LoginBgColor field if non-nil, zero value otherwise.

### GetLoginBgColorOk

`func (o *WhitelabelSettings) GetLoginBgColorOk() (*string, bool)`

GetLoginBgColorOk returns a tuple with the LoginBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBgColor

`func (o *WhitelabelSettings) SetLoginBgColor(v string)`

SetLoginBgColor sets LoginBgColor field to given value.

### HasLoginBgColor

`func (o *WhitelabelSettings) HasLoginBgColor() bool`

HasLoginBgColor returns a boolean if a field has been set.

### GetOverrideCss

`func (o *WhitelabelSettings) GetOverrideCss() string`

GetOverrideCss returns the OverrideCss field if non-nil, zero value otherwise.

### GetOverrideCssOk

`func (o *WhitelabelSettings) GetOverrideCssOk() (*string, bool)`

GetOverrideCssOk returns a tuple with the OverrideCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCss

`func (o *WhitelabelSettings) SetOverrideCss(v string)`

SetOverrideCss sets OverrideCss field to given value.

### HasOverrideCss

`func (o *WhitelabelSettings) HasOverrideCss() bool`

HasOverrideCss returns a boolean if a field has been set.

### GetCopyrightString

`func (o *WhitelabelSettings) GetCopyrightString() string`

GetCopyrightString returns the CopyrightString field if non-nil, zero value otherwise.

### GetCopyrightStringOk

`func (o *WhitelabelSettings) GetCopyrightStringOk() (*string, bool)`

GetCopyrightStringOk returns a tuple with the CopyrightString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightString

`func (o *WhitelabelSettings) SetCopyrightString(v string)`

SetCopyrightString sets CopyrightString field to given value.

### HasCopyrightString

`func (o *WhitelabelSettings) HasCopyrightString() bool`

HasCopyrightString returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *WhitelabelSettings) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *WhitelabelSettings) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *WhitelabelSettings) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *WhitelabelSettings) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *WhitelabelSettings) GetPrivacyPolicy() string`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *WhitelabelSettings) GetPrivacyPolicyOk() (*string, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *WhitelabelSettings) SetPrivacyPolicy(v string)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *WhitelabelSettings) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetSupportMenuLinks

`func (o *WhitelabelSettings) GetSupportMenuLinks() []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner`

GetSupportMenuLinks returns the SupportMenuLinks field if non-nil, zero value otherwise.

### GetSupportMenuLinksOk

`func (o *WhitelabelSettings) GetSupportMenuLinksOk() (*[]ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner, bool)`

GetSupportMenuLinksOk returns a tuple with the SupportMenuLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMenuLinks

`func (o *WhitelabelSettings) SetSupportMenuLinks(v []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner)`

SetSupportMenuLinks sets SupportMenuLinks field to given value.

### HasSupportMenuLinks

`func (o *WhitelabelSettings) HasSupportMenuLinks() bool`

HasSupportMenuLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


