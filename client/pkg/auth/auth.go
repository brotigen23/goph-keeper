package auth

import "os"

func SaveToken(token, path string) error {
	file, err := os.OpenFile(path, 0, 0600)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(token))
	return err
}
