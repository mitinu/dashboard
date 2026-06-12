package application

import (
	"backend/src/internal/DTO"
	"backend/src/internal/config"
	"backend/src/internal/repository/postgres"
	"backend/src/pkg/logger"
)

func CreateSuperadmin(cfg *config.Config, db *postgres.PostgresDB) {
	accessStatusUser := postgres.NewAccessStatusUserRepository(db)
	userRepository := postgres.NewUserRepository(db, cfg)
	reqAndAutoService := NewReqAndAutoService(userRepository, cfg)
	userRepository.DeleteSuperadmin()
	accessStatusId, err := accessStatusUser.GetIdByTitle(cfg.TitleSadminInDB)
	if err != nil {
		logger.Error.Println(err)
	}
	userReq := DTO.UserReq{
		Login:          cfg.Login,
		Password:       cfg.Password,
		AccessStatusId: accessStatusId,
	}

	err = reqAndAutoService.Register(userReq)
	if err != nil {
		logger.Error.Println(err)
	}
}
