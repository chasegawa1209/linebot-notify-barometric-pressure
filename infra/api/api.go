package api

import (
	"encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/model"
)

// APIInterface APIのインターフェース
type APIInterface interface {
    GetBarometricPressureByZutool() (*model.BarometricPressuresByZutool, error)
}

// API APIの実装
type API struct {
    placeID string
}

// NewAPI コンストラクタ
func NewAPI(placeID string) *API {
    return &API{
        placeID: placeID,
    }
}

//
func (a *API) GetBarometricPressureByZutool() (*model.BarometricPressuresByZutool, error) {
    resp, err := http.Get("https://zutool.jp/api/getweatherstatus/" + a.placeID)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result *model.BarometricPressuresByZutool
    if err = json.Unmarshal(body, &result); err != nil {
        return nil, err
    }

    return result, nil
}
