package gojikan

import "errors"

func (ths *jikanClient) checkStatusError(status int) error {
	switch status {
	case 400:
		return errors.New("Invalid or incomplete request. Please double check the request documentation")
	case 404:
		return errors.New("Resource does not exist")
	case 405:
		return errors.New("Method is not allowed for this resource")
	case 429:
		return errors.New("Too many request sent. Rate limited by the source")
	case 500:
		return errors.New("Something is not working in Jikan API")
	case 503:
		return errors.New("Something is not working in MyAnimeList")
	}

	return nil
}
