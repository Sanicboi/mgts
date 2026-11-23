package models

import "time"

// Сань комментарии специально для тебя оставлю, чтобы ты понимал че к чему относится

// Detection — результат распознавания одного объекта на фото (придет от ML)
type Detection struct {
	ObjectType string  `json:"object_type"`
	Condition  string  `json:"condition"`
	DefectType string  `json:"defect_type"`
	Confidence float64 `json:"confidence"`
	BBox       [4]int  `json:"bbox"`
}

// PhotoResult — итог анализа одной фотографии с привязкой к координатам
type PhotoResult struct {
	FileName       string      `json:"file_name"`
	ImageURL       string      `json:"image_url"`
	Lat            float64     `json:"lat"`
	Lon            float64     `json:"lon"`
	Status         string      `json:"status"`
	Recommendation string      `json:"recommendation"`
	Detections     []Detection `json:"detections"`
}

// Batch — пакет загруженных фотографий (один выезд или одна загрузка)
type Batch struct {
	ID        string        `json:"id"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	Total     int           `json:"total"`
	Results   []PhotoResult `json:"results"`
}
