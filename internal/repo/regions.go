package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type RegionsRepo struct {
	Regions []*model.Region
}

func LoadRegionsRepo() *RegionsRepo {
	//загрузка данных
	file, err := os.Open("data/regions.csv")
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

	var res []*model.Region

	for i, record := range records {
		if i == 0 {
			continue // Пропускаем заголовок
		}

		// Проверяем, что в строке достаточно полей
		if len(record) < 2 {
			fmt.Printf("Not enough fields in line %d\n", i+1)
			continue
		}

		// Парсим ID региона
		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing region ID on line %d: %v\n", i+1, err)
			continue
		}

		res = append(res, &model.Region{
			ID:   id,
			Name: record[1],
		})
	}

	return &RegionsRepo{
		Regions: res,
	}
}

func (r *RegionsRepo) GetAll() []*model.Region {
	return r.Regions
}
