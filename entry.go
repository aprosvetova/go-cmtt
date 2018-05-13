package cmtt

import (
	"fmt"
	"strconv"
	"strings"
)

//GetEntry returns entry by ID
func (api *Cmtt) GetEntry(entryID int) (*Entry, error) {
	res, err := api.execute("entry/{id}", map[string]string{
		"id": strconv.Itoa(entryID),
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &Entry{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetEntryComments returns array of comments for the specified entry
func (api *Cmtt) GetEntryComments(entryID int, hash, sorting string) (*[]Comment, error) {
	res, err := api.execute("entry/{id}/comments/{sorting}", map[string]string{
		"id":      strconv.Itoa(entryID),
		"sorting": sorting,
	}, map[string]string{
		"hash": hash,
	}, nil)
	if err != nil {
		return nil, err
	}
	result := &[]Comment{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetFlashEntries returns urgent news. Usually contains only one element
func (api *Cmtt) GetFlashEntries() (*[]Entry, error) {
	res, err := api.execute("getflashholdedentry", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &[]Entry{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetEntryBundles returns short info about entries specified
func (api *Cmtt) GetEntryBundles(IDs []int) (*map[string]EntryBundle, error) {
	res, err := api.execute("entry/bundle", nil, map[string]string{
		"ids": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(IDs)), ","), "[]"),
	}, nil)
	if err != nil {
		return nil, err
	}
	result := &map[string]EntryBundle{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetRecentNews returns list of entries shown as news on VC
func (api *Cmtt) GetRecentNews(count, offset int) (*[]Entry, error) {
	res, err := api.execute("news/default/recent", nil, map[string]string{
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
	}, nil)
	if err != nil {
		return nil, err
	}
	result := &[]Entry{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//LikeEntry likes the specified entry
func (api *Cmtt) LikeEntry(entryID int) (*Likes, error) {
	return api.rate(entryID, 1)
}

//DislikeEntry dislikes the specified entry
func (api *Cmtt) DislikeEntry(entryID int) (*Likes, error) {
	return api.rate(entryID, -1)
}

//CommentEntry adds a comment to the specified entry
func (api *Cmtt) CommentEntry(entryID int, text string, replyTo int, attachments map[string]string) (*Comment, error) {
	body := map[string]string{
		"text":     text,
		"reply_to": strconv.Itoa(replyTo),
	}
	if attachments != nil {
		for k, v := range attachments {
			body[k] = v
		}
	}
	res, err := api.execute("entry/{id}/comments", map[string]string{
		"id": strconv.Itoa(entryID),
	}, nil, body)
	if err != nil {
		return nil, err
	}
	result := &Comment{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (api *Cmtt) rate(entryID, sign int) (*Likes, error) {
	res, err := api.execute("entry/{id}/likes", map[string]string{
		"id": strconv.Itoa(entryID),
	}, nil, map[string]string{
		"sign": strconv.Itoa(sign),
	})
	if err != nil {
		return nil, err
	}
	result := &Likes{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
