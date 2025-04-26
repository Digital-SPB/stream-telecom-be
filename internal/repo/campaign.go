package repo

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type CampaignRepo struct {
	Campaignes *[]model.Campaign
}

func LoadCampaignRepo() *CampaignRepo {
	//загрузка данных
	file, err := os.Open("data/campaign.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ',' // Указываем разделитель (по умолчанию ',')
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	var res []model.Campaign

	for i, record := range records {
		if i == 0 {
			continue // Пропускаем заголовок
		}

		// Проверяем, что в строке достаточно полей
		if len(record) < 2 {
			fmt.Printf("Not enough fields in line %d\n", i+1)
			continue
		}

		// Парсим ID
		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing ID on line %d: %v\n", i+1, err)
			continue
		}

		// Парсим дату (предполагаем формат "2006-01-02 15:04:05")
		createdAt, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			fmt.Printf("Error parsing CreatedAt on line %d: %v\n", i+1, err)
			continue
		}

		res = append(res, model.Campaign{
			ID:        id,
			CreatedAt: createdAt,
		})
	}

	

	return &CampaignRepo{
		Campaignes: &res,
	}
}

func (r *CampaignRepo) GetCampainActivity(id int, ctx context.Context) {}
