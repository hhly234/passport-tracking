package types

type Employee struct {
	ID             string `json:"id,omitempty" bson:"_id"`
	Name           string `json:"name,omitempty" bson:"_name"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	PassportID     string `json:"passport_id,omitempty"`
	Role           string `json:"role,omitempty"`
	Project        string `json:"project,omitempty"`
	DeliveryCenter string `json:"delivery_center,omitempty"`
	EmailManager   string `json:"email_manager,omitempty"`
}
