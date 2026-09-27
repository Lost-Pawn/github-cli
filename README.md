# github-cli

A command-line tool to track GitHub activity — view a user's recent public events, or poll a specific repository for new activity in real time.

Built as part of the [roadmap.sh backend beginner projects](https://roadmap.sh/projects/github-user-activity).

## Features

- **`github-activity`** — fetch and display a GitHub user's recent public activity (pushes, forks, issues, pull requests, comments, and more)
- **`git-watch`** — poll a specific repository at a configurable interval and print new activity as it happens
- **`help`** — list available commands

## Installation

Clone the repo and install the binary:

```bash
git clone https://github.com/<your-username>/github-cli.git
cd github-cli
go install .
```

This installs a `github-cli` binary to your `$GOPATH/bin` (make sure that's on your `PATH`).

## Usage

### View a user's recent activity

```bash
github-cli github-activity <username>
```

Example:

```bash
github-cli github-activity lost-pawn
```

Sample output:

```
Recent GitHub activities for user 'torvalds':
- Pushed to repository 'torvalds/linux' at 2026-09-20T10:15:00Z
  Commit message: fix off-by-one in scheduler
- Forked repository 'someuser/somerepo' to 'torvalds/somerepo' at 2026-09-19T08:00:00Z
- Issue 'Build failing on ARM' (number 42) in repository 'torvalds/linux', Action: opened at 2026-09-18T14:30:00Z
```

### Watch a repository for new activity

```bash
github-cli git-watch -repo <owner>/<repo> -interval <seconds>
```

- `-repo` (required) — the target repository, in `owner/repo` format
- `-interval` (optional, default `60`) — how often (in seconds) to poll for new events

Example:

```bash
github-cli git-watch -repo golang/go -interval 30
```

The command runs indefinitely, polling on the given interval and printing only newly seen events. Stop it with `Ctrl+C`.

### Help

```bash
github-cli help
```

Lists all available commands and basic usage examples.

## How it works

- Uses the public GitHub Events API (`/users/{username}/events` and `/repos/{owner}/{repo}/events`) — no authentication required, subject to GitHub's [unauthenticated rate limits](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api).
- `git-watch` deduplicates events between polls by comparing each event's unique `id` field against the most recent one seen on the prior poll.

## Project structure

```
github-cli/
├── main.go              # entry point
├── cmd/
│   └── root.go          # command routing (github-activity, git-watch, help)
└── internals/
    └── activity.go       # GitHub API types, fetch logic, display logic, watch/poll logic
```

## Limitations

- Only tracks *public* events (GitHub's Events API doesn't surface private activity without authentication).
- Polling, not true webhooks — no server or public endpoint required, but updates are only as fresh as your `-interval`.
- Subject to GitHub's unauthenticated API rate limits (60 requests/hour per IP); frequent polling on a short interval can hit this quickly.