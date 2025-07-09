package handlers

import (
	"RSSGator/commands"
	"context"
	"fmt"
)

func HandlerReset(s *commands.State, _ commands.Command) error {
	ctx := context.Background()
	if err := s.Db.DeleteUsers(ctx); err != nil {
		return err
	}
	fmt.Println("All users deleted.")
	return nil
}
