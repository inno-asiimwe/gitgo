// namespace is main
package main

// standard imports 
import (
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
	"net/http"
)

// constants 
const (
	apiURL = "https://api.github.com"
	userEndpoint = "/users/"
)

// User struct uses JSON-to-GO
type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             string      `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

/// getUsers function 
func getUsers(name string) User {
	resp, err := http.Get(apiURL + userEndpoint + name)

	// Incase of errir, throw the error and quit the application
	if err != nil {
		log.Fatalf("Error retrieving data: %s \n", err)
	}

	// defer is always executed if application crashes or function finishes successfully
	// defer closing the response body
	defer resp.Body.Close()

	// read response body and handle any errors during reading 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s \n", err)
	}

	// create a user variable of type "User"
	var user User
	json.Unmarshal(body, &user)

	return user
}
