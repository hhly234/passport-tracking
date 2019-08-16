package types

type employee struct {
	ID           string `json:"id,omitempty" bson:"_id"`
	Name         string `json:"name,omitempty" bson:"_name"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PassportID   string `json:"passportid,omitempty"`
	Role         string `json:"role,omitempty"`
	Project      string `json:"project,omitempty"`
	DC           string `json:"dc,omitempty"`
	EmailManager string `json:"emailmanager,omitempty"`
}
