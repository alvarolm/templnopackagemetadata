# pkg0

to reproduce: try to rename the name argument from the Greet method in x.templ

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


templ.log:
```json
{"time":"2025-11-25T23:32:10.610131735-03:00","level":"INFO","msg":"lsp: starting up..."}
{"time":"2025-11-25T23:32:10.610223871-03:00","level":"INFO","msg":"lsp: starting gopls..."}
{"time":"2025-11-25T23:32:10.610284035-03:00","level":"INFO","msg":"creating gopls client"}
{"time":"2025-11-25T23:32:10.610304105-03:00","level":"INFO","msg":"creating proxy"}
{"time":"2025-11-25T23:32:10.610309269-03:00","level":"INFO","msg":"creating templ server"}
{"time":"2025-11-25T23:32:10.610315562-03:00","level":"INFO","msg":"starting debug http server","addr":"localhost:7575"}
{"time":"2025-11-25T23:32:10.610347682-03:00","level":"INFO","msg":"listening"}
{"time":"2025-11-25T23:32:10.614179029-03:00","level":"INFO","msg":"client -> server: Initialize"}
{"time":"2025-11-25T23:32:10.679350298-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata"}
{"time":"2025-11-25T23:32:10.679578826-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git"}
{"time":"2025-11-25T23:32:10.679625159-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/HEAD"}
{"time":"2025-11-25T23:32:10.679664477-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/branches"}
{"time":"2025-11-25T23:32:10.679693788-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/config"}
{"time":"2025-11-25T23:32:10.679706684-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/description"}
{"time":"2025-11-25T23:32:10.679775138-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks"}
{"time":"2025-11-25T23:32:10.679788775-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/applypatch-msg.sample"}
{"time":"2025-11-25T23:32:10.679807573-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/commit-msg.sample"}
{"time":"2025-11-25T23:32:10.679816404-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/fsmonitor-watchman.sample"}
{"time":"2025-11-25T23:32:10.679824338-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/post-update.sample"}
{"time":"2025-11-25T23:32:10.679833284-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-applypatch.sample"}
{"time":"2025-11-25T23:32:10.679842225-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-commit.sample"}
{"time":"2025-11-25T23:32:10.679852819-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-merge-commit.sample"}
{"time":"2025-11-25T23:32:10.679860966-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-push.sample"}
{"time":"2025-11-25T23:32:10.679870572-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-rebase.sample"}
{"time":"2025-11-25T23:32:10.679878589-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/pre-receive.sample"}
{"time":"2025-11-25T23:32:10.679887867-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/prepare-commit-msg.sample"}
{"time":"2025-11-25T23:32:10.679898426-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/push-to-checkout.sample"}
{"time":"2025-11-25T23:32:10.679908177-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/hooks/update.sample"}
{"time":"2025-11-25T23:32:10.679918533-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/index"}
{"time":"2025-11-25T23:32:10.679949455-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/info"}
{"time":"2025-11-25T23:32:10.679981359-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/info/exclude"}
{"time":"2025-11-25T23:32:10.680002857-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs"}
{"time":"2025-11-25T23:32:10.680012315-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/HEAD"}
{"time":"2025-11-25T23:32:10.68003411-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs"}
{"time":"2025-11-25T23:32:10.680055197-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs/heads"}
{"time":"2025-11-25T23:32:10.680063572-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs/heads/main"}
{"time":"2025-11-25T23:32:10.680083403-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs/remotes"}
{"time":"2025-11-25T23:32:10.680103814-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs/remotes/origin"}
{"time":"2025-11-25T23:32:10.680114976-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/logs/refs/remotes/origin/main"}
{"time":"2025-11-25T23:32:10.680141111-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects"}
{"time":"2025-11-25T23:32:10.6801709-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/07"}
{"time":"2025-11-25T23:32:10.680186648-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/07/518be9954a9b0c2ad14b8b9f64dc18e5f6a827"}
{"time":"2025-11-25T23:32:10.680217722-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/24"}
{"time":"2025-11-25T23:32:10.680230883-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/24/6fd530dbfc616602e358851092766e513a9525"}
{"time":"2025-11-25T23:32:10.680260131-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/3a"}
{"time":"2025-11-25T23:32:10.680271023-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/3a/bed9c3753aee765c6a6f21f446b3225ad792e8"}
{"time":"2025-11-25T23:32:10.680297223-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/40"}
{"time":"2025-11-25T23:32:10.680308781-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/40/01de0fe9d2f2127e909508f6667f12b70d01f9"}
{"time":"2025-11-25T23:32:10.680335196-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/4b"}
{"time":"2025-11-25T23:32:10.680347755-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/4b/825dc642cb6eb9a060e54bf8d69288fbee4904"}
{"time":"2025-11-25T23:32:10.680375564-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/6f"}
{"time":"2025-11-25T23:32:10.680385905-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/6f/e4aac816f6e294b62b70ee65279d29f4fd89a7"}
{"time":"2025-11-25T23:32:10.68041078-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/89"}
{"time":"2025-11-25T23:32:10.680422935-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/89/2079492ffd51e52b81fb18953c5b1162a3dc5f"}
{"time":"2025-11-25T23:32:10.680448878-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/bc"}
{"time":"2025-11-25T23:32:10.680470123-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/bc/83599ced512399c04c0b168f1570eb3d5a6d87"}
{"time":"2025-11-25T23:32:10.680498664-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/c6"}
{"time":"2025-11-25T23:32:10.680519322-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/c6/d445c7b834b566ef7d8736a338116e8cf25a51"}
{"time":"2025-11-25T23:32:10.680549201-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/ce"}
{"time":"2025-11-25T23:32:10.680569897-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/ce/fdacd40c8032fa99ffc8e6d218c493789063a7"}
{"time":"2025-11-25T23:32:10.680603022-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/de"}
{"time":"2025-11-25T23:32:10.680624052-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/de/bf5b7ddbfe842bc86923fa449aeab0960c1c93"}
{"time":"2025-11-25T23:32:10.68065655-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/e7"}
{"time":"2025-11-25T23:32:10.680672479-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/e7/1211bbd3578a337c023ac589f8e488ac17c581"}
{"time":"2025-11-25T23:32:10.680706077-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/info"}
{"time":"2025-11-25T23:32:10.68073257-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/objects/pack"}
{"time":"2025-11-25T23:32:10.680761717-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs"}
{"time":"2025-11-25T23:32:10.680793382-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/heads"}
{"time":"2025-11-25T23:32:10.68080685-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/heads/main"}
{"time":"2025-11-25T23:32:10.680834518-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/remotes"}
{"time":"2025-11-25T23:32:10.680859133-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/remotes/origin"}
{"time":"2025-11-25T23:32:10.680871801-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/remotes/origin/main"}
{"time":"2025-11-25T23:32:10.680903412-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/.git/refs/tags"}
{"time":"2025-11-25T23:32:10.6809162-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/README.md"}
{"time":"2025-11-25T23:32:10.680927805-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/go.mod"}
{"time":"2025-11-25T23:32:10.680942337-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/go.sum"}
{"time":"2025-11-25T23:32:10.68095744-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/main.go"}
{"time":"2025-11-25T23:32:10.680969579-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:10.68185196-03:00","level":"INFO","msg":"setting source map cache contents","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:10.681935589-03:00","level":"INFO","msg":"found file","path":"/home/alvarolm/temp/templnopackagemetadata/x_templ.go"}
{"time":"2025-11-25T23:32:10.681950994-03:00","level":"INFO","msg":"client -> server: Initialize end"}
{"time":"2025-11-25T23:32:10.70016551-03:00","level":"INFO","msg":"client -> server: Initialized"}
{"time":"2025-11-25T23:32:10.700281863-03:00","level":"INFO","msg":"client -> server: Initialized end"}
{"time":"2025-11-25T23:32:10.70071257-03:00","level":"INFO","msg":"client <- server: WorkDoneProgressCreate"}
{"time":"2025-11-25T23:32:10.715098818-03:00","level":"INFO","msg":"client -> server: DidOpen","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:10.715374057-03:00","level":"INFO","msg":"setting source map cache contents","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:10.715431963-03:00","level":"INFO","msg":"client -> server: DidOpen end"}
{"time":"2025-11-25T23:32:11.099973518-03:00","level":"INFO","msg":"client -> server: DocumentSymbol"}
{"time":"2025-11-25T23:32:11.100254987-03:00","level":"INFO","msg":"client <- server: Progress"}
{"time":"2025-11-25T23:32:11.100351629-03:00","level":"INFO","msg":"client <- server: Configuration"}
{"time":"2025-11-25T23:32:11.152258277-03:00","level":"INFO","msg":"client <- server: LogMessage","message":"2025/11/25 23:32:11 Created View (#1)\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tview_type=\"GoMod\"\n\troot_dir=\"file:///home/alvarolm/temp/templnopackagemetadata\"\n\tgo_version=\"go version go1.25.4 linux/amd64\"\n\tbuild_flags=[]\n\tenv={GOOS:linux GOARCH:amd64 GOCACHE:/home/alvarolm/.cache/go-build GOMODCACHE:/home/alvarolm/go/pkg/mod GOPATH:/home/alvarolm/go GOPRIVATE: GOFLAGS: GO111MODULE: GOTOOLCHAIN:auto GOROOT:/usr/local/go GoVersion:25 GoVersionOutput:go version go1.25.4 linux/amd64\n ExplicitGOWORK: EffectiveGOPACKAGESDRIVER:}\n\tenv_overlay=[]\n"}
{"time":"2025-11-25T23:32:11.281827902-03:00","level":"INFO","msg":"client <- server: LogMessage","message":"2025/11/25 23:32:11 go/packages.Load #1\n\tview_id=\"1\"\n\tsnapshot=0\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tquery=[/home/alvarolm/temp/templnopackagemetadata/... builtin]\n\tpackages=2\n\tduration=129.04397ms\n"}
{"time":"2025-11-25T23:32:11.311840481-03:00","level":"INFO","msg":"client <- server: Progress"}
{"time":"2025-11-25T23:32:11.31284017-03:00","level":"INFO","msg":"client <- server: RegisterCapability"}
{"time":"2025-11-25T23:32:11.513952156-03:00","level":"INFO","msg":"client <- server: RegisterCapability"}
{"time":"2025-11-25T23:32:11.517803871-03:00","level":"INFO","msg":"client -> server: DocumentSymbol end"}
{"time":"2025-11-25T23:32:11.517910041-03:00","level":"INFO","msg":"client -> server: CodeAction","params":{"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"context":{"diagnostics":[]},"range":{"start":{"line":2,"character":14},"end":{"line":2,"character":14}}}}
{"time":"2025-11-25T23:32:11.518636785-03:00","level":"INFO","msg":"client <- server: PublishDiagnostics"}
{"time":"2025-11-25T23:32:11.539163025-03:00","level":"INFO","msg":"client -> server: CodeAction end"}
{"time":"2025-11-25T23:32:11.539275962-03:00","level":"INFO","msg":"client -> server: DocumentLink","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:11.539289394-03:00","level":"INFO","msg":"client -> server: DocumentLink end"}
{"time":"2025-11-25T23:32:11.539327892-03:00","level":"INFO","msg":"client -> server: DocumentSymbol"}
{"time":"2025-11-25T23:32:11.53971459-03:00","level":"INFO","msg":"client -> server: DocumentSymbol end"}
{"time":"2025-11-25T23:32:11.539806265-03:00","level":"INFO","msg":"client -> server: SetTrace"}
{"time":"2025-11-25T23:32:11.539828432-03:00","level":"INFO","msg":"client -> server: SetTrace end"}
{"time":"2025-11-25T23:32:11.682763347-03:00","level":"INFO","msg":"client -> server: CodeLens"}
{"time":"2025-11-25T23:32:11.683590505-03:00","level":"INFO","msg":"client -> server: CodeLens end"}
{"time":"2025-11-25T23:32:11.948023979-03:00","level":"INFO","msg":"client -> server: CodeAction","params":{"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"context":{"diagnostics":[]},"range":{"start":{"line":2,"character":14},"end":{"line":2,"character":14}}}}
{"time":"2025-11-25T23:32:11.949024432-03:00","level":"INFO","msg":"client -> server: CodeAction end"}
{"time":"2025-11-25T23:32:11.949095445-03:00","level":"INFO","msg":"client -> server: CodeAction","params":{"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"context":{"diagnostics":[]},"range":{"start":{"line":2,"character":14},"end":{"line":2,"character":14}}}}
{"time":"2025-11-25T23:32:11.950024203-03:00","level":"INFO","msg":"client -> server: CodeAction end"}
{"time":"2025-11-25T23:32:11.950084225-03:00","level":"INFO","msg":"client -> server: CodeAction","params":{"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"context":{"diagnostics":[]},"range":{"start":{"line":2,"character":14},"end":{"line":2,"character":14}}}}
{"time":"2025-11-25T23:32:11.950787641-03:00","level":"INFO","msg":"client -> server: CodeAction end"}
{"time":"2025-11-25T23:32:12.876698019-03:00","level":"INFO","msg":"client -> server: DocumentLink","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"}
{"time":"2025-11-25T23:32:12.876728413-03:00","level":"INFO","msg":"client -> server: DocumentLink end"}
{"time":"2025-11-25T23:32:13.117503084-03:00","level":"INFO","msg":"client -> server: CodeAction","params":{"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"context":{"diagnostics":[]},"range":{"start":{"line":2,"character":14},"end":{"line":2,"character":14}}}}
{"time":"2025-11-25T23:32:13.118291414-03:00","level":"INFO","msg":"client -> server: CodeAction end"}
{"time":"2025-11-25T23:32:19.003324568-03:00","level":"INFO","msg":"client -> server: PrepareRename"}
{"time":"2025-11-25T23:32:19.003362223-03:00","level":"INFO","msg":"updatePosition: found","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ","fromTempl":"2:14","toGo":"9:13"}
{"time":"2025-11-25T23:32:19.003870428-03:00","level":"WARN","msg":"go->templ: range not found in sourcemap","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}
{"time":"2025-11-25T23:32:19.003898609-03:00","level":"INFO","msg":"client -> server: PrepareRename end"}
{"time":"2025-11-25T23:32:19.021933908-03:00","level":"INFO","msg":"client -> server: Implementation"}
{"time":"2025-11-25T23:32:19.021963553-03:00","level":"INFO","msg":"updatePosition: found","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ","fromTempl":"2:12","toGo":"9:11"}
{"time":"2025-11-25T23:32:19.02231868-03:00","level":"INFO","msg":"client -> server: Implementation end"}
{"time":"2025-11-25T23:32:19.038381815-03:00","level":"INFO","msg":"client -> server: Definition"}
{"time":"2025-11-25T23:32:19.038414654-03:00","level":"INFO","msg":"updatePosition: found","uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ","fromTempl":"2:12","toGo":"9:11"}
{"time":"2025-11-25T23:32:19.038901693-03:00","level":"INFO","msg":"client -> server: Definition end"}
{"time":"2025-11-25T23:32:25.116366949-03:00","level":"INFO","msg":"client -> server: Rename"}
{"time":"2025-11-25T23:32:25.174300447-03:00","level":"INFO","msg":"client <- server: LogMessage","message":"2025/11/25 23:32:25 go/packages.Load #2\n\tview_id=\"1\"\n\tsnapshot=2\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tquery=[file=/home/alvarolm/temp/templnopackagemetadata/x.templ]\n\tpackages=0\n\tduration=57.196726ms\n"}
{"time":"2025-11-25T23:32:25.174327474-03:00","level":"INFO","msg":"client -> server: Rename end"}
{"time":"2025-11-25T23:32:25.367168278-03:00","level":"INFO","msg":"client -> server: CodeLens"}
{"time":"2025-11-25T23:32:25.368038287-03:00","level":"INFO","msg":"client -> server: CodeLens end"}
```

gopls.log:

```
[Trace - 23:32:10.675 PM] Sending request 'initialize - (1)'.
Params: {"processId":1127212,"clientInfo":{"name":"Visual Studio Code","version":"1.106.2"},"locale":"en","rootPath":"/home/alvarolm/temp/templnopackagemetadata","rootUri":"file:///home/alvarolm/temp/templnopackagemetadata","capabilities":{"workspace":{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"textOnlyTransactional","resourceOperations":["create","rename","delete"],"normalizesLineEndings":true,"changeAnnotationSupport":{"groupsOnLabel":true}},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26]},"tagSupport":{"valueSet":[1]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true,"semanticTokens":{"refreshSupport":true},"codeLens":{"refreshSupport":true},"fileOperations":{"dynamicRegistration":true,"didCreate":true,"willCreate":true,"didRename":true,"willRename":true,"didDelete":true,"willDelete":true}},"textDocument":{"synchronization":{"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true,"didSave":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["markdown","plaintext"],"deprecatedSupport":true,"preselectSupport":true,"tagSupport":{"valueSet":[1]},"insertReplaceSupport":true,"resolveSupport":{"properties":["documentation","detail","additionalTextEdits"]},"insertTextModeSupport":{"valueSet":[1,2]}},"completionItemKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["markdown","plaintext"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["markdown","plaintext"],"parameterInformation":{"labelOffsetSupport":true},"activeParameterSupport":true},"contextSupport":true},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26]},"hierarchicalDocumentSymbolSupport":true,"tagSupport":{"valueSet":[1]},"labelSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["","quickfix","refactor","refactor.extract","refactor.inline","refactor.rewrite","source","source.organizeImports"]}},"isPreferredSupport":true,"disabledSupport":true,"dataSupport":true,"resolveSupport":{"properties":["edit"]},"honorsChangeAnnotations":true},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true,"tooltipSupport":true},"colorProvider":{"dynamicRegistration":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"publishDiagnostics":{"relatedInformation":true,"tagSupport":{"valueSet":[1,2]},"codeDescriptionSupport":true,"dataSupport":true},"rename":{"dynamicRegistration":true,"prepareSupport":true,"prepareSupportDefaultBehavior":1,"honorsChangeAnnotations":true},"selectionRange":{"dynamicRegistration":true},"callHierarchy":{"dynamicRegistration":true},"semanticTokens":{"dynamicRegistration":true,"requests":{"range":true,"full":{"delta":true}},"tokenTypes":["namespace","type","class","enum","interface","struct","typeParameter","parameter","variable","property","enumMember","event","function","method","macro","keyword","modifier","comment","string","number","regexp","operator","decorator"],"tokenModifiers":["declaration","definition","readonly","static","deprecated","abstract","async","modification","documentation","defaultLibrary"],"formats":["relative"]},"linkedEditingRange":{"dynamicRegistration":true}},"window":{"workDoneProgress":true,"showMessage":{"messageActionItem":{"additionalPropertiesSupport":true}},"showDocument":{"support":true}},"general":{"regularExpressions":{"engine":"ECMAScript","version":"ES2020"},"markdown":{"parser":"marked","version":"1.1.0"}}},"trace":"off","workspaceFolders":[{"uri":"file:///home/alvarolm/temp/templnopackagemetadata","name":"templnopackagemetadata"}]}


[Trace - 23:32:10.677 PM] Received response 'initialize - (1)' in 1ms.
Result: {"capabilities":{"textDocumentSync":{"openClose":true,"change":2,"save":{}},"completionProvider":{"triggerCharacters":["."]},"hoverProvider":true,"signatureHelpProvider":{"triggerCharacters":["(",","]},"definitionProvider":true,"typeDefinitionProvider":true,"implementationProvider":true,"referencesProvider":true,"documentHighlightProvider":true,"documentSymbolProvider":true,"codeActionProvider":{"codeActionKinds":["gopls.doc.features","quickfix","refactor.extract.constant","refactor.extract.constant-all","refactor.extract.function","refactor.extract.method","refactor.extract.toNewFile","refactor.extract.variable","refactor.extract.variable-all","refactor.inline.call","refactor.inline.variable","refactor.rewrite.changeQuote","refactor.rewrite.fillStruct","refactor.rewrite.fillSwitch","refactor.rewrite.invertIf","refactor.rewrite.joinLines","refactor.rewrite.removeUnusedParam","refactor.rewrite.splitLines","source.assembly","source.doc","source.fixAll","source.freesymbols","source.organizeImports","source.splitPackage"],"resolveProvider":true},"codeLensProvider":{},"documentLinkProvider":{},"workspaceSymbolProvider":true,"documentFormattingProvider":true,"renameProvider":{"prepareProvider":true},"foldingRangeProvider":true,"selectionRangeProvider":true,"executeCommandProvider":{"commands":["gopls.add_dependency","gopls.add_import","gopls.add_telemetry_counters","gopls.add_test","gopls.apply_fix","gopls.assembly","gopls.change_signature","gopls.check_upgrades","gopls.client_open_url","gopls.diagnose_files","gopls.doc","gopls.edit_go_directive","gopls.extract_to_new_file","gopls.fetch_vulncheck_result","gopls.free_symbols","gopls.gc_details","gopls.generate","gopls.go_get_package","gopls.list_imports","gopls.list_known_packages","gopls.maybe_prompt_for_telemetry","gopls.mem_stats","gopls.modify_tags","gopls.modules","gopls.package_symbols","gopls.packages","gopls.regenerate_cgo","gopls.remove_dependency","gopls.reset_go_mod_diagnostics","gopls.run_go_work_command","gopls.run_govulncheck","gopls.run_tests","gopls.scan_imports","gopls.split_package","gopls.start_debugging","gopls.start_profile","gopls.stop_profile","gopls.tidy","gopls.update_go_sum","gopls.upgrade_dependency","gopls.vendor","gopls.views","gopls.vulncheck","gopls.workspace_stats"]},"callHierarchyProvider":true,"semanticTokensProvider":{"legend":{"tokenTypes":["namespace","type","typeParameter","parameter","variable","function","method","macro","keyword","comment","string","number","operator","label"],"tokenModifiers":["definition","readonly","defaultLibrary","array","bool","chan","format","interface","map","number","pointer","signature","slice","string","struct"]},"range":true,"full":true},"typeHierarchyProvider":true,"inlayHintProvider":{},"workspace":{"workspaceFolders":{"supported":true,"changeNotifications":"workspace/didChangeWorkspaceFolders"},"fileOperations":{"didCreate":{"filters":[{"scheme":"file","pattern":{"glob":"**/*.go"}}]}}}},"serverInfo":{"name":"gopls","version":"{\"GoVersion\":\"go1.25.4\",\"Path\":\"golang.org/x/tools/gopls\",\"Main\":{\"Path\":\"golang.org/x/tools/gopls\",\"Version\":\"v0.20.0\",\"Sum\":\"h1:fxOYZXKl6IsOTKIh6IgjDbIDHlr5btOtOUkrGOgFDB4=\"},\"Deps\":[{\"Path\":\"github.com/BurntSushi/toml\",\"Version\":\"v1.5.0\",\"Sum\":\"h1:W5quZX/G/csjUnuI8SUYlsHs9M38FC7znL0lIO+DvMg=\"},{\"Path\":\"github.com/fatih/camelcase\",\"Version\":\"v1.0.0\",\"Sum\":\"h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=\"},{\"Path\":\"github.com/fatih/gomodifytags\",\"Version\":\"v1.17.1-0.20250423142747-f3939df9aa3c\",\"Sum\":\"h1:dDSgAjoOMp8da3egfz0t2S+t8RGOpEmEXZubcGuc0Bg=\"},{\"Path\":\"github.com/fatih/structtag\",\"Version\":\"v1.2.0\",\"Sum\":\"h1:/OdNE99OxoI/PqaW/SuSK9uxxT3f/tcSZgon/ssNSx4=\"},{\"Path\":\"github.com/fsnotify/fsnotify\",\"Version\":\"v1.9.0\",\"Sum\":\"h1:2Ml+OJNzbYCTzsxtv8vKSFD9PbJjmhYF14k/jKC7S9k=\"},{\"Path\":\"github.com/google/go-cmp\",\"Version\":\"v0.7.0\",\"Sum\":\"h1:wk8382ETsv4JYUZwIsn6YpYiWiBsYLSJiTsyBybVuN8=\"},{\"Path\":\"golang.org/x/exp/typeparams\",\"Version\":\"v0.0.0-20250620022241-b7579e27df2b\",\"Sum\":\"h1:KdrhdYPDUvJTvrDK9gdjfFd6JTk8vA1WJoldYSi0kHo=\"},{\"Path\":\"golang.org/x/mod\",\"Version\":\"v0.26.0\",\"Sum\":\"h1:EGMPT//Ezu+ylkCijjPc+f4Aih7sZvaAr+O3EHBxvZg=\"},{\"Path\":\"golang.org/x/sync\",\"Version\":\"v0.16.0\",\"Sum\":\"h1:ycBJEhp9p4vXvUZNszeOq0kGTPghopOL8q0fq3vstxw=\"},{\"Path\":\"golang.org/x/sys\",\"Version\":\"v0.34.0\",\"Sum\":\"h1:H5Y5sJ2L2JRdyv7ROF1he/lPdvFsd0mJHFw2ThKHxLA=\"},{\"Path\":\"golang.org/x/telemetry\",\"Version\":\"v0.0.0-20250710130107-8d8967aff50b\",\"Sum\":\"h1:DU+gwOBXU+6bO0sEyO7o/NeMlxZxCZEvI7v+J4a1zRQ=\"},{\"Path\":\"golang.org/x/text\",\"Version\":\"v0.27.0\",\"Sum\":\"h1:4fGWRpyh641NLlecmyl4LOe6yDdfaYNrGb2zdfo4JV4=\"},{\"Path\":\"golang.org/x/tools\",\"Version\":\"v0.35.1-0.20250728180453-01a3475a31bc\",\"Sum\":\"h1:ZRKyKRJl/YEWl9ScZwd6Ua6xSt7DE6tHp1I3ucMroGM=\"},{\"Path\":\"golang.org/x/vuln\",\"Version\":\"v1.1.4\",\"Sum\":\"h1:Ju8QsuyhX3Hk8ma3CesTbO8vfJD9EvUBgHvkxHBzj0I=\"},{\"Path\":\"honnef.co/go/tools\",\"Version\":\"v0.7.0-0.dev.0.20250523013057-bbc2f4dd71ea\",\"Sum\":\"h1:fj8r9irJSpolAGUdZBxJIRY3lLc4jH2Dt4lwnWyWwpw=\"},{\"Path\":\"mvdan.cc/gofumpt\",\"Version\":\"v0.8.0\",\"Sum\":\"h1:nZUCeC2ViFaerTcYKstMmfysj6uhQrA2vJe+2vwGU6k=\"},{\"Path\":\"mvdan.cc/xurls/v2\",\"Version\":\"v2.6.0\",\"Sum\":\"h1:3NTZpeTxYVWNSokW3MKeyVkz/j7uYXYiMtXRUfmjbgI=\"}],\"Settings\":[{\"Key\":\"-buildmode\",\"Value\":\"exe\"},{\"Key\":\"-compiler\",\"Value\":\"gc\"},{\"Key\":\"DefaultGODEBUG\",\"Value\":\"containermaxprocs=0,decoratemappings=0,tlssha1=1,updatemaxprocs=0,x509sha256skid=0\"},{\"Key\":\"CGO_ENABLED\",\"Value\":\"1\"},{\"Key\":\"CGO_CFLAGS\"},{\"Key\":\"CGO_CPPFLAGS\"},{\"Key\":\"CGO_CXXFLAGS\"},{\"Key\":\"CGO_LDFLAGS\"},{\"Key\":\"GOARCH\",\"Value\":\"amd64\"},{\"Key\":\"GOOS\",\"Value\":\"linux\"},{\"Key\":\"GOAMD64\",\"Value\":\"v1\"}],\"Version\":\"v0.20.0\"}"}}


[Trace - 23:32:10.700 PM] Sending notification 'initialized'.
Params: {}


[Trace - 23:32:10.700 PM] Sending notification 'textDocument/didOpen'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","languageId":"go","version":1,"text":"// Code generated by templ - DO NOT EDIT.\n\npackage main\n\n//lint:file-ignore SA4006 This context is only used if a nested component is present.\n\nimport \"github.com/a-h/templ\"\nimport templruntime \"github.com/a-h/templ/runtime\"\n\nfunc Greet(name string) templ.Component {\n\treturn templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {\n\t\ttempl_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context\n\t\tif templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {\n\t\t\treturn templ_7745c5c3_CtxErr\t\t}\n\t\ttempl_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)\n\t\tif !templ_7745c5c3_IsBuffer {\n\t\t\tdefer func() {\n\t\t\t\ttempl_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)\n\t\t\t\tif templ_7745c5c3_Err == nil {\n\t\t\t\t\ttempl_7745c5c3_Err = templ_7745c5c3_BufErr\n\t\t\t\t}\n\t\t\t}()\n\t\t}\n\t\tctx = templ.InitializeContext(ctx)\n\t\ttempl_7745c5c3_Var1 := templ.GetChildren(ctx)\n\t\tif templ_7745c5c3_Var1 == nil {\n\t\t\ttempl_7745c5c3_Var1 = templ.NopComponent\n\t\t}\n\t\tctx = templ.ClearChildren(ctx)\n\t\ttempl_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, \"\u003cdiv\u003eHello, \")\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\tvar templ_7745c5c3_Var2 string\n\t\ttempl_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(name)\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn\ttempl.Error{Err: templ_7745c5c3_Err, FileName: ``, Line: 4, Col: 19}\n\t\t}\n\t\t_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\ttempl_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, \"\u003c/div\u003e\")\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\treturn nil\n\t})\n}\nvar _ = templruntime.GeneratedTemplate"}}


[Trace - 23:32:10.700 PM] Received request 'window/workDoneProgress/create - (1)'.
Params: {"token":"1205925143784098141"}


[Trace - 23:32:10.715 PM] Sending notification 'textDocument/didOpen'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","languageId":"templ","version":1,"text":"// Code generated by templ - DO NOT EDIT.\n\npackage main\n\n//lint:file-ignore SA4006 This context is only used if a nested component is present.\n\nimport \"github.com/a-h/templ\"\nimport templruntime \"github.com/a-h/templ/runtime\"\n\nfunc Greet(name string) templ.Component {\n\treturn templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {\n\t\ttempl_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context\n\t\tif templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {\n\t\t\treturn templ_7745c5c3_CtxErr\t\t}\n\t\ttempl_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)\n\t\tif !templ_7745c5c3_IsBuffer {\n\t\t\tdefer func() {\n\t\t\t\ttempl_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)\n\t\t\t\tif templ_7745c5c3_Err == nil {\n\t\t\t\t\ttempl_7745c5c3_Err = templ_7745c5c3_BufErr\n\t\t\t\t}\n\t\t\t}()\n\t\t}\n\t\tctx = templ.InitializeContext(ctx)\n\t\ttempl_7745c5c3_Var1 := templ.GetChildren(ctx)\n\t\tif templ_7745c5c3_Var1 == nil {\n\t\t\ttempl_7745c5c3_Var1 = templ.NopComponent\n\t\t}\n\t\tctx = templ.ClearChildren(ctx)\n\t\ttempl_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, \"\u003cdiv\u003eHello, \")\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\tvar templ_7745c5c3_Var2 string\n\t\ttempl_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(name)\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn\ttempl.Error{Err: templ_7745c5c3_Err, FileName: ``, Line: 4, Col: 19}\n\t\t}\n\t\t_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\ttempl_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, \"\u003c/div\u003e\")\n\t\tif templ_7745c5c3_Err != nil {\n\t\t\treturn templ_7745c5c3_Err\n\t\t}\n\t\treturn nil\n\t})\n}\nvar _ = templruntime.GeneratedTemplate"}}


[Trace - 23:32:11.099 PM] Sending response 'window/workDoneProgress/create - (1)' in 399ms.
Result: 


[Trace - 23:32:11.100 PM] Sending request 'textDocument/documentSymbol - (2)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"}}


[Trace - 23:32:11.100 PM] Received notification '$/progress'.
Params: {"token":"1205925143784098141","value":{"kind":"begin","title":"Setting up workspace","message":"Loading packages..."}}


[Trace - 23:32:11.100 PM] Received request 'workspace/configuration - (2)'.
Params: {"items":[{"scopeUri":"file:///home/alvarolm/temp/templnopackagemetadata","section":"gopls"}]}


[Trace - 23:32:11.140 PM] Sending response 'workspace/configuration - (2)' in 40ms.
Result: [{}]


[Trace - 23:32:11.151 PM] Received notification 'window/logMessage'.
Params: {"type":3,"message":"2025/11/25 23:32:11 Created View (#1)\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tview_type=\"GoMod\"\n\troot_dir=\"file:///home/alvarolm/temp/templnopackagemetadata\"\n\tgo_version=\"go version go1.25.4 linux/amd64\"\n\tbuild_flags=[]\n\tenv={GOOS:linux GOARCH:amd64 GOCACHE:/home/alvarolm/.cache/go-build GOMODCACHE:/home/alvarolm/go/pkg/mod GOPATH:/home/alvarolm/go GOPRIVATE: GOFLAGS: GO111MODULE: GOTOOLCHAIN:auto GOROOT:/usr/local/go GoVersion:25 GoVersionOutput:go version go1.25.4 linux/amd64\n ExplicitGOWORK: EffectiveGOPACKAGESDRIVER:}\n\tenv_overlay=[]\n"}


[Trace - 23:32:11.281 PM] Received notification 'window/logMessage'.
Params: {"type":3,"message":"2025/11/25 23:32:11 go/packages.Load #1\n\tview_id=\"1\"\n\tsnapshot=0\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tquery=[/home/alvarolm/temp/templnopackagemetadata/... builtin]\n\tpackages=2\n\tduration=129.04397ms\n"}


[Trace - 23:32:11.311 PM] Received notification '$/progress'.
Params: {"token":"1205925143784098141","value":{"kind":"end","message":"Finished loading packages."}}


[Trace - 23:32:11.312 PM] Received request 'client/registerCapability - (3)'.
Params: {"registrations":[{"id":"workspace/didChangeWatchedFiles-0","method":"workspace/didChangeWatchedFiles","registerOptions":{"watchers":[{"globPattern":"**/*.{mod,work}","kind":7},{"globPattern":"/home/alvarolm/temp/templnopackagemetadata","kind":7},{"globPattern":"/home/alvarolm/temp/templnopackagemetadata/**/*.{go,mod,sum,work}","kind":7}]}}]}


[Trace - 23:32:11.513 PM] Sending response 'client/registerCapability - (3)' in 201ms.
Result: 


[Trace - 23:32:11.513 PM] Received request 'client/registerCapability - (4)'.
Params: {"registrations":[{"id":"workspace/didChangeConfiguration","method":"workspace/didChangeConfiguration"}]}


[Trace - 23:32:11.514 PM] Sending response 'client/registerCapability - (4)' in 1ms.
Result: 


[Trace - 23:32:11.517 PM] Received response 'textDocument/documentSymbol - (2)' in 417ms.
Result: [{"name":"Greet","detail":"func(name string) templ.Component","kind":12,"range":{"start":{"line":9,"character":0},"end":{"line":48,"character":1}},"selectionRange":{"start":{"line":9,"character":5},"end":{"line":9,"character":10}}}]


[Trace - 23:32:11.517 PM] Sending request 'textDocument/codeAction - (3)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"context":{"diagnostics":[]},"range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}}


[Trace - 23:32:11.518 PM] Received notification 'textDocument/publishDiagnostics'.
Params: {"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","version":1,"diagnostics":[]}


[Trace - 23:32:11.538 PM] Received response 'textDocument/codeAction - (3)' in 20ms.
Result: [{"title":"Organize Imports","kind":"source.organizeImports","edit":{"documentChanges":[{"textDocument":{"version":1,"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"edits":[{"range":{"start":{"line":6,"character":7},"end":{"line":6,"character":7}},"newText":"(\n\t"},{"range":{"start":{"line":7,"character":0},"end":{"line":7,"character":7}},"newText":"\t"},{"range":{"start":{"line":7,"character":50},"end":{"line":7,"character":50}},"newText":"\n)"}]}]}},{"title":"Browse amd64 assembly for Greet","kind":"source.assembly","command":{"title":"Browse amd64 assembly for Greet","command":"gopls.assembly","arguments":["1","pkg0","main.Greet"]}},{"title":"Browse documentation for package main","kind":"source.doc","command":{"title":"Browse documentation for package main","command":"gopls.doc","arguments":[{"Location":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}},"ShowDocument":true}]}},{"title":"Split package \"main\"","kind":"source.splitPackage","command":{"title":"Split package \"main\"","command":"gopls.split_package","arguments":["1","pkg0"]}},{"title":"Show compiler optimization details for \"templnopackagemetadata\"","kind":"source.toggleCompilerOptDetails","command":{"title":"Show compiler optimization details for \"templnopackagemetadata\"","command":"gopls.gc_details","arguments":["file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"]}},{"title":"Browse gopls feature documentation","kind":"gopls.doc.features","command":{"title":"Browse gopls feature documentation","command":"gopls.client_open_url","arguments":["https://github.com/golang/tools/blob/master/gopls/doc/features/index.md"]}}]


[Trace - 23:32:11.539 PM] Sending request 'textDocument/documentSymbol - (4)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"}}


[Trace - 23:32:11.539 PM] Received response 'textDocument/documentSymbol - (4)' in 0ms.
Result: [{"name":"Greet","detail":"func(name string) templ.Component","kind":12,"range":{"start":{"line":9,"character":0},"end":{"line":48,"character":1}},"selectionRange":{"start":{"line":9,"character":5},"end":{"line":9,"character":10}}}]


[Trace - 23:32:11.539 PM] Sending notification '$/setTrace'.
Params: {"value":"off"}


[Trace - 23:32:11.683 PM] Sending request 'textDocument/codeLens - (5)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"}}


[Trace - 23:32:11.683 PM] Received response 'textDocument/codeLens - (5)' in 0ms.
Result: null


[Trace - 23:32:11.948 PM] Sending request 'textDocument/codeAction - (6)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"context":{"diagnostics":[]},"range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}}


[Trace - 23:32:11.948 PM] Received response 'textDocument/codeAction - (6)' in 0ms.
Result: [{"title":"Organize Imports","kind":"source.organizeImports","edit":{"documentChanges":[{"textDocument":{"version":1,"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"edits":[{"range":{"start":{"line":6,"character":7},"end":{"line":6,"character":7}},"newText":"(\n\t"},{"range":{"start":{"line":7,"character":0},"end":{"line":7,"character":7}},"newText":"\t"},{"range":{"start":{"line":7,"character":50},"end":{"line":7,"character":50}},"newText":"\n)"}]}]}},{"title":"Browse amd64 assembly for Greet","kind":"source.assembly","command":{"title":"Browse amd64 assembly for Greet","command":"gopls.assembly","arguments":["1","pkg0","main.Greet"]}},{"title":"Browse documentation for package main","kind":"source.doc","command":{"title":"Browse documentation for package main","command":"gopls.doc","arguments":[{"Location":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}},"ShowDocument":true}]}},{"title":"Split package \"main\"","kind":"source.splitPackage","command":{"title":"Split package \"main\"","command":"gopls.split_package","arguments":["1","pkg0"]}},{"title":"Show compiler optimization details for \"templnopackagemetadata\"","kind":"source.toggleCompilerOptDetails","command":{"title":"Show compiler optimization details for \"templnopackagemetadata\"","command":"gopls.gc_details","arguments":["file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"]}},{"title":"Browse gopls feature documentation","kind":"gopls.doc.features","command":{"title":"Browse gopls feature documentation","command":"gopls.client_open_url","arguments":["https://github.com/golang/tools/blob/master/gopls/doc/features/index.md"]}}]


[Trace - 23:32:11.949 PM] Sending request 'textDocument/codeAction - (7)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"context":{"diagnostics":[]},"range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}}


[Trace - 23:32:11.949 PM] Received response 'textDocument/codeAction - (7)' in 0ms.
Result: [{"title":"Organize Imports","kind":"source.organizeImports","edit":{"documentChanges":[{"textDocument":{"version":1,"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"edits":[{"range":{"start":{"line":6,"character":7},"end":{"line":6,"character":7}},"newText":"(\n\t"},{"range":{"start":{"line":7,"character":0},"end":{"line":7,"character":7}},"newText":"\t"},{"range":{"start":{"line":7,"character":50},"end":{"line":7,"character":50}},"newText":"\n)"}]}]}},{"title":"Browse amd64 assembly for Greet","kind":"source.assembly","command":{"title":"Browse amd64 assembly for Greet","command":"gopls.assembly","arguments":["1","pkg0","main.Greet"]}},{"title":"Browse documentation for package main","kind":"source.doc","command":{"title":"Browse documentation for package main","command":"gopls.doc","arguments":[{"Location":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}},"ShowDocument":true}]}},{"title":"Split package \"main\"","kind":"source.splitPackage","command":{"title":"Split package \"main\"","command":"gopls.split_package","arguments":["1","pkg0"]}},{"title":"Show compiler optimization details for \"templnopackagemetadata\"","kind":"source.toggleCompilerOptDetails","command":{"title":"Show compiler optimization details for \"templnopackagemetadata\"","command":"gopls.gc_details","arguments":["file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"]}},{"title":"Browse gopls feature documentation","kind":"gopls.doc.features","command":{"title":"Browse gopls feature documentation","command":"gopls.client_open_url","arguments":["https://github.com/golang/tools/blob/master/gopls/doc/features/index.md"]}}]


[Trace - 23:32:11.950 PM] Sending request 'textDocument/codeAction - (8)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"context":{"diagnostics":[]},"range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}}


[Trace - 23:32:11.950 PM] Received response 'textDocument/codeAction - (8)' in 0ms.
Result: [{"title":"Organize Imports","kind":"source.organizeImports","edit":{"documentChanges":[{"textDocument":{"version":1,"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"edits":[{"range":{"start":{"line":6,"character":7},"end":{"line":6,"character":7}},"newText":"(\n\t"},{"range":{"start":{"line":7,"character":0},"end":{"line":7,"character":7}},"newText":"\t"},{"range":{"start":{"line":7,"character":50},"end":{"line":7,"character":50}},"newText":"\n)"}]}]}},{"title":"Browse amd64 assembly for Greet","kind":"source.assembly","command":{"title":"Browse amd64 assembly for Greet","command":"gopls.assembly","arguments":["1","pkg0","main.Greet"]}},{"title":"Browse documentation for package main","kind":"source.doc","command":{"title":"Browse documentation for package main","command":"gopls.doc","arguments":[{"Location":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}},"ShowDocument":true}]}},{"title":"Split package \"main\"","kind":"source.splitPackage","command":{"title":"Split package \"main\"","command":"gopls.split_package","arguments":["1","pkg0"]}},{"title":"Show compiler optimization details for \"templnopackagemetadata\"","kind":"source.toggleCompilerOptDetails","command":{"title":"Show compiler optimization details for \"templnopackagemetadata\"","command":"gopls.gc_details","arguments":["file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"]}},{"title":"Browse gopls feature documentation","kind":"gopls.doc.features","command":{"title":"Browse gopls feature documentation","command":"gopls.client_open_url","arguments":["https://github.com/golang/tools/blob/master/gopls/doc/features/index.md"]}}]


[Trace - 23:32:13.117 PM] Sending request 'textDocument/codeAction - (9)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"context":{"diagnostics":[]},"range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}}


[Trace - 23:32:13.118 PM] Received response 'textDocument/codeAction - (9)' in 0ms.
Result: [{"title":"Organize Imports","kind":"source.organizeImports","edit":{"documentChanges":[{"textDocument":{"version":1,"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"edits":[{"range":{"start":{"line":6,"character":7},"end":{"line":6,"character":7}},"newText":"(\n\t"},{"range":{"start":{"line":7,"character":0},"end":{"line":7,"character":7}},"newText":"\t"},{"range":{"start":{"line":7,"character":50},"end":{"line":7,"character":50}},"newText":"\n)"}]}]}},{"title":"Browse amd64 assembly for Greet","kind":"source.assembly","command":{"title":"Browse amd64 assembly for Greet","command":"gopls.assembly","arguments":["1","pkg0","main.Greet"]}},{"title":"Browse documentation for package main","kind":"source.doc","command":{"title":"Browse documentation for package main","command":"gopls.doc","arguments":[{"Location":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":13},"end":{"line":9,"character":13}}},"ShowDocument":true}]}},{"title":"Split package \"main\"","kind":"source.splitPackage","command":{"title":"Split package \"main\"","command":"gopls.split_package","arguments":["1","pkg0"]}},{"title":"Show compiler optimization details for \"templnopackagemetadata\"","kind":"source.toggleCompilerOptDetails","command":{"title":"Show compiler optimization details for \"templnopackagemetadata\"","command":"gopls.gc_details","arguments":["file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"]}},{"title":"Browse gopls feature documentation","kind":"gopls.doc.features","command":{"title":"Browse gopls feature documentation","command":"gopls.client_open_url","arguments":["https://github.com/golang/tools/blob/master/gopls/doc/features/index.md"]}}]


[Trace - 23:32:19.003 PM] Sending request 'textDocument/prepareRename - (10)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"position":{"line":9,"character":13}}


[Trace - 23:32:19.003 PM] Received response 'textDocument/prepareRename - (10)' in 0ms.
Result: {"range":{"start":{"line":9,"character":11},"end":{"line":9,"character":15}},"placeholder":"name"}


[Trace - 23:32:19.022 PM] Sending request 'textDocument/implementation - (11)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"position":{"line":9,"character":11}}


[Error - Received] 23:32:19.022 PM #11 name is a var, not a type


[Trace - 23:32:19.038 PM] Sending request 'textDocument/definition - (12)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"},"position":{"line":9,"character":11}}


[Trace - 23:32:19.038 PM] Received response 'textDocument/definition - (12)' in 0ms.
Result: [{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go","range":{"start":{"line":9,"character":11},"end":{"line":9,"character":15}}}]


[Trace - 23:32:25.116 PM] Sending request 'textDocument/rename - (13)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x.templ"},"position":{"line":2,"character":14},"newName":"name2"}


[Error - Received] 23:32:25.174 PM #13 no package metadata for file file:///home/alvarolm/temp/templnopackagemetadata/x.templ


[Trace - 23:32:25.174 PM] Received notification 'window/logMessage'.
Params: {"type":3,"message":"2025/11/25 23:32:25 go/packages.Load #2\n\tview_id=\"1\"\n\tsnapshot=2\n\tdirectory=/home/alvarolm/temp/templnopackagemetadata\n\tquery=[file=/home/alvarolm/temp/templnopackagemetadata/x.templ]\n\tpackages=0\n\tduration=57.196726ms\n"}


[Trace - 23:32:25.367 PM] Sending request 'textDocument/codeLens - (14)'.
Params: {"textDocument":{"uri":"file:///home/alvarolm/temp/templnopackagemetadata/x_templ.go"}}


[Trace - 23:32:25.367 PM] Received response 'textDocument/codeLens - (14)' in 0ms.
Result: null
```
