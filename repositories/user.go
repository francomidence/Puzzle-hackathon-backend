package repositories

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"time"

	"puzzle-hackathon-backend/models"

	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUsers(ctx context.Context) (*[]models.User, error)
	GetUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) (*models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	supabase := InitializeSupabaseClient()

	user.ID = uuid.NewV4()

	user.CreatedAt = time.Now()
	user.Active = true

	var results []map[string]interface{}
	err := supabase.DB.From("Users").Insert(user).Execute(&results)
	if err != nil {
		fmt.Println("The error is: ", err.Error())
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		// Convert the first element of 'results' to JSON
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var newUser models.User // Use a single instance of 'models.User'
		err = json.Unmarshal(jsonData, &newUser)
		if err != nil {
			panic(err)
		}

		return &newUser, nil
	}

	return nil, fmt.Errorf("no user data returned after insert operation")
}

func (r *userRepository) GetUsers(ctx context.Context) (*[]models.User, error) {
	supabase := InitializeSupabaseClient()

	var results []map[string]interface{} // Change to a slice of maps
	err := supabase.DB.From("Users").Select("*").Eq("active", "true").Execute(&results)
	if err != nil {
		fmt.Println("The error is ", err)
		fmt.Println(err.Error())
		panic(err)
	}

	// Convert to JSON and then to User struct
	jsonData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	var users *[]models.User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		panic(err)
	}

	return users, nil
}

func (r *userRepository) GetUser(ctx context.Context, user *models.User) (*models.User, error) {
	supabase := InitializeSupabaseClient()

	var results []map[string]interface{} // Change to a slice of maps
	err := supabase.DB.From("Users").Select("*").Eq("id", user.ID.String()).Execute(&results)
	if err != nil {
		panic(err)
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedUser models.User
		err = json.Unmarshal(jsonData, &updatedUser)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedUser, nil
	}

	return nil, fmt.Errorf("no user")
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	supabase := InitializeSupabaseClient()

	// Assign current time to UpdatedAt
	user.UpdatedAt = time.Now()

	// Create a map of fields to update, excluding 'CreatedAt' and any other fields you don't want to update
	updateFields := map[string]interface{}{
		"updated_at": user.UpdatedAt,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"role":       user.Role,
		"code":       user.Code,
		"active":     user.Active,
	}

	var results []map[string]interface{}
	err := supabase.DB.From("Users").Update(updateFields).Eq("id", user.ID.String()).Execute(&results)
	if err != nil {
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedUser models.User
		err = json.Unmarshal(jsonData, &updatedUser)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedUser, nil
	}

	return nil, fmt.Errorf("no user data returned after update operation")
}

func (r *userRepository) DeleteUser(ctx context.Context, user *models.User) (*models.User, error) {
	supabase := InitializeSupabaseClient()

	// Assign current time to DeletedAt
	user.DeletedAt = time.Now()

	updateFields := map[string]interface{}{
		"deleted_at": user.DeletedAt,
		"active":     false,
	}

	var results []map[string]interface{}
	err := supabase.DB.From("Users").Update(updateFields).Eq("id", user.ID.String()).Execute(&results)
	if err != nil {
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedUser models.User
		err = json.Unmarshal(jsonData, &updatedUser)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedUser, nil
	}

	return nil, fmt.Errorf("no user data returned after update operation")
}
