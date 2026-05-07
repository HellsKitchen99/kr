package build

import (
	"github.com/HellsKitchen99/kr/app/internal/config"
	"github.com/HellsKitchen99/kr/app/internal/delivery/http"
	"github.com/HellsKitchen99/kr/app/internal/server"
	"github.com/HellsKitchen99/kr/app/internal/usecase"
)

func Build() error {
	infoPathConfig, err := config.LoadInfoPaths()
	if err != nil {
		return err
	}
	service := usecase.NewService(infoPathConfig.CpuInfoPath, infoPathConfig.MemInfoPath, infoPathConfig.MemInfoLimitPath)
	handler := http.NewHandler(service)
	server := server.NewServer(":1111", handler)
	if err := server.Start(); err != nil {
		return err
	}
	return nil
}
