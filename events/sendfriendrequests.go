package events

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
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

			time.Sleep(time.Second * 10)
		}
	}

	ms := []int{}

	// id 524147853 is used as a head id, and id 632047853 as tail id
	// to create a range of people who registered in VK from 2019 to 2021

	for _, m := range members {
		if m > 524147853 && m < 632047853 {
			ms = append(ms, m)
		}
	}

	members = ms

	// change it to reverse slice of members
	slices.SortFunc(members, func(a, b int) int {
		return a - b
	})

	slog.Info(
		fmt.Sprintf("Members gathered from TargetGroups (%d). Start sending requests", len(members)),
	)

	// the depricated method. Changed to execute method in order to prevent a flood control error

	// for _, member := range members {
	// 	if !slices.Contains(b.friendRequests, member) && !slices.Contains(b.friends, member) {
	// 		b.Vk.FriendsAdd(methods.FriendsAdd{
	// 			UserID: member,
	// 		})
	//
	// 		b.AddFriendRequests(member)
	//
	// 		// this is one and only line of code, that keeps VKAPI from blocking an account
	// 		// a truly divine technique
	// 		// UPD: a deprecated information
	// 		time.Sleep(time.Second * time.Duration(b.Config.RequestWait))
	// 	}
	// }

	membersQueue := []int{}
	code := []string{}

	for _, member := range members {
		if !slices.Contains(b.friendRequests, member) && !slices.Contains(b.friends, member) {
			code = append(
				code,
				fmt.Sprintf("results.push(API.friends.add({\"user_id\": %d}));", member),
			)
			membersQueue = append(membersQueue, member)
		}

		if len(membersQueue) >= 5 {
			b.Vk.Execute(methods.Execute{
				Code: fmt.Sprintf("var results = [];\n%s\nreturn results;", strings.Join(code, "\n")),
			})

			b.AddFriendRequests(membersQueue...)

			membersQueue = []int{}
			code = []string{}

			time.Sleep(time.Duration(b.Config.RequestWait) * time.Second * 5)
		}
	}

	go func() {
		slog.Warn("Friend requests are sent for an each member of groups represented in targetGroups. If you want to proceed, add more groups")
		time.Sleep(5 * time.Minute)
	}()
}
