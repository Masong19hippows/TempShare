package main

import (
	ses "github.com/masong19hippows/tempShare/session"
)

func main() {
	test := ses.CreateSession("garten323@hotmail.com", "timer placeholder")
	test.Storage.Create("test file", &test.Storage.Folders[0], []byte("test string\n"), false)
	test.Storage.Create("New Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("Another Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("testing", &test.Storage.Folders[2], []byte("test string\n"), false)
	test.Share([]string{"tylonggarten@gmail.com", "masongart@hotmail.com"}, nil)
	// test.Share([]string{"masongarten@gmail.com", "tylonggarten@hotmail.com"}, nil)
}
