package controllers

import (

    "context"
    "SoundGadget-API/configs"
    "SoundGadget-API/models"
    "SoundGadget-API/responses"
    "net/http"
    "time"
  
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

)



var ApiCollection *mongo.Collection = configs.GetCollection(configs.DB, "Mixers")

var validate = validator.New()


func CreateMixer(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var mixer models.Mixers

	defer cancel()

	if err := c.BodyParser(&mixer); err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}


	if validErr := validate.Struct(&mixer); validErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ApiResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validErr.Error()}})
	}

	newMixer := models.Mixers{
		Id : primitive.NewObjectID(),
		Brand: "Allen&Heath",
		Name: "QU-32",
		Price: 2350,
		MixerDetails: "This is a user friendly console. it has 32 input and 16 outputs",
	}

	result, errInsert := ApiCollection.InsertOne(ctx, newMixer)

	if errInsert != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message:"error", Data: &fiber.Map{"data": errInsert.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ApiResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}

// get user

func GetMixer(c *fiber.Ctx) error {


    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    Id := c.Params("Id")

    var themixer models.Mixers

    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(Id)

    errFind := ApiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&themixer)
    if errFind != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errFind.Error()}})
    }

    return c.Status(http.StatusOK).JSON(responses.ApiResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": themixer}})
}