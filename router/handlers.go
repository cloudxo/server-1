package router

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/riesinger/freecloud/auth"

	"github.com/riesinger/freecloud/config"
	"github.com/riesinger/freecloud/models"
	log "gopkg.in/clog.v1"
	macaron "gopkg.in/macaron.v1"
)

const oneGigabyte = 1024 * 1024 * 1024 * 1024

func (server) FileUpload(c *macaron.Context) {
	// Parse the multipart form in the request
	err := c.Req.ParseMultipartForm(config.GetInt64("http.upload_limit") * oneGigabyte)
	if err != nil {
		c.Error(http.StatusInternalServerError, "File upload failed:", err.Error())
		return
	}

	multiform := c.Req.MultipartForm

	// Get the *fileheaders
	files := multiform.File["files"]
	for i := range files {
		// For each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		if err != nil {
			c.Error(http.StatusInternalServerError, "Could not open file:", err.Error())
			return
		}
		defer file.Close()

		// Create the destination file making sure the path is writeable.
		dst, err := filesystem.NewFileHandle(files[i].Filename)
		if err != nil {
			c.Error(http.StatusInternalServerError, "Could not open file for writing:", err.Error())
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			c.Error(http.StatusInternalServerError, "Could not copy the file:", err.Error())
			return
		}
	}

	c.HTML(http.StatusCreated, "files/upload", "Upload successful!")
}

func (server) SignupPageHandler(c *macaron.Context) {
	c.HTML(200, "auth/signup")
}

// SignupHandler handles the /signup route.
// If GETting the URL, the signup page is shown, if POSTing to this URL, the signup
// process is initiated.
func (server) SignupHandler(c *macaron.Context) {
	if c.Req.Request.Body == nil {
		log.Warn("No user data received while signing up")
		c.WriteHeader(http.StatusBadRequest)
		return
	}
	defer c.Req.Request.Body.Close()
	// Deserialize user
	log.Trace("Deserializing user")

	var user models.User

	err := json.NewDecoder(c.Req.Request.Body).Decode(&user)
	if err != nil {
		log.Error(0, "Could not decode user data: %v", err)
		c.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Trace("Signing up user: %s %s, uid %d, email %s", user.FirstName, user.LastName, user.ID, user.Email)
	session, err := auth.NewUser(&user)
	if err != nil {
		c.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO: send the session ID and user ID as cookies
	c.Write([]byte(session))
}

// IndexHandler handles the / route, which is only GETtable.
// Note that this handler is not called if the user is not signed in. The /login handler
// will be called instaead.
func (server) IndexHandler(c *macaron.Context) {
	files, err := filesystem.ListFiles(".")
	if err != nil {
		c.Error(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(200, "index", struct {
		Files       []os.FileInfo
		CurrentUser models.User
	}{
		files,
		models.User{SignedIn: false},
	})
}

// func (server) ListUsersHandler(c *macaron.Context) {
// 	users := db.
// 	c.HTML(200, "listUsers", struct{
// 		Users []*model.User
// 	}{
// 		users,
// 	})
// }