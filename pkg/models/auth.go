package models

// CurrentUser -
type CurrentUser struct {
	User   *UserModel   `json:"user"`
	Person *PersonModel `json:"person"`
}

// AuthSuccessResponse -
type AuthSuccessResponse struct {
	Token           string           `json:"token"`
	CurrentUser     *CurrentUser     `json:"currentUser"`
	CurrentBusiness *BusinessModel   `json:"currentBusiness"`
	UserBusinesses  []*BusinessModel `json:"userBusinesses"`
}
