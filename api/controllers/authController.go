package controllers


import ()

func Login(c *fiber.Ctx) {
	// Get the request body
	var loginData LoginData
	if err := c.BodyParser(&loginData); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return
	}

	// Validate the login data
	if err := validateLoginData(loginData); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid login data",
			"errors":  err,
		})
		return
	}

	// Find the user by email
	user, err := models.FindUserByEmail(loginData.Email)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
		return
	}

	// Return the token
	c.JSON(fiber.Map{
		"token": token,
	})
}