package gojikan

import "errors"

const (
	// InvalidRequestError is an error message for status code 400
	InvalidRequestError = "Invalid or incomplete request. Please double check the request documentation"

	// ResourceNotFoundError is an error message for status code 404
	ResourceNotFoundError = "Resource does not exist"

	// MethodNotAllowedError is an error message for status code 403
	MethodNotAllowedError = "Method is not allowed for this resource"

	// RateLimitedError is an error message for status code 429
	RateLimitedError = "Too many request sent. Rate limited by the source"

	// JikanAPIError is an error message for status code 500
	JikanAPIError = "Something is not working in Jikan API"

	// MyAnimeListError is an error message for status code 503
	MyAnimeListError = "Something is not working in MyAnimeList"
)

func (ths *jikanClient) checkStatusError(status int) error {
	switch status {
	case 400:
		return errors.New(InvalidRequestError)
	case 404:
		return errors.New(ResourceNotFoundError)
	case 405:
		return errors.New(MethodNotAllowedError)
	case 429:
		return errors.New(RateLimitedError)
	case 500:
		return errors.New(JikanAPIError)
	case 503:
		return errors.New(MyAnimeListError)
	}

	return nil
}
