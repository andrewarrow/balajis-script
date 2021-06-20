package main

import "fmt"

func HandleFollower2Followed(follower, followed []byte) {
	fmt.Println("HandleFollower2Followed", len(follower), len(followed))
}
func HandleFollowed2Follower(followed, follower []byte) {
	fmt.Println("HandleFollowed2Follower", len(followed), len(follower))
}
