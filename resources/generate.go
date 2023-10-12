//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/skin.png -output=./images/skin.go  -var=SkinPng
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/characters.png -output=./images/characters.go -var=CharactersPng
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/spritesheet.png -output=./images/spritesheet.go -var=SpritesheetPng
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=images -input=./images/walls.png -output=./images/walls.go -var=WallsPng
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@v1.0.0 -package=fonts -input=./fonts/arcade-n.ttf -output=./fonts/arcade-n.go -var=ArcadeTTF
//go:generate gofmt -s -w .

// Resources package contains font, image resources needed by the game
package resources
