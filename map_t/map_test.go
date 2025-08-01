package map_t

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

type A struct {
	a int
}

func TestMap(t *testing.T) {
	b := []A{{1}, A{2}, A{3}, {4}}
	bMap := make(map[int]*A)
	for k, bi := range b {
		println(&bi)
		println(&k)
		bMap[k] = &bi
	}
	fmt.Printf("%+v", bMap)

}

func TestMapForRandom(t *testing.T) {
	aMap := map[int]struct{}{}
	aMap[1] = struct{}{}
	aMap[2] = struct{}{}
	aMap[3] = struct{}{}
	aMap[4] = struct{}{}
	aMap[5] = struct{}{}
	aMap[6] = struct{}{}
	aMap[7] = struct{}{}
	for k, _ := range aMap {
		println(k)
	}

}

func TestMapAppend(t *testing.T) {
	a := make([]map[string]interface{}, 0)
	for i := 0; i < 10; i++ {
		t := map[string]interface{}{"a": i}
		a = append(a, t)
	}
	fmt.Println(a)
}

var EcsInstanseId = []string{
	"i-t4n317o78w5nmej9jvgu",
	"i-t4n317o78w5nmej9jvgt",
	"i-t4n7bqi0f5qntqwd09so",
	"i-t4n7bqi0f5qntqwd09sp",
	"i-t4nb2htfd1lvjs335e1c",
	"i-t4n0fketd6fmatdcvdrc",
	"i-t4n0pyibci9cyq59njka",
	"i-t4n4d6gh513ac69lxynj",
	"i-t4n0fketd6fmatdcvdra",
	"i-t4n57y7u94v29li2u9by",
	"i-t4n0fketd6fmatdcvdrb",
	"i-t4n1mp121g05m0ohlf0q",
}

func TestMapAppend2(t *testing.T) {
	arr := make([]map[string]string, 0)

	for _, instanceId := range EcsInstanseId {
		t := map[string]string{"instanceId": instanceId}
		arr = append(arr, t)
	}
	v, _ := json.Marshal(arr)
	fmt.Println(string(v))
}

func TestGetUnExistsElement(t *testing.T) {
	arr := make(map[string]string)

	fmt.Println(arr["a"])
}

// 测试 map[string]map[string]int 结构添加过程中是否需要创建
func TestMapAdd(t *testing.T) {
	println("23:58:59.000" > "24:58:39.000")
	for i := 0; i < 10; i++ {
		randN := rand.Intn(1000)
		println(randN)

	}
}

func TestMapAdd2(t *testing.T) {
	m := make(map[string][]int)
	for i := 0; i < 10; i++ {
		m["a"] = append(m["a"], i)
	}
	for i := 0; i < 10; i++ {
		m["b"] = append(m["b"], i)
	}
	fmt.Printf("%+v", m)

}

type HomeModuleType int

const (
	KingKong               HomeModuleType = iota // king_kong
	HotSeries                                    // hot_series
	LikeStory                                    // like_story
	ReadToday                                    // read_today
	FeaturedList                                 // featured_list
	RecommendBookList                            // recommend_book_list
	RecommendLong                                // recommend_long
	EveryoneWatchTitle                           // everyone_watch_title
	EveryoneWatch                                // everyone_watch
	ActivityBanners                              // activity_banners
	FeaturedListNew                              // featured_list_new
	ExcellentRadio                               // excellent_radio
	ListenToday                                  // listen_today
	EveryoneListenTitle                          // everyone_listen_title
	EveryoneListen                               // everyone_listen
	EveryoneListenAudio                          // everyone_listen_audio
	EveryoneListenBooklist                       // everyone_listen_booklist
	PinFeed                                      // pin_feed
	PinFeedback                                  // pin_feedback
	PinTopList                                   // pin_top_list
	PinReadToday                                 // pin_read_today
	VipPurchaseBanner                            // vip_purchase_banner
	PinFreeZone                                  // pin_free_zone
	AudioTabHackData                             // audio_tab_hack_data
	VipInstabook                                 // vip_instabook
	FeaturedListV3                               // featured_list_v3

	FollowWorkPublish    // follow_work_publish
	FollowPinPublish     // follow_pin_publish
	FollowBookListUpdate // follow_book_list_update
	FollowWorkUpdate     // follow_work_update
	FollowAuthorTop      // follow_author_top
	FollowLiveRoom       // follow_live_room

	PinLiveRoomCard //pin_live_room_card

	ListenTodayV2            // listen_today_v2
	AudioBookTopBanner       // audio_book_top_banner
	EveryoneListenAudioV2    // everyone_listen_audio_v2
	ActivityPage             // activity_page
	EveryoneListenBooklistV2 // everyone_listen_booklist_v2
	NoticeBar                //notice_bar
	VipPinLiveRoomCard       //vip_pin_live_room_card
	EditorRecommend          //editor_recommend 编辑推荐
	LatestUpdate             //latest_update 最近更新
	TopList                  //top_list 榜单
)

// HomeModuleType Key Value Map
var homeModuleTypeKeyMap = map[HomeModuleType]string{
	KingKong:                 "king_kong",
	HotSeries:                "hot_series",
	LikeStory:                "like_story",
	ReadToday:                "read_today",
	FeaturedList:             "featured_list",
	RecommendBookList:        "recommend_book_list",
	RecommendLong:            "recommend_long",
	EveryoneWatchTitle:       "everyone_watch_title",
	EveryoneWatch:            "everyone_watch",
	ActivityBanners:          "activity_banners",
	FeaturedListNew:          "featured_list_new",
	ExcellentRadio:           "excellent_radio",
	ListenToday:              "listen_today",
	EveryoneListenTitle:      "everyone_listen_title",
	EveryoneListen:           "everyone_listen",
	EveryoneListenAudio:      "everyone_listen_audio",
	EveryoneListenBooklist:   "everyone_listen_booklist",
	PinFeed:                  "pin_feed",
	PinFeedback:              "pin_feedback",
	PinTopList:               "pin_top_list",
	PinReadToday:             "pin_read_today",
	VipPurchaseBanner:        "vip_purchase_banner",
	PinFreeZone:              "pin_free_zone",
	AudioTabHackData:         "audio_tab_hack_data",
	VipInstabook:             "vip_instabook",
	FeaturedListV3:           "featured_list_v3",
	FollowWorkPublish:        "follow_work_publish",
	FollowPinPublish:         "follow_pin_publish",
	FollowBookListUpdate:     "follow_book_list_update",
	FollowWorkUpdate:         "follow_work_update",
	FollowAuthorTop:          "follow_author_top",
	FollowLiveRoom:           "follow_live_room",
	PinLiveRoomCard:          "pin_live_room_card",
	ListenTodayV2:            "listen_today_v2",
	AudioBookTopBanner:       "audio_book_top_banner",
	EveryoneListenAudioV2:    "everyone_listen_audio_v2",
	ActivityPage:             "activity_page",
	EveryoneListenBooklistV2: "everyone_listen_booklist_v2",
	NoticeBar:                "notice_bar",
	VipPinLiveRoomCard:       "vip_pin_live_room_card",
	EditorRecommend:          "editor_recommend",
	LatestUpdate:             "latest_update",
	TopList:                  "top_list",
}

func HomeModuleTypeByKey(key string) (HomeModuleType, error) {
	target := strings.ToLower(key)
	for k, v := range homeModuleTypeKeyMap {
		if v == target {
			return k, nil
		}
	}
	return 0, fmt.Errorf("invalid key: %s for HomeModuleType", key)

}

func TestM(t *testing.T) {
	fmt.Println(HomeModuleTypeByKey(" editor_recommend "))
}
