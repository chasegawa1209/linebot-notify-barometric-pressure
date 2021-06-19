package usecase

import (
	"fmt"
    "time"

	"github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/model"
	"github.com/chasegawa1209/linebot-notify-barometric-pressure/domain/repository"
)

// UsecaseInterface ユースケースのインターフェース
type UsecaseInterface interface {
    Exec() *UsecaseResult
}

// Usecase ユースケースの実装
type Usecase struct {
    repository repository.RepositoryInterface
}

// NewUsecase コンストラクタ
func NewUsecase(repository repository.RepositoryInterface) *Usecase {
    return &Usecase{
        repository: repository,
    }
}

// UsecaseResult ユースケースの結果
type UsecaseResult struct {
    Err            error
    ProcessingTime float64
}

// Exec ユースケース
func (u *Usecase) Exec() *UsecaseResult {
    result := &UsecaseResult{}
    now := time.Now()
    defer func() {
        result.ProcessingTime = time.Since(now).Seconds()
    }()

    // 気圧取得
    barometricPressure, err := u.repository.GetBarometricPressure(now.Hour())
    if err != nil {
        result.Err = err
        return result
    }

    // 警戒レベルがあるかチェック
    if barometricPressure.NowLevel > model.PressureLevelIntSomewhatWarning || barometricPressure.After1HourLevel > model.PressureLevelIntSomewhatWarning || barometricPressure.After2HourLevel > model.PressureLevelIntSomewhatWarning {
        // メッセージ作成
        message := createMessage(barometricPressure)

        // LINEBotで送信
        if err := u.repository.PostMessage(message); err != nil {
            result.Err = err
            return result
        }
    }

    return result
}

func createMessage(barometricPressure *model.BarometricPressure) string {
    levelStr := convertPressureLevel(barometricPressure)
    messageFormat := "気圧が下がってきてるまる…！\n薬の準備をして、無理せず休むまるよ〜。。。\n\n"
    messageFormat = messageFormat + "現在：%s\n1時間後：%s\n2時間後：%s"

    message := fmt.Sprintf(
        messageFormat,
        levelStr[0],
        levelStr[1],
        levelStr[2],
    )
    return message
}

func convertPressureLevel(barometricPressure *model.BarometricPressure) []string {
    var levelStr []string
    levelStr = append(levelStr, model.PressureLevelMap[barometricPressure.NowLevel])
    levelStr = append(levelStr, model.PressureLevelMap[barometricPressure.After1HourLevel])
    levelStr = append(levelStr, model.PressureLevelMap[barometricPressure.After2HourLevel])
    return levelStr
}
