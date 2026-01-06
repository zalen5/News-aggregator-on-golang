package storage

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDB_News(t *testing.T) {
	random := rand.Intn(100000)
	posts := []Post{
		{
			Title: "Test Post",
			Link:  strconv.Itoa(random),
		},
	}
	db, err := New()
	if err != nil {
		t.Fatal(err)
	}
	err = db.StoreNews(posts)
	if err != nil {
		t.Fatal(err)
	}
	news, err := db.News(2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", news)
}
