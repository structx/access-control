package domain

// Publish
type Publish string

// Subscriptions
type Subscriptions string

// RequestResponse
type RequestResponse string

const (

	// UserAccessGranted
	UserAccessGranted Subscriptions = "user_access_granted"
	// UserAccessDenied
	UserAccessDenied Subscriptions = "user_access_denied"
	// ServiceAccessGranted
	ServiceAccessGranted Subscriptions = "service_access_granted"
	// ModifyAccessControlList
	ModifyAccessControlList Subscriptions = "modify_access_control_list"

	// VerifyUserAcess
	VerifyUserAccess RequestResponse = "verify_user_access"
	// VerifyServiceAccess
	VerifyServiceAccess RequestResponse = "verify_service_access"
)

// String
func (p Publish) String() string {
	return string(p)
}

// List
func (p Publish) List() []Publish {
	return []Publish{}
}

// String
func (s Subscriptions) String() string {
	return string(s)
}

// ListSubscriptions
func ListSubscriptions() []Subscriptions {
	return []Subscriptions{
		UserAccessDenied,
		UserAccessGranted,
		ServiceAccessGranted,
		ModifyAccessControlList,
	}
}

// String
func (rr RequestResponse) String() string {
	return string(rr)
}

// List
func (rr RequestResponse) List() []RequestResponse {
	return []RequestResponse{VerifyUserAccess}
}
