package authorization

import (
	"github.com/gitpods/gitpods"
	"github.com/gitpods/gitpods/session"
	"golang.org/x/crypto/bcrypt"
)

// Service authenticates users and creates sessions for them.
type Service interface {
	AuthenticateUser(email, password string) (*gitpods.User, error)
	CreateSession(string, string) (*session.Session, error)
}

// Store finds users by emails.
type Store interface {
	FindUserByEmail(string) (*gitpods.User, error)
}

// NewService takes a store to find users by their email and
// takes a session service to create sessions for them once authenticated.
func NewService(store Store, sessions session.Service) Service {
	return &service{store: store, sessions: sessions}
}

type service struct {
	store    Store
	sessions session.Service
}

// AuthenticateUser by hashing the given password an comparing it against the one stored.
func (s *service) AuthenticateUser(email, password string) (*gitpods.User, error) {
	user, err := s.store.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) CreateSession(userID, userUsername string) (*session.Session, error) {
	return s.sessions.CreateSession(userID, userUsername)
}