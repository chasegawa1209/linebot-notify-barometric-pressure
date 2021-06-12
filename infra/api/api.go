package api

import (
	"encoding/json"
    "io/ioutil"
	"strconv"
    "net/http"

    "github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/model"

    "go.uber.org/zap"
)

// APIInterface APIのインターフェース
type APIInterface interface {

}

// API APIの実装
type API struct {
    logger  *zap.Logger
    placeID int
}

// NewAPI コンストラクタ
func NewAPI(logger *zap.Logger, placeID int) *API {
    return &API{
        logger:  logger,
        placeID: placeID,
    }
}

//
func (a *API) GetBarometricPressureByZutool() (*model.BarometricPressuresByZutool, error) {
    placeIDStr := strconv.Itoa(a.placeID)
    resp, err := http.Get("https://zutool.jp/api/getweatherstatus/" + placeIDStr)
    if err != nil {
        a.logger.Error(err.Error())
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        a.logger.Error(err.Error())
        return nil, err
    }

    var result *model.BarometricPressuresByZutool
    if err = json.Unmarshal(body, &result); err != nil {
        a.logger.Error(err.Error())
        return nil, err
    }

    return result, nil
}
