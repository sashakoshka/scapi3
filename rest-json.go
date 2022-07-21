package scapi3

/* HealthResponse contains data about the health of the Scratch website.
 */
type HealthResponse struct {
	Version string      `json:"version"`
	Uptime  int         `json:"uptime"`
	Load    []float64   `json:"load"`
	SQL struct {
		Main                HealthResponseSQL `json:"main"`
		ProjectComments     HealthResponseSQL `json:"project_comments"`
		GalleryComments     HealthResponseSQL `json:"gallery_comments"`
		UserprofileComments HealthResponseSQL `json:"userprofile_comments"`
		Timestamp           uint64            `json:"timestamp"`
	} `json:"sql"`
	Cache struct {
		Connected bool `json:"connected"`
		Ready     bool `json:"ready"`
	} `json:"cache"`
}

/* HealthResponseSQL represents an database in the site SQL health data.
 */
type HealthResponseSQL struct {
	Primary HealthResponseSQLStatistic `json:"primary"`
	Replica HealthResponseSQLStatistic `json:"replica"`
}

/* HealthResponseSQL represents a database's health data in the site SQL
 * health data.
 */
type HealthResponseSQLStatistic struct {
	SSL             bool `json:"ssl"`
	Destroyed       bool `json:"destroyed"`
	Min             int  `json:"min"`
	Max             int  `json:"max"`
	NumUsed         int  `json:"numUsed"`
	NumFree         int  `json:"numFree"`
	PendingAcquires int  `json:"pendingAcquires"`
	PendingCreates  int  `json:"pendingCreates"`
}

/* NewsResponse represents a list of recent news from the Scratch website.
 */
type NewsResponse []NewsResponseItem

/* NewsResponseItem represents a single news article
 */
type NewsResponseItem struct {
	Id       uint64 `json:"id"`
	Stamp    string `json:"stamp"`
	Headline string `json:"headline"`
	URL      string `json:"url"`
	Image    string `json:"image"`
	Copy     string `json:"copy"`
}

/* CountResponse contains a simple count.
 */
type CountResponse struct {
	Count uint64 `json:"count"`
}

/* ProjectResponse holds information about a project.
 */
type ProjectResponse struct {
	ID              uint64       `json:"id"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	Instructions    string       `json:"instructions"`
	Visibility      string       `json:"visibility"`
	Public          bool         `json:"public"`
	CommentsAllowed bool         `json:"comments_allowed"`
	IsPublished     bool         `json:"is_published"`
	Author          UserResponse `json:"author"`
	Image           string       `json:"image"`
	Images struct {
		Size218px string `json:"282x218"`
		Size163px string `json:"216x163"`
		Size200px string `json:"200x200"`
		Size108px string `json:"144x108"`
		Size102px string `json:"135x102"`
		Size80px  string `json:"100x80"`
	} `json:"images"`
	History struct {
		Created  string `json:"created"`
		Modified string `json:"modified"`
		Shared   string `json:"shared"`
	} `json:"history"`
	Stats struct {
		Views     int `json:"views"`
		Loves     int `json:"loves"`
		Favorites int `json:"favorites"`
		Remixes   int `json:"remixes"`
	} `json:"stats"`
	Remix struct {
		Parent uint64 `json:"parent"`
		Root   uint64 `json:"root"`
	} `json:"remix"`
	ProjectToken string `json:"project_token"`
}

/* StudioResponse holds information about a studio. 
 */
type StudioResponse struct {
	ID              uint64 `json:"id"`
	Title           string `json:"title"`
	Host            uint64 `json:"host"`
	Description     string `json:"description"`
	Visibility      string `json:"visibility"`
	Public          bool   `json:"public"`
	OpenToAll       bool   `json:"open_to_all"`
	CommentsAllowed bool   `json:"comments_allowed"`
	Image           string `json:"image"`
	History struct {
		Created  string `json:"created"`
		Modified string `json:"modified"`
	} `json:"history"`
	Stats struct {
		Comments  int `json:"comments"`
		Followers int `json:"followers"`
		Managers  int `json:"managers"`
		Projects  int `json:"projects"`
	} `json:"stats"`
}

/* StudioProjectsResponse represents a list of projects inside of a studio.
 */
type StudioProjectsResponse []StudioProjectsResponseItem

/* StudioProjectsResponseItem represents one project in a studio.
 */
type StudioProjectsResponseItem struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	CreatorID uint64 `json:"creator_id"`
	Username  string `json:"username"`
	Avatar struct {
		Size90px string `json:"90x90"`
		Size60px string `json:"60x60"`
		Size55px string `json:"55x55"`
		Size50px string `json:"50x50"`
		Size32px string `json:"32x32"`
	} `json:"avatar"`
	ActorID uint64 `json:"actor_id"`
}

/* StudioActivityResponse represents a list of recent activity in a studio.
 */
type StudioActivityResponse []StudioActivityResponseItem

/* StudioActivityResponseItem represents a single activity in a studio.
 */
type StudioActivityResponseItem struct {
	DateTimeCreated string `json:"datetime_created"`
	ID              string `json:"id"`
	ActorID         uint64 `json:"actor_id"`
	ProjectID       uint64 `json:"project_id"`
	ProjectTitle    string `json:"project_title"`
	Type            string `json:"type"`
	ActorUsername   string `json:"actor_username"`
}

/* CommentResponse represents information about a single comment
 */
type CommentResponse struct {
	ID               uint64 `json:"id"`
	ParentID         uint64 `json:"parent_id"`
	CommenteeID      uint64 `json:"commentee_id"`
	Content          string `json:"content"`
	DateTimeCreated  string `json:"datetime_created"`
	DateTimeModified string `json:"datetime_modified"`
	Visibility       string `json:"visibility"`
	Author struct {
		ID          uint64 `json:"id"`
		Username    string `json:"username"`
		ScratchTeam bool   `json:"scratchteam"`
		Image       string `json:"image"`
	} `json:"author"`
	ReplyCount int `json:"reply_count"`
}

/* FeaturedResponse contains information about what's on the front page of the
 * Scratch website.
 */
type FeaturedResponse struct {
	CommunityNewestProjects      []FeaturedResponseItem `json:"community_newest_projects"`
	CommunityMostRemixedProjects []FeaturedResponseItem `json:"community_most_remixed_projects"`
	CommunityMostLovedProjects   []FeaturedResponseItem `json:"community_most_loved_projects"`
	CommunityFeaturedStudios     []FeaturedResponseItem `json:"community_featured_studios"`
	CommunityFeaturedProjects    []FeaturedResponseItem `json:"community_featured_projects"`
	ScratchDesignStudio          []FeaturedResponseItem `json:"scratch_design_studio"`
	CuratorTopProjects           []FeaturedResponseItem `json:"curator_top_projects"`
}

/* FeaturedResponseItem contains information about a front paged project.
 */
type FeaturedResponseItem struct {
	ThumbnailUrl string `json:"thumbnail_url"`
	Title        string `json:"title"`
	Creator      string `json:"string"`
	Type         string `json:"type"`
	ID           uint64 `json:"id"`
	LoveCount    int    `json:"love_count"`
}

/* UserResponse holds information about a user.
 */
type UserResponse struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	ScratchTeam bool   `json:"scratchteam"`
	History struct {
		Joined string `json:"joined"`
	} `json:"history"`
	Profile struct {
		ID uint64 `json:"id"`
		Images struct {
			Size90px string `json:"90x90"`
			Size60px string `json:"60x60"`
			Size55px string `json:"55x55"`
			Size50px string `json:"50x50"`
			Size32px string `json:"32x32"`
		} `json:"images"`
		Status  string `json:"status"`
		Bio     string `json:"bio"`
		Country string `json:"country"`
	} `json:"profile"`
}
