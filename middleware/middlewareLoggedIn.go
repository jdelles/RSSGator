package middleware

import (
	"context"
	"fmt"

	"RSSGator/commands"
	"RSSGator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *commands.State, cmd commands.Command, user database.User) error) func(*commands.State, commands.Command) error {
	return func(s *commands.State, cmd commands.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not logged in: %w", err)

		}
		return handler(s, cmd, user)
	}
}
