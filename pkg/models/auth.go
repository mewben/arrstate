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
func NewCurrentUser(userID, businessID string) *CurrentUser {
	return &CurrentUser{
		User:   NewUserModel(),
		Person: NewPersonModel(userID, businessID),
	}
}

// NewAuthSuccessResponse -
func NewAuthSuccessResponse(userID, businessID string) *AuthSuccessResponse {
	return &AuthSuccessResponse{
		CurrentUser:     NewCurrentUser(userID, businessID),
		CurrentBusiness: NewBusinessModel(),
		UserBusinesses:  make([]*BusinessModel, 0),
	}
}
