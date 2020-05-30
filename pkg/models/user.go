package models

// UserModel -
type UserModel struct {
	ID            string `bson:"_id,omitempty" json:"_id,omitempty"`
	Email         string `bson:"email" json:"email"`
	Password      string `bson:"password" json:"-"`
	AccountStatus string `bson:"accountStatus" json:"accountStatus"`
}

// NewUserModel -
func NewUserModel() *UserModel {
	return &UserModel{}
}
