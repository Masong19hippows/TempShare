package main

import (
	"time"

	ses "github.com/masong19hippows/TempShare/hosting/session"
)

func main() {
	test := ses.CreateSession("garten323@hotmail.com", 8*time.Second)
	test.Storage.Create("test file", &test.Storage.Folders[0], []byte("test string\n"), false)
	test.Storage.Create("New Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("Another Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("testing", &test.Storage.Folders[2], []byte("test string\n"), false)
	test.Share([]string{"tylonggarten@gmail.com", "masongart@hotmail.com"}, nil)
	time.Sleep(10 * time.Second)
	test.Storage.Delete()
	// test.Share([]string{"masongarten@gmail.com", "tylonggarten@hotmail.com"}, nil)
}
