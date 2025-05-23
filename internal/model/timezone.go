package model

// RegionTimezones содержит смещения часовых поясов для регионов России относительно UTC
var RegionTimezones = map[int64]int{
	// UTC+2
	82: 2, // Республика Крым
	92: 2, // Севастополь

	// UTC+3 (Москва)
	77: 3,  // Москва
	78: 3,  // Санкт-Петербург
	31: 3,  // Белгородская область
	32: 3,  // Брянская область
	33: 3,  // Владимирская область
	36: 3,  // Воронежская область
	37: 3,  // Ивановская область
	40: 3,  // Калужская область
	44: 3,  // Костромская область
	46: 3,  // Курская область
	48: 3,  // Липецкая область
	50: 3,  // Московская область
	57: 3,  // Орловская область
	62: 3,  // Рязанская область
	67: 3,  // Смоленская область
	68: 3,  // Тамбовская область
	69: 3,  // Тверская область
	71: 3,  // Тульская область
	76: 3,  // Ярославская область
	
	// UTC+4
	1:  4, // Республика Адыгея
	5:  4, // Республика Дагестан
	6:  4, // Ингушетия
	7:  4, // Кабардино-Балкарская Республика
	8:  4, // Республика Калмыкия
	9:  4, // Карачаево-Черкесская Республика
	15: 4, // Северная Осетия - Алания
	20: 4, // Чеченская Республика
	23: 4, // Краснодарский край
	26: 4, // Ставропольский край
	30: 4, // Астраханская область
	34: 4, // Волгоградская область
	61: 4, // Ростовская область

	// UTC+5
	2:  5, // Республика Башкортостан
	16: 5, // Татарстан
	18: 5, // Удмуртская Республика
	43: 5, // Кировская область
	56: 5, // Оренбургская область
	58: 5, // Пензенская область
	59: 5, // Пермский край
	63: 5, // Самарская область
	64: 5, // Саратовская область
	73: 5, // Ульяновская область
	74: 5, // Челябинская область

	// UTC+6
	45: 6, // Курганская область
	55: 6, // Омская область
	72: 6, // Тюменская область
	86: 6, // Ханты-Мансийский автономный округ - Югра
	89: 6, // Ямало-Ненецкий автономный округ

	// UTC+7
	4:  7, // Республика Алтай
	17: 7, // Республика Тыва
	19: 7, // Республика Хакасия
	22: 7, // Алтайский край
	24: 7, // Красноярский край
	42: 7, // Кемеровская область
	54: 7, // Новосибирская область
	70: 7, // Томская область

	// UTC+8
	3:  8, // Республика Бурятия
	38: 8, // Иркутская область
	101: 8, // Забайкальский край

	// UTC+9
	14: 9, // Республика Саха (Якутия)
	28: 9, // Амурская область
	79: 9, // Еврейская автономная область
	27: 9, // Хабаровский край

	// UTC+10
	25: 10, // Приморский край
	65: 10, // Сахалинская область

	// UTC+11
	41: 11, // Камчатский край
	49: 11, // Магаданская область

	// UTC+12
	87: 12, // Чукотский автономный округ

	// Значение по умолчанию для неизвестных регионов
	0: 3, // UTC+3 (московское время)
}

// GetRegionOffset возвращает смещение времени в часах для заданного региона
func GetRegionOffset(regionID int64) int {
	if offset, ok := RegionTimezones[regionID]; ok {
		return offset
	}
	return 3 // По умолчанию возвращаем московское время
} 