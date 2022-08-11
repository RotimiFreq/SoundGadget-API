package controllers

import (
	"SoundGadget-API/configs"
	"SoundGadget-API/models"
	"SoundGadget-API/responses"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// function create a new collection in out
func createColl(collectionName string) *mongo.Collection {

	createdColl := configs.GetCollection(configs.DB, collectionName)

	return createdColl

}

var newColl = createColl("Sound-Gadget")

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	if validErr := validate.Struct(&product); validErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validErr.Error()}})
	}
	// data body for products

	newProduct := models.Product{
		Id:                   primitive.NewObjectID(),
		ProductType:          product.ProductType,
		Brand:                product.Brand,
		Name:                 product.Name,
		Price:                product.Price,
		MixerDetails:         product.MixerDetails,
		PickupPattern:        product.PickupPattern,
		MicType:              product.MicType,
		Wireless:             product.Wireless,
		ProcessorType:        product.ProcessorDetails,
		ProcessorDetails:     product.ProcessorDetails,
		SpeakerPeakpower:     product.SpeakerPeakpower,
		SpeakerProgrampower:  product.SpeakerProgrampower,
		SpeakerType:          product.SpeakerType,
		SpeakerAbout:         product.SpeakerAbout,
		AmplifierType:        product.AmplifierType,
		AmplifierPowerRating: product.AmplifierPowerRating,
		AmplifierAbout:       product.AmplifierAbout,
	}

	result, errInsert := newColl.InsertOne(ctx, newProduct)

	if errInsert != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errInsert.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ApiResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}

func GetProductByName(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	productName := c.Params("name")
	fmt.Println(productName)

	var retProduct models.Product

	defer cancel()

	errFind := newColl.FindOne(ctx, bson.M{"name": productName}).Decode(&retProduct)
	if errFind != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errFind.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": retProduct}})
}

func GetProductByPriceandType(c *fiber.Ctx) error {

	var retProduct []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	nameType := c.Params("*")

	fmt.Println(nameType)

	np_split := strings.Split(nameType, "/")

	np_type := np_split[0]
	np_price := np_split[1]
	queryPrice, _ := strconv.Atoi(np_price)

	cursor, errTypePrice := newColl.Find(ctx, bson.M{"producttype": np_type, "price": bson.M{"$lte": queryPrice}})

	if errTypePrice != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errTypePrice.Error()}})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tp models.Product
		if errDecode := cursor.Decode(&tp); errDecode != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errDecode.Error()}})
		}

		retProduct = append(retProduct, tp)

	}

	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": retProduct}})

}

func GetProductByTypeAndBrand(c *fiber.Ctx) error {

	var retProduct []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	nameType := c.Params("*")

	np_split := strings.Split(nameType, "/")

	np_type := np_split[0]
	np_brand := np_split[1]

	cursor, errTypeBrand := newColl.Find(ctx, bson.M{"$and": []bson.M{{"producttype": np_type}, {"brand": np_brand}}})

	if errTypeBrand != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errTypeBrand.Error()}})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tp models.Product
		if errDecode := cursor.Decode(&tp); errDecode != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errDecode.Error()}})
		}

		retProduct = append(retProduct, tp)

	}

	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": retProduct}})

}

func DeleteProductByName(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	ProductName := c.Params("prtodeleted")

	defer cancel()

	result, errDelete := newColl.DeleteOne(ctx, bson.M{"name": ProductName})

	if errDelete != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": errDelete.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": "Gadget with the name is not found!"}})

	}

	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Product as been deleted"}})
}

func EditProductByName(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	ProductName := c.Params("prtoedit")

	defer cancel()

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	if validErr := validate.Struct(&product); validErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validErr.Error()}})
	}

	update := bson.M{

		"ProductType":          product.ProductType,
		"Brand":                product.Brand,
		"Name":                 product.Name,
		"Price":                product.Price,
		"MixerDetails":         product.MixerDetails,
		"PickupPattern":        product.PickupPattern,
		"MicType":              product.MicType,
		"Wireless":             product.Wireless,
		"ProcessorType":        product.ProcessorDetails,
		"ProcessorDetails":     product.ProcessorDetails,
		"SpeakerPeakpower":     product.SpeakerPeakpower,
		"SpeakerProgrampower":  product.SpeakerProgrampower,
		"SpeakerType":          product.SpeakerType,
		"SpeakerAbout":         product.SpeakerAbout,
		"AmplifierType":        product.AmplifierType,
		"AmplifierPowerRating": product.AmplifierPowerRating,
		"AmplifierAbout":       product.AmplifierAbout,
	}

	result, errUpdate := newColl.UpdateOne(ctx, bson.M{"name": ProductName}, bson.M{"$set": update})

	if errUpdate != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": errUpdate.Error()}})
	}

	var UpdatedProduct models.Product

	if result.MatchedCount == 1 {
		errUpdatePr := newColl.FindOne(ctx, bson.M{"name": ProductName}).Decode(&UpdatedProduct)
		if errUpdatePr != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "Error", Data: &fiber.Map{"data": errUpdatePr.Error()}})
		}
	}
	return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": UpdatedProduct}})
}
