package fs

import (
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"time"

	"github.com/freecloudio/freecloud/config"
	"github.com/freecloudio/freecloud/utils"

	"github.com/freecloudio/freecloud/models"
	"github.com/mholt/archiver"
	log "gopkg.in/clog.v1"
)

// DiskFilesystem implements the Filesystem interface, writing actual files to the disk
type DiskFilesystem struct {
	base    string
	tmpName string
	done    chan struct{}
}

// NewDiskFilesystem sets up a disk filesystem and returns it
func NewDiskFilesystem(baseDir string) (dfs *DiskFilesystem, err error) {
	base, err := filepath.Abs(baseDir)
	if err != nil {
		log.Error(0, "Could not initialize filesystem: %v", err)
		return nil, err
	}

	// Check if the base directory does not exist. If it doesn't, create it.
	baseInfo, err := os.Stat(base)
	if os.IsNotExist(err) {
		log.Info("Base directory does not exist, creating it now")
		err = os.Mkdir(base, 0755)
		if err != nil {
			log.Error(0, "Could not create base directory: %v", err)
			return
		}
	} else if !baseInfo.IsDir() {
		log.Fatal(0, "Base directory does exist but is not a directory")
		return
	} else if err != nil {
		log.Fatal(0, "Could not check if base directory exists: %v", err)
		return
	}

	log.Info("Initialized filesystem at base directory %s", base)
	dfs = &DiskFilesystem{base: base, tmpName: config.GetString("fs.tmp_folder_name"), done: make(chan struct{})}

	go dfs.cleanupTempFolderRoutine(time.Hour * time.Duration(config.GetInt("fs.tmp_data_expiry")))

	return
}

func (dfs *DiskFilesystem) Close() {
	dfs.done <- struct{}{}
}

func (dfs *DiskFilesystem) cleanupTempFolderRoutine(interval time.Duration) {
	log.Trace("Starting temp folder cleaner, running now and every %v", interval)
	dfs.cleanupTempFolder()

	ticker := time.NewTicker(interval)
	for {
		select {
		case <-dfs.done:
			return
		case <-ticker.C:
			dfs.cleanupTempFolder()
		}
	}
}

func (dfs *DiskFilesystem) cleanupTempFolder() {
	log.Trace("Cleaning temp folder")

	infoList, err := ioutil.ReadDir(dfs.base)
	if err != nil {
		log.Warn("Cleaning temp folder failed: %v", err)
	}

	for _, info := range infoList {
		if !info.IsDir() {
			continue
		}

		tmpFolderPath := filepath.Join(dfs.base, info.Name(), dfs.tmpName)
		tmpInfoList, err := ioutil.ReadDir(tmpFolderPath)
		if err != nil {
			log.Warn("Error reading temp folder in %v during temp cleanup: %v", tmpFolderPath, err)
		}

		for _, tmpInfo := range tmpInfoList {
			if time.Now().After(tmpInfo.ModTime().Add(time.Hour * time.Duration(config.GetInt("fs.tmp_data_expiry")))) {
				err = os.RemoveAll(filepath.Join(tmpFolderPath, tmpInfo.Name()))
				if err != nil {
					log.Warn("Error deleting file %v in temp folder %v during temp cleanup: %v", tmpInfo.Name(), tmpFolderPath, err)
					continue
				}
			}
		}
	}
}

// NewFileHandle opens an *os.File handle for writing to.
// Before opening the file, it check the path for sanity.
func (dfs *DiskFilesystem) NewFileHandle(path string) (*os.File, error) {
	if !utils.ValidatePath(path) {
		return nil, ErrForbiddenPathName
	}
	f, err := os.Create(filepath.Join(dfs.base, path))
	if err != nil {
		log.Error(0, "Could not create file %s: %v", path, err)
		return nil, err
	}
	return f, nil
}

// CreateDirectory creates a new directory at "path".
// Before doing so, it check the path for sanity.
func (dfs *DiskFilesystem) CreateDirectory(path string) error {
	log.Trace("Path for new directory is '%s'", path)
	if !utils.ValidatePath(path) {
		return ErrForbiddenPathName
	}
	err := os.MkdirAll(filepath.Join(dfs.base, path), 0755)
	if err != nil {
		log.Error(0, "Could not create directory %s: %v", path, err)
	}
	return err
}

// CreateDirIfNotExist checks whether directory exists and creates it otherwise
func (dfs *DiskFilesystem) CreateDirIfNotExist(path string) (created bool, err error) {
	_, fileErr := os.Stat(filepath.Join(dfs.base, path))
	if os.IsNotExist(fileErr) {
		log.Info("Directory does not exist, creating it now")
		err = dfs.CreateDirectory(path)
		return true, err
	} else if fileErr != nil {
		err = fileErr
		log.Warn("Could not check if directory exists, assuming it does: %v", err)
		return false, err
	}
	return false, nil
}

// GetOSFileInfo checks whether a path exists and returns the os file info for it
func (dfs *DiskFilesystem) GetOSFileInfo(path string) (osFileInfo os.FileInfo, err error) {
	osFileInfo, err = os.Stat(filepath.Join(dfs.base, path))
	if os.IsNotExist(err) {
		err = ErrFileNotExist
		return
	} else if err != nil {
		err = fmt.Errorf("Error resolving file path: %v", err)
		return
	}

	return
}

// GetDirectoryContent returns a list of all files and folders in the given "path" (relative to the user's directory).
// Before doing so, it checks the path for sanity.
func (dfs *DiskFilesystem) GetDirectoryContent(userPath, path string) ([]*models.FileInfo, error) {
	if !utils.ValidatePath(path) {
		return nil, ErrForbiddenPathName
	}

	info, err := ioutil.ReadDir(filepath.Join(dfs.base, userPath, path))
	if err != nil {
		log.Error(0, "Could not list files in %s: %v", path, err)
		return nil, err
	}

	if path == "" {
		path = "/"
	}
	path = utils.ConvertToSlash(path)

	fileInfos := make([]*models.FileInfo, len(info), len(info))
	for i, f := range info {
		fileInfos[i] = &models.FileInfo{
			Path:        path,
			Name:        f.Name(),
			IsDir:       f.IsDir(),
			Size:        f.Size(),
			LastChanged: f.ModTime(),
			MimeType:    mime.TypeByExtension(filepath.Ext(f.Name())),
		}
	}
	return fileInfos, nil
}

// ZipFiles zips all given absolute paths to a zip archive with the given path in the temp folder
func (dfs *DiskFilesystem) ZipFiles(paths []string, outputPath string) (zipPath string, err error) {
	for it := 0; it < len(paths); it++ {
		paths[it] = filepath.Join(dfs.base, paths[it])
	}

	fullZipPath := filepath.Join(dfs.base, outputPath)
	if err != nil {
		return
	}

	err = archiver.Zip.Make(fullZipPath, paths)

	return
}

func (dfs *DiskFilesystem) MoveFile(oldPath, newPath string) (fileInfo *models.FileInfo, err error) {
	if !utils.ValidatePath(oldPath) {
		err = ErrForbiddenPathName
		return
	}
	if !utils.ValidatePath(newPath) {
		err = ErrForbiddenPathName
		return
	}

	err = os.Rename(filepath.Join(dfs.base, oldPath), filepath.Join(dfs.base, newPath))
	if err != nil {
		err = fmt.Errorf("Moving %v to %v failed", oldPath, newPath)
		return
	}

	return
}
