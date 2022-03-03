package repository

type UserRepository interface {
	ChangeRole(int64, int) error
	Create(int64) error
}
