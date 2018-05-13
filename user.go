package cmtt

import (
	"strconv"
)

//GetMe returns User object for authenticated user
func (api *Cmtt) GetMe() (*User, error) {
	res, err := api.execute("user/me", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &User{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetMineComments returns comments that were posted by authenticated user
func (api *Cmtt) GetMineComments(count, offset int) (*[]Comment, error) {
	res, err := api.execute("user/me/comments", nil, map[string]string{
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
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

//GetMineEntries returns entries that were posted by authenticated user
func (api *Cmtt) GetMineEntries(count, offset int) (*[]Entry, error) {
	res, err := api.execute("user/me/entries", nil, map[string]string{
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

//GetMineFavoriteComments returns comments that were favored by authenticated user
func (api *Cmtt) GetMineFavoriteComments(count, offset int) (*[]Comment, error) {
	res, err := api.execute("user/me/favorites/comments", nil, map[string]string{
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
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

//GetMineFavoriteEntries returns entries that were favored by authenticated user
func (api *Cmtt) GetMineFavoriteEntries(count, offset int) (*[]Entry, error) {
	res, err := api.execute("user/me/favorites/entries", nil, map[string]string{
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

//GetUser returns User object for the specified user
func (api *Cmtt) GetUser(userID int) (*User, error) {
	res, err := api.execute("user/{id}", map[string]string{
		"id": strconv.Itoa(userID),
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &User{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetUserComments returns comments that were posted by the specified user
func (api *Cmtt) GetUserComments(userID, count, offset int) (*[]Comment, error) {
	res, err := api.execute("user/{id}/comments", map[string]string{
		"id": strconv.Itoa(userID),
	}, map[string]string{
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
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

//GetUserEntries returns entries that were posted by the specified user
func (api *Cmtt) GetUserEntries(userID, count, offset int) (*[]Entry, error) {
	res, err := api.execute("user/{id}/entries", map[string]string{
		"id": strconv.Itoa(userID),
	}, map[string]string{
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

//GetUserFavoriteComments returns comments that were favored by the specified user
func (api *Cmtt) GetUserFavoriteComments(userID, count, offset int) (*[]Comment, error) {
	res, err := api.execute("user/{id}/favorites/comments", map[string]string{
		"id": strconv.Itoa(userID),
	}, map[string]string{
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
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

//GetUserFavoriteEntries returns entries that were favored by authenticated user
func (api *Cmtt) GetUserFavoriteEntries(userID, count, offset int) (*[]Entry, error) {
	res, err := api.execute("user/{id}/favorites/entries", map[string]string{
		"id": strconv.Itoa(userID),
	}, map[string]string{
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

//GetUserPushTopic returns... I don't really know. Didn't test it
func (api *Cmtt) GetUserPushTopic() (string, error) {
	//TODO: fix description
	res, err := api.execute("user/push/topic", nil, nil, nil)
	if err != nil {
		return "", err
	}
	var result string
	err = castStruct(res, &result)
	if err != nil {
		return "", err
	}
	return result, err
}
