# pkg0

to reproduce:

```
[Info  - 20:29:02] 2025/11/25 20:29:02 Created View (#1)
	directory=/home/alvarolm/temp/pkg0
	view_type="GoMod"
	root_dir="file:///home/alvarolm/temp/pkg0"
	go_version="go version go1.25.4 linux/amd64"
	build_flags=[]
	env={GOOS:linux GOARCH:amd64 GOCACHE:/home/alvarolm/.cache/go-build GOMODCACHE:/home/alvarolm/go/pkg/mod GOPATH:/home/alvarolm/go GOPRIVATE: GOFLAGS: GO111MODULE: GOTOOLCHAIN:auto GOROOT:/usr/local/go GoVersion:25 GoVersionOutput:go version go1.25.4 linux/amd64
 ExplicitGOWORK: EffectiveGOPACKAGESDRIVER:}
	env_overlay=[]

[Info  - 20:29:02] 2025/11/25 20:29:02 go/packages.Load #1
	view_id="1"
	snapshot=0
	directory=/home/alvarolm/temp/pkg0
	query=[/home/alvarolm/temp/pkg0/... builtin]
	packages=3
	duration=90.208676ms

[Error - 20:29:04] Request textDocument/implementation failed.
  Message: X is a var, not a type
  Code: 0 
[Info  - 20:29:06] 2025/11/25 20:29:06 go/packages.Load #2
	view_id="1"
	snapshot=2
	directory=/home/alvarolm/temp/pkg0
	query=[file=/home/alvarolm/temp/pkg0/pkg1/a.templ]
	packages=0
	duration=56.402668ms

[Error - 20:29:06] Request textDocument/rename failed.
  Message: no package metadata for file file:///home/alvarolm/temp/pkg0/pkg1/a.templ
  Code: 0 

```

```bash
$ templ info
(✓) os [ goos=linux goarch=amd64 ]
(✓) go [ location=/usr/local/go/bin/go version=go version go1.25.4 linux/amd64 ]
(✓) gopls [ location=/home/alvarolm/go/bin/gopls version=golang.org/x/tools/gopls v0.20.0 ]
(✓) templ [ location=/home/alvarolm/go/bin/templ version=v0.3.960 ]
```

vscode version:
```
Version: 1.106.2
Commit: 1e3c50d64110be466c0b4a45222e81d2c9352888
Date: 2025-11-19T16:56:50.023Z
Electron: 37.7.0
ElectronBuildId: 12781156
Chromium: 138.0.7204.251
Node.js: 22.20.0
V8: 13.8.258.32-electron.0
OS: Linux x64 6.1.0-40-amd64
```

vscode extension:
```
Identifier  a-h.templ
Version     0.0.35
```