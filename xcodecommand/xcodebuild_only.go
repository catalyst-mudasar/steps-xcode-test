package xcodecommand

import (
	"bytes"
	"time"

	"github.com/bitrise-io/go-utils/progress"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
)

var xcodeCommandEnvs = []string{"NSUnbufferedIO=YES"}

type rawXcodeCommand struct {
	logger         log.Logger
	commandFactory command.Factory
}

func NewRawCommandRunner(logger log.Logger, commandFactory command.Factory) Runner {
	return &rawXcodeCommand{
		logger:         logger,
		commandFactory: commandFactory,
	}
}

func (c *rawXcodeCommand) Run(workDir string, args []string, _ []string, envVars string) (Output, error) {
	var (
		outBuffer bytes.Buffer
		err       error
		exitCode  int
	)

	command := c.commandFactory.Create("xcodebuild", args, &command.Opts{
		Stdout: &outBuffer,
		Stderr: &outBuffer,
		Env:    xcodeCommandEnvs,
		Dir:    workDir,
	})

	c.logger.TPrintf("$ %s", command.PrintableCommandArgs())

	progress.SimpleProgress(".", time.Minute, func() {
		exitCode, err = command.RunAndReturnExitCode()
	})

	return Output{
		RawOut:   outBuffer.Bytes(),
		ExitCode: exitCode,
	}, err
}
