package main

import (
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
	"testing"
)

func Test_FetchUserByUsername(t *testing.T) {
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

//func BenchmarkCase1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//	}
//}
//
//func BenchmarkCase2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//	}
//}
