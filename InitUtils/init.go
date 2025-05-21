package InitUtils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
)

type DatabaseCredentials struct {
	Host                           string `json:"host"`
	Port                           int    `json:"port"`
	User                           string `json:"user"`
	Password                       string `json:"password"`
	DBName                         string `json:"dbName"`
	DontGeneratePreparedStatements bool   `json:"-"`
	SkipDefaultTransaction         bool   `json:"-"`
}

func ReadCredentialsFromFile(filePath string) (DatabaseCredentials, error) {
	var credentials DatabaseCredentials
	file, err := os.Open(filePath)
	if err != nil {
		return credentials, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&credentials)
	return credentials, err
}

func WriteCredentialsToFile(filePath string, credentials DatabaseCredentials) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: Pretty-print JSON
	return encoder.Encode(credentials)
}

type UserPassword struct {
	Email    string
	Password string
}

func WriteUsersAndPasswords(filePath string, users []UserPassword) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: Pretty-print JSON
	return encoder.Encode(users)
}

func CreateCredentials() DatabaseCredentials {
	credentials := &DatabaseCredentials{
		Host:   "localhost",
		Port:   3306,
		DBName: "biblioteca_test",
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter database host (default: localhost)> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		credentials.Host = input
	}

	fmt.Print("Enter database name (default: biblioteca_test)> ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		credentials.DBName = input
	}

	fmt.Print("Enter port (default: 3306)> ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		_, err := fmt.Sscanf(input, "%d", &credentials.Port)
		if err != nil {
			log.Fatalf("Error reading port: %s", err)
			return DatabaseCredentials{}
		}
	}

	fmt.Print("Enter database user > ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		credentials.User = input
	}

	fmt.Print("Enter database password > ")
	if term.IsTerminal(int(os.Stdin.Fd())) {
		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			log.Fatalln("Error reading password:", err)
		}
		fmt.Println() // Print a newline after password input
		credentials.Password = string(passwordBytes)
	} else {
		fmt.Print("\nWARN: Your terminal does not support secure data inputs, using legacy input instead.\n > ")
		input, _ = reader.ReadString('\n')
		credentials.Password = strings.TrimSpace(input)
	}

	return *credentials
}
