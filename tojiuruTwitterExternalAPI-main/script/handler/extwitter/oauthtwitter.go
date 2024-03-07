package extwitter

import (
	"flag"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/joho/godotenv"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
}

type TwitterOauth struct {
	Client *twitter.Client
}

func OAuthConfig() *oauth1.Config {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatalln(err)
	}

	config := &TwitterConfig{
		ConsumerKey:    os.Getenv("Twitter_Consumer"),
		ConsumerSecret: os.Getenv("Twitter_Secret"),
	}
	// allow consumer credential flags to override config fields
	consumerKey := flag.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flag.String("consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()
	if *consumerKey != "" {
		config.ConsumerKey = *consumerKey
	}
	if *consumerSecret != "" {
		config.ConsumerSecret = *consumerSecret
	}
	if config.ConsumerKey == "" {
		log.Fatal("Missing Twitter Consumer Key")
	}
	if config.ConsumerSecret == "" {
		log.Fatal("Missing Twitter Consumer Secret")
	}

	//このコンフィグを使用してloginとcallbackのハンドラーが動く
	oauth1Config := &oauth1.Config{
		ConsumerKey:    config.ConsumerKey,
		ConsumerSecret: config.ConsumerSecret,
		CallbackURL:    "http://localhost:8080/twitter/callback",
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}

	return oauth1Config

}
