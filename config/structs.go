package config

import (
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}
