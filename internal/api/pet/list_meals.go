package pet

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"xgmdr.com/pad/internal/api"
	"xgmdr.com/pad/internal/logger"
	"xgmdr.com/pad/internal/storage"
	pb "xgmdr.com/pad/proto"
)

// GetMeals returns a list of meals that the user has access to.
func (m *Api) GetMeals(request *pb.ListMealsRequest, petsvc pb.PetService_GetMealsServer) error {
	logger.Get().Debug("List meals request received")

	authentication := api.ExtractAuthentication(petsvc.Context())

	if request.GetPetId() == "" {
		return status.Errorf(codes.InvalidArgument, "Pet ID is required")
	}

	validPetId, err := uuid.Parse(request.PetId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid Pet ID format")
	}

	// TODO: check access rights

	logger.Get().Debug(
		"Listing meals for:",
		zap.String("owner.email", authentication.Email),
		zap.String("petId", request.PetId),
	)

	var result = storage.FindMealsByPetID(validPetId)

	logger.Get().Debug("Meals count matching pet ", zap.Int("count", len(result)))

	//meals, err := storage.ListMeals(authentication)
	//if err != nil {
	//	return nil, err
	//}
	//
	//response := &pb.ListMealsResponse{
	//	Meals: make([]*pb.Meal, 0),
	//}
	//
	//for _, meal := range meals {
	//	response.Meals = append(response.Meals, &pb.Meal{
	//		Id:      meal.ID.String(),
	//		Name:    meal.Name,
	//		Comment: meal.Comment,
	//	})
	//}

	return nil
}
