package linestore

import (
	"go.uber.org/zap"
    "github.com/line/line-bot-sdk-go/linebot"
)

// LineStoreInterface LINE接続のインターフェース
type LineStoreInterface interface {
    Post(message string) error
}

// LineStore LineStoreの実装
type LineStore struct {
    logger *zap.Logger
    bot    linebot.Client
    roomID string
}

// NewLineStore コンストラクタ
func NewLineStore(logger *zap.Logger, accessToken string, secret string, roomID string) (*LineStore, error) {
    bot, err := linebot.New(secret, accessToken)
    if err != nil {
        return nil, err
    }
    return &LineStore{
        logger: logger,
        bot:    *bot,
        roomID: roomID,
    }, nil
}

func (s *LineStore) Post(message string) error {
    _, err := s.bot.PushMessage(s.roomID, linebot.NewTextMessage(message)).Do()
    if err != nil {
        s.logger.Error(err.Error())
        return err
    }
    return nil
}
