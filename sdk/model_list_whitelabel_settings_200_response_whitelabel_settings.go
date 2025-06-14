/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the ListWhitelabelSettings200ResponseWhitelabelSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWhitelabelSettings200ResponseWhitelabelSettings{}

// ListWhitelabelSettings200ResponseWhitelabelSettings struct for ListWhitelabelSettings200ResponseWhitelabelSettings
type ListWhitelabelSettings200ResponseWhitelabelSettings struct {
	Enabled                   *bool                                                                      `json:"enabled,omitempty"`
	ApplianceName             *string                                                                    `json:"applianceName,omitempty"`
	DisableSupportMenu        *bool                                                                      `json:"disableSupportMenu,omitempty"`
	HeaderLogo                *string                                                                    `json:"headerLogo,omitempty"`
	FooterLogo                *string                                                                    `json:"footerLogo,omitempty"`
	LoginLogo                 *string                                                                    `json:"loginLogo,omitempty"`
	Favicon                   *string                                                                    `json:"favicon,omitempty"`
	HeaderBgColor             *string                                                                    `json:"headerBgColor,omitempty"`
	HeaderFgColor             *string                                                                    `json:"headerFgColor,omitempty"`
	NavBgColor                *string                                                                    `json:"navBgColor,omitempty"`
	NavFgColor                *string                                                                    `json:"navFgColor,omitempty"`
	NavHoverColor             *string                                                                    `json:"navHoverColor,omitempty"`
	PrimaryButtonBgColor      *string                                                                    `json:"primaryButtonBgColor,omitempty"`
	PrimaryButtonFgColor      *string                                                                    `json:"primaryButtonFgColor,omitempty"`
	PrimaryButtonHoverBgColor *string                                                                    `json:"primaryButtonHoverBgColor,omitempty"`
	PrimaryButtonHoverFgColor *string                                                                    `json:"primaryButtonHoverFgColor,omitempty"`
	FooterBgColor             *string                                                                    `json:"footerBgColor,omitempty"`
	FooterFgColor             *string                                                                    `json:"footerFgColor,omitempty"`
	LoginBgColor              *string                                                                    `json:"loginBgColor,omitempty"`
	OverrideCss               *string                                                                    `json:"overrideCss,omitempty"`
	CopyrightString           *string                                                                    `json:"copyrightString,omitempty"`
	TermsOfUse                *string                                                                    `json:"termsOfUse,omitempty"`
	PrivacyPolicy             *string                                                                    `json:"privacyPolicy,omitempty"`
	SupportMenuLinks          []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner `json:"supportMenuLinks,omitempty"`
	AdditionalProperties      map[string]interface{}                                                     `json:",remain"`
}

type _ListWhitelabelSettings200ResponseWhitelabelSettings ListWhitelabelSettings200ResponseWhitelabelSettings

// NewListWhitelabelSettings200ResponseWhitelabelSettings instantiates a new ListWhitelabelSettings200ResponseWhitelabelSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWhitelabelSettings200ResponseWhitelabelSettings() *ListWhitelabelSettings200ResponseWhitelabelSettings {
	this := ListWhitelabelSettings200ResponseWhitelabelSettings{}
	return &this
}

// NewListWhitelabelSettings200ResponseWhitelabelSettingsWithDefaults instantiates a new ListWhitelabelSettings200ResponseWhitelabelSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWhitelabelSettings200ResponseWhitelabelSettingsWithDefaults() *ListWhitelabelSettings200ResponseWhitelabelSettings {
	this := ListWhitelabelSettings200ResponseWhitelabelSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApplianceName returns the ApplianceName field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetApplianceName() string {
	if o == nil || IsNil(o.ApplianceName) {
		var ret string
		return ret
	}
	return *o.ApplianceName
}

// GetApplianceNameOk returns a tuple with the ApplianceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetApplianceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplianceName) {
		return nil, false
	}
	return o.ApplianceName, true
}

// IsSetApplianceName returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetApplianceName() bool {
	if o != nil && !IsNil(o.ApplianceName) {
		return true
	}

	return false
}

// SetApplianceName gets a reference to the given string and assigns it to the ApplianceName field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetApplianceName(v string) {
	o.ApplianceName = &v
}

// GetDisableSupportMenu returns the DisableSupportMenu field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetDisableSupportMenu() bool {
	if o == nil || IsNil(o.DisableSupportMenu) {
		var ret bool
		return ret
	}
	return *o.DisableSupportMenu
}

// GetDisableSupportMenuOk returns a tuple with the DisableSupportMenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetDisableSupportMenuOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSupportMenu) {
		return nil, false
	}
	return o.DisableSupportMenu, true
}

// IsSetDisableSupportMenu returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetDisableSupportMenu() bool {
	if o != nil && !IsNil(o.DisableSupportMenu) {
		return true
	}

	return false
}

// SetDisableSupportMenu gets a reference to the given bool and assigns it to the DisableSupportMenu field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetDisableSupportMenu(v bool) {
	o.DisableSupportMenu = &v
}

// GetHeaderLogo returns the HeaderLogo field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderLogo() string {
	if o == nil || IsNil(o.HeaderLogo) {
		var ret string
		return ret
	}
	return *o.HeaderLogo
}

// GetHeaderLogoOk returns a tuple with the HeaderLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderLogoOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderLogo) {
		return nil, false
	}
	return o.HeaderLogo, true
}

// IsSetHeaderLogo returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetHeaderLogo() bool {
	if o != nil && !IsNil(o.HeaderLogo) {
		return true
	}

	return false
}

// SetHeaderLogo gets a reference to the given string and assigns it to the HeaderLogo field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetHeaderLogo(v string) {
	o.HeaderLogo = &v
}

// GetFooterLogo returns the FooterLogo field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterLogo() string {
	if o == nil || IsNil(o.FooterLogo) {
		var ret string
		return ret
	}
	return *o.FooterLogo
}

// GetFooterLogoOk returns a tuple with the FooterLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterLogoOk() (*string, bool) {
	if o == nil || IsNil(o.FooterLogo) {
		return nil, false
	}
	return o.FooterLogo, true
}

// IsSetFooterLogo returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetFooterLogo() bool {
	if o != nil && !IsNil(o.FooterLogo) {
		return true
	}

	return false
}

// SetFooterLogo gets a reference to the given string and assigns it to the FooterLogo field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetFooterLogo(v string) {
	o.FooterLogo = &v
}

// GetLoginLogo returns the LoginLogo field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetLoginLogo() string {
	if o == nil || IsNil(o.LoginLogo) {
		var ret string
		return ret
	}
	return *o.LoginLogo
}

// GetLoginLogoOk returns a tuple with the LoginLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetLoginLogoOk() (*string, bool) {
	if o == nil || IsNil(o.LoginLogo) {
		return nil, false
	}
	return o.LoginLogo, true
}

// IsSetLoginLogo returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetLoginLogo() bool {
	if o != nil && !IsNil(o.LoginLogo) {
		return true
	}

	return false
}

// SetLoginLogo gets a reference to the given string and assigns it to the LoginLogo field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetLoginLogo(v string) {
	o.LoginLogo = &v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFavicon() string {
	if o == nil || IsNil(o.Favicon) {
		var ret string
		return ret
	}
	return *o.Favicon
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFaviconOk() (*string, bool) {
	if o == nil || IsNil(o.Favicon) {
		return nil, false
	}
	return o.Favicon, true
}

// IsSetFavicon returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetFavicon() bool {
	if o != nil && !IsNil(o.Favicon) {
		return true
	}

	return false
}

// SetFavicon gets a reference to the given string and assigns it to the Favicon field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetFavicon(v string) {
	o.Favicon = &v
}

// GetHeaderBgColor returns the HeaderBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderBgColor() string {
	if o == nil || IsNil(o.HeaderBgColor) {
		var ret string
		return ret
	}
	return *o.HeaderBgColor
}

// GetHeaderBgColorOk returns a tuple with the HeaderBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderBgColor) {
		return nil, false
	}
	return o.HeaderBgColor, true
}

// IsSetHeaderBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetHeaderBgColor() bool {
	if o != nil && !IsNil(o.HeaderBgColor) {
		return true
	}

	return false
}

// SetHeaderBgColor gets a reference to the given string and assigns it to the HeaderBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetHeaderBgColor(v string) {
	o.HeaderBgColor = &v
}

// GetHeaderFgColor returns the HeaderFgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderFgColor() string {
	if o == nil || IsNil(o.HeaderFgColor) {
		var ret string
		return ret
	}
	return *o.HeaderFgColor
}

// GetHeaderFgColorOk returns a tuple with the HeaderFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetHeaderFgColorOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderFgColor) {
		return nil, false
	}
	return o.HeaderFgColor, true
}

// IsSetHeaderFgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetHeaderFgColor() bool {
	if o != nil && !IsNil(o.HeaderFgColor) {
		return true
	}

	return false
}

// SetHeaderFgColor gets a reference to the given string and assigns it to the HeaderFgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetHeaderFgColor(v string) {
	o.HeaderFgColor = &v
}

// GetNavBgColor returns the NavBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavBgColor() string {
	if o == nil || IsNil(o.NavBgColor) {
		var ret string
		return ret
	}
	return *o.NavBgColor
}

// GetNavBgColorOk returns a tuple with the NavBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.NavBgColor) {
		return nil, false
	}
	return o.NavBgColor, true
}

// IsSetNavBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetNavBgColor() bool {
	if o != nil && !IsNil(o.NavBgColor) {
		return true
	}

	return false
}

// SetNavBgColor gets a reference to the given string and assigns it to the NavBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetNavBgColor(v string) {
	o.NavBgColor = &v
}

// GetNavFgColor returns the NavFgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavFgColor() string {
	if o == nil || IsNil(o.NavFgColor) {
		var ret string
		return ret
	}
	return *o.NavFgColor
}

// GetNavFgColorOk returns a tuple with the NavFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavFgColorOk() (*string, bool) {
	if o == nil || IsNil(o.NavFgColor) {
		return nil, false
	}
	return o.NavFgColor, true
}

// IsSetNavFgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetNavFgColor() bool {
	if o != nil && !IsNil(o.NavFgColor) {
		return true
	}

	return false
}

// SetNavFgColor gets a reference to the given string and assigns it to the NavFgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetNavFgColor(v string) {
	o.NavFgColor = &v
}

// GetNavHoverColor returns the NavHoverColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavHoverColor() string {
	if o == nil || IsNil(o.NavHoverColor) {
		var ret string
		return ret
	}
	return *o.NavHoverColor
}

// GetNavHoverColorOk returns a tuple with the NavHoverColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetNavHoverColorOk() (*string, bool) {
	if o == nil || IsNil(o.NavHoverColor) {
		return nil, false
	}
	return o.NavHoverColor, true
}

// IsSetNavHoverColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetNavHoverColor() bool {
	if o != nil && !IsNil(o.NavHoverColor) {
		return true
	}

	return false
}

// SetNavHoverColor gets a reference to the given string and assigns it to the NavHoverColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetNavHoverColor(v string) {
	o.NavHoverColor = &v
}

// GetPrimaryButtonBgColor returns the PrimaryButtonBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonBgColor() string {
	if o == nil || IsNil(o.PrimaryButtonBgColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonBgColor
}

// GetPrimaryButtonBgColorOk returns a tuple with the PrimaryButtonBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonBgColor) {
		return nil, false
	}
	return o.PrimaryButtonBgColor, true
}

// IsSetPrimaryButtonBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetPrimaryButtonBgColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonBgColor) {
		return true
	}

	return false
}

// SetPrimaryButtonBgColor gets a reference to the given string and assigns it to the PrimaryButtonBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetPrimaryButtonBgColor(v string) {
	o.PrimaryButtonBgColor = &v
}

// GetPrimaryButtonFgColor returns the PrimaryButtonFgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonFgColor() string {
	if o == nil || IsNil(o.PrimaryButtonFgColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonFgColor
}

// GetPrimaryButtonFgColorOk returns a tuple with the PrimaryButtonFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonFgColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonFgColor) {
		return nil, false
	}
	return o.PrimaryButtonFgColor, true
}

// IsSetPrimaryButtonFgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetPrimaryButtonFgColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonFgColor) {
		return true
	}

	return false
}

// SetPrimaryButtonFgColor gets a reference to the given string and assigns it to the PrimaryButtonFgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetPrimaryButtonFgColor(v string) {
	o.PrimaryButtonFgColor = &v
}

// GetPrimaryButtonHoverBgColor returns the PrimaryButtonHoverBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonHoverBgColor() string {
	if o == nil || IsNil(o.PrimaryButtonHoverBgColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonHoverBgColor
}

// GetPrimaryButtonHoverBgColorOk returns a tuple with the PrimaryButtonHoverBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonHoverBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonHoverBgColor) {
		return nil, false
	}
	return o.PrimaryButtonHoverBgColor, true
}

// IsSetPrimaryButtonHoverBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetPrimaryButtonHoverBgColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonHoverBgColor) {
		return true
	}

	return false
}

// SetPrimaryButtonHoverBgColor gets a reference to the given string and assigns it to the PrimaryButtonHoverBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetPrimaryButtonHoverBgColor(v string) {
	o.PrimaryButtonHoverBgColor = &v
}

// GetPrimaryButtonHoverFgColor returns the PrimaryButtonHoverFgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonHoverFgColor() string {
	if o == nil || IsNil(o.PrimaryButtonHoverFgColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonHoverFgColor
}

// GetPrimaryButtonHoverFgColorOk returns a tuple with the PrimaryButtonHoverFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrimaryButtonHoverFgColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonHoverFgColor) {
		return nil, false
	}
	return o.PrimaryButtonHoverFgColor, true
}

// IsSetPrimaryButtonHoverFgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetPrimaryButtonHoverFgColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonHoverFgColor) {
		return true
	}

	return false
}

// SetPrimaryButtonHoverFgColor gets a reference to the given string and assigns it to the PrimaryButtonHoverFgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetPrimaryButtonHoverFgColor(v string) {
	o.PrimaryButtonHoverFgColor = &v
}

// GetFooterBgColor returns the FooterBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterBgColor() string {
	if o == nil || IsNil(o.FooterBgColor) {
		var ret string
		return ret
	}
	return *o.FooterBgColor
}

// GetFooterBgColorOk returns a tuple with the FooterBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.FooterBgColor) {
		return nil, false
	}
	return o.FooterBgColor, true
}

// IsSetFooterBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetFooterBgColor() bool {
	if o != nil && !IsNil(o.FooterBgColor) {
		return true
	}

	return false
}

// SetFooterBgColor gets a reference to the given string and assigns it to the FooterBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetFooterBgColor(v string) {
	o.FooterBgColor = &v
}

// GetFooterFgColor returns the FooterFgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterFgColor() string {
	if o == nil || IsNil(o.FooterFgColor) {
		var ret string
		return ret
	}
	return *o.FooterFgColor
}

// GetFooterFgColorOk returns a tuple with the FooterFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetFooterFgColorOk() (*string, bool) {
	if o == nil || IsNil(o.FooterFgColor) {
		return nil, false
	}
	return o.FooterFgColor, true
}

// IsSetFooterFgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetFooterFgColor() bool {
	if o != nil && !IsNil(o.FooterFgColor) {
		return true
	}

	return false
}

// SetFooterFgColor gets a reference to the given string and assigns it to the FooterFgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetFooterFgColor(v string) {
	o.FooterFgColor = &v
}

// GetLoginBgColor returns the LoginBgColor field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetLoginBgColor() string {
	if o == nil || IsNil(o.LoginBgColor) {
		var ret string
		return ret
	}
	return *o.LoginBgColor
}

// GetLoginBgColorOk returns a tuple with the LoginBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetLoginBgColorOk() (*string, bool) {
	if o == nil || IsNil(o.LoginBgColor) {
		return nil, false
	}
	return o.LoginBgColor, true
}

// IsSetLoginBgColor returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetLoginBgColor() bool {
	if o != nil && !IsNil(o.LoginBgColor) {
		return true
	}

	return false
}

// SetLoginBgColor gets a reference to the given string and assigns it to the LoginBgColor field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetLoginBgColor(v string) {
	o.LoginBgColor = &v
}

// GetOverrideCss returns the OverrideCss field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetOverrideCss() string {
	if o == nil || IsNil(o.OverrideCss) {
		var ret string
		return ret
	}
	return *o.OverrideCss
}

// GetOverrideCssOk returns a tuple with the OverrideCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetOverrideCssOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideCss) {
		return nil, false
	}
	return o.OverrideCss, true
}

// IsSetOverrideCss returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetOverrideCss() bool {
	if o != nil && !IsNil(o.OverrideCss) {
		return true
	}

	return false
}

// SetOverrideCss gets a reference to the given string and assigns it to the OverrideCss field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetOverrideCss(v string) {
	o.OverrideCss = &v
}

// GetCopyrightString returns the CopyrightString field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetCopyrightString() string {
	if o == nil || IsNil(o.CopyrightString) {
		var ret string
		return ret
	}
	return *o.CopyrightString
}

// GetCopyrightStringOk returns a tuple with the CopyrightString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetCopyrightStringOk() (*string, bool) {
	if o == nil || IsNil(o.CopyrightString) {
		return nil, false
	}
	return o.CopyrightString, true
}

// IsSetCopyrightString returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetCopyrightString() bool {
	if o != nil && !IsNil(o.CopyrightString) {
		return true
	}

	return false
}

// SetCopyrightString gets a reference to the given string and assigns it to the CopyrightString field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetCopyrightString(v string) {
	o.CopyrightString = &v
}

// GetTermsOfUse returns the TermsOfUse field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetTermsOfUse() string {
	if o == nil || IsNil(o.TermsOfUse) {
		var ret string
		return ret
	}
	return *o.TermsOfUse
}

// GetTermsOfUseOk returns a tuple with the TermsOfUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetTermsOfUseOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfUse) {
		return nil, false
	}
	return o.TermsOfUse, true
}

// IsSetTermsOfUse returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetTermsOfUse() bool {
	if o != nil && !IsNil(o.TermsOfUse) {
		return true
	}

	return false
}

// SetTermsOfUse gets a reference to the given string and assigns it to the TermsOfUse field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetTermsOfUse(v string) {
	o.TermsOfUse = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrivacyPolicy() string {
	if o == nil || IsNil(o.PrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicy) {
		return nil, false
	}
	return o.PrivacyPolicy, true
}

// IsSetPrivacyPolicy returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetPrivacyPolicy() bool {
	if o != nil && !IsNil(o.PrivacyPolicy) {
		return true
	}

	return false
}

// SetPrivacyPolicy gets a reference to the given string and assigns it to the PrivacyPolicy field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetPrivacyPolicy(v string) {
	o.PrivacyPolicy = &v
}

// GetSupportMenuLinks returns the SupportMenuLinks field value if set, zero value otherwise.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetSupportMenuLinks() []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner {
	if o == nil || IsNil(o.SupportMenuLinks) {
		var ret []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner
		return ret
	}
	return o.SupportMenuLinks
}

// GetSupportMenuLinksOk returns a tuple with the SupportMenuLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) GetSupportMenuLinksOk() ([]ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner, bool) {
	if o == nil || IsNil(o.SupportMenuLinks) {
		return nil, false
	}
	return o.SupportMenuLinks, true
}

// IsSetSupportMenuLinks returns a boolean if a field has been set.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) IsSetSupportMenuLinks() bool {
	if o != nil && !IsNil(o.SupportMenuLinks) {
		return true
	}

	return false
}

// SetSupportMenuLinks gets a reference to the given []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner and assigns it to the SupportMenuLinks field.
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) SetSupportMenuLinks(v []ListWhitelabelSettings200ResponseWhitelabelSettingsSupportMenuLinksInner) {
	o.SupportMenuLinks = v
}

func (o ListWhitelabelSettings200ResponseWhitelabelSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWhitelabelSettings200ResponseWhitelabelSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ApplianceName) {
		toSerialize["applianceName"] = o.ApplianceName
	}
	if !IsNil(o.DisableSupportMenu) {
		toSerialize["disableSupportMenu"] = o.DisableSupportMenu
	}
	if !IsNil(o.HeaderLogo) {
		toSerialize["headerLogo"] = o.HeaderLogo
	}
	if !IsNil(o.FooterLogo) {
		toSerialize["footerLogo"] = o.FooterLogo
	}
	if !IsNil(o.LoginLogo) {
		toSerialize["loginLogo"] = o.LoginLogo
	}
	if !IsNil(o.Favicon) {
		toSerialize["favicon"] = o.Favicon
	}
	if !IsNil(o.HeaderBgColor) {
		toSerialize["headerBgColor"] = o.HeaderBgColor
	}
	if !IsNil(o.HeaderFgColor) {
		toSerialize["headerFgColor"] = o.HeaderFgColor
	}
	if !IsNil(o.NavBgColor) {
		toSerialize["navBgColor"] = o.NavBgColor
	}
	if !IsNil(o.NavFgColor) {
		toSerialize["navFgColor"] = o.NavFgColor
	}
	if !IsNil(o.NavHoverColor) {
		toSerialize["navHoverColor"] = o.NavHoverColor
	}
	if !IsNil(o.PrimaryButtonBgColor) {
		toSerialize["primaryButtonBgColor"] = o.PrimaryButtonBgColor
	}
	if !IsNil(o.PrimaryButtonFgColor) {
		toSerialize["primaryButtonFgColor"] = o.PrimaryButtonFgColor
	}
	if !IsNil(o.PrimaryButtonHoverBgColor) {
		toSerialize["primaryButtonHoverBgColor"] = o.PrimaryButtonHoverBgColor
	}
	if !IsNil(o.PrimaryButtonHoverFgColor) {
		toSerialize["primaryButtonHoverFgColor"] = o.PrimaryButtonHoverFgColor
	}
	if !IsNil(o.FooterBgColor) {
		toSerialize["footerBgColor"] = o.FooterBgColor
	}
	if !IsNil(o.FooterFgColor) {
		toSerialize["footerFgColor"] = o.FooterFgColor
	}
	if !IsNil(o.LoginBgColor) {
		toSerialize["loginBgColor"] = o.LoginBgColor
	}
	if !IsNil(o.OverrideCss) {
		toSerialize["overrideCss"] = o.OverrideCss
	}
	if !IsNil(o.CopyrightString) {
		toSerialize["copyrightString"] = o.CopyrightString
	}
	if !IsNil(o.TermsOfUse) {
		toSerialize["termsOfUse"] = o.TermsOfUse
	}
	if !IsNil(o.PrivacyPolicy) {
		toSerialize["privacyPolicy"] = o.PrivacyPolicy
	}
	if !IsNil(o.SupportMenuLinks) {
		toSerialize["supportMenuLinks"] = o.SupportMenuLinks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListWhitelabelSettings200ResponseWhitelabelSettings) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
