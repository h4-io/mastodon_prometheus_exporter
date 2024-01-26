package mastodon

type nestedStats struct {
	UserCount   int `json:"user_count"`
	StatusCount int `json:"status_count"`
	DomainCount int `json:"domain_count"`
}

type GenericStats struct {
	Title   string      `json:"title"`
	Version string      `json:"version"`
	Stats   nestedStats `json:"stats"`
}

type Activities struct {
	Week          string `json:"week"`
	Statuses      string `json:"statuses"`
	Logins        string `json:"logins"`
	Registrations string `json:"registrations"`
}

type nodeInfoUsers struct {
	ActiveMonth    int `json:"activeMonth"`
	ActiveHalfYear int `json:"activeHalfYear"`
}

type nodeInfoUsage struct {
	Users nodeInfoUsers `json:"users"`
}

type NodeInfo struct {
	Usage nodeInfoUsage `json:"usage"`
}
