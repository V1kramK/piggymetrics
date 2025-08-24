package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_accountname "github.com/input-api/mcp-server/tools/accountname"
	tools_current "github.com/input-api/mcp-server/tools/current"
	tools_recipients "github.com/input-api/mcp-server/tools/recipients"
	tools_users "github.com/input-api/mcp-server/tools/users"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_accountname.CreateGet_accountnameTool(cfg),
		tools_accountname.CreatePut_accountnameTool(cfg),
		tools_current.CreateGet_currentTool(cfg),
		tools_recipients.CreateGet_recipientsTool(cfg),
		tools_users.CreateGet_usersTool(cfg),
	}
}
