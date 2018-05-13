package cmtt

import (
	"fmt"
	"strconv"
	"strings"
)

//GetCommentLikers returns a map with people who liked/disliked the specified comment. Key is UserID
func (api *Cmtt) GetCommentLikers(commentID int) (*map[int]Liker, error) {
	res, err := api.execute("comment/likers/{commentId}", map[string]string{
		"commentId": strconv.Itoa(commentID),
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &map[int]Liker{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetCommentBundles returns short info about comments specified
func (api *Cmtt) GetCommentBundles(IDs []int) (*map[string]CommentBundle, error) {
	res, err := api.execute("comment/bundle", nil, map[string]string{
		"ids": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(IDs)), ","), "[]"),
	}, nil)
	if err != nil {
		return nil, err
	}
	result := &map[string]CommentBundle{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//EditComment edits already posted comment
func (api *Cmtt) EditComment(entryID, commentID int, text string) error {
	//TODO: Ask Cmtt for regular response format
	_, err := api.execute("comment/edit/{commentId}/{entryId}", map[string]string{
		"commentId": strconv.Itoa(commentID),
		"entryId":   strconv.Itoa(entryID),
	}, nil, map[string]string{})
	if err != nil {
		return err
	}
	return nil
}
