# gh-rebot 项目说明文档

gh-rebot 是一个针对 sealos 项目的 GitHub rebot，用于自动执行一些常见操作，如生成变更日志、发布新版本等。本文档将介绍该 rebot 的配置文件，并提供相应的使用指南。

## 配置文件

下面是 gh-rebot 项目的配置文件：

```yaml
version: v1
debug: true
bot:
  prefix: /sealos
  spe: _
  allowOps:
    - sealos-ci-robot
    - sealos-release-robot
  email: sealos-ci-robot@sealos.io
  username: sealos-ci-robot
repo:
  org: true
  name: labring-actions/sealos
  fork: cuisongliu/sealos

changelog:
  title: "docs: Automated Changelog Update for {{.ReleaseVersion}}"
  body: |
    🤖 add release changelog using rebot.<br/>
    copilot:all
  script: scripts/changelog.sh
  allowOps:
    - cuisongliu
  reviewers:
    - cuisongliu

release:
  retry: 15s
  action: Release
  allowOps:
    - cuisongliu

message:
  success: |
    🤖 says: The action {{.Body}} finished successfully 🎉
  format_error: |
    🤖 says: ‼️ The action format error, please check the format of this action.
  permission_error: |
    🤖 says: ‼️ The action no has permission to trigger.
  release_error: |
    🤖 says: ‼️ The action release error.
  changelog_error: |
    🤖 says: ‼️ The action changelog error.

```

### 配置文件详解

- `version` - 版本标识，当前为 v1。
- `debug` - 是否开启调试模式，设置为 true 时开启。
- `bot` \- 机器人配置。
   - `prefix` - 机器人命令前缀，用于识别命令。默认值 `/`,如果设置为`/` 则 `spe` 失效。命令为`/release`
   - `spe` - 机器人命令分隔符，用于识别命令。默认值 `_`
   - `allowOps` - 允许操作的用户名列表。
   - `email` - 机器人邮箱。
   - `username` - 机器人用户名。
- `repo` \- 仓库配置。
   - `org` - 是否为组织仓库，设置为 true 时表示是组织仓库。
   - `name` - 仓库名称。
   - `fork` - fork 的仓库名称。
- `changelog` \- 变更日志配置。
   - `title` - 变更日志标题模板。`ReleaseVersion`为当前版本号
   - `body` - 变更日志内容模板。`ReleaseVersion`为当前版本号
   - `script` - 生成变更日志的脚本。默认值 `scripts/changelog.sh`,可使用模板渲染。
   - `allowOps` - 允许触发变更日志操作的用户名列表。
   - `reviewers` - 审核者列表。
- `release` \- 发布配置。
   - `retry` - 重试间隔，例如：15s。
   - `action` - 执行动作，例如：Release。
   - `allowOps` - 允许触发发布操作的用户名列表。
- `message` \- 消息配置。
   - `success` - 成功消息模板。
   - `format_error` - 格式错误消息模板。
   - `permission_error` - 权限错误消息模板。
   - `release_error` - 发布错误消息模板。
   - `changelog_error` - 变更日志错误消息模板。

## 使用文档

使用 gh-rebot 时，需要遵循以下步骤：

1. 将配置文件添加到项目的`.github`目录` gh-bot.yml `文件。
2. 确保配置文件中的用户名、仓库名称等信息与实际情况相符。
3. 根据配置文件中的命令前缀（如本例中的 `/sealos`）在 GitHub 仓库的 issue 或 PR 中发表评论，以触发相应的操作。

### 变更日志操作

如果需要生成变更日志，请在 issue 或 PR 中使用以下命令：

```bash
/sealos_changelog
```

此命令会触发配置文件中定义的脚本（如本例中的 `scripts/changelog.sh`）来生成变更日志。需要注意的是，只有在 `changelog` 配置节中的 `allowOps` 列表中的用户才有权限触发此操作。

### 发布操作

如果需要发布新版本，请在 issue 或 PR 中使用以下命令：

```
/sealos_release
```

### 错误处理

根据配置文件中的消息模板，gh-rebot 会在执行操作过程中遇到错误时返回相应的提示消息。例如：

- 格式错误：‼️ 机器人说：操作格式错误，请检查此操作的格式。
- 权限错误：‼️ 机器人说：操作无权限触发。
- 发布错误：‼️ 机器人说：操作发布错误。
- 变更日志错误：‼️ 机器人说：操作变更日志错误。

在遇到错误时，请根据提示信息进行相应的调整。


## Roadmap

- [ ] 支持label操作
- [ ] 支持里程碑操作
- [ ] 支持pr的code review操作
- [ ] 支持pr的merge操作
- [ ] 支持pr的close操作
- [ ] 支持pr的reopen操作
- [ ] 支持pr的comment操作
- [ ] 支持pr和issue的assign操作
