package clever

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const debug = false

// API supports basic auth with an API key or bearer auth with a token
type Auth struct {
	APIKey, Token string
}

type Clever struct {
	Auth
	Url string
}

// Creates a new clever object to make requests with. URL must be a valid base url, e.g. "https://api.getclever.com"
func New(auth Auth, url string) *Clever {
	return &Clever{auth, url}
}

type CleverError struct {
	Code    string
	Message string `json:"error"`
}

func (err *CleverError) Error() string {
	if err.Code == "" {
		return err.Message
	}
	return fmt.Sprintf("%s (%s)", err.Error, err.Code)
}

type Paging struct {
	Count   int
	Current int
	Total   int
}

type Link struct {
	Rel string
	Uri string
}

type DistrictsResp struct {
	Districts []DistrictResp `json:"data"`
	Links     []Link
	Paging
}

type DistrictResp struct {
	District District `json:"data"`
	Links    []Link
	Uri      string
}

type District struct {
	Id   string
	Name string
}

type SchoolsResp struct {
	Links []Link
	Paging
	Schools []SchoolResp `json:"data"`
}

type SchoolResp struct {
	Links  []Link
	School School `json:"data"`
	Uri    string
}

type School struct {
	Created      string
	District     string
	HighGrade    string `json:"high_grade"`
	Id           string
	LastModified string `json:"last_modified"`
	Location
	LowGrade     string `json:"low_grade"`
	Name         string
	NcesId       string `json:"nces_id"`
	Phone        string
	SchoolNumber string `json:"school_number"`
	SisId        string `json:"sis_id"`
	StateId      string `json:"state_id"`
}

type TeachersResp struct {
	Links []Link
	Paging
	Teachers []TeacherResp `json:"data"`
}

type TeacherResp struct {
	Links   []Link
	Teacher Teacher `json:"data"`
	Uri     string
}

type Teacher struct {
	Created      string
	District     string
	Email        string
	Id           string
	LastModified string `json:"last_modified"`
	Name
	School        string
	SisId         string `json:"sis_id"`
	TeacherNumber string `json:"teacher_number"`
	Title         string
}

type StudentsResp struct {
	Links []Link
	Paging
	Students []StudentResp `json:"data"`
}

type StudentResp struct {
	Links   []Link
	Student Student `json:"data"`
	Uri     string
}

type Student struct {
	Created           string
	District          string
	Dob               string
	Email             string
	FrlStatus         string `json:"frl_status"`
	Gender            string
	Grade             string
	HispanicEthnicity string `json:"hispanic_ethnicity"`
	Id                string
	LastModified      string `json:"last_modified"`
	Location
	Name
	Race          string
	School        string
	SisId         string `json:"sis_id"`
	StateId       string `json:"state_id"`
	StudentNumber string `json:"student_number"`
}

type SectionsResp struct {
	Links []Link
	Paging
	Sections []SectionResp `json:"data"`
}

type SectionResp struct {
	Links   []Link
	Section Section `json:"data"`
	Uri     string
}

type Section struct {
	CourseName   string `json:"course_name"`
	CourseNumber string `json:"course_number"`
	Created      string
	District     string
	Grade        string
	Id           string `json:"id"`
	LastModified string `json:"last_modified"`
	Name
	School   string
	SisId    string `json:"sis_id"`
	Students []string
	Subject  string
	Teacher  string
	Term
}

type Location struct {
	Address string
	City    string
	State   string
	Zip     string
}

type Name struct {
	First  string
	Middle string
	Last   string
}

type Term struct {
	Name      string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (clever *Clever) query(path string, params map[string]string, resp interface{}) error {
	v := url.Values{}
	for key, val := range params {
		v.Set(key, val)
	}
	url := fmt.Sprintf("%s%s?%s", clever.Url, path, v.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if clever.Auth.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", clever.Auth.Token))
	} else if clever.Auth.APIKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(clever.Auth.APIKey+":"))))
	} else {
		return fmt.Errorf("Must provide either API key or bearer token")
	}
	if debug {
		log.Printf("get { %v } -> {\n", url)
	}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}
	if r.StatusCode != 200 {
		var error CleverError
		json.NewDecoder(r.Body).Decode(&error)
		return &error
	}
	err = json.NewDecoder(r.Body).Decode(resp)
	return err
}
