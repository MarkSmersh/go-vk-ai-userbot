package events

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"time"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) SendFriendRequests() {
	if len(b.TargetGroups) <= 0 {
		return
	}

	members := []int{}

	for _, g := range b.TargetGroups {
		for i := 0; ; i++ {
			ms := b.Vk.GroupsGetMembers(methods.GroupsGetMembers{
				GroupID: strconv.Itoa(g),
				Sort:    "id_desc",
				Offset:  i * 1000,
			}).Items

			for _, m := range ms {
				members = append(members, m.ID)
			}

			if len(ms) < 1000 {
				break
			}

			time.Sleep(time.Second * 5)
		}
	}

	slices.SortFunc(members, func(a, b int) int {
		return b - a
	})

	ms := []int{}

	// id 524147853 is used as a head id, and id 632047853 as tail id
	// to create a range of people who registered in VK from 2019 to 2021

	for _, m := range members {
		if m > 524147853 && m < 632047853 {
			ms = append(ms, m)
		}
	}

	members = ms

	slog.Info(
		fmt.Sprintf("Members gathered from TargetGroups (%d). Start sending requests", len(members)),
	)

	for _, m := range members {
		if !slices.Contains(b.FriendsAdded, m) {
			b.Vk.FriendsAdd(methods.FriendsAdd{
				UserID: m,
			})

			b.FriendsAdded = append(b.FriendsAdded, m)

			// this is one and only line of code, that keeps VKAPI from blocking account
			// a truly divine technique
			time.Sleep(time.Minute)
		}
	}

	slog.Info("Friend requests are sent for each member of groups represented in TargetGroups. If you want to proceed, add more groups")
}
