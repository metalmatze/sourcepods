package user

// Service handles all interactions with users.
type Service interface {
	FindAll() ([]*User, error)
	FindByUsername(string) (*User, error)
	Create(*User) (*User, error)
	Update(string, *User) (*User, error)
	Delete(string) error
}

// Store users after manipulation or read them.
type Store interface {
	FindAll() ([]*User, error)
	Find(string) (*User, error)
	FindByUsername(string) (*User, error)
	Create(*User) (*User, error)
	Update(string, *User) (*User, error)
	Delete(string) error
}

type service struct {
	users Store
}

// NewService returns a Service that handles all interactions with users.
func NewService(users Store) Service {
	return &service{users: users}
}

func (s *service) FindAll() ([]*User, error) {
	return s.users.FindAll()
}

func (s *service) FindByUsername(username string) (*User, error) {
	return s.users.FindByUsername(username)
}

func (s *service) Create(user *User) (*User, error) {
	return user, nil
}

func (s *service) Update(username string, user *User) (*User, error) {
	return s.users.Update(username, user)
}

func (s *service) Delete(username string) error {
	return nil
}
