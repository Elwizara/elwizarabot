// main is an example web app using Login with Twitter.
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/dghubble/sessions"
	"github.com/tarekbadrshalaan/anaconda"
)

const (
	sessionName    = "example-twtter-app"
	sessionSecret  = "example cookie signing secret"
	sessionUserKey = "twitterID"
)

// sessionStore encodes and decodes session data stored in signed cookies
var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// Config configures the main ServeMux.
type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
}

// New returns a new ServeMux with app routes.
func New(config *Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomeHandler)
	mux.Handle("/profile", requireLogin(http.HandlerFunc(profileHandler)))
	mux.HandleFunc("/logout", logoutHandler)
	// 1. Register Twitter login and callback handlers
	oauth1Config := &oauth1.Config{
		ConsumerKey:    config.TwitterConsumerKey,
		ConsumerSecret: config.TwitterConsumerSecret,
		CallbackURL:    "http://localhost:8080/twitter/callback",
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
	mux.Handle("/twitter/login", twitter.LoginHandler(oauth1Config, nil))
	mux.Handle("/twitter/callback", twitter.CallbackHandler(oauth1Config, issueSession(), nil))
	return mux
}

// issueSession issues a cookie session after successful Twitter login
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		twitterUser, err := twitter.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. Implement a success handler to issue some form of session
		session := sessionStore.New(sessionName)
		session.Values[sessionUserKey] = twitterUser.ID
		session.Save(w)
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

// welcomeHandler shows a welcome message and login button.
func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	if isAuthenticated(req) {
		http.Redirect(w, req, "/profile", http.StatusFound)
		return
	}
	page, _ := ioutil.ReadFile("home.html")
	fmt.Fprintf(w, string(page))
}

// profileHandler shows protected user content.
func profileHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, `<p>You are logged in!</p><form action="/logout" method="post"><input type="submit" value="Logout"></form>`)
}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, sessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// requireLogin redirects unauthenticated users to the login route.
func requireLogin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if !isAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

// isAuthenticated returns true if the user has a signed session cookie.
func isAuthenticated(req *http.Request) bool {
	if _, err := sessionStore.Get(req, sessionName); err == nil {
		return true
	}
	return false
}

func twitterUploadMedia(api *anaconda.TwitterApi, ImagePath string) (string, error) {
	data, err := ioutil.ReadFile(ImagePath)
	if os.IsNotExist(err) {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	// when we need some shit we do shit.
	api.HttpClient.Timeout = (time.Minute * 5)

	res, err := api.UploadMedia(encoded)
	if err != nil {
		return "", err
	}
	return res.MediaIDString, nil
}

// main creates and starts a Server listening.
func main() {
	const address = "localhost:8080"
	// read credentials from environment variables if available

	//tarek
	TwitterConsumerKey := "TwitterConsumerKey"
	TwitterConsumerSecret := "TwitterConsumerSecret"
	//FkexheLB7z51
	TwitterConsumerKey = "TwitterConsumerKey"
	TwitterConsumerSecret = "TwitterConsumerSecret"

	config := &Config{
		TwitterConsumerKey:    TwitterConsumerKey,
		TwitterConsumerSecret: TwitterConsumerSecret,
	}

	//elwizara-FkexheLB7z51
	accessToken := "accessToken"
	accessSecret := "accessSecret"
	//

	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, TwitterConsumerKey, TwitterConsumerSecret)
	imageID, err := twitterUploadMedia(api, "2.png")
	if err != nil {
		panic(err)
	}

	u := url.Values{"media_ids": []string{imageID}}
	_, err = api.PostTweet("Hello WOrld", u)
	if err != nil {
		panic(err)
	}

	log.Printf("Starting Server listening on %s\n", address)
	err = http.ListenAndServe(address, New(config))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
