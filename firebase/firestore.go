package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitializeFirestore() *firestore.Client {
	projectid := "restapi-using-golang-crud"
	ctx := context.Background()
	opt := option.WithCredentialsFile("restapi-using-golang-crud-firebase-adminsdk-8t3ja-430.json")
	client, err := firestore.NewClient(ctx, projectid, opt)
	if err != nil {
		fmt.Println("failed to create client", err)
	}
	return client

}
