package cmtt

//GetVacancies returns vacancies from VC widget
func (api *Cmtt) GetVacancies() (*[]Vacancy, error) {
	res, err := api.execute("vacancies/widget", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &[]Vacancy{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetRates returns currency rates from VC
func (api *Cmtt) GetRates() (*Rates, error) {
	res, err := api.execute("rates", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	result := &Rates{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

//GetPaymentsCheck returns... I don't really know. Didn't test it
func (api *Cmtt) GetPaymentsCheck() (*PaymentsCheck, error) {
	//TODO: fix description
	res, err := api.execute("payments/check", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	paymentsCheck := &PaymentsCheck{}
	err = castStruct(res, paymentsCheck)
	if err != nil {
		return nil, err
	}
	return paymentsCheck, err
}

/*
TODO: Ask Cmtt for regular response format
Dear Cmtt,
Stop using DIFFERENT RESPONSE formats.
I'm expecting {"result": ...} in every method.
Thanks.
func (api *Cmtt) GetTweets(mode string, count, offset int) (*[]Tweet, error) {
	res, err := api.execute("tweets/{mode}", map[string]string {
		"mode": mode,
	}, map[string]string {
		"count": strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
	}, nil)
	if err != nil {
		return nil, err
	}
	result := &[]Tweet{}
	err = castStruct(res, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
*/
