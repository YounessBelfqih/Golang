package main
import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

type Tweet struct {
	ID int64 `datastore:"-"`
	Username string
	Text string
	Time time.Time
}

func getTweets(ctx context.Context) ([]*Tweet, error) {
	return getUserTweets(ctx, "")
}

func getUserTweets(ctx context.Context, username string)  ([]*Tweet, error) {
	q := datastore.NewQuery("Tweet")
	if username != "" {
		q = q.Filter("Username =", username)
	}
	q = q.Order("-Time").Limit(10)
	var tweets []*Tweet
	keys, err:= q.GetAll(ctx, &tweets)
	if err != nil {
		return nil, err
	}
	for i := 0; i<len(tweets); i++ {
		tweets[i].ID = keys[i].IntID()
	}
	return tweets, nil
}

func makeTweet(ctx context.Context, tweet *Tweet) error{
	key:= datastore.NewIncompleteKey(ctx, "Tweet", nil)
	key, err := datastore.Put(ctx, key, tweet)
	if err == nil {
		tweet.ID = key.IntID()
	}
	return err
}