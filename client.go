package nuapiclient

import (
	"encoding/json"
	"net/http"
)

const baseURL = "http://api.asg.northwestern.edu/"

// A Term is a term.
type Term struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// A School is a school.
type School struct {
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
}

// A Subject is a subject.
type Subject struct {
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
}

// A Course is a course.
type Course struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Term        string `json:"term,omitempty"`
	Instructor  string `json:"instructor,omitempty"`
	Subject     string `json:"subject,omitempty"`
	CatalogNum  string `json:"catalog_num,omitempty"`
	Section     string `json:"section,omitempty"`
	Room        string `json:"room,omitempty"`
	MeetingDays string `json:"meeting_days,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	Seats       int    `json:"seats,omitempty"`
	Topic       string `json:"topic,omitempty"`
	Component   string `json:"component,omitempty"`
	ClassNum    int    `json:"class_num,omitempty"`
	CourseID    int    `json:"course_id,omitempty"`
}

// An Instructor is an instructor.
type Instructor struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Bio         string `json:"bio,omitempty"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	OfficeHours string `json:"office_hours,omitempty"`
	// Subjects refers to any subjects in which instructors have ever taught courses.
	Subjects []string `json:"subjects,omitempty"`
}

// A Building is a building.
type Building struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Lat        float64 `json:"lat,omitempty"`
	Lon        float64 `json:"lon,omitempty"`
	NUMapsLink string  `json:"nu_maps_link,omitempty"`
}

// A Room is a room.
type Room struct {
	ID         int    `json:"id,omitempty"`
	BuildingID int    `json:"building_id,omitempty"`
	Name       string `json:"name,omitempty"`
}

// A Client is a Northwestern Course Data API client.
type Client struct {
	key    string
	client *http.Client
}

// Terms returns a list of terms for which course data is available.
func (c *Client) Terms() ([]*Term, error) {
	req, err := http.NewRequest("GET", baseURL+"terms", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var terms []*Term
	if err := json.NewDecoder(resp.Body).Decode(&terms); err != nil {
		return nil, err
	}
	return terms, nil
}

// Schools lists all schools at Northwestern University.
func (c *Client) Schools() ([]*School, error) {
	req, err := http.NewRequest("GET", baseURL+"schools", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var schools []*School
	if err := json.NewDecoder(resp.Body).Decode(&schools); err != nil {
		return nil, err
	}
	return schools, nil
}

// A SubjectsConfig is a config for a Subjects request.
type SubjectsConfig struct {
	Term   string
	School string
}

// Subjects returns a list of subjects.
//
// You can filter by `term` and `school` - filtering by `term` is recommended because subjects have changed over the years.
func (c *Client) Subjects(config SubjectsConfig) ([]*Subject, error) {
	req, err := http.NewRequest("GET", baseURL+"subjects", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	if len(config.Term) > 0 {
		q.Add("term", config.Term)
	}
	if len(config.School) > 0 {
		q.Add("school", config.School)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var subjects []*Subject
	if err := json.NewDecoder(resp.Body).Decode(&subjects); err != nil {
		return nil, err
	}
	return subjects, nil
}

// A CoursesConfig is a config for a Courses request.
type CoursesConfig struct {
	// ID is the id of the course. Uniquely idenfities a course; used only by this API.
	ID string `json:"id,omitempty"`
	// Term is the id of the term.
	Term string `json:"term,omitempty"`
	// Subject is a symbol representing the subject of the course e.g. BIOL_SCI, PHIL.
	Subject string `json:"subject,omitempty"`
	// Instructor is the id of the instructor who teaches the course.
	Instructor string `json:"instructor,omitempty"`
	// Room is the id of the room in which the course is held.
	Room string `json:"room,omitempty"`
	// CatalogNum is the departmental catalog number of the course, including the sequence number.
	CatalogNum string `json:"catalog_num,omitempty"`
	// MeetingDays is the days of the week when the course meets.
	MeetingDays string `json:"meeting_days,omitempty"`
	// StartTime is the HH:MM (24-hour clock) time of day when the course starts.
	StartTime string `json:"start_time,omitempty"`
	// EndTime is the HH:MM (24-hour clock) time of day when the course stops.
	EndTime string `json:"end_time,omitempty"`
	// StartDate is the YYYY-MM-DD date when the course starts.
	StartDate string `json:"start_date,omitempty"`
	// EndDate is the YYYY-MM-DD date when the course stops.
	EndDate string `json:"end_date,omitempty"`
	// Seats is the number of seats available.
	Seats string `json:"seats,omitempty"`
	// Component is what part of the course it is e.g. LEC, LAB.
	Component string `json:"component,omitempty"`
	// Section is a particular occurrence of this course. Used to help identify the course if multiple professors or times are offered in a given term.
	Section string `json:"section,omitempty"`
	// ClassNum is the University's class number for the course (not unique).
	ClassNum string `json:"class_num,omitempty"`
	// CourseID is the University's id number for the course (not unique).
	CourseID string `json:"course_id,omitempty"`
}

// Courses returns a list of courses.
//
// To query the /courses endpoint, you must include a minimum set of parameters. These combinations are listed here:
// - `instructor`
// - `id` (multiple accepted up to 200)
// - `term`, `subject`
// - `term`, `room`
// You can include additional parameters to narrow down your search.
// The `start_time`, `end_time`, `start_date`, `end_date`, and `seats` expect exact values by default. You can append `__lt`, `__gt`, `__lte`, or `__gte` to each of them to filter by values that are less than or greater than and/or equal to the value you specify.
func (c *Client) Courses(config CoursesConfig) ([]*Course, error) {
	req, err := http.NewRequest("GET", baseURL+"courses", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	if len(config.ID) > 0 {
		q.Add("id", config.ID)
	}
	if len(config.Term) > 0 {
		q.Add("term", config.Term)
	}
	if len(config.Subject) > 0 {
		q.Add("subject", config.Subject)
	}
	if len(config.Instructor) > 0 {
		q.Add("instructor", config.Instructor)
	}
	if len(config.Room) > 0 {
		q.Add("room", config.Room)
	}
	if len(config.CatalogNum) > 0 {
		q.Add("catalog_num", config.CatalogNum)
	}
	if len(config.MeetingDays) > 0 {
		q.Add("meeting_days", config.MeetingDays)
	}
	if len(config.StartTime) > 0 {
		q.Add("start_time", config.StartTime)
	}
	if len(config.EndTime) > 0 {
		q.Add("end_time", config.EndTime)
	}
	if len(config.StartDate) > 0 {
		q.Add("start_date", config.StartDate)
	}
	if len(config.EndDate) > 0 {
		q.Add("end_date", config.EndDate)
	}
	if len(config.Seats) > 0 {
		q.Add("seats", config.Seats)
	}
	if len(config.Component) > 0 {
		q.Add("component", config.Component)
	}
	if len(config.Section) > 0 {
		q.Add("section", config.Section)
	}
	if len(config.ClassNum) > 0 {
		q.Add("class_num", config.ClassNum)
	}
	if len(config.CourseID) > 0 {
		q.Add("course_id", config.CourseID)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var courses []*Course
	if err := json.NewDecoder(resp.Body).Decode(&courses); err != nil {
		return nil, err
	}
	return courses, nil
}

// Instructors returns a list of instructors.
//
// Must include the `subject` parameter.
func (c *Client) Instructors(subject string) ([]*Instructor, error) {
	req, err := http.NewRequest("GET", baseURL+"instructors", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	q.Add("subject", subject)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var instructors []*Instructor
	if err := json.NewDecoder(resp.Body).Decode(&instructors); err != nil {
		return nil, err
	}
	return instructors, nil
}

// A BuildingsConfig is a config for a Buildings request.
type BuildingsConfig struct {
	ID  string `json:"id,omitempty"`
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}

// Buildings returns a list of buildings.
//
// When queried without parameters, returns all buildings. Most buildings have latitude and longitude coordinates.
// You can get information about particular buildings by including one or more `id` parameters, or you can filter using `lon` and `lat` with the `__lt`, `__gt`, `__lte`, and `__gte` suffixes. The Northwestern maps link may be null.
func (c *Client) Buildings(config BuildingsConfig) ([]*Building, error) {
	req, err := http.NewRequest("GET", baseURL+"buildings", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	if len(config.ID) > 0 {
		q.Add("id", config.ID)
	}
	if len(config.Lat) > 0 {
		q.Add("lat", config.Lat)
	}
	if len(config.Lon) > 0 {
		q.Add("lon", config.Lon)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buildings []*Building
	if err := json.NewDecoder(resp.Body).Decode(&buildings); err != nil {
		return nil, err
	}
	return buildings, nil
}

// A RoomsConfig is a config for a Rooms request.
type RoomsConfig struct {
	ID       string `json:"id,omitempty"`
	Building string `json:"building,omitempty"`
}

// Rooms returns a list of rooms.
//
// You must include a building id using the `building` parameter, or you can get details about specific rooms by including one or more `id` parameters.
func (c *Client) Rooms(config RoomsConfig) ([]*Room, error) {
	req, err := http.NewRequest("GET", baseURL+"rooms", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", c.key)
	if len(config.ID) > 0 {
		q.Add("id", config.ID)
	}
	if len(config.Building) > 0 {
		q.Add("building", config.Building)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rooms []*Room
	if err := json.NewDecoder(resp.Body).Decode(&rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

// NewClient returns a new Client using key.
func NewClient(key string) *Client {
	return &Client{
		key:    key,
		client: &http.Client{},
	}
}
