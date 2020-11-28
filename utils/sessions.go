package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//RandomBytes ... //
func randomBytes(length int) []byte {
    token := make([]byte, length)
    rand.Read(token)
    return token
}


func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

//GetInfo ... //
func GetInfo(USERNAME string) []byte {
    name := filepath.Join(".", "loginSession/")
    name = filepath.Join(name, USERNAME +".data")
    datar := []byte("")
    if fileExists(name) {
        data, err := ioutil.ReadFile(name)
        if err != nil {
            fmt.Printf("\n%s %s", err, "\n failed to load login session details from: " + name)
            return []byte(randomBytes(12))
        }
    datar = data;

        // fmt.Print(string(data))
    } else {
    fmt.Println("Generating new login session details.")
    return []byte(randomBytes(12))
    }
    fmt.Println("Loggin with session details from: " + name)
    return datar;

}

//SetInfo ... //
func SetInfo(USERNAME string, SESSION []byte)  {

    ////make login session directory
    dir := filepath.Join(".", "loginSession/")
    file := filepath.Join(dir, USERNAME +".data")
    if !fileExists(dir) {
        path := filepath.Join(".", "loginSession")
        os.MkdirAll(path, os.ModePerm)
    }


    if fileExists(file) {
        f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil {
            panic(err)
        }
        defer f.Close()
    
        if _, err = f.WriteString(fmt.Sprintf("%s", SESSION)); err != nil {
            panic(err)
        }
    } else {
        mydata := SESSION
        err := ioutil.WriteFile(file, mydata, 0777)
        // handle this error
        if err != nil {
            // print it out
            fmt.Println(err)
        }
    }

}