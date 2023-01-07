package model

// 站点导航
type ActionLink struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// 积分配置
type ScoreConfig struct {
	PostTopicScore   int64 `json:"postTopicScore"`   // 发帖获得积分
	PostCommentScore int64 `json:"postCommentScore"` // 跟帖获得积分
	CheckInScore     int64 `json:"checkInScore"`     // 签到积分
}

// vip积分配置
type VipConfig struct {
	Vip1 int64 `json:"vip1"` // vip1每日积分
	Vip2 int64 `json:"vip2"` // vip2每日积分
	Vip3 int64 `json:"vip3"` // vip3每日积分
	Vip4 int64 `json:"vip4"` // vip4每日积分
	Vip5 int64 `json:"vip5"` // vip5每日积分
	Vip6 int64 `json:"vip6"` // vip6每日积分
}

type LoginMethod struct {
	Password bool `json:"password"`
	QQ       bool `json:"qq"`
	Github   bool `json:"github"`
	Osc      bool `json:"osc"`
}

// SysConfigResponse
//
//	配置返回结构体
type SysConfigResponse struct {
	SiteTitle                  string         `json:"siteTitle"`
	SiteDescription            string         `json:"siteDescription"`
	SiteKeywords               []string       `json:"siteKeywords"`
	SiteNavs                   []ActionLink   `json:"siteNavs"`
	SiteNotification           string         `json:"siteNotification"`
	RecommendTags              []string       `json:"recommendTags"`
	UrlRedirect                bool           `json:"urlRedirect"`
	ScoreConfig                ScoreConfig    `json:"scoreConfig"`
	VipConfig                  VipConfig      `json:"vipConfig"`
	DefaultNodeId              int64          `json:"defaultNodeId"`
	ArticlePending             bool           `json:"articlePending"`
	TopicCaptcha               bool           `json:"topicCaptcha"`
	UserObserveSeconds         int            `json:"userObserveSeconds"`
	TokenExpireDays            int            `json:"tokenExpireDays"`
	LoginMethod                LoginMethod    `json:"loginMethod"`
	CreateTopicEmailVerified   bool           `json:"createTopicEmailVerified"`
	CreateArticleEmailVerified bool           `json:"createArticleEmailVerified"`
	CreateCommentEmailVerified bool           `json:"createCommentEmailVerified"`
	EnableHideContent          bool           `json:"enableHideContent"`
	Modules                    []ModuleConfig `json:"modules"`
	EmailWhitelist             []string       `json:"emailWhitelist"` // 邮箱白名单
}

// Moduleconfig
//
//	模块配置
type ModuleConfig struct {
	Module  string `json:"module"`
	Enabled bool   `json:"enabled"`
}
