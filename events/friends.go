package events

import "slices"

func (b *VKAIUserBot) AddFriends(friends ...int) {
	for _, friend := range friends {
		if !slices.Contains(b.friends, friend) {
			b.friends = append(b.friends, friend)
		}
	}
}

func (b *VKAIUserBot) RemoveFriends(friends ...int) {
	b.friends = slices.DeleteFunc(b.friends, func(f int) bool {
		return slices.Contains(friends, f)
	})
}

func (b *VKAIUserBot) AddFriendRequests(requests ...int) {
	for _, request := range requests {
		if !slices.Contains(b.friendRequests, request) {
			b.friendRequests = append(b.friendRequests, request)
		}
	}
}

func (b *VKAIUserBot) RemoveFriendRequests(requests ...int) {
	b.friendRequests = slices.DeleteFunc(b.friendRequests, func(r int) bool {
		return slices.Contains(requests, r)
	})
}
