# Guidelines for Junie AI Agent

## About the project

This project is a CLI/TUI tool that allows users to browser and edit Google Firestore data.
It can call Firestore API in two modes:

- Using Firebase Admin API (_using service account_)
- Using Firestore Client API (_using Firebase user token obtained through OAuth2_)

## Tech stack

- `github.com/charmbracelet/bubbletea` as UI framework
- `github.com/charmbracelet/bubbles` as basic UI components
- `cloud.google.com/go/firestore` - for reading and writing to a Cloud Firestore database
- `google.golang.org/api/firebase` - for calling Firebase Admin API

## Project structure

- CLI entry point is in `fsv` package
    - Should be minimal
- The implementation of Firestore Viewer is in `fsviewer` package
- Code related to authentication with Firebase is in `fbauth` package
    - UI for displaying list of added service accounts and
        - choose an active one
        - deleting a service account
    - UI for adding a new service account (name and path to JSON file)
    - UI for displaying list of authenticated Firebase users and
        - choosing an active one
        - Signing out
        - Signing in into a new user account using OAuth2 in browser

## Persisted data

Some data are persisted. That way we can restore state of the app on app restart.

### Data persisted to local storage

- List of service accounts
  ```go
  type ServiceAccount struct {
    Name string `json:"name"` // Human readable alias for the service account
    Path string `json:"path"` // Path to the JSON file with Firebase service account credentials
  }
  ```
- Firestore viewer UI state
    - ID of active service account (can be empty if not selected yet or if active account deleted)
    - ID of active Firebase project
    - ID of ative Firestore collection

## UI

- For menus we use `bubbles.List`
- For tables we use bubbles table
- If not in edit mode pressing ESC returns user to a parent screen (_if applicable_)

### UI Compoments

- [Breadcrumbs](./ui-breadcrumbs.md)
- [Navigation breadcrumbs](./ui-nav-breadcrumbs)

### Bubbletea UI models

- `firestoreViewerState`
    - `activeServiceAccount ServiceAccount`
    - `activeFirebaseUserAccount FirebaseUserAccount`
- `firestoreViewerTopMenu` - Main menu of next items:
    - Service accounts: {numberOfServiceAccount}<, active: {firestoreViewerState.activeServiceAccount.Name}>
    - Collections
    - [About Firestore Viewer](https://github.com/datatug/firestore-viewer)

#### Projects

At top shows [navigation breadcrumbs](./ui-nav-breadcrumbs) like: `{activeServiceAccount.Name} > Projects`

Underneath shows list of Firebase projects available to the active service account:

- Project A
- Project B
- ...
  When a project selected goes to the project view.

#### Project

At top show navigation breadcrumbs like `{activeServiceAccount.Name} > Projects > {Project.Name}`.

Underneath shows project menu:

- Collections
- Indexes

#### Collections

At top show [navigation breadcrumbs](./ui-nav-breadcrumbs) like
`{activeServiceAccount.Name} > Projects > {Project.Name}`.

Underneath shows list of root Firestore collections:

- Collection A
- Collection B
- ...
  When a collection is selected, user goes to the collection view.

#### Collection

At top show [navigation breadcrumbs](./ui-nav-breadcrumbs) like
`{activeServiceAccount.Name} > Projects > {Project.Name} > Collections`.

If collection has parent underneath show [collection full path breadcrumbs](./ui-col-breadcrumbs.md) .

Below a table of records from the current colletion is shown.

## [Code style and best practices](./CODE_STYLE.md)

You must learn and understand [code style and best practices](./CODE_STYLE.md).

## Final validation before submitting changes:

- Run `go vet ./...` and make sure it passes
- Run `golangci-lint run ./...` and make sure all reported issues are fixed
- Run `go build ./...` and make sure it passes
- Run `go test ./... --cover` and make sure there is no failing tests
