package hashes

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
)

type VerificationHash struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewHash(str string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(str))
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash, nil
}

func SaveHash(email, hash string) error {
	data := VerificationHash{
		Email: email,
		Hash:  hash,
	}

	file, err := os.OpenFile("../internal/storage/database.json", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var database []VerificationHash

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() != 0 {
		err = json.NewDecoder(file).Decode(&database)
		if err != nil {
			return err
		}
	}

	database = append(database, data)

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(database)
	return err
}

func VerifyAndDeleteHash(hash string) (bool, error) {
	var database []VerificationHash

	file, err := os.OpenFile("../internal/storage/database.json", os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}

	err = json.NewDecoder(file).Decode(&database)

	if err != nil {
		return false, err
	}

	for idx, data := range database {
		if data.Hash == hash {
			database = append(database[:idx], database[idx+1:]...)

			err = file.Truncate(0)
			if err != nil {
				return false, err
			}
			_, err = file.Seek(0, 0)
			if err != nil {
				return false, err
			}

			err = json.NewEncoder(file).Encode(&database)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}
