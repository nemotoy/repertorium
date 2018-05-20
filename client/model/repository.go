package model

import "time"

// Repository ...
type Repository struct {
	Owner struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		HTMLURL           string `json:"html_url"`
		GravatarID        string `json:"gravatar_id"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
		URL               string `json:"url"`
		EventsURL         string `json:"events_url"`
		FollowingURL      string `json:"following_url"`
		FollowersURL      string `json:"followers_url"`
		GistsURL          string `json:"gists_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
	} `json:"owner"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	CodeOfConduct struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"code_of_conduct"`
	DefaultBranch    string    `json:"default_branch"`
	CreatedAt        time.Time `json:"created_at"`
	PushedAt         time.Time `json:"pushed_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	HTMLURL          string    `json:"html_url"`
	CloneURL         string    `json:"clone_url"`
	GitURL           string    `json:"git_url"`
	SSHURL           string    `json:"ssh_url"`
	SvnURL           string    `json:"svn_url"`
	Language         string    `json:"language"`
	Fork             bool      `json:"fork"`
	ForksCount       int       `json:"forks_count"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Size             int       `json:"size"`
	Private          bool      `json:"private"`
	HasIssues        bool      `json:"has_issues"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasProjects      bool      `json:"has_projects"`
	HasDownloads     bool      `json:"has_downloads"`
	Archived         bool      `json:"archived"`
	URL              string    `json:"url"`
	ArchiveURL       string    `json:"archive_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BlobsURL         string    `json:"blobs_url"`
	BranchesURL      string    `json:"branches_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	CommentsURL      string    `json:"comments_url"`
	CommitsURL       string    `json:"commits_url"`
	CompareURL       string    `json:"compare_url"`
	ContentsURL      string    `json:"contents_url"`
	ContributorsURL  string    `json:"contributors_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	DownloadsURL     string    `json:"downloads_url"`
	EventsURL        string    `json:"events_url"`
	ForksURL         string    `json:"forks_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	IssuesURL        string    `json:"issues_url"`
	KeysURL          string    `json:"keys_url"`
	LabelsURL        string    `json:"labels_url"`
	LanguagesURL     string    `json:"languages_url"`
	MergesURL        string    `json:"merges_url"`
	MilestonesURL    string    `json:"milestones_url"`
	NotificationsURL string    `json:"notifications_url"`
	PullsURL         string    `json:"pulls_url"`
	ReleasesURL      string    `json:"releases_url"`
	StargazersURL    string    `json:"stargazers_url"`
	StatusesURL      string    `json:"statuses_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	TagsURL          string    `json:"tags_url"`
	TreesURL         string    `json:"trees_url"`
	TeamsURL         string    `json:"teams_url"`
}