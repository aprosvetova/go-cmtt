package cmtt

func (api *Cmtt) authQr(token string) (*User, error) {
	res, err := api.execute("auth/qr", nil, nil, map[string]string{
		"token": token,
	})
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
