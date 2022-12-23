package main

import (
	"time"

	ses "github.com/masong19hippows/TempShare/hosting/session"
	"github.com/masong19hippows/TempShare/hosting/web"
)

func main() {
	// portHTTP := flag.Int("portHTTP", 80, "Select the port that you wish the http server to run on")
	// portHTTPS := flag.Int("portHTTPS", 443, "Select the port that you wish the https server to run on")
	// vaultPath := flag.String("vault-path", "", "Choose the path of the vault to store sessions in")
	// flag.Parse()
	go web.Start()
	test, err := ses.Create("garten323@hotmail.com", 15*time.Second)
	if err != nil {
		panic(err)
	}
	test.Storage.Create("test file", &test.Storage.Folders[0], []byte("test string\n"), false)
	test.Storage.Create("New Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("Another Folder", &test.Storage.Folders[0], nil, true)
	test.Storage.Create("testing", &test.Storage.Folders[2], []byte("test string\n"), false)
	time.Sleep(15 * time.Minute)

	test.Share([]string{"tylonggarten@gmail.com", "masongart@hotmail.com"}, nil)
	// time.Sleep(10 * time.Second)
	// test.Storage.Delete()
	// test.Share([]string{"masongarten@gmail.com", "tylonggarten@hotmail.com"}, nil)
}
