package domain

// Publish application publish topics
type Publish string

// Subscriptions application subscribe topics
type Subscriptions string

// RequestResponse application request/response topics
type RequestResponse string

const (
	// UserAccessGranted user access granted
	UserAccessGranted Subscriptions = "user_access_granted"
	// UserAccessDenied user access denied
	UserAccessDenied Subscriptions = "user_access_denied"
	// ServiceAccessGranted service access granted
	ServiceAccessGranted Subscriptions = "service_access_granted"
	// ModifyAccessControlList modify access control list
	ModifyAccessControlList Subscriptions = "modify_access_control_list"

	// VerifyUserAccess verify user access
	VerifyUserAccess RequestResponse = "verify_user_access"
	// VerifyServiceAccess verify service access
	VerifyServiceAccess RequestResponse = "verify_service_access"
)

// String stringify topic
func (p Publish) String() string {
	return string(p)
}

// ListPublish topics
func ListPublish() []Publish {
	return []Publish{}
}

// String stringify topic
func (s Subscriptions) String() string {
	return string(s)
}

// ListSubscriptions topics
func ListSubscriptions() []Subscriptions {
	return []Subscriptions{
		UserAccessDenied,
		UserAccessGranted,
		ServiceAccessGranted,
		ModifyAccessControlList,
	}
}

// String stringify topic
func (rr RequestResponse) String() string {
	return string(rr)
}

// List request/response topics
func (rr RequestResponse) List() []RequestResponse {
	return []RequestResponse{VerifyUserAccess}
}
