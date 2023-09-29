package controller

import (
	"context"
	"gowithcurd/firebase"
	"gowithcurd/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItem(c *gin.Context) {
	ctx := context.Background()
	ItemCollection := firebase.FirestoreClient.Collection("items")
	querySnapshot, err := ItemCollection.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Failed to get Item from Firestore %v", err)
		c.JSON(http.StatusNotFound, gin.H{"Error": "nothing in database"})
		return
	}
	var items []model.Person
	for _, doc := range querySnapshot {
		var item model.Person
		if err := doc.DataTo(&item); err != nil {
			log.Printf("Failed to parse item data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

// func ItemAll(c *gin.Context) {

// 	ctx := context.Background()
// 	itemstore := firebase.FirestoreClient.Collection("items")
// 	query, err := itemstore.Documents(ctx).GetAll()
// 	if err != nil {
// 		log.Printf("failed at snapshot %v", err)
// 		c.JSON(http.StatusNotFound, gin.H{"Error": "while storeing data in query"})
// 		return
// 	}

// 	var items []model.Person
// 	for _, doc := range query {
// 		var item model.Person
// 		if err := doc.DataTo(&item); err != nil {
// 			log.Printf("failed at parsing item from query %v", err)
// 			c.JSON(http.StatusNotFound, gin.H{"Error": "while parsing the data"})
// 		}
// 		items = append(items, item)
// 	}
// 	c.JSON(http.StatusOK, items)
// }

func GetItem(c *gin.Context) {
	itemId := c.Param("id")
	docref := firebase.FirestoreClient.Collection("items").Doc(itemId)
	doc, err := docref.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get item from Firestore: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	var item model.Person
	if err := doc.DataTo(&item); err != nil {
		log.Printf("Failed to parse item data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func getbole(c *gin.Context) {
	itemid := c.Param("id")

	docref := firebase.FirestoreClient.Collection("item").Doc(itemid)
	doc, err := docref.Get(context.Background())
	if err != nil {
		log.Printf("while getting the data from firestore %v", err)
		c.JSON(http.StatusNotFound, gin.H{"Error": "white getting the document"})
		return
	}
	var item model.Person
	if err := doc.DataTo(&item); err != nil {
		log.Printf("while parsing the data into struct %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "parsing the data"})
		return
	}
	c.JSON(http.StatusOK, item)

}

func CreateItem(c *gin.Context) {
	var newPerson model.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	_, err := firebase.FirestoreClient.Collection("items").Doc(newPerson.Id).Set(context.Background(), newPerson)
	if err != nil {
		log.Fatalf("Failed to create item in Firestore: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error is got"})
		return
	}
	c.JSON(http.StatusCreated, newPerson)
}

func create(c *gin.Context) {
	var newperson model.Person

	if err := c.ShouldBindJSON(&newperson); err != nil {

		log.Printf("error while parsing data %v", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return

	}

	_, err := firebase.FirestoreClient.Collection("items").Doc(newperson.Id).Set(context.Background(), newperson)
	if err != nil {
		log.Printf("eroor while seting data %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, newperson)
}

func UpdateItem(c *gin.Context) {
	itemid := c.Param("id")
	var updatePerson model.Person
	if err := c.ShouldBindJSON(&updatePerson); err != nil {
		log.Printf("Error white parsing data %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	updatePerson.Id = itemid
	docRef := firebase.FirestoreClient.Collection("items").Doc(itemid)
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get item from Firestore: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	if !docSnapshot.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	_, err = docRef.Set(context.Background(), updatePerson)
	if err != nil {
		// Handle errors related to updating the Firestore document
		log.Fatalf("Failed to update item in Firestore: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, updatePerson)
}
func DeleteItem(c *gin.Context) {
	itemId := c.Param("id")
	docRef := firebase.FirestoreClient.Collection("items").Doc(itemId)

	snapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get item from Firestore: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in while getting"})
		return
	}

	if !snapshot.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in database"})
		return
	}

	_, err = docRef.Delete(context.Background())
	if err != nil {
		log.Printf("Failed to delete item from Firestore: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
