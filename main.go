package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// }
func main() {

	// setEnv()
	sa := option.WithCredentialsFile("./ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	dsnap := client.Collection("user").Where("role", "==", "user").Documents(context.Background())
	for {
		doc, err := dsnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		data := doc.Data()
		nomorhp := data["nomorhp"]

		fmt.Println(data["nomorhp"])
		fmt.Println(reflect.TypeOf(nomorhp))
		sendMessage("S", nomorhp.(string))
	}
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()
	fmt.Println(os.Getenv("TWILIO_ACCOUNT_SID"))

}

func sendMessage(message, phone string) {
	godotenv.Load("go.env")
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		})
	a := &openapi.CreateMessageParams{}
	a.SetTo("whatsapp:" + phone)
	a.SetFrom("whatsapp:+14155238886")
	a.SetBody(message)

	_, err := client.Api.CreateMessage(a)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}
