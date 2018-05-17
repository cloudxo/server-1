package httpRouter

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/freecloudio/freecloud/config"
	"github.com/freecloudio/freecloud/models"
	log "gopkg.in/clog.v1"
	"gopkg.in/macaron.v1"
)

const oneGigabyte = 1024 * 1024 * 1024 * 1024

func (s ServerHandler) UploadHandler(c *macaron.Context) {
	user := c.Data["user"].(*models.User)
	path := c.Data["path"].(string)

	// Parse the multipart form in the request
	err := c.Req.ParseMultipartForm(config.GetInt64("http.upload_limit") * oneGigabyte)
	if err != nil {
		log.Error(0, "File upload failed: %v", err)
		c.Data["response"] = fmt.Errorf("file upload failed: %v", err)
		return
	}

	multiform := c.Req.MultipartForm

	// Get the *fileheaders
	files, ok := multiform.File["files"]
	if !ok {
		log.Error(0, "No 'files' form field, aborting file upload")
		c.Data["response"] = fmt.Errorf("no 'files' form field, aborting file upload")
		c.Data["responseCode"] = http.StatusBadRequest
		return
	}
	for i := range files {
		// For each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		if err != nil {
			log.Error(0, "Could not open file: %v", err)
			c.Data["response"] = fmt.Errorf("could not open file: %v", err)
			return
		}
		defer file.Close()

		// Create the destination file making sure the path is writeable.
		filePath := filepath.Join(path, files[i].Filename)
		dst, err := s.filesystem.NewFileHandleForUser(user, filePath)
		if err != nil {
			log.Error(0, "Could not open file for writing: %v", err)
			c.Data["response"] = fmt.Errorf("could not open file for writing: %v", err)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			log.Error(0, "Could not copy the file: %v", err)
			c.Data["response"] = fmt.Errorf("could not copy the file: %v", err)
			return
		}

		err = s.filesystem.FinishNewFile(user, filePath)
		if err != nil {
			log.Error(0, "Could not finish new file: %v", err)
			c.Data["response"] = fmt.Errorf("could not finish new file: %v", err)
			return
		}
	}

	c.Data["response"] = struct {
		Message string `json:"message"`
	}{
		"file uploaded",
	}
	c.Data["responseCode"] = http.StatusCreated
}

func (s ServerHandler) DownloadHandler(c *macaron.Context) {
	user := c.Data["user"].(*models.User)
	path := c.Data["path"].(string)
	fullPath, filename, err := s.filesystem.GetDownloadPath(user, path)
	if err != nil || filename == "" {
		log.Error(0, "Could not resolve filepath for download: %v", err)
		c.Data["responseCode"] = http.StatusNotFound
		c.Data["response"] = struct {
			Message string `json:"message"`
		}{
			"file not found",
		}
		return
	}
	c.ServeFile(fullPath, filename)
}

func (s ServerHandler) NotFoundHandler(c *macaron.Context) {
	c.Redirect("/#/404")
}