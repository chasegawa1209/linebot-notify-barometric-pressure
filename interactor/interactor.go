package interactor

import (
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/api"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/repository"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/linestore"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/usecase"
	"go.uber.org/zap"
)

// Interactor インタラクタ
type Interactor struct {
    logger          *zap.Logger
    lineAccessToken string
    lineSecret      string
    lineRoomID      string
    placeID         string
}

// NewInteractor コンストラクタ
func NewInteractor(
    logger *zap.Logger,
    lineAccessToken string,
    lineSecret string,
    lineRoomID string,
    placeID    string,
) *Interactor {
    return &Interactor{
        logger:          logger,
        lineAccessToken: lineAccessToken,
        lineSecret:      lineSecret,
        lineRoomID:      lineRoomID,
        placeID:         placeID,
    }
}

// NewUsecase ユースケース
func (i *Interactor) NewUsecase() usecase.UsecaseInterface {
    return usecase.NewUsecase(
        i.NewRepository(),
    )
}

// NewRepository リポジトリ
func (i *Interactor) NewRepository() repository.RepositoryInterface {
    return repository.NewRepository(
       i.logger,
       i.NewLineStore(),
       i.NewAPI(),
    )
}

// NewLineStore LINE接続
func (i *Interactor) NewLineStore() linestore.LineStoreInterface {
    linestore, err := linestore.NewLineStore(
        i.logger,
        i.lineAccessToken,
        i.lineSecret,
        i.lineRoomID,
    )
    if err != nil {
        panic(err)
    }
    return linestore
}

// NewAPI API接続
func (i *Interactor) NewAPI() api.APIInterface {
    return api.NewAPI(
        i.placeID,
    )
}
