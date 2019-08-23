package types

type Passport struct {
	ID           string `json:"id,omitempty" bson:"_id"`
	FullName     string `json:"fullname,omitempty"`
	Nationality  string `json:"nationality,omitempty"`
	DateOfBirth  string `json:"date_of_birth,omitempty"`
	PlaceOfBirth string `json:"place_of_birth,omitempty"`
	Sex          string `json:"sex,omitempty"`
	IDCardN      string `json:"id_card_n,omitempty"`
	DateOfIssue  string `json:"date_of_issue,omitempty"`
	DateOfExpiry string `json:"date_of_expiry,omitempty"`
}
