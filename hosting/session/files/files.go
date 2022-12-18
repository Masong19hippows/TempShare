package files

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	rand1 "math/rand"
	"os"
	"path/filepath"
	"time"
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
	root    string
	Files   []file
	Folders []folder
}

func (s *Storage) Init(user string) {
	rand1.Seed(time.Now().UnixNano())
	s.root = encryptAES(user)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	s.Create(s.root, &folder{RelPath: "", absPath: filepath.Join(filepath.Dir(ex), "session", "files", "vault")}, nil, true)

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
	for i := 0; i < len(s.Files); i++ {
		s.Files[i].Delete()
	}
	// Deletes the root folder last
	s.Folders[0].Delete()
	return nil
}

// Encrypt given string in AES with random key. We don't want the user account stored on the Server, thats dumb
func encryptAES(plaintext string) string {
	// create cipher
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var key string
	for i := 0; i < 32; i++ {
		key += string(charset[rand1.Intn(len(charset))])
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(plaintext), nil))
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
	e := os.Remove(f.absPath)
	if e != nil {
		panic(e)
	}

}

// Total Deletion of Folder
func (f *folder) Delete() {
	err := os.RemoveAll(f.absPath)
	if err != nil {
		panic(err)
	}
}
