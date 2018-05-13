package cmtt

import (
	"github.com/mitchellh/mapstructure"
	"gopkg.in/resty.v1"
	"reflect"
	"time"
)

//Cmtt is the main type that allows you to interact with Cmtt services' API (tjournal, vc, dtf)
type Cmtt struct {
	client   *resty.Client
	token    string
	platform string
}

//User is an object that contains information about user
type User struct {
	ID             int             `mapstructure:"id"`
	URL            string          `mapstructure:"url"`
	Created        time.Time       `mapstructure:"created"`
	Name           string          `mapstructure:"name"`
	AvatarURL      string          `mapstructure:"avatar_url"`
	Karma          int             `mapstructure:"karma"`
	SocialAccounts []SocialAccount `mapstructure:"social_accounts"`
	Counters       struct {
		Comments  int `mapstructure:"comments"`
		Entries   int `mapstructure:"entries"`
		Favorites int `mapstructure:"favorites"`
	} `mapstructure:"counters"`
	PushTopic      string `mapstructure:"push_topic"`
	AdvancedAccess struct {
		Hash                string `mapstructure:"hash"`
		NeedsAdvancedAccess bool   `mapstructure:"is_needs_advanced_access"`
		Subscription        struct {
			Active      bool      `mapstructure:"is_active"`
			ActiveUntil time.Time `mapstructure:"active_until"`
		} `mapstructure:"tj_subscription"`
		Actions struct {
			ReadComments  bool `mapstructure:"read_comments"`
			WriteComments bool `mapstructure:"write_comments"`
		} `mapstructure:"actions"`
	} `mapstructure:"advanced_access"`
	Cover UserCover `mapstructure:"cover"`
}

//UserCover is an object that contains information about background image used in user's profile
type UserCover struct {
	ID     string `mapstructure:"id"`
	UUID   string `mapstructure:"uuid"`
	Width  int    `mapstructure:"width"`
	Height int    `mapstructure:"height"`
	Y      string `mapstructure:"y"`
	Type   int    `mapstructure:"type"`
	URL    string `mapstructure:"cover_url"`
}

//Cover is an object that contains information about background image used in the entry
type Cover struct {
	Size         Size   `mapstructure:"size"`
	SizeSimple   string `mapstructure:"size_simple"`
	ThumbnailURL string `mapstructure:"thumbnailUrl"`
	Type         int    `mapstructure:"type"`
	URL          string `mapstructure:"url"`
}

//SocialAccount is an object that contains information about social network account attached to user
type SocialAccount struct {
	ID       string `mapstructure:"id"`
	Type     int    `mapstructure:"type"`
	Username string `mapstructure:"username"`
	URL      string `mapstructure:"url"`
}

//Author is an object that contains information about author of the entry
type Author struct {
	ID             int             `mapstructure:"id"`
	URL            string          `mapstructure:"url"`
	Created        time.Time       `mapstructure:"created"`
	Name           string          `mapstructure:"name"`
	FirstName      string          `mapstructure:"first_name"`
	LastName       string          `mapstructure:"last_name"`
	AvatarURL      string          `mapstructure:"avatar_url"`
	Karma          int             `mapstructure:"karma"`
	Gender         int             `mapstructure:"gender"`
	SocialAccounts []SocialAccount `mapstructure:"social_accounts"`
	Work           interface{}     `mapstructure:"work"`
}

//SimpleAuthor is an object that contains information about author of the comment
type SimpleAuthor struct {
	ID        int    `mapstructure:"id"`
	Name      string `mapstructure:"name"`
	AvatarURL string `mapstructure:"avatar_url"`
}

//CommentEntry is an object that contains information about entry the comment was posted for
type CommentEntry struct {
	ID     int       `mapstructure:"id"`
	Date   time.Time `mapstructure:"date"`
	Badges []Badge   `mapstructure:"badges"`
	Title  string    `mapstructure:"title"`
}

//Entry is an object that contains information about entry
type Entry struct {
	ID                       int           `mapstructure:"id"`
	Author                   SimpleAuthor  `mapstructure:"author"`
	Badges                   []Badge       `mapstructure:"badges"`
	CommentatorsAvatars      []string      `mapstructure:"commentatorsAvatars"`
	CommentsCount            int           `mapstructure:"commentsCount"`
	CommentsPreview          []Comment     `mapstructure:"commentsPreview"`
	Cover                    *Cover        `mapstructure:"cover"`
	Date                     time.Time     `mapstructure:"date"`
	EntryContent             EntryContent  `mapstructure:"entryContent"`
	EntryJSON                string        `mapstructure:"entryJSON"`
	FavoriteType             int           `mapstructure:"favoriteTy[e"`
	ForcedToMainPage         int           `mapstructure:"forcedToMainPage"`
	HitsCount                int           `mapstructure:"hitsCount"`
	Intro                    string        `mapstructure:"intro"`
	IntroInFeed              interface{}   `mapstructure:"introInFeed"`
	IsAdvertisement          bool          `mapstructure:"isAdvertisement"`
	IsBigPicture             bool          `mapstructure:"isBigPicture"`
	IsDraft                  bool          `mapstructure:"isDraft"`
	IsEnabledComments        bool          `mapstructure:"isEnabledComments"`
	IsEnabledInstantArticles bool          `mapstructure:"isEnabledInstantArticles"`
	IsEnabledLikes           bool          `mapstructure:"isEnabledLikes"`
	IsFavorited              bool          `mapstructure:"isFavorited"`
	IsNewEditor              bool          `mapstructure:"isNewEditor"`
	IsPro                    bool          `mapstructure:"isPro"`
	IsRemoved                bool          `mapstructure:"isRemoved"`
	IsWebPushSent            bool          `mapstructure:"isWebPushSent"`
	LastComments             []LastComment `mapstructure:"lastComments"`
	LastModificationDate     time.Time     `mapstructure:"lastModificationDate"`
	Likes                    Likes         `mapstructure:"likes"`
	RawIntro                 string        `mapstructure:"rawIntro"`
	RemovedReasonID          int           `mapstructure:"removedReasonId"`
	RemovedReasonText        string        `mapstructure:"removedReasonText"`
	Similar                  []Similar     `mapstructure:"similar"`
	Title                    string        `mapstructure:"title"`
	Type                     int           `mapstructure:"type"`
	URL                      string        `mapstructure:"url"`
	WebViewURL               string        `mapstructure:"webviewUrl"`
}

//EntryContent is an object that contains the entrie's body
type EntryContent struct {
	HTML    string `mapstructure:"html"`
	Version string `mapstructure:"version"`
}

type LastComment struct {
	ID          int       `mapstructure:"id"`
	Date        time.Time `mapstructure:"date"`
	Entry       string    `mapstructure:"entry"`
	EntryID     int       `mapstructure:"entryid"`
	EntryURL    string    `mapstructure:"entryurl"`
	Text        string    `mapstructure:"text"`
	Username    string    `mapstructure:"username"`
	UserPic     string    `mapstructure:"userpic"`
	UserPicMode int       `mapstructure:"userpic_mode"`
	UserURL     string    `mapstructure:"userurl"`
}

//Similar is a similar entry for the outer entry
type Similar struct {
	ID              int       `mapstructure:"id"`
	Date            time.Time `mapstructure:"date"`
	IsAdvertisement bool      `mapstructure:"isAdvertisement"`
	Title           string    `mapstructure:"title"`
	URL             string    `mapstructure:"url"`
}

//Badge is a comment author badge on VC
type Badge struct {
	Background string `mapstructure:"background"`
	Border     string `mapstructure:"border"`
	Color      string `mapstructure:"color"`
	Text       string `mapstructure:"text"`
	Type       string `mapstructure:"type"`
}

//Comment is a entry's comment object
type Comment struct {
	ID            int           `mapstructure:"id"`
	Author        SimpleAuthor  `mapstructure:"author"`
	CommentBundle CommentBundle `mapstructure:"commentBundle"`
	Date          time.Time     `mapstructure:"date"`
	Entry         CommentEntry  `mapstructure:"entry"`
	IsFavorited   bool          `mapstructure:"isFavorited"`
	IsPinned      bool          `mapstructure:"is_pinned"`
	Level         int           `mapstructure:"level"`
	Likes         Likes         `mapstructure:"likes"`
	Media         []Medium      `mapstructure:"media"`
	ReplyTo       int           `mapstructure:"replyTo"`
	SourceID      int           `mapstructure:"source_id"`
	Text          string        `mapstructure:"text"`
}

type Medium struct {
	AdditionalData AdditionalData `mapstructure:"additionalData"`
	IframeURL      string         `mapstructure:"iframeUrl"`
	ImageURL       string         `mapstructure:"imageUrl"`
	Service        string         `mapstructure:"service"`
	Size           Size           `mapstructure:"size"`
	Type           int            `mapstructure:"type"`
}

type AdditionalData struct {
	URL  string `mapstructure:"url"`
	UUID string `mapstructure:"uuid"`
	Size int    `mapstructure:"size"`
	Type string `mapstructure:"type"`
}

type Size struct {
	Width  int     `mapstructure:"width"`
	Height int     `mapstructure:"height"`
	Ratio  float64 `mapstructure:"ratio"`
}

type EntryBundle struct {
	CommentsCount   int       `mapstructure:"commentsCount"`
	CommentsPreview []Comment `mapstructure:"commentsPreview"`
	HitsCount       int       `mapstructure:"hitsCount"`
	IsFavorited     bool      `mapstructure:"isFavorited"`
	Likes           Likes     `mapstructure:"likes"`
}

type CommentBundle struct {
	IsFavorited bool  `mapstructure:"isFavorited"`
	Likes       Likes `mapstructure:"likes"`
}

type Likes struct {
	Count    int              `mapstructure:"count"`
	IsHidden bool             `mapstructure:"is_hidden"`
	IsLiked  bool             `mapstructure:"is_liked"`
	Likers   map[string]Liker `mapstructure:"likers"`
	Sum      int              `mapstructure:"summ"`
}

type Liker struct {
	ID        int    `mapstructure:"id"`
	Name      string `mapstructure:"name"`
	AvatarURL string `mapstructure:"avatar_url"`
	Sign      int    `mapstructure:"sign"`
}

type Vacancy struct {
	Area    int64  `mapstructure:"area"`
	City    string `mapstructure:"city_name"`
	Company struct {
		ID   int64  `mapstructure:"id"`
		Logo string `mapstructure:"logo"`
		Name string `mapstructure:"name"`
		URL  string `mapstructure:"url"`
	} `mapstructure:"company"`
	EntryID    int64  `mapstructure:"entry_id"`
	ID         int64  `mapstructure:"id"`
	SalaryFrom string `mapstructure:"salary_from"`
	SalaryTo   string `mapstructure:"salary_to"`
	Title      string `mapstructure:"title"`
}

type Tweet struct {
	Date          time.Time `mapstructure:"created_at"`
	FavoriteCount int64     `mapstructure:"favorite_count"`
	HasMedia      bool      `mapstructure:"has_media"`
	ID            string    `mapstructure:"id"`
	Media         []struct {
		MediaURL        string  `mapstructure:"media_url"`
		Ratio           float64 `mapstructure:"ratio"`
		ThumbnailWidth  int64   `mapstructure:"thumbnail_width"`
		ThumbnailHeight int64   `mapstructure:"thumbnail_height"`
		ThumbnailURL    string  `mapstructure:"thumbnail_url"`
		Type            int64   `mapstructure:"type"`
	} `mapstructure:"media"`
	RetweetCount int64  `mapstructure:"retweet_count"`
	Text         string `mapstructure:"text"`
	User         struct {
		Date                  time.Time `mapstructure:"created_at"`
		FollowersCount        int64     `mapstructure:"followers_count"`
		FriendsCount          int64     `mapstructure:"friends_count"`
		ID                    int64     `mapstructure:"id"`
		Name                  string    `mapstructure:"name"`
		ProfileImageURL       string    `mapstructure:"profile_image_url"`
		ProfileImageURLBigger string    `mapstructure:"profile_image_url_bigger"`
		ScreenName            string    `mapstructure:"screen_name"`
		StatusesCount         int64     `mapstructure:"statuses_count"`
	} `mapstructure:"user"`
}

type Rates struct {
	Btc Rate `mapstructure:"BTC"`
	Eth Rate `mapstructure:"ETH"`
	Eur Rate `mapstructure:"EUR"`
	Usd Rate `mapstructure:"USD"`
}

type Rate struct {
	Change float64 `mapstructure:"change"`
	Rate   string  `mapstructure:"rate"`
	Symbol string  `mapstructure:"sym"`
}

type PaymentsCheck struct {
	Result bool `mapstructure:"result"`
}

type responseResult struct {
	Result interface{} `json:"result"`
}

type responseError struct {
	Error struct {
		Code int         `json:"code"`
		Info interface{} `json:"info"`
	}
	Message string `json:"message"`
}

func castStruct(input interface{}, output interface{}) error {
	config := mapstructure.DecoderConfig{
		DecodeHook:       unixTimeDecodeHook,
		Result:           &output,
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	err = decoder.Decode(input)
	if err != nil {
		return err
	}
	return nil
}

func unixTimeDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.Float64 {
		return data, nil
	}
	if t != reflect.TypeOf(time.Time{}) {
		return data, nil
	}
	return time.Unix(int64(data.(float64)), 0), nil
}
