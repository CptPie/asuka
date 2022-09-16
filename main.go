package main

import (
	"asuka/providers"
	"asuka/utils"
)

func main() {

	aniList := providers.NewAniListProvider()
	results, err := aniList.Search("gate", false)
	if err != nil {
		panic(err.Error())
		return
	}
	for _, anime := range results {
		utils.PrettyPrint(anime.Title)
	}

	//views.MainView().Exec()

}
