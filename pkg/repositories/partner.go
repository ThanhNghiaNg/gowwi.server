package repositories

import (
	"context"
	databases "owwi/pkg/database"
	"owwi/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func convertToPartner(partnerRepo models.PartnerRepository) *models.Partner {
	return &models.Partner{
		ID:          partnerRepo.ID.Hex(),
		User:        partnerRepo.User.Hex(),
		Name:        partnerRepo.Name,
		Description: partnerRepo.Description,
		Type:        partnerRepo.Type.Hex(),
		UsedTime:    partnerRepo.UsedTime,
		CreatedAt:   partnerRepo.CreatedAt,
		UpdatedAt:   partnerRepo.UpdatedAt,
	}
}

func createPartner(partner models.CreatePartner) error {
	// Convert the partner to a CreatePartnerRepository type
	userId, errP1 := bson.ObjectIDFromHex(partner.User)
	if errP1 != nil {
		return errP1
	}

	typeId, errP2 := bson.ObjectIDFromHex(partner.Type)
	if errP2 != nil {
		return errP2
	}

	partnerRepo := models.CreatePartnerRepository{
		User:        userId,
		Name:        partner.Name,
		Description: partner.Description,
		Type:        typeId,
		UsedTime:    partner.UsedTime,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	_, err := databases.Client.Collection("partners").InsertOne(context.TODO(), partnerRepo)
	return err
}

func updatePartner(partner models.UpdatePartner) error {
	partnerId, errP0 := bson.ObjectIDFromHex(partner.ID)
	if errP0 != nil {
		return errP0
	}

	userId, errP1 := bson.ObjectIDFromHex(partner.User)
	if errP1 != nil {
		return errP1
	}

	typeId, errP2 := bson.ObjectIDFromHex(partner.Type)
	if errP2 != nil {
		return errP2
	}

	partnerRepo := models.CreatePartnerRepository{
		User:        userId,
		Name:        partner.Name,
		Description: partner.Description,
		Type:        typeId,
		UsedTime:    partner.UsedTime,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	_, err := databases.Client.Collection("partners").UpdateByID(
		context.TODO(),
		partnerId,
		map[string]interface{}{
			"$set": partnerRepo,
		},
	)
	return err
}

func getPartnerByID(id string) (*models.Partner, error) {
	var partnerRepo models.PartnerRepository
	partnerId, errP := bson.ObjectIDFromHex(id)
	if errP != nil {
		return nil, errP
	}

	err := databases.Client.Collection("partners").FindOne(context.TODO(), bson.M{"_id": partnerId}).Decode(&partnerRepo)
	if err != nil {
		return nil, err
	}

	return convertToPartner(partnerRepo), nil
}

func getAllPartnersByUser(userID string) ([]models.Partner, error) {
	var partners []models.Partner
	objectId, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := databases.Client.Collection("partners").Find(context.TODO(), bson.M{"user": objectId})
	if err != nil {
		return nil, err
	}

	var results []models.PartnerRepository
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	for _, partnerRepo := range results {
		partners = append(partners, *convertToPartner(partnerRepo))
	}

	return partners, nil
}

func deletePartner(id string) error {
	objectId, errP := bson.ObjectIDFromHex(id)
	if errP != nil {
		return errP
	}

	_, err := databases.Client.Collection("partners").DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}

var PartnerRepository = struct {
	CreatePartner         func(models.CreatePartner) error
	UpdatePartner         func(models.UpdatePartner) error
	GetPartnerByID        func(string) (*models.Partner, error)
	GetAllPartnersByUser func(string) ([]models.Partner, error)
	DeletePartner         func(string) error
}{
	CreatePartner:         createPartner,
	UpdatePartner:         updatePartner,
	GetPartnerByID:        getPartnerByID,
	GetAllPartnersByUser: getAllPartnersByUser,
	DeletePartner:         deletePartner,
}
