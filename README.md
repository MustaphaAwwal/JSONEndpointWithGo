# JSON Endpoint with Go

## Objective

Create and host an endpoint using the Go programming language that takes two GET request query parameters and returns specific information in JSON format.

## Requirements

The information required includes:

- Slack name
- Current day of the week
- Current UTC time (with validation of +/-2 minutes)
- Track
- The GitHub URL of the file being run
- The GitHub URL of the full source code
- A Status Code of Success (200)

### JSON Format

```json
{
  "slack_name": "example_name",
  "current_day": "Monday",
  "utc_time": "2023-08-21T15:04:05Z",
  "track": "backend",
  "github_file_url": "https://github.com/username/repo/blob/main/file_name.ext",
  "github_repo_url": "https://github.com/username/repo",
  "status_code": 200
}
