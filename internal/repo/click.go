package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
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

/*
сколько кликов по компании в день
% от клиеов по компании за все время
сколько кликов в месяц
% от кликов за все время
*/
func (r *ClickRepo) GetClickDynamic(id int64) (*model.CampaignStats, error) {
	// Фильтруем клики по campaignID
	var filteredClicks []*model.Click
	totalClicks := 0

	for _, click := range r.Clicks {
		if click.CampaignID == id {
			filteredClicks = append(filteredClicks, click)
			totalClicks++
		}
	}

	if totalClicks == 0 {
		return nil, fmt.Errorf("no clicks found for campaign ID %d", id)
	}

	// Группируем клики по дням и месяцам
	dailyCounts := make(map[time.Time]int)
	monthlyCounts := make(map[time.Time]int)

	for _, click := range filteredClicks {
		// Для дневной статистики
		day := time.Date(
			click.ClickDate.Year(),
			click.ClickDate.Month(),
			click.ClickDate.Day(),
			0, 0, 0, 0,
			click.ClickDate.Location(),
		)
		dailyCounts[day]++

		// Для месячной статистики (первый день месяца)
		month := time.Date(
			click.ClickDate.Year(),
			click.ClickDate.Month(),
			1,
			0, 0, 0, 0,
			click.ClickDate.Location(),
		)
		monthlyCounts[month]++
	}

	// Формируем дневную статистику
	var dailyStats []*model.DailyStat
	for date, count := range dailyCounts {
		percentage := float64(count) / float64(totalClicks) * 100
		dailyStats = append(dailyStats, &model.DailyStat{
			Date:        date,
			ClicksCount: count,
			Percentage:  percentage,
		})
	}

	// Сортируем по дате
	sort.Slice(dailyStats, func(i, j int) bool {
		return dailyStats[i].Date.Before(dailyStats[j].Date)
	})

	// Формируем месячную статистику
	var monthlyStats []*model.MonthlyStat
	for month, count := range monthlyCounts {
		percentage := float64(count) / float64(totalClicks) * 100
		monthlyStats = append(monthlyStats, &model.MonthlyStat{
			Month:       month,
			ClicksCount: count,
			Percentage:  percentage,
		})
	}

	// Сортируем по месяцу
	sort.Slice(monthlyStats, func(i, j int) bool {
		return monthlyStats[i].Month.Before(monthlyStats[j].Month)
	})

	return &model.CampaignStats{
		DailyStats:   dailyStats,
		MonthlyStats: monthlyStats,
		TotalClicks:  totalClicks,
	}, nil
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
