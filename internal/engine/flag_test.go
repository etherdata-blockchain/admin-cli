package engine

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"cli/internal/clients"
	"cli/internal/constants"
	customerror "cli/internal/errors"
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

func TestFlagEngine_Init_Environment_Beta(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		etdClient.EXPECT().SetNodeId(gomock.Any())
		etdClient.EXPECT().SetPassword(gomock.Any())
		etdClient.EXPECT().SetURL(gomock.Any())

		os.Args = []string{"cmd", "-template=mock", "-password=mock", "-environment=beta"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should not return a error", func() {
				So(err, ShouldBeNil)
				So(engine.Config.DefaultEndpoint, ShouldEqual, constants.BetaEndpoint)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_Environment_Production(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		etdClient.EXPECT().SetNodeId(gomock.Any())
		etdClient.EXPECT().SetPassword(gomock.Any())
		etdClient.EXPECT().SetURL(gomock.Any())

		os.Args = []string{"cmd", "-template=mock", "-password=mock", "-environment=production"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should not return a error", func() {
				So(err, ShouldBeNil)
				So(engine.Config.DefaultEndpoint, ShouldEqual, constants.ProductionEndpoint)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_Environment_Local(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		etdClient.EXPECT().SetNodeId(gomock.Any())
		etdClient.EXPECT().SetPassword(gomock.Any())
		etdClient.EXPECT().SetURL(gomock.Any())

		os.Args = []string{"cmd", "-template=mock", "-password=mock", "-environment=local"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should not return a error", func() {
				So(err, ShouldBeNil)
				So(engine.Config.DefaultEndpoint, ShouldEqual, constants.LocalEndpoint)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_invalid_template_ErrorCase(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		os.Args = []string{"cmd", "-password=mock"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should return a error", func() {
				So(errors.Is(err, customerror.NewInvalidTemplateIdError("")), ShouldBeTrue)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_invalid_password_ErrorCase(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		os.Args = []string{"cmd", "-template=mock"}
		t.Setenv(constants.NodeIDKey, "mock_id")
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should return a error", func() {
				So(errors.Is(err, customerror.NewInvalidPasswordError("")), ShouldBeTrue)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_invalid_node_id_ErrorCase(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		os.Args = []string{"cmd", "-template=mock", "-password=mock"}
		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should return a error", func() {
				So(errors.Is(err, customerror.NewInvalidNodeIdError("")), ShouldBeTrue)
			})
		})

		Reset(func() {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			defer mockcontroller.Finish()
		})
	})
}

func TestFlagEngine_Init_invalid_environment_ErrorCase(t *testing.T) {
	Convey("Given a flag engine", t, func() {
		mockcontroller := gomock.NewController(t)

		etdClient := clients.NewMockETDInterface(mockcontroller)
		zipClient := clients.NewMockZipInterface(mockcontroller)

		os.Args = []string{"cmd", "-template=mock", "-password=mock", "-environment=test"}
		t.Setenv(constants.NodeIDKey, "mock_id")

		engine := NewFlagEngine(etdClient, zipClient)
		Convey("When calling init function", func() {
			err := engine.Init()
			Convey("Should return a error", func() {
				So(errors.Is(err, customerror.NewInvalidEnvironmentError("")), ShouldBeTrue)
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
		zipClient.EXPECT().Remove(gomock.Any())

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
