.PHONY: assets
assets :
	go-bindata -o assets/assets.go -pkg assets img/* 

.PHONY: mobile
mobile :
	env GO111MODULE=off ebitenmobile bind -target android -javapkg cafe.gophers.ebitengame -o ~/Desktop/splendidmobile.aar github.com/kyeett/splendid/splendidmobile