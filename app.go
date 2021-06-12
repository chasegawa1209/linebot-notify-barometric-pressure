package main

import (
    "os"

    "github.com/chasegawa1209/linebot-notify-barometric-pressure/infra/logging"
    "github.com/chasegawa1209/linebot-notify-barometric-pressure/interactor"
)

func main() {
    LINE_ACCESS_TOKEN := os.Getenv("LINE_ACCESS_TOKEN")
    LINE_SECRET := os.Getenv("LINE_SECRET")
    LINE_ROOM_ID := os.Getenv("LINE_ROOM_ID")
    PLACE_ID := os.Getenv("PLACE_ID")

    // logger
    isDebug := true
    logger := logging.NewZapLogger(isDebug)

    i := interactor.NewInteractor(logger, LINE_ACCESS_TOKEN, LINE_SECRET, LINE_ROOM_ID, PLACE_ID)

    result := i.NewUsecase().Exec()
    logger.Sugar().Infof("ProcessingTime: %f[s]", result.ProcessingTime)
    if result.Err != nil {
        logger.Sugar().Fatal("failed to NotifyBarometricPressureBatch: %s", result.Err.Error())
    }
    logger.Sugar().Infof("success to NotifyBarometricPressureBatch")
}
