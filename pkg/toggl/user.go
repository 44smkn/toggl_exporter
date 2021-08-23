package toggl

import "time"

type User struct {
	Since int  `json:"since"`
	Data  Data `json:"data"`
}
type NewBlogPost struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
type TimeEntries struct {
	ID          int       `json:"id"`
	Wid         int       `json:"wid"`
	Billable    bool      `json:"billable"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    int       `json:"duration"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	At          time.Time `json:"at"`
}
type Projects struct {
	ID       int       `json:"id"`
	Wid      int       `json:"wid"`
	Name     string    `json:"name"`
	Billable bool      `json:"billable"`
	Active   bool      `json:"active"`
	At       time.Time `json:"at"`
	Color    string    `json:"color"`
}
type Tags struct {
	ID   int       `json:"id"`
	Wid  int       `json:"wid"`
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}
type Workspaces struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Premium bool      `json:"premium"`
	At      time.Time `json:"at"`
}
type Clients struct {
	ID   int       `json:"id"`
	Wid  int       `json:"wid"`
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}
type Data struct {
	ID                    int           `json:"id"`
	APIToken              string        `json:"api_token"`
	DefaultWid            int           `json:"default_wid"`
	Email                 string        `json:"email"`
	Fullname              string        `json:"fullname"`
	JqueryTimeofdayFormat string        `json:"jquery_timeofday_format"`
	JqueryDateFormat      string        `json:"jquery_date_format"`
	TimeofdayFormat       string        `json:"timeofday_format"`
	DateFormat            string        `json:"date_format"`
	StoreStartAndStopTime bool          `json:"store_start_and_stop_time"`
	BeginningOfWeek       int           `json:"beginning_of_week"`
	Language              string        `json:"language"`
	ImageURL              string        `json:"image_url"`
	SidebarPiechart       bool          `json:"sidebar_piechart"`
	At                    time.Time     `json:"at"`
	Retention             int           `json:"retention"`
	RecordTimeline        bool          `json:"record_timeline"`
	RenderTimeline        bool          `json:"render_timeline"`
	TimelineEnabled       bool          `json:"timeline_enabled"`
	TimelineExperiment    bool          `json:"timeline_experiment"`
	NewBlogPost           NewBlogPost   `json:"new_blog_post"`
	TimeEntries           []TimeEntries `json:"time_entries"`
	Projects              []Projects    `json:"projects"`
	Tags                  []Tags        `json:"tags"`
	Workspaces            []Workspaces  `json:"workspaces"`
	Clients               []Clients     `json:"clients"`
}
