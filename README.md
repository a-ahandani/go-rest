
```
gorest
├─ .env
├─ .git
│  ├─ COMMIT_EDITMSG
│  ├─ HEAD
│  ├─ branches
│  ├─ config
│  ├─ description
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  ├─ sendemail-validate.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ main
│  │     └─ remotes
│  │        └─ origin
│  │           ├─ HEAD
│  │           └─ main
│  ├─ objects
│  │  ├─ 06
│  │  │  └─ 410454e5399e6c64d3bc16f0409c8d79698f4a
│  │  ├─ 1f
│  │  │  └─ a0ebf2b0887941663db0a3c89ea58fd4a765c0
│  │  ├─ 21
│  │  │  └─ 06ad155a55f8a60e3068de7f0b070aeb0343c4
│  │  ├─ 54
│  │  │  └─ 77b621fcea31fc799b954199289dcca65c26a2
│  │  ├─ 66
│  │  │  └─ eab0085d1eff31544e4e035723ad79a7ce96ee
│  │  ├─ 6c
│  │  │  └─ 407b308e08eb238d28ac28d2f0aaaccc1e8d0d
│  │  ├─ 76
│  │  │  └─ 651c48a3d54d5febdc87a3ad2bc056d6eb7ef4
│  │  ├─ 9a
│  │  │  └─ 8deb0e91d7b8244f1840720ec020ebeb595910
│  │  ├─ e9
│  │  │  └─ 03cd638aefe745774af3b0c578fd91428322a0
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-c24294e244650b3d59a2735a444cb381e9032efd.idx
│  │     ├─ pack-c24294e244650b3d59a2735a444cb381e9032efd.pack
│  │     └─ pack-c24294e244650b3d59a2735a444cb381e9032efd.rev
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ main
│     ├─ remotes
│     │  └─ origin
│     │     ├─ HEAD
│     │     └─ main
│     └─ tags
├─ .gitignore
├─ Dockerfile
├─ casbin_model.conf
├─ config
│  └─ config.go
├─ database
│  └─ connect.go
├─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ init-scripts
│  ├─ 01-create-db.sql
│  └─ 02-create-user.sql
├─ internal
│  ├─ handlers
│  │  ├─ note
│  │  │  └─ note.go
│  │  └─ user
│  │     └─ user.go
│  ├─ models
│  │  ├─ note.go
│  │  └─ user.go
│  └─ validators
│     └─ validators.go
├─ main.go
├─ middlewares
│  ├─ auth.go
│  └─ casbin.go
└─ router
   ├─ note
   │  └─ note.go
   ├─ router.go
   └─ user
      └─ user.go

```
```
gorest
├─ .env
├─ .git
│  ├─ COMMIT_EDITMSG
│  ├─ HEAD
│  ├─ branches
│  ├─ config
│  ├─ description
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  ├─ sendemail-validate.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ main
│  │     └─ remotes
│  │        └─ origin
│  │           ├─ HEAD
│  │           └─ main
│  ├─ objects
│  │  ├─ 06
│  │  │  └─ 410454e5399e6c64d3bc16f0409c8d79698f4a
│  │  ├─ 1f
│  │  │  └─ a0ebf2b0887941663db0a3c89ea58fd4a765c0
│  │  ├─ 21
│  │  │  └─ 06ad155a55f8a60e3068de7f0b070aeb0343c4
│  │  ├─ 54
│  │  │  └─ 77b621fcea31fc799b954199289dcca65c26a2
│  │  ├─ 66
│  │  │  └─ eab0085d1eff31544e4e035723ad79a7ce96ee
│  │  ├─ 6c
│  │  │  └─ 407b308e08eb238d28ac28d2f0aaaccc1e8d0d
│  │  ├─ 76
│  │  │  └─ 651c48a3d54d5febdc87a3ad2bc056d6eb7ef4
│  │  ├─ 9a
│  │  │  └─ 8deb0e91d7b8244f1840720ec020ebeb595910
│  │  ├─ e9
│  │  │  └─ 03cd638aefe745774af3b0c578fd91428322a0
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-c24294e244650b3d59a2735a444cb381e9032efd.idx
│  │     ├─ pack-c24294e244650b3d59a2735a444cb381e9032efd.pack
│  │     └─ pack-c24294e244650b3d59a2735a444cb381e9032efd.rev
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ main
│     ├─ remotes
│     │  └─ origin
│     │     ├─ HEAD
│     │     └─ main
│     └─ tags
├─ .gitignore
├─ .vscode
│  └─ settings.json
├─ Dockerfile
├─ README.md
├─ casbin_model.conf
├─ config
│  └─ config.go
├─ database
│  └─ connect.go
├─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ init-scripts
│  ├─ 01-create-db.sql
│  └─ 02-create-user.sql
├─ internal
│  ├─ handlers
│  │  ├─ note
│  │  │  └─ note.go
│  │  └─ user
│  │     └─ user.go
│  ├─ models
│  │  ├─ note.go
│  │  └─ user.go
│  └─ validators
│     └─ validators.go
├─ main.go
├─ middlewares
│  ├─ auth.go
│  └─ casbin.go
└─ router
   ├─ note
   │  └─ note.go
   ├─ router.go
   └─ user
      └─ user.go

```