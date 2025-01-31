/*
Copyright 2023 cuisongliu@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/cuisongliu/logger"
	"github.com/labring-actions/gh-rebot/pkg/types"
)

func preCheck() error {
	var err error
	types.GlobalsGithubVar, err = types.GetGHEnvToVar()
	if err != nil {
		return err
	}
	logger.Debug("github env to var: %v", types.GlobalsGithubVar)
	if err = checkGithubEnv(); err != nil {
		return err
	}
	return nil
}

func checkGithubEnv() error {
	if types.GlobalsGithubVar.RunnerID == "" {
		return fmt.Errorf("error: GITHUB_RUN_ID is not set. Please set the GITHUB_RUN_ID environment variable")
	}
	if types.GlobalsGithubVar.SafeRepo == "" {
		return fmt.Errorf("error: not found repository.full_name in github event")
	}
	if types.GlobalsGithubVar.CommentBody == "" {
		return fmt.Errorf("error: not found comment.body in github event")
	}
	if types.GlobalsGithubVar.IssueOrPRNumber == 0 {
		return fmt.Errorf("error: not found issue.number or pull_request.number in github event")
	}
	return nil
}
