package main

import "os"

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsExist(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsExist(err) {
			return false, nil
		}
		return false, err
	}
	f.Close()

	return true, nil
}
