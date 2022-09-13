package model

import "strings"

const (
	Name    = "name"
	KeyId   = "keyid"
	Country = "country"
	Email   = "email"
)

// FFCUser A collection of attributes that can affect flag evaluation, usually corresponding to a user of your application.
// The only mandatory property is the key, which must uniquely identify each user;
// this could be a username or email address for authenticated users, or a ID for anonymous users.
// All other built-in properties are optional,
// it's strongly recommended to set userName in order to search your user quickly
// You may also define custom properties with arbitrary names and values.
type FFCUser struct {
	UserName string
	Email    string
	Key      string
	Country  string
	Custom   map[string]string
}

// FFUserBuilder  A builder  that helps construct FFCClient objects. Builder calls can be chained, supporting the following pattern:
//
//	fb := NewFFUserBuilder().
//		UserName("userName").
//		Country("country").
//		Email("email").
//		Custom("key", "value").Build()
//
type FFUserBuilder struct {
	FFCUser
}

// Of build a FFCUser from {@link UserTag}
// @Param tags a UserTag map
// @Return a FFCUser
func (f *FFCUser) Of(tags map[UserTag]string) FFCUser {

	fb := NewFFUserBuilder()
	for tag, value := range tags {

		tagLower := strings.ToLower(tag.UserProperty)
		switch tagLower {

		case KeyId:
			fb.Key(value)
			continue
		case Name:
			fb.UserName(value)
			continue
		case Email:
			fb.Email(value)
			continue
		case Country:
			fb.Country(value)
			continue
		default:
			if len(tag.UserProperty) > 0 {
				fb.Custom(tag.UserProperty, value)
			} else {
				fb.Custom(tag.RequestProperty, value)
			}
		}
	}
	if len(fb.FFCUser.Key) == 0 {
		return FFCUser{}
	}
	return fb.Build()
}

// GetProperty Gets the value of a user attribute, if present.
// This can be either a built-in attribute or a custom one
// @Param attribute – the attribute to get
// @Return the attribute value or nil
func (f *FFCUser) GetProperty(attribute string) string {

	if strings.ToLower(attribute) == Name {
		return f.UserName
	}

	if strings.ToLower(attribute) == Email {
		return f.Email
	}

	if strings.ToLower(attribute) == KeyId {
		return f.Key
	}

	if strings.ToLower(attribute) == Country {
		return f.Country
	}
	return f.Custom[attribute]
}

func (f *FFCUser) IsEmpty() bool {
	return f == nil
}

func NewFFUserBuilder() *FFUserBuilder {

	customMap := make(map[string]string, 0)
	fb := new(FFUserBuilder)
	fb.FFCUser.Custom = customMap
	return fb

}

// Key Changes the user's key.
// @Return the builder
func (fb *FFUserBuilder) Key(key string) *FFUserBuilder {
	fb.FFCUser.Key = key
	return fb
}

// Email set the user's email.
// @Return the builder
func (fb *FFUserBuilder) Email(email string) *FFUserBuilder {
	fb.FFCUser.Email = email
	return fb
}

// UserName   set the user's userName.
// @Return the builder
func (fb *FFUserBuilder) UserName(userName string) *FFUserBuilder {
	fb.FFCUser.UserName = userName
	return fb
}

// Country  set the user's country.
// @Return the builder
func (fb *FFUserBuilder) Country(country string) *FFUserBuilder {
	fb.FFCUser.Country = country
	return fb
}

// Custom   Adds a String-valued custom attribute. When set to one of the built-in user attribute keys
// the key/value pair will be ignored.
// @Return the builder
func (fb *FFUserBuilder) Custom(key string, value string) *FFUserBuilder {

	if len(key) > 0 && len(value) > 0 {
		fb.FFCUser.Custom[key] = value
	}
	return fb
}

// Build Builds the configured FFCUser object.
// @Return the FFCUser configured by this builder
func (fb *FFUserBuilder) Build() FFCUser {

	if len(fb.FFCUser.Key) == 0 {
		return FFCUser{}
	}
	return fb.FFCUser
}
