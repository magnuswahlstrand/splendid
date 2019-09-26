package client

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/denisbrodbeck/machineid"
	"github.com/google/uuid"
	"github.com/kyeett/highscore-server/model"
)

var _ HighscoreClient = &BasicClient{}

type HighscoreClient interface {
	AddSimple(score float64) error
	ListSimple() ([]*model.Score, error)
}

func New(url string, gameName string) (*BasicClient, error) {
	userID, err := machineid.ProtectedID(gameName)
	if err != nil {
		return nil, errors.New("could not create unique client ID: " + err.Error())
	}

	return &BasicClient{
		baseURL: url,
		client:  http.DefaultClient,
		game: model.Game{
			Name: gameName,
		},
		user: model.User{
			ID: userID,
		},
	}, nil
}

type BasicClient struct {
	baseURL string
	client  *http.Client
	game    model.Game
	user    model.User
}

func (c *BasicClient) AddSimple(score float64) error {
	u := c.baseURL + "/highscore/"

	m := model.Score{
		ID:    uuid.New(),
		Score: score,
		Game:  c.game,
		User:  c.user,
	}

	b, err := json.MarshalIndent(&m, "", "  ")
	if err != nil {
		return err
	}

	resp, err := c.client.Post(u, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		b, _ = ioutil.ReadAll(resp.Body)
		return fmt.Errorf("expected status code %d, got %s: %s", http.StatusCreated, resp.Status, string(b))
	}

	return nil
}

func (c *BasicClient) ListSimple() ([]*model.Score, error) {
	url := c.baseURL + "/highscore/" + url.PathEscape(c.game.Name)
	fmt.Println(url)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	var score []*model.Score
	if err := json.NewDecoder(resp.Body).Decode(&score); err != nil {
		return nil, err
	}

	return score, nil
}

func encodedUUIDFromBytes(input []byte) (uuid.UUID, error) {
	s := fmt.Sprintf("%x", md5.Sum(input))

	id := fmt.Sprintf("%s-%s-%s-%s-%s",
		s[0:8], s[8:12], s[12:16], s[16:20], s[20:])

	return uuid.Parse(id)
}
