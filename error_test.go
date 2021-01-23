package gojikan

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckStatusError(t *testing.T) {
	Convey("Testing Check Status Error Method", t, func() {
		client := NewJikanClient().(*jikanClient)

		Convey("checkStatusError() with status 404 should return error ResourcesNotFoundError", func() {
			err := client.checkStatusError(404)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, ResourceNotFoundError)
		})

		Convey("checkStatusError() with status 400 should return error InvalidRequestError", func() {
			err := client.checkStatusError(400)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, InvalidRequestError)
		})

		Convey("checkStatusError() with status 405 should return error MethodNotAllowedError", func() {
			err := client.checkStatusError(405)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, MethodNotAllowedError)
		})

		Convey("checkStatusError() with status 429 should return error RateLimitedError", func() {
			err := client.checkStatusError(429)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, RateLimitedError)
		})

		Convey("checkStatusError() with status 500 should return error JikanAPIError", func() {
			err := client.checkStatusError(500)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, JikanAPIError)
		})

		Convey("checkStatusError() with status 503 should return error MyAnimeListError", func() {
			err := client.checkStatusError(503)

			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, MyAnimeListError)
		})

		Convey("checkStatusError() with status 200 should return nil", func() {
			err := client.checkStatusError(200)

			So(err, ShouldBeNil)
		})
	})
}
