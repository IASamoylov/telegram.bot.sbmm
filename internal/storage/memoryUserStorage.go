package storage

import (
	"fmt"
	"log"

	"ia.samoylov/telegram.bot.sbmm/internal/models"
)

// Storage contains structure of data for storing user data.
type Storage struct {
	items map[int]*models.User
}

// NewUserStorage returns a new user memroy storage.
func NewUserStorage() *Storage {
	var storage = Storage{}

	storage.items = make(map[int]*models.User)

	return &storage
}

//Get returns User from storage or error if user not found.
func (stg *Storage) Get(ID int) (models.User, error) {
	userPtr, isExist := stg.items[ID]

	if !isExist {
		return models.User{
			ID: ID,
		}, fmt.Errorf(`User %v not found`, ID)
	}

	log.Println("GET", userPtr)
	return *userPtr, nil
}

// Add creates a new User in storage if the user already exit returns error.
func (stg *Storage) Add(u models.User) error {
	_, isExist := stg.items[u.ID]

	if isExist {
		return nil
	}

	stg.items[u.ID] = &u

	log.Println("Add", stg.items[u.ID])
	return nil
}

// Update updates user in storage.
func (stg *Storage) Update(u models.User) error {
	userPtr, isExist := stg.items[u.ID]

	log.Println("Update", u)
	if !isExist {
		return fmt.Errorf(`User %v not found`, u.ID)
	}

	userPtr.Language = u.Language
	userPtr.Platform = u.Platform
	userPtr.LastGameTime = u.LastGameTime

	return nil
}

// Delete user from storage.
func (stg *Storage) Delete(ID int) {
	delete(stg.items, ID)
}
