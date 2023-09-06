package dto

import (
	"github.com/qahta0/tasitracker/model"
)

func ToModel(dto *CompanyDTO) model.Company {
	var points []model.TodayPoints
	if dto.TodayPoints != nil {
		points = append(points, *dto.TodayPoints)
	}
	return model.Company{
		GUID:        dto.GUID,
		CompanyName: dto.CompanyName,
		ShortName:   dto.ShortName,
		SymbolCode:  dto.SymbolCode,
		Address:     dto.Address,
		Info:        dto.Info,
		MarketGUID:  dto.MarketGUID,
		SectorGUID:  dto.SectorGUID,
		TodayPoints: points,
		IsFavorite:  dto.IsFavorite,
	}
}

type ResponseDTO struct {
	Items []CompanyDTO `json:"items"`
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
	Pages int          `json:"pages"`
}

type CompanyDTO struct {
	GUID        string             `json:"guid"`
	CompanyName string             `json:"company_name"`
	ShortName   string             `json:"short_name"`
	SymbolCode  string             `json:"symbol_code"`
	Address     string             `json:"address"`
	Info        string             `json:"info"`
	MarketGUID  string             `json:"market_guid"`
	SectorGUID  string             `json:"sector_guid"`
	TodayPoints *model.TodayPoints `json:"today_points"`
	IsFavorite  *bool              `json:"is_favorite"`
}
