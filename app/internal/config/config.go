package config

import (
	"os"

	"github.com/HellsKitchen99/kr/app/internal/domain"
)

func LoadInfoPaths() (InfoPathConfig, error) {
	cpuInfoPath := os.Getenv("CPU_INFO_PATH")
	if cpuInfoPath == "" {
		return InfoPathConfig{}, domain.ErrEmptyCpuInfoPath
	}
	memInfoPath := os.Getenv("MEM_INFO_PATH")
	if memInfoPath == "" {
		return InfoPathConfig{}, domain.ErrEmptyMemInfoPath
	}

	memInfoLimitPath := os.Getenv("MEM_INFO_LIMIT_PATH")
	if memInfoLimitPath == "" {
		return InfoPathConfig{}, domain.ErrEmptyMemInfoLimitPath
	}

	infoPathConfig := InfoPathConfig{
		CpuInfoPath:      cpuInfoPath,
		MemInfoPath:      memInfoPath,
		MemInfoLimitPath: memInfoLimitPath,
	}

	return infoPathConfig, nil
}
