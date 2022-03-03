package bot

type UserService interface {
	ChangeRole(int64, int) error
	Create(int64) error
}
