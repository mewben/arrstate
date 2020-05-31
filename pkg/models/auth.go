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

// NewCurrentUser -
func NewCurrentUser() *CurrentUser {
	return &CurrentUser{
		User:   NewUserModel(),
		Person: NewPersonModel(),
	}
}

// NewAuthSuccessResponse -
func NewAuthSuccessResponse() *AuthSuccessResponse {
	return &AuthSuccessResponse{
		CurrentUser:     NewCurrentUser(),
		CurrentBusiness: NewBusinessModel(),
		UserBusinesses:  make([]*BusinessModel, 0),
	}
}
