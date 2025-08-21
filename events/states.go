package events

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

func (b *VKAIUserBot) IsTyping(id int) (bool, error) {
	typing, err := b.Rdb.Get(
		context.Background(),
		fmt.Sprintf("typing:%d", id),
	).Result()

	if err != nil {
		return false, err
	}

	if typing == "true" {
		return true, nil
	}

	if typing == "false" {
		return false, nil
	}

	err = errors.New(
		fmt.Sprintf("Invalid format for data should be (true/false) got %s", typing),
	)

	slog.Error(err.Error())

	return false, err
}

func (b *VKAIUserBot) TypingSet(id int, state bool) error {
	_, err := b.Rdb.Set(
		context.Background(),
		fmt.Sprintf("typing:%d", id),
		fmt.Sprintf("%t", state),
		5*time.Minute,
	).Result()

	if err != nil {
		slog.Error(err.Error())
	}

	return err
}

func (b *VKAIUserBot) ProgressGet(id int) (int, error) {
	progressStr, err := b.Rdb.Get(
		context.Background(),
		fmt.Sprintf("progress:%d", id),
	).Result()

	if err != nil {
		return 0, err
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

func (b *VKAIUserBot) ProgressSet(id int, progress int) error {
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
