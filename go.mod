module luoxy.xyz/winmgr

go 1.16

require (
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	luoxy.xyz/winmgr/app => ./app
	luoxy.xyz/winmgr/common => ./common
	luoxy.xyz/winmgr/plugins => ./plugins
)
