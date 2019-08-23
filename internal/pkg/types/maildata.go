package types

type EmailData struct {
	Name         string `json:"name,omitempty"`
	PassportID   string `json:"passport_id,omitempty"`
	DateOfExpiry string `json:"date_of_expiry,omitempty"`
}

type NotificationData struct {
	Data []EmailData
}
