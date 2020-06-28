package bilibiliapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//FollowList 返回的追番/追剧列表
type FollowList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}

//Data FollowList的子项
type Data struct {
	List  []List `json:"list"`
	Pn    int    `json:"pn"`
	Ps    int    `json:"ps"`
	Total uint16 `json:"total"`
}

//List Data的子项
type List struct {
	SeasonID       int    `json:"season_id"`
	MediaID        int    `json:"media_id"`
	SeasonType     int    `json:"season_type"`
	SeasonTypeName string `json:"season_type_name"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	TotalCount     int    `json:"total_count"`
	IsFinish       int    `json:"is_finish"`
	IsStarted      int    `json:"is_started"`
	IsPlay         int    `json:"is_play"`
	Badge          string `json:"badge"`
	BadgeType      int    `json:"badge_type"`
	SquareCover    string `json:"square_cover"`
	SeasonStatus   int    `json:"season_status"`
	SeasonTitle    string `json:"season_title"`
	BadgeEp        string `json:"badge_ep"`
	MediaAttr      int    `json:"media_attr"`
	SeasonAttr     int    `json:"season_attr"`
	Evaluate       string `json:"evaluate"`
	Subtitle       string `json:"subtitle"`
	FirstEp        int    `json:"first_ep"`
	CanWatch       int    `json:"can_watch"`
	Mode           int    `json:"mode"`
	URL            string `json:"url"`
	RenewalTime    string `json:"renewal_time,omitempty"`
	FollowStatus   int    `json:"follow_status"`
	IsNew          int    `json:"is_new"`
	Progress       string `json:"progress"`
	BothFollow     bool   `json:"both_follow"`
}

//GetFollowListAll 获取某个用户追番/追剧的详细列表
func GetFollowListAll(vmid uint32, pn uint16, ps uint8, flag uint8) (*FollowList, error) {
	if flag != 1 && flag != 2 {
		return nil, errors.New("错误的查询类型!\n请输入1或2")
	}
	resp, err := http.Get("http://api.bilibili.com/x/space/bangumi/follow/list?vmid=" + string(vmid) + "&pn=" + string(pn) + "&ps=" + string(ps) + "&type=" + string(flag))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	miao := &FollowList{}
	err = json.Unmarshal(body, miao)
	if err != nil {
		return nil, err
	}
	return miao, nil
}
