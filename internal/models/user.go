package models

import "fmt"

// User holds information about bot user.
type User struct {
	ID           int
	Language     string
	Platform     string
	UserName     string
	LastGameTime string
}

// UserStorage defines all methods for working with user.
type UserStorage interface {
	Get(ID int) (User, error)
	Add(u User) error
	Update(u User) error
	Delete(ID int)
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (u *User) String() string {
	return fmt.Sprintf(`
{
	ID: %d,
	LanguageCode: %s,
	Platform: %s,
}`,
		u.ID,
		u.Language,
		u.Platform,
	)
}
