package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type ClickRepo struct {
	Clicks []*model.Click
}

func LoadClickRepo() *ClickRepo {
	//загрузка данных
	file, err := os.Open("data/clicks.csv")
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

	var clicks []*model.Click

	for i, record := range records {
		if i == 0 {
			continue // Пропускаем заголовок
		}

		// Проверяем, что в строке достаточно полей
		if len(record) < 11 {
			fmt.Printf("Not enough fields in line %d\n", i+1)
			continue
		}

		// Парсим дату клика (предполагаем формат "2006-01-02")
		clickDate, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			fmt.Printf("Error parsing ClickDate on line %d: %v\n", i+1, err)
			continue
		}

		// Парсим время клика (предполагаем формат "15:04:05")
		clickTime, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			fmt.Printf("Error parsing ClickTime on line %d: %v\n", i+1, err)
			continue
		}

		uid := record[2]

		// Парсим числовые ID
		memberID, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing MemberID on line %d: %v\n", i+1, err)
			continue
		}

		campaignID, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing CampaignID on line %d: %v\n", i+1, err)
			continue
		}

		regionID, err := strconv.ParseInt(record[5], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing RegionID on line %d: %v\n", i+1, err)
			continue
		}

		// Создаем структуру и добавляем в слайс
		click := &model.Click{
			ID:         uid,
			ClickDate:  clickDate,
			ClickTime:  clickTime,
			MemberID:   memberID,
			CampaignID: campaignID,
			RegionID:   regionID,
			OS:         record[6],
			Browser:    record[7],
			UserAgent:  record[8],
			Language:   record[9],
			Device:     record[10],
		}

		clicks = append(clicks, click)
	}

	return &ClickRepo{
		Clicks: clicks,
	}
}

func (r *ClickRepo) GetByCampaignID(id int64) []*model.Click {
	var clicks []*model.Click
	for _, click := range r.Clicks {
		if click.CampaignID == id {
			clicks = append(clicks, click)
		}
	}
	return clicks
}

func (r *ClickRepo) GetAll() []*model.Click {
	return r.Clicks
}
