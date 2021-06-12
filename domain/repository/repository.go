package repository

import (
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/model"
	"go.uber.org/zap"
    "github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/linestore"
)

// RepositoryInterface リポジトリのインターフェース
type RepositoryInterface interface {
    PostMessage(message string) error
    GetBarometricPressure() (*model.BarometricPressure, error)
}

// Repository Repositoryの実装
type Repository struct {
    logger    *zap.Logger
    lineStore linestore.LineStoreInterface
}

// NewRepository コンストラクタ
func NewRepository(logger *zap.Logger, lineStore linestore.LineStoreInterface) *Repository {
    return &Repository{
        logger:    logger,
        lineStore: lineStore,
    }
}

// GetBarometricPressure 気圧情報を取得
func (r *Repository) GetBarometricPressure() (*model.BarometricPressure, error) {
    return nil, nil
}

// PostMessage LINEBotで投稿
func (r *Repository) PostMessage(message string) error {
    if err := r.lineStore.Post(message); err != nil {
        r.logger.Error(err.Error())
        return err
    }
    return nil
}
