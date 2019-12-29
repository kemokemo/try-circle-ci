package main

import "fmt"

func greeter(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is empty")
	}
	return fmt.Sprintf("Hello %s!", name), nil
}
