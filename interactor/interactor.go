package interactor

import (
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
    placeID         int
}

// NewInteractor コンストラクタ
func NewInteractor(
    logger *zap.Logger,
    lineAccessToken string,
    lineSecret string,
    lineRoomID string,
) *Interactor {
    return &Interactor{
        logger:          logger,
        lineAccessToken: lineAccessToken,
        lineSecret:      lineSecret,
        lineRoomID:      lineRoomID,
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
