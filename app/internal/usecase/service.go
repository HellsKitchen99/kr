package usecase

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/HellsKitchen99/kr/app/internal/domain"
)

type Service struct {
	cpuInfoPath         string
	memoryInfoPath      string
	memoryLimitInfoPath string
}

func NewService(cpuInfoPath, memoryInfoPath, memoryLimitInfoPath string) *Service {
	return &Service{
		cpuInfoPath:         cpuInfoPath,
		memoryInfoPath:      memoryInfoPath,
		memoryLimitInfoPath: memoryLimitInfoPath,
	}
}

func (s *Service) GetInfo() domain.Info {
	cpuModel := s.getCpuModel()
	memTotalKb := s.getMemoryTotalKb()
	hostname := s.getHostname()
	memoryLimit := s.getMemoryLimit()
	return domain.Info{
		CpuModel:      cpuModel,
		MemoryTotalKb: memTotalKb,
		Hostname:      hostname,
		MemoryLimit:   memoryLimit,
	}
}

func (s *Service) getCpuModel() string {
	file, err := os.Open(s.cpuInfoPath)
	if err != nil {
		return "unavailable"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	modelName := "unavailable"
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "model name") {
			parts := strings.Split(text, ":")
			modelName = parts[len(parts)-1]
			return modelName
		}
	}

	if err := scanner.Err(); err != nil {
		return "unavailable"
	}

	return "unavailable"
}

func (s *Service) getMemoryTotalKb() string {
	file, err := os.Open(s.memoryInfoPath)
	if err != nil {
		return "unavailable"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memTotalKb := "unavailable"
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "MemTotal") {
			_, value, _ := strings.Cut(text, ":")
			memTotalKb = strings.Fields(value)[0]
			return memTotalKb
		}
	}

	return "unavailable"
}

func (s *Service) getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unavailable"
	}
	return hostname
}

func (s *Service) getMemoryLimit() string {
	file, err := os.Open(s.memoryLimitInfoPath)
	if err != nil {
		return "unavailable"
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "unavailable"
	}
	mem, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return "unavailable"
	}
	return strconv.Itoa(mem / 1024)
}
