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

package types

import "fmt"

type Bot struct {
	Prefix   string   `json:"prefix"`
	Spe      string   `json:"spe"`
	AllowOps []string `json:"allowOps"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
}

type Repo struct {
	Org        bool   `json:"org"`
	OrgCommand string `json:"-"`
	Name       string `json:"name"`
	Fork       string `json:"fork"`
}

type Changelog struct {
	AllowOps  []string `json:"allowOps"`
	Reviewers []string `json:"reviewers"`
	Script    string   `json:"script"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
}

type Release struct {
	Retry    string   `json:"retry"`
	Action   string   `json:"action"`
	AllowOps []string `json:"allowOps"`
}

type Config struct {
	Version string            `json:"version"`
	Debug   bool              `json:"debug"`
	Bot     Bot               `json:"bot"`
	Repo    Repo              `json:"repo"`
	Message map[string]string `json:"message"`
	Token   string            `json:"-"`

	Changelog Changelog `json:"changelog"`
	Release   Release   `json:"release"`
}

func (r *Config) Validate() error {
	if r.Bot.Username == "" {
		return fmt.Errorf("bot username is required")
	}
	if r.Bot.Email == "" {
		return fmt.Errorf("bot email is required")
	}
	if r.Repo.Name == "" {
		return fmt.Errorf("repo name is required")
	}
	if r.Repo.Fork == "" {
		return fmt.Errorf("repo fork is required")
	}
	if r.Changelog.Body == "" {
		return fmt.Errorf("changelog body is required")
	}
	if r.Changelog.Title == "" {
		return fmt.Errorf("changelog title is required")
	}
	if r.Release.Action == "" {
		return fmt.Errorf("release action is required")
	}
	if r.Release.Retry == "" {
		return fmt.Errorf("release retry is required")
	}
	return nil
}

// GetPrefix returns the prefix for the bot
func (r *Config) GetPrefix() string {
	if r.Bot.Prefix == "" {
		return "/"
	}
	return r.Bot.Prefix
}

// GetSpe returns the spe for the bot
func (r *Config) GetSpe() string {
	if r.Bot.Spe == "" {
		return "_"
	}
	return r.Bot.Spe
}

// GetBotAllowOps returns the triggers for the bot
func (r *Config) GetBotAllowOps() []string {
	return r.Bot.AllowOps
}

// GetEmail returns the email for the bot
func (r *Config) GetEmail() string {
	return r.Bot.Email
}

// GetUsername returns the username for the bot
func (r *Config) GetUsername() string {
	return r.Bot.Username
}

// GetOrgCommand returns the org command for the repo
func (r *Config) GetOrgCommand() string {
	return r.Repo.OrgCommand
}

// GetRepoName returns the name for the repo
func (r *Config) GetRepoName() string {
	return r.Repo.Name
}

// GetForkName returns the fork for the repo
func (r *Config) GetForkName() string {
	return r.Repo.Fork
}

// GetDebug returns the debug for the config
func (c *Config) GetDebug() bool {
	return c.Debug
}

// GetToken returns the token for the config
func (c *Config) GetToken() string {
	return c.Token
}

func (c *Config) GetMessage(key string) string {
	if c.Message[key] != "" {
		return c.Message[key]
	}
	return ""
}

func (c *Config) GetChangelogScript() string {
	if c.Changelog.Script == "" {
		return "scripts/changelog.sh"
	}
	return c.Changelog.Script
}
