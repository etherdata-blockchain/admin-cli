package engine

import (
	"flag"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"cli/internal/clients"
	"cli/internal/constants"
)

func TestFlagEngine_Init(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		etdClient.EXPECT().SetNodeId(gomock.Any())
		etdClient.EXPECT().SetPassword(gomock.Any())
		etdClient.EXPECT().SetURL(gomock.Any())

		os.Args = []string{"cmd", "-template=mock", "-password=mock"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should not return a error", func() {
				So(err, ShouldBeNil)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Run(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		//oldArgs := os.Args
		mockcontroller := gomock.NewController(t)
		defer mockcontroller.Finish()

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		etdClient.EXPECT().SetNodeId(gomock.Any())
		etdClient.EXPECT().SetPassword(gomock.Any())
		etdClient.EXPECT().SetURL(gomock.Any())
		etdClient.EXPECT().VerifyPassword().Return(nil)
		etdClient.EXPECT().GetTemplate(gomock.Any()).Return(nil)

		zipClient.EXPECT().UnZip(gomock.Any())

		os.Args = []string{"cmd", "-template=mock", "-password=mock"}

		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)

		Convey("When calling run function", func() {
			err := engine.Init()
			err = engine.Run()
			Convey("Should not return a error", func() {
				So(err, ShouldBeNil)
			})
		})
		Reset(func() {
			defer mockcontroller.Finish()
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		})
	})
}
