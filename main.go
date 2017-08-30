package main

import (
	"fmt"
	"github.com/alexneamtu/goftp-server"
	"os"
	"path/filepath"
)

func main() {
	cwd, _ := os.Getwd()
	factory, _ := MyDriverFactory()
	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     11990,
		PassivePorts: "20000-20100",
		TLS:      true,
		CertFile: filepath.Join(cwd, "src/ftpgo/certs/server.crt"),
		KeyFile:  filepath.Join(cwd, "src/ftpgo/certs/server.key"),
		Auth:     auth{},
	}

	session := server.NewServer(opts)
	if err := session.ListenAndServe(); err != nil {
		fmt.Printf("%v", err)

	}

}

type auth struct{}

func (a auth) CheckPasswd(user string, password string) (bool, error) {
	if user == "test" && password == "test" {
		return true, nil
	}
	return false, nil
}

func MyDriverFactory() (factory server.DriverFactory, err error) {
	factory = &FileDriverFactory{"/home/test", server.NewSimplePerm("root", "root")}
	return
}
