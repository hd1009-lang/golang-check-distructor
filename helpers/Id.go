package helpers

import (
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

func GenerateId() int64 {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0
	}
	return node.Generate().Int64()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Santize(password)), 4)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Santize(password)))
	return err == nil
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
