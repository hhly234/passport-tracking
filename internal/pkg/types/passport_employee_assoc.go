package types

type PassportEmployeeAssoc struct {
	ID                   string `json:"id,omitempty" bson:"_id"`
	PassportID           string `json:"passport_id,omitempty"`
	EmployeeID           string `json:"employee_id,omitempty"`
	EmployeeName         string `json:"employee_name,omitempty"`
	DateOfExpiryPassport string `json:"date_of_expiry_passport,omitempty"`
}
