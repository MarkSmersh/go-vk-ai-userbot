package events

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
)

func (b *VKAIUserBot) IsTyping(id int) bool {
	return b.typing.Get(id)
}

func (b *VKAIUserBot) SetTyping(id int, state bool) {
	b.typing.Set(id, state)
}

func (b *VKAIUserBot) GetProgress(id int) (int, error) {
	progressStr, err := b.Rdb.Get(
		context.Background(),
		fmt.Sprintf("progress:%d", id),
	).Result()

	if err != nil {
		return -1, err
	}

	progress, err := strconv.Atoi(progressStr)

	if err != nil {
		err = errors.New(
			fmt.Sprintf("Invalid format for data should be 0...2 got %s", progressStr),
		)
		slog.Error(err.Error())
		return -1, err
	}

	return progress, err
}

func (b *VKAIUserBot) SetProgress(id int, progress int) error {
	_, err := b.Rdb.Set(
		context.Background(),
		fmt.Sprintf("progress:%d", id),
		strconv.Itoa(progress),
		0,
	).Result()

	if err != nil {
		slog.Error(err.Error())
	}

	return err
}
