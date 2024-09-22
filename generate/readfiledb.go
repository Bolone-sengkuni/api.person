package generate

import (
    "os"
    "path/filepath"
)

func findDBFiles(rootDir string, dbname string) (string, error) {
    var dbFiles string
    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            if filepath.Base(path) == dbname {
                dbFiles = path
            }
        }
        return nil
    })

    
    if err != nil {
        return "", err
    }
    return dbFiles, nil
}


func ReadFileDb(dbname string) string {
	currentDir, _ := os.Getwd()
    dbFiles, _ := findDBFiles(currentDir, dbname)
    return dbFiles
}