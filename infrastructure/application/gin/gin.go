package infragin

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DolkMd/LineBotSample/domain/errs"
	"github.com/DolkMd/LineBotSample/domain/interfaces"
	"github.com/DolkMd/LineBotSample/usecase"
	"github.com/fcfcqloow/go-advance/log"
	"github.com/gin-gonic/gin"
)

// type router struct {
// 	httpMethod   string
// 	relativePath string
// 	handler      gin.HandlerFunc
// }

type app struct {
	engine *gin.Engine
	ucase  usecase.UseCase
}

func NewApplication(ucase usecase.UseCase) interfaces.Application {
	handleError := func(fn func(c *gin.Context) error) func(c *gin.Context) {
		return func(c *gin.Context) {
			if err := fn(c); err != nil {
				log.Error(err)
				switch {
				case errs.IsApplicationError(err):
					ucase.NotifyError(c, "エラーがあるで")
					fallthrough
				case errs.IsStreamerError(err):
					switch errs.StreamerErrorType(err) {
					case errs.StreamerErrorRequest:
						c.JSON(http.StatusBadRequest, map[string]interface{}{
							"error message": "BadRequest",
						})
					case errs.StreamerErrorExternal:
						fallthrough
					case errs.StreamerErrorUnKnown:
						c.JSON(http.StatusInternalServerError, map[string]interface{}{
							"error message": "unkown type error",
						})
					}
				default:
					ucase.NotifyError(c, "エラーがあるで")
					c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"error message": "InternalServerError",
					})
				}
			}
		}
	}
	engine := gin.Default()
	instance := &app{engine: engine, ucase: ucase}
	engine.Use(func(c *gin.Context) {
		log.Debug(c.Request.URL)
	})
	engine.Handle(http.MethodGet, "/", handleError(instance.index))
	engine.Handle(http.MethodPost, "/weather/auto", handleError(instance.postWeatherForStreamer))
	engine.Handle(http.MethodPost, "/callback", handleError(instance.handleStreamer))

	switch log.GetLevel() {
	case log.LOG_LEVEL_DEBUG:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)

	}

	return instance
}

func (a *app) Run(ctx context.Context, port int) error {
	log.Info("PORT", port)
	if err := a.engine.Run(":" + strconv.Itoa(port)); err != nil {
		return errs.ApplicationError{OriginErr: err}
	}

	return nil
}

func (a *app) index(c *gin.Context) error {
	c.String(http.StatusOK, "OK")
	return nil
}

func (a *app) postWeatherForStreamer(c *gin.Context) error {
	if err := a.ucase.PostWeather(c); err != nil {
		return err
	}

	c.String(http.StatusOK, "OK")

	return nil
}

func (a *app) handleStreamer(c *gin.Context) error {
	if err := a.ucase.HandleStreamerCallback(c, c.Request); err != nil {
		return err
	}

	c.String(http.StatusOK, "OK")

	return nil
}
