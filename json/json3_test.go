package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var json03Str = `


`

type ContentStatisticResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		Statistics struct {
			UpVoteCount       int `json:"up_vote_count"`
			DownVoteCount     int `json:"down_vote_count"`
			LikeCount         int `json:"like_count"`
			CommentCount      int `json:"comment_count"`
			ShareCount        int `json:"share_count"`
			PlayCount         int `json:"play_count"`
			InterestPlayCount int `json:"interest_play_count"`
			FavoriteCount     int `json:"favorite_count"`
			CardShowCount     int `json:"card_show_count"`
			PvCount           int `json:"pv_count"`
			BulletCount       int `json:"bullet_count"`
			ContentLength     int `json:"content_length"`
			RepostCount       int `json:"repost_count"`
			ReactionCount     int `json:"reaction_count"`
			FollowCount       int `json:"follow_count"`
			ReadFinishedCount int `json:"read_finished_count"`
		} `json:"statistics"`
	} `json:"data"`
	Version string `json:"version"`
}

func Test0301(t *testing.T) {
	res := new(ContentStatisticResponse)
	json.Unmarshal([]byte(json03Str), &res)
	fmt.Printf("%+v\n", res)

}
