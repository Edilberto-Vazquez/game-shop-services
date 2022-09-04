package domains

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/drivers/database"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/repository"
)

// type UserDomain struct {
// 	repository.UserRepository
// 	User models.User
// }

func NewUserDomain() {
	db := database.NewMongoDB(config.MongoConfig())
	repository.SetRepository(db)

}
