package storage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// File type
type file struct {
	RelPath string
	absPath string
	Name    string
}

// folder type
type folder struct {
	RelPath string
	absPath string
	Name    string
	Files   []file
}

// How packages interact with Storage. The root is the root folder and this contains all folders and files
type Storage struct {
	ID      string
	Folders []folder
}

func (s *Storage) Init(id string) {
	s.ID = id

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	s.Create(id, &folder{RelPath: "", absPath: filepath.Join(filepath.Dir(ex), "session", "storage", "vault")}, nil, true)

}

// Searches folder for user and returns the folder or file object
func (f *Storage) Search(user string, isFolder bool) (*folder, *file) {
	if isFolder {
		var test *folder
		return test, nil
	} else {
		var test *file
		return nil, test
	}

}

// Print the hiearchy of Users Storage.
func (s *Storage) Print() {
	fmt.Println("TODO")
	// for _,i := range s.Folders {
	// 	fmt.Printf("-%v\n", i.RelPath)
	// 	for _,k := range i.
	// }
}

// Deltes Session folders and data, permantley
func (s *Storage) Delete() error {
	// Loops through folders and deletes files and folders recurslvley and permanatley
	for i := 1; i < len(s.Folders); i++ {
		for j := 0; j < len(s.Folders[i].Files); j++ {
			s.Folders[i].Files[j].Delete()
		}
		s.Folders[i].Delete()
	}
	// Loops through root folder's files
	for i := 0; i < len(s.Folders[0].Files); i++ {
		s.Folders[0].Files[i].Delete()
	}
	// Deletes the root folder last
	s.Folders[0].Delete()
	return nil
}

// Create File or Folder given the name, parent directory, data (if file) and if its a folder
func (f *Storage) Create(name string, fol *folder, fileData []byte, isFolder bool) {
	if isFolder {
		if err := os.Mkdir(filepath.Join(fol.absPath, name), os.ModePerm); err != nil {
			panic(err)
		}
		f.Folders = append(f.Folders, folder{RelPath: filepath.Join(fol.RelPath, name), absPath: filepath.Join(fol.absPath, name), Name: name})
	} else {
		newFile, err := os.Create(filepath.Join(fol.absPath, name))
		if err != nil {
			panic(err)
		}
		newFile.Write(fileData)
		newFile.Close()
		fol.Files = append(fol.Files, file{RelPath: filepath.Join(fol.RelPath, name), absPath: filepath.Join(fol.absPath, name), Name: name})
	}

}

// Total Deletion of File
func (f *file) Delete() {
	cmd := exec.Command("shred", "-zu", "-n 5", f.absPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd)
		panic(err)
	}
}

// Total Deletion of Folder
func (f *folder) Delete() {

	err := os.RemoveAll(f.absPath)
	if err != nil {
		panic(err)
	}
}
