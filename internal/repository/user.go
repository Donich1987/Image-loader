package repository

import (
	"Image-loader/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	fileName string
}

func NewUserRepository(fileName string) *UserRepository {
	return &UserRepository{fileName: fileName}
}

func (u *UserRepository) AddUser(user model.User) error {
	repoUser := User(user)

	file, err := os.OpenFile(u.fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln(err.Error())
		}
	}(file)

	initialBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	users := make([]User, 0)

	if len(initialBytes) != 0 {
		err = json.Unmarshal(initialBytes, &users)
		if err != nil {
			return fmt.Errorf("failed to unmarshal jile: %w", err)
		}
	}

	users = append(users, repoUser)

	b, err := json.MarshalIndent(&users, "\t", "")
	if err != nil {
		return fmt.Errorf("failed to marshal users: %w", err)
	}

	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to seek beginning of the file: %w", err)
	}

	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func (u *UserRepository) GetUser(id int) (model.User, error) {
	file, err := os.Open(u.fileName)
	if err != nil {
		return model.User{}, fmt.Errorf("couldn't open file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		return model.User{}, fmt.Errorf("failed to decode json token: %w", err)
	}

	for decoder.More() {
		var user User
		if err := decoder.Decode(&user); err != nil {
			return model.User{}, err
		}
		if user.Id == id {
			return model.User(user), nil
		}
	}

	if _, err := decoder.Token(); err != nil {
		return model.User{}, err
	}

	return model.User{}, fmt.Errorf("couldn't find user")
}
