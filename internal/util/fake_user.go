package util

import (
    "douq.merouaneamqor.com/internal/model"
    "github.com/brianvoe/gofakeit/v6"
    "time"
)

func GenerateFakeUser() model.User {
    var user model.User

    // Manually set fields to avoid overwriting the ID
    user.Name = gofakeit.Name()
    user.Email = gofakeit.Email()
    user.Password = gofakeit.Password(true, true, true, true, false, 8)
    user.Birthdate = time.Time(gofakeit.Date())
    user.IsActive = gofakeit.Bool()

    return user
}
