package infraline

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DolkMd/LineBotSample/domain/errs"
	"github.com/DolkMd/LineBotSample/domain/interfaces"
	"github.com/line/line-bot-sdk-go/linebot"
)

type lineBot struct {
	client  *linebot.Client
	groupId string
}

func NewStreamer(channelToken, channelSecret, groupId string) (interfaces.Streamer, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, errs.StreamerError{OriginErr: fmt.Errorf("fail new line bot instance: %w", err)}
	}

	return &lineBot{client: bot, groupId: groupId}, nil

}

func (l *lineBot) SendText(ctx context.Context, text string) error {
	message := linebot.NewTextMessage(text)
	if _, err := l.client.PushMessage(l.groupId, message).Do(); err != nil {
		return errs.StreamerError{OriginErr: err, ErrType: errs.StreamerErrorExternal}
	}

	return nil
}

func (l *lineBot) HandleCallback(ctx context.Context, request *http.Request, cmdCallback func(string) (string, error)) error {
	events, err := l.client.ParseRequest(request)
	if err != nil {
		return errs.StreamerError{OriginErr: err, ErrType: errs.StreamerErrorRequest}
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "groupID", "groupId", "group", "gpId":
					if _, err = l.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.Source.GroupID)).Do(); err != nil {
						return errs.StreamerError{OriginErr: err, ErrType: errs.StreamerErrorExternal}
					}
				case "test":
					if _, err = l.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(time.Now().String())).Do(); err != nil {
						return errs.StreamerError{OriginErr: err, ErrType: errs.StreamerErrorExternal}
					}
				default:
					value, err := cmdCallback(message.Text)
					if err != nil {
						if _, terr := l.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("失敗するよね")).Do(); terr != nil {
							return errs.StreamerError{OriginErr: errors.New(err.Error() + " / " + terr.Error()), ErrType: errs.StreamerErrorExternal}
						}
						return err
					}

					if _, err = l.client.PushMessage(l.groupId, linebot.NewTextMessage(value)).Do(); err != nil {
						return errs.StreamerError{OriginErr: err, ErrType: errs.StreamerErrorExternal}
					}
				}
			default:
			}
		}
	}

	return nil
}
