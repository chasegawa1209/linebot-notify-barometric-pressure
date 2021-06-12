package repository

import (
	"strconv"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/api"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/model"
	"go.uber.org/zap"
    "github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/linestore"
)

// RepositoryInterface リポジトリのインターフェース
type RepositoryInterface interface {
    PostMessage(message string) error
    GetBarometricPressure(hour int) (*model.BarometricPressure, error)
}

// Repository Repositoryの実装
type Repository struct {
    logger    *zap.Logger
    lineStore linestore.LineStoreInterface
    api       api.APIInterface
}

// NewRepository コンストラクタ
func NewRepository(logger *zap.Logger, lineStore linestore.LineStoreInterface, api api.APIInterface) *Repository {
    return &Repository{
        logger:    logger,
        lineStore: lineStore,
        api:       api,
    }
}

// GetBarometricPressure 気圧情報を取得
func (r *Repository) GetBarometricPressure(hour int) (*model.BarometricPressure, error) {
    barometricPressure, err := r.api.GetBarometricPressureByZutool()
    if err != nil {
        r.logger.Error(err.Error())
        return nil, err
    }

    var nowLevel, After1HourLevel, After2HourLevel string
    for _, v := range barometricPressure.Today {
        intTime, err := strconv.Atoi(v.Time)
        if err != nil {
            return nil, err
        }
        switch intTime {
        case hour:
            nowLevel = v.PressureLevel
        case hour+1:
            After1HourLevel = v.PressureLevel
        case hour+2:
            After2HourLevel = v.PressureLevel
        }
    }

    result := &model.BarometricPressure{
        NowLevel:        nowLevel,
        After1HourLevel: After1HourLevel,
        After2HourLevel: After2HourLevel,
    }
    return result, nil
}

// PostMessage LINEBotで投稿
func (r *Repository) PostMessage(message string) error {
    if err := r.lineStore.Post(message); err != nil {
        r.logger.Error(err.Error())
        return err
    }
    return nil
}
