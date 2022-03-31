package clients

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestETD_GetTemplate(t *testing.T) {

}

func TestETD_ListTemplate(t *testing.T) {

}

func TestETD_SetNodeId(t *testing.T) {
	Convey("Given a etd client", t, func() {
		client := ETD{}
		Convey("When calling set node id", func() {
			client.SetNodeId("mock_id")
			Convey("Should update the client's field", func() {
				So(client.nodeId, ShouldEqual, "mock_id")
			})
		})
	})

}

func TestETD_SetPassword(t *testing.T) {
	Convey("Given a etd client", t, func() {
		client := ETD{}
		Convey("When calling set password", func() {
			client.SetPassword("mock_password")
			Convey("Should update the client's field", func() {
				So(client.password, ShouldEqual, "mock_password")
			})
		})
	})
}

func TestETD_SetURL(t *testing.T) {
	Convey("Given a etd client", t, func() {
		client := ETD{}
		Convey("When calling set url", func() {
			client.SetURL("https://node_url")
			Convey("Should update the client's field", func() {
				So(client.url, ShouldEqual, "https://node_url")
			})
		})
	})

}

func TestETD_VerifyPassword(t *testing.T) {

}
