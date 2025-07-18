/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the EducationUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EducationUser{}

// EducationUser An extension of user with education-specific attributes
type EducationUser struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Set to \"true\" when the account is enabled.
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $orderby.
	DisplayName *string `json:"displayName,omitempty"`
	// A collection of drives available for this user. Read-only.
	Drives []Drive `json:"drives,omitempty"`
	Drive *Drive `json:"drive,omitempty"`
	// Identities associated with this account.
	Identities []ObjectIdentity `json:"identities,omitempty"`
	// The SMTP address for the user, for example, 'jeff@contoso.opencloud.com'. Returned by default.
	Mail *string `json:"mail,omitempty"`
	// Groups that this user is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
	MemberOf []Group `json:"memberOf,omitempty"`
	// Contains the on-premises SAM account name synchronized from the on-premises directory. Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// The user's surname (family name or last name). Returned by default.
	Surname *string `json:"surname,omitempty"`
	// The user's givenName. Returned by default.
	GivenName *string `json:"givenName,omitempty"`
	// The user`s default role. Such as \"student\" or \"teacher\"
	PrimaryRole *string `json:"primaryRole,omitempty"`
	// The user`s type. This can be either \"Member\" for regular user, \"Guest\" for guest users or \"Federated\" for users imported from a federated instance.
	UserType *string `json:"userType,omitempty"`
}

// NewEducationUser instantiates a new EducationUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationUser() *EducationUser {
	this := EducationUser{}
	return &this
}

// NewEducationUserWithDefaults instantiates a new EducationUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationUserWithDefaults() *EducationUser {
	this := EducationUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EducationUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EducationUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EducationUser) SetId(v string) {
	o.Id = &v
}

// GetAccountEnabled returns the AccountEnabled field value if set, zero value otherwise.
func (o *EducationUser) GetAccountEnabled() bool {
	if o == nil || IsNil(o.AccountEnabled) {
		var ret bool
		return ret
	}
	return *o.AccountEnabled
}

// GetAccountEnabledOk returns a tuple with the AccountEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetAccountEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountEnabled) {
		return nil, false
	}
	return o.AccountEnabled, true
}

// HasAccountEnabled returns a boolean if a field has been set.
func (o *EducationUser) HasAccountEnabled() bool {
	if o != nil && !IsNil(o.AccountEnabled) {
		return true
	}

	return false
}

// SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.
func (o *EducationUser) SetAccountEnabled(v bool) {
	o.AccountEnabled = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EducationUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EducationUser) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EducationUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *EducationUser) GetDrives() []Drive {
	if o == nil || IsNil(o.Drives) {
		var ret []Drive
		return ret
	}
	return o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetDrivesOk() ([]Drive, bool) {
	if o == nil || IsNil(o.Drives) {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *EducationUser) HasDrives() bool {
	if o != nil && !IsNil(o.Drives) {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []Drive and assigns it to the Drives field.
func (o *EducationUser) SetDrives(v []Drive) {
	o.Drives = v
}

// GetDrive returns the Drive field value if set, zero value otherwise.
func (o *EducationUser) GetDrive() Drive {
	if o == nil || IsNil(o.Drive) {
		var ret Drive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetDriveOk() (*Drive, bool) {
	if o == nil || IsNil(o.Drive) {
		return nil, false
	}
	return o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *EducationUser) HasDrive() bool {
	if o != nil && !IsNil(o.Drive) {
		return true
	}

	return false
}

// SetDrive gets a reference to the given Drive and assigns it to the Drive field.
func (o *EducationUser) SetDrive(v Drive) {
	o.Drive = &v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *EducationUser) GetIdentities() []ObjectIdentity {
	if o == nil || IsNil(o.Identities) {
		var ret []ObjectIdentity
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetIdentitiesOk() ([]ObjectIdentity, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *EducationUser) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []ObjectIdentity and assigns it to the Identities field.
func (o *EducationUser) SetIdentities(v []ObjectIdentity) {
	o.Identities = v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *EducationUser) GetMail() string {
	if o == nil || IsNil(o.Mail) {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetMailOk() (*string, bool) {
	if o == nil || IsNil(o.Mail) {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *EducationUser) HasMail() bool {
	if o != nil && !IsNil(o.Mail) {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *EducationUser) SetMail(v string) {
	o.Mail = &v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise.
func (o *EducationUser) GetMemberOf() []Group {
	if o == nil || IsNil(o.MemberOf) {
		var ret []Group
		return ret
	}
	return o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetMemberOfOk() ([]Group, bool) {
	if o == nil || IsNil(o.MemberOf) {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *EducationUser) HasMemberOf() bool {
	if o != nil && !IsNil(o.MemberOf) {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given []Group and assigns it to the MemberOf field.
func (o *EducationUser) SetMemberOf(v []Group) {
	o.MemberOf = v
}

// GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field value if set, zero value otherwise.
func (o *EducationUser) GetOnPremisesSamAccountName() string {
	if o == nil || IsNil(o.OnPremisesSamAccountName) {
		var ret string
		return ret
	}
	return *o.OnPremisesSamAccountName
}

// GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetOnPremisesSamAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.OnPremisesSamAccountName) {
		return nil, false
	}
	return o.OnPremisesSamAccountName, true
}

// HasOnPremisesSamAccountName returns a boolean if a field has been set.
func (o *EducationUser) HasOnPremisesSamAccountName() bool {
	if o != nil && !IsNil(o.OnPremisesSamAccountName) {
		return true
	}

	return false
}

// SetOnPremisesSamAccountName gets a reference to the given string and assigns it to the OnPremisesSamAccountName field.
func (o *EducationUser) SetOnPremisesSamAccountName(v string) {
	o.OnPremisesSamAccountName = &v
}

// GetPasswordProfile returns the PasswordProfile field value if set, zero value otherwise.
func (o *EducationUser) GetPasswordProfile() PasswordProfile {
	if o == nil || IsNil(o.PasswordProfile) {
		var ret PasswordProfile
		return ret
	}
	return *o.PasswordProfile
}

// GetPasswordProfileOk returns a tuple with the PasswordProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetPasswordProfileOk() (*PasswordProfile, bool) {
	if o == nil || IsNil(o.PasswordProfile) {
		return nil, false
	}
	return o.PasswordProfile, true
}

// HasPasswordProfile returns a boolean if a field has been set.
func (o *EducationUser) HasPasswordProfile() bool {
	if o != nil && !IsNil(o.PasswordProfile) {
		return true
	}

	return false
}

// SetPasswordProfile gets a reference to the given PasswordProfile and assigns it to the PasswordProfile field.
func (o *EducationUser) SetPasswordProfile(v PasswordProfile) {
	o.PasswordProfile = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *EducationUser) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *EducationUser) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *EducationUser) SetSurname(v string) {
	o.Surname = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *EducationUser) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *EducationUser) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *EducationUser) SetGivenName(v string) {
	o.GivenName = &v
}

// GetPrimaryRole returns the PrimaryRole field value if set, zero value otherwise.
func (o *EducationUser) GetPrimaryRole() string {
	if o == nil || IsNil(o.PrimaryRole) {
		var ret string
		return ret
	}
	return *o.PrimaryRole
}

// GetPrimaryRoleOk returns a tuple with the PrimaryRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetPrimaryRoleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryRole) {
		return nil, false
	}
	return o.PrimaryRole, true
}

// HasPrimaryRole returns a boolean if a field has been set.
func (o *EducationUser) HasPrimaryRole() bool {
	if o != nil && !IsNil(o.PrimaryRole) {
		return true
	}

	return false
}

// SetPrimaryRole gets a reference to the given string and assigns it to the PrimaryRole field.
func (o *EducationUser) SetPrimaryRole(v string) {
	o.PrimaryRole = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *EducationUser) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationUser) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *EducationUser) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *EducationUser) SetUserType(v string) {
	o.UserType = &v
}

func (o EducationUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EducationUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountEnabled) {
		toSerialize["accountEnabled"] = o.AccountEnabled
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Drives) {
		toSerialize["drives"] = o.Drives
	}
	if !IsNil(o.Drive) {
		toSerialize["drive"] = o.Drive
	}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}
	if !IsNil(o.Mail) {
		toSerialize["mail"] = o.Mail
	}
	if !IsNil(o.MemberOf) {
		toSerialize["memberOf"] = o.MemberOf
	}
	if !IsNil(o.OnPremisesSamAccountName) {
		toSerialize["onPremisesSamAccountName"] = o.OnPremisesSamAccountName
	}
	if !IsNil(o.PasswordProfile) {
		toSerialize["passwordProfile"] = o.PasswordProfile
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.PrimaryRole) {
		toSerialize["primaryRole"] = o.PrimaryRole
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	return toSerialize, nil
}

type NullableEducationUser struct {
	value *EducationUser
	isSet bool
}

func (v NullableEducationUser) Get() *EducationUser {
	return v.value
}

func (v *NullableEducationUser) Set(val *EducationUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationUser(val *EducationUser) *NullableEducationUser {
	return &NullableEducationUser{value: val, isSet: true}
}

func (v NullableEducationUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


