// repository/repository.go

package repository

// Repository represents your data storage or database interaction interface.
type Repository interface {
    GetUserByID(userID int) (error)
    // Add more repository methods as needed
}
