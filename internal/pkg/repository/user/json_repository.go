package user

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"refactoring/internal/pkg/entity/user"
	"strconv"
)

type JsonRepository struct {
	Storage  store
	JsonLink string
}

type store struct {
	Increment int                  `json:"increment"`
	List      map[string]user.User `json:"list"`
}

var (
	NotFound = errors.New("user_not_found")
)

func MakeJsonRepository(StorageLink string) (*JsonRepository, error) {
	file, err := ioutil.ReadFile(StorageLink)
	if err != nil {
		return nil, err
	}
	storage := store{}

	err = json.Unmarshal(file, &storage)
	if err != nil {
		return nil, err
	}

	return &JsonRepository{
		JsonLink: StorageLink,
		Storage:  storage,
	}, nil
}

func (repo *JsonRepository) updateStorage() {
	b, _ := json.Marshal(&repo.Storage)
	_ = ioutil.WriteFile(repo.JsonLink, b, fs.ModePerm)
}

func (repo *JsonRepository) Add(u *user.User) (string, error) {
	id := strconv.Itoa(repo.Storage.Increment)
	repo.Storage.List[id] = *u

	repo.updateStorage()
	return id, nil
}

func (repo *JsonRepository) GetAll() (map[string]user.User, error) {
	return repo.Storage.List, nil
}

func (repo *JsonRepository) Find(id string) (user.User, error) {
	if _, ok := repo.Storage.List[id]; !ok {
		return user.User{}, NotFound
	}

	return repo.Storage.List[id], nil
}

func (repo *JsonRepository) Inc() error {
	repo.Storage.Increment++

	return nil
}

func (repo *JsonRepository) Update(id string, user *user.User) error {
	if _, ok := repo.Storage.List[id]; !ok {
		return NotFound
	}

	repo.Storage.List[id] = *user

	repo.updateStorage()
	return nil
}

func (repo *JsonRepository) Delete(id string) error {
	if _, ok := repo.Storage.List[id]; !ok {
		return NotFound
	}

	delete(repo.Storage.List, id)
	repo.updateStorage()

	return nil
}
