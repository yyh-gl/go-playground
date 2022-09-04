package github

import (
	"net/http"
	"testing"
	"time"

	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func Test_FetchUserByUsernameWithoutVRC(t *testing.T) {
	c := NewGitHubClient(http.DefaultClient)

	got, err := c.FetchUserByUsername("yyh-gl")
	t.Run("外部APIからGitHubユーザー情報を取得できる", func(t *testing.T) {
		if got.ID == 0 || got.Name == "" || err != nil {
			t.Error("ユーザー情報の取得に失敗", got, err)
		}
	})
}

func Test_FetchUserByUsernameWithVRC(t *testing.T) {
	r, err := recorder.New("fixtures/Test_FetchUserByUsername")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	c := NewGitHubClient(r.GetDefaultClient())

	got, err := c.FetchUserByUsername("yyh-gl")
	t.Run("外部APIからGitHubユーザ情報を取得できる", func(t *testing.T) {
		if got.ID == 0 || got.Name == "" || err != nil {
			t.Error("ユーザ情報の取得に失敗", got, err)
		}
	})
}

func Test_FetchUserByUsernameWithVRC2(t *testing.T) {
	r, err := recorder.New("fixtures/customization")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	c := NewGitHubClient(&http.Client{
		Timeout:   10 * time.Second, // Timeoutをカスタマイズ
		Transport: r,
	})

	got, err := c.FetchUserByUsername("yyh-gl")
	t.Run("外部APIからGitHubユーザ情報を取得できる", func(t *testing.T) {
		if got.ID == 0 || got.Name == "" || err != nil {
			t.Error("ユーザ情報の取得に失敗", got, err)
		}
	})
}
