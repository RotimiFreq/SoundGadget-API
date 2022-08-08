package controllers

import (
	"SoundGadget-API/configs"
	"SoundGadget-API/models"
	"SoundGadget-API/responses"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// validator variable to validate our string

var validate = validator.New()

// variable for our data models
var product models.Product

// function create a new collection in out database
func createColl(collectionName string) *mongo.Collection {

	createdColl := configs.GetCollection(configs.DB, collectionName)

	return createdColl

}

var newColl = createColl("Mixers")

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	if validErr := validate.Struct(&product); validErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validErr.Error()}})
	}

	// data body for the products

	newMixer := models.Product{
		Id:           primitive.NewObjectID(),
		ProductType:  product.ProductType,
		Brand:        product.Brand,
		Name:         product.Name,
		Price:        product.Price,
		MixerDetails: product.MixerDetails,
	}

	newMic := models.Product{
		Id:            primitive.NewObjectID(),
		ProductType:   product.ProductType,
		Brand:         product.Brand,
		Name:          product.Name,
		Price:         product.Price,
		PickupPattern: product.PickupPattern,
		MicType:       product.MicType,
		Wireless:      product.Wireless,
	}
	newProcessor := models.Product{
		Id:               primitive.NewObjectID(),
		ProductType:      product.ProductType,
		Brand:            product.Brand,
		Name:             product.Name,
		Price:            product.Price,
		ProcessorType:    product.ProcessorDetails,
		ProcessorDetails: product.ProcessorDetails,
	}

	newSpeaker := models.Product{
		Id:                  primitive.NewObjectID(),
		ProductType:         product.ProductType,
		Brand:               product.Brand,
		Name:                product.Name,
		Price:               product.Price,
		SpeakerPeakpower:    product.SpeakerPeakpower,
		SpeakerProgrampower: product.SpeakerProgrampower,
		SpeakerType:         product.SpeakerType,
		SpeakerAbout:        product.SpeakerAbout,
	}

	newAmp := models.Product{
		Id:                   primitive.NewObjectID(),
		ProductType:          product.ProductType,
		Brand:                product.Brand,
		Name:                 product.Name,
		Price:                product.Price,
		AmplifierType:        product.AmplifierType,
		AmplifierPowerRating: product.AmplifierPowerRating,
		AmplifierAbout:       product.AmplifierAbout,
	}

	//inserting the data body into the database

	result, errInsert := newColl.InsertOne(ctx, newMixer)

	if errInsert != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errInsert.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ApiResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}

// get user

func GetMixer(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	Id := c.Params("Id")

	var themixer models.Product

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(Id)

	errFind := newColl.FindOne(ctx, bson.M{"id": objId}).Decode(&themixer)
	if errFind != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errFind.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": themixer}})
}
