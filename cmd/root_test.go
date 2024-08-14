package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedCommandRun struct {
	mock.Mock
}

func (m *MockedCommandRun) RunE(cmd *cobra.Command, args []string) error {
	margs := m.Called(cmd, args)
	return margs.Error(0)
}

func setupMockRunE() map[string]*MockedCommandRun {
	mocks := make(map[string]*MockedCommandRun)

	for _, cmd := range rootCmd.Commands() {
		mockedCommandRunE := new(MockedCommandRun)
		mockedCommandRunE.On("RunE", mock.Anything, mock.Anything).Return(nil)
		mocks[cmd.Name()] = mockedCommandRunE
		cmd.RunE = mockedCommandRunE.RunE
	}

	return mocks
}

func assertCommandCalled(t *testing.T, command string, mocked map[string]*MockedCommandRun) {
	targetRunE := "RunE"

	for name, cmd := range mocked {
		if name == command {
			cmd.AssertCalled(t, targetRunE, mock.Anything, mock.Anything)
		} else {
			cmd.AssertNotCalled(t, targetRunE)
		}
	}
}

func TestRootCommand(t *testing.T) {
	commands := []string{"migrate", "serve"}
	for _, command := range commands {
		t.Run(command+" by setting storage.database to postgres", func(t *testing.T) {
			rootCmd.SetArgs([]string{command, "--storage.database=postgres"})

			mocked := setupMockRunE()
			assert.Nil(t, rootCmd.Execute())

			assertCommandCalled(t, command, mocked)
			assert.Equal(t, "postgres", viper.GetString("storage.database"))
		})
	}
}