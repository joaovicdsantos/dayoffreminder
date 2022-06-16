package model

type SlackRequest struct {
	Token               string `form:"token"`
	TeamId              string `form:"team_id"`
	TeamDomain          string `form:"team_domain"`
	ChannelId           string `form:"channel_id"`
	ChannelName         string `form:"channel_name"`
	UserId              string `form:"user_id"`
	UserName            string `form:"user_name"`
	Command             string `form:"command"`
	Text                string `form:"text"`
	ApiAppId            string `form:"api_app_id"`
	IsEnterpriseInstall string `form:"is_enterprise_install"`
	ResponseUrl         string `form:"response_url"`
	TriggerId           string `form:"trigger_id"`
}
