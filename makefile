android:
	export GOPROXY=https://proxy.golang.org,direct;unset GOSUMDB
	go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -androidapi 21 -javapkg com.mikelangelon.dutchrpg -o ./mobile/android/dutchrpg/dutchrpg.aar ./mobile
