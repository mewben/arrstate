package auth

import "github.com/gofiber/fiber"

// SignupPayload -
type SignupPayload struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Business   string `json:"business"`
	Domain     string `json:"domain"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

// Signup -
func (h *Handler) Signup(c *fiber.Ctx) {
	payload := &SignupPayload{}
	if err := c.BodyParser(&payload); err != nil {
		c.Next(err)
		return
	}

	// 1. Create Business
	// 2. Create User
	// 3. Create Person
	// 4. Return AuthSuccessResponse

	c.Status(200).JSON(payload)
	c.Next()
	return
}
