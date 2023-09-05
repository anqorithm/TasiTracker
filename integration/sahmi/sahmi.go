package sahmi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/qahta0/tasitracker/dto"
	"gorm.io/gorm"
)

func FetchStocks() ([]dto.CompanyDTO, error) {
	var companies []dto.CompanyDTO
	for i := 1; i < 4; i++ {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://service.sahmi.sa/api/v1/stocks/?order_by=-close&page=%d&size=%d", i, 100), nil)
		if err != nil {
			log.Fatal("err: ", err)
		}
		req.Header.Add("authority", "service.sahmi.sa")
		req.Header.Add("accept", "application/json, text/plain, */*")
		req.Header.Add("accept-language", "ar")
		req.Header.Add("origin", "https://sahmi.sa")
		req.Header.Add("referer", "https://sahmi.sa/")
		req.Header.Add("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
		req.Header.Add("sec-ch-ua-mobile", "?1")
		req.Header.Add("sec-ch-ua-platform", "Android")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-site", "same-site")
		req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Mobile Safari/537.36")
		req.Header.Add("x-sahm-token", "6MEAbzx9qW")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("err: ", err)
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal("err: ", err)
			return nil, err
		}
		var response *dto.ResponseDTO
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		companies = append(companies, response.Items...)
	}
	return companies, nil
}

func SaveStocks(companies *[]dto.CompanyDTO, db *gorm.DB) {
	for _, company := range *companies {
		companyModel := dto.ToModel(&company)
		result := db.First(&companyModel, "guid = ?", company.GUID)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				db.Create(&companyModel)
			} else {
				fmt.Println("err: ", result.Error)
			}
		} else {
			db.Model(&companyModel).Association("TodayPoints").Append(&company.TodayPoints)
		}
	}
}
