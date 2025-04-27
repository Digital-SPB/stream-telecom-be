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

func (r *RegionsRepo) GetRegionsInfo() []*model.RegionInfo {
	return []*model.RegionInfo{
		{1, "Republic of Adygea", "UTC+3"},
		{2, "Republic of Altai", "UTC+7"},
		{3, "Altai Krai", "UTC+7"},
		{4, "Amur Oblast", "UTC+9"},
		{5, "Arkhangelsk Oblast", "UTC+3"},
		{6, "Astrakhan Oblast", "UTC+4"},
		{7, "Republic of Bashkortostan", "UTC+5"},
		{8, "Belgorod Oblast", "UTC+3"},
		{9, "Bryansk Oblast", "UTC+3"},
		{10, "Republic of Buryatia", "UTC+8"},
		{11, "Chechen Republic", "UTC+3"},
		{12, "Chelyabinsk Oblast", "UTC+5"},
		{13, "Chukotka Autonomous Okrug", "UTC+12"},
		{14, "Chuvash Republic", "UTC+3"},
		{15, "Republic of Dagestan", "UTC+3"},
		{16, "Republic of Ingushetia", "UTC+3"},
		{17, "Irkutsk Oblast", "UTC+8"},
		{18, "Ivanovo Oblast", "UTC+3"},
		{19, "Jewish Autonomous Oblast", "UTC+10"},
		{20, "Kabardino-Balkar Republic", "UTC+3"},
		{21, "Kaliningrad Oblast", "UTC+2"},
		{22, "Kalmykia", "UTC+3"},
		{23, "Kaluga Oblast", "UTC+3"},
		{24, "Kamchatka Krai", "UTC+12"},
		{25, "Karachay-Cherkess Republic", "UTC+3"},
		{26, "Republic of Karelia", "UTC+3"},
		{27, "Kemerovo Oblast", "UTC+7"},
		{28, "Khabarovsk Krai", "UTC+10"},
		{29, "Republic of Khakassia", "UTC+7"},
		{30, "Khanty-Mansi Autonomous Okrug", "UTC+5"},
		{31, "Kirov Oblast", "UTC+3"},
		{32, "Komi Republic", "UTC+3"},
		{33, "Kostroma Oblast", "UTC+3"},
		{34, "Krasnodar Krai", "UTC+3"},
		{35, "Krasnoyarsk Krai", "UTC+7"},
		{36, "Kurgan Oblast", "UTC+5"},
		{37, "Kursk Oblast", "UTC+3"},
		{38, "Leningrad Oblast", "UTC+3"},
		{39, "Lipetsk Oblast", "UTC+3"},
		{40, "Magadan Oblast", "UTC+11"},
		{41, "Mari El Republic", "UTC+3"},
		{42, "Republic of Mordovia", "UTC+3"},
		{43, "Moscow", "UTC+3"},
		{44, "Moscow Oblast", "UTC+3"},
		{45, "Murmansk Oblast", "UTC+3"},
		{46, "Nenets Autonomous Okrug", "UTC+3"},
		{47, "Nizhny Novgorod Oblast", "UTC+3"},
		{48, "Novgorod Oblast", "UTC+3"},
		{49, "Novosibirsk Oblast", "UTC+7"},
		{50, "Omsk Oblast", "UTC+6"},
		{51, "Orenburg Oblast", "UTC+5"},
		{52, "Oryol Oblast", "UTC+3"},
		{53, "Penza Oblast", "UTC+3"},
		{54, "Perm Krai", "UTC+5"},
		{55, "Primorsky Krai", "UTC+10"},
		{56, "Pskov Oblast", "UTC+3"},
		{57, "Rostov Oblast", "UTC+3"},
		{58, "Ryazan Oblast", "UTC+3"},
		{59, "Saint Petersburg", "UTC+3"},
		{60, "Sakha Republic (Yakutia)", "UTC+9"},
		{61, "Sakhalin Oblast", "UTC+11"},
		{62, "Samara Oblast", "UTC+4"},
		{63, "Saratov Oblast", "UTC+4"},
		{64, "Republic of North Ossetia-Alania", "UTC+3"},
		{65, "Smolensk Oblast", "UTC+3"},
		{66, "Stavropol Krai", "UTC+3"},
		{67, "Sverdlovsk Oblast", "UTC+5"},
		{68, "Tambov Oblast", "UTC+3"},
		{69, "Republic of Tatarstan", "UTC+3"},
		{70, "Tomsk Oblast", "UTC+7"},
		{71, "Tula Oblast", "UTC+3"},
		{72, "Tver Oblast", "UTC+3"},
		{73, "Tyumen Oblast", "UTC+5"},
		{74, "Republic of Tyva", "UTC+7"},
		{75, "Udmurt Republic", "UTC+4"},
		{76, "Ulyanovsk Oblast", "UTC+4"},
		{77, "Vladimir Oblast", "UTC+3"},
		{78, "Volgograd Oblast", "UTC+3"},
		{79, "Vologda Oblast", "UTC+3"},
		{80, "Voronezh Oblast", "UTC+3"},
		{81, "Yamalo-Nenets Autonomous Okrug", "UTC+5"},
		{82, "Yaroslavl Oblast", "UTC+3"},
		{83, "Zabaykalsky Krai", "UTC+9"},
		{84, "Crimea Republic", "UTC+3"},
		{85, "Sevastopol", "UTC+3"},
	}
}
