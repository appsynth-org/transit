# Transit

Trans(late) it

## Prerequisites

1. Create service account in GCP console. Generate a new key and export the json file
2. Encrypted the service_account.json file with base64 and store with `SERVICE_ACCOUNT_BASE64` key in .env
3. Get the spreadsheet id and store with `GOOGLE_SHEET_ID` key in .env

### Command

```shell
# Read the google sheet, then generated the output.json
go run main.go
```

### References & Guides

- [How to create a new Service account](https://robocorp.com/docs/development-guide/google-sheets/interacting-with-google-sheets#create-a-new-google-sheet-and-add-the-service-account-as-an-editor-to-it)

- [Google Sheet SDK docs](https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/get)

### Release
1. git checkout main
2. git tag v*, i.e. v.0.1, v.1.0, v.10.0
3. git tag origin [tag]