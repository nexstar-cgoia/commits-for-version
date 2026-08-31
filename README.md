# Commits for Version

Small project that use Jira api to get all the ticket that have a specific fix version assigned to get all the commits in order for those tickets using Github api.

# Development

1. run `mise install` to install dependecies
2. run `go get .` to install go dependencies
3. run `eval $(mise activate)` to start mise environment in the current shell session
4. run `mise run start` to run the project

## Requirements

- need [mise](https://mise.jdx.dev/) in order to setup dependencies

## Environment Variables

```
JIRA_TOKEN=<jira_generated_token>
JIRA_USER=<user_email>
GITHUB_TOKEN=<github_generated_token>
```
