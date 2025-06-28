package commands

import (
	"RSSGator/internal/config"
	"RSSGator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
