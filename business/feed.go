package business

import "github.com/aljorhythm/sapere-server/model"

type Tag struct {
}

type Reading struct {
	UserId int64 `json:"userId"`
}

type Feed struct {
	List []Reading
}

func presentReading(reading model.Reading) Reading {
	return Reading{
		reading.GetUserId(),
	}
}

func presentFeed(readingList model.ReadingList) Feed {
	feed := Feed{}
	for _, reading := range readingList.GetList() {
		feed.List = append(feed.List, presentReading(reading))
	}
	return feed
}
