## List Command

- Creating migration file
  goose create nama_table sql --dir=./db/migration
- Up Migration
  1. Set Goose Config in .env
  2. Command
  ```bash
  goose up
  ```
- Down Migration (by 1 last version)
  1. Set Goose Config in .env
  2. Command
  ```bash
  goose down
  ```
