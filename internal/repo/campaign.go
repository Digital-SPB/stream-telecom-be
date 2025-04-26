package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/pkg/generator"
)

type CampaignRepo struct {
	Campaignes []*model.Campaign
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

	var res []*model.Campaign

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

		res = append(res, &model.Campaign{
			ID:        id,
			CreatedAt: createdAt,
			Name:      generator.GenerateCampaignName(id),
		})
	}

	return &CampaignRepo{
		Campaignes: res,
	}
}

func (r *CampaignRepo) GetByID(id int64) (*model.Campaign, error) {
	for _, campaign := range r.Campaignes {
		if campaign.ID == id {
			return campaign, nil
		}
	}
	return nil, fmt.Errorf("campaign not found")
}

func (r *CampaignRepo) GetAllCampaigns() []*model.Campaign {
	return r.Campaignes
}

func (r *CampaignRepo) GetCreationDynamic(start time.Time, end time.Time, intervalType string) ([]*model.IntervalResult, error) {
	// Инициализируем карту для подсчета
	resultMap := make(map[time.Time]int)

	// Перебираем все кампании
	for _, campaign := range r.Campaignes {
		createdAt := campaign.CreatedAt

		// Проверяем, попадает ли дата создания в заданный интервал
		if (createdAt.Equal(start) || createdAt.After(start)) && (createdAt.Equal(end) || createdAt.Before(end)) {
			var key time.Time

			// Определяем ключ для группировки в зависимости от типа интервала
			if intervalType == "day" {
				key = time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, createdAt.Location())
			} else { // month
				key = time.Date(createdAt.Year(), createdAt.Month(), 1, 0, 0, 0, 0, createdAt.Location())
			}

			resultMap[key]++
		}
	}

	// Если нет данных в интервале, возвращаем пустой результат
	if len(resultMap) == 0 {
		return []*model.IntervalResult{}, nil
	}

	// Преобразуем карту в слайс IntervalResult
	var result []*model.IntervalResult
	for date, count := range resultMap {
		result = append(result, &model.IntervalResult{
			Date:  date,
			Count: count,
		})
	}

	// Сортируем результат по дате
	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result, nil
}
