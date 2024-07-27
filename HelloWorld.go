package main

import (
	"awesomeProject/DataBases"
	"awesomeProject/FightingGame"
	"awesomeProject/GoBasics"
	"fmt"
	"sync"
)

func main() {
	//LESSON1: Execute Some Basic Go Code
	//GoBasics.PrintSection("LESSON1: Execute Some Basic Go Code")
	//GoBasics.ExecuteBasics()
	//
	//LESSON2: Execute Builder Pattern
	//GoBasics.PrintSection("LESSON2: Execute Builder Pattern")
	//pc, err := new(DesignPatterns.ComputerBuilder).
	//	SetRAM("16GB").
	//	SetCPU("i7").
	//	SetGPU("GTX 1080").
	//	SetStorage("1TB SSD").
	//	Build()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(pc.ToString())

	//LESSON3: Execute Builder Pattern
	//GoBasics.PrintSection("LESSON3: Concurrent Read Test")

	//GoBasics.Label("Concurrent Read Test")
	//GoBasics.ExecuteConcurrentReadTest()
	//
	//GoBasics.Label("Concurrent Write Test")
	//GoBasics.ExecuteConcurrentWriteTest()

	//GoBasics.Label("Concurrent Read Write/Test Fighting Game Roster")
	////startDevelopmentCycle()
	//
	//builder := FightingGame.Builder{}
	//builder.
	//	AddCharacter(FightingGame.Character{Name: "Ryu"}).
	//	AddCharacter(FightingGame.Character{Name: "Ken"})
	//
	//game1, _ := builder.
	//	SetTitle("SF6").
	//	Build()
	//game2, _ := builder.
	//	SetTitle("MK1").
	//	Build()
	//
	//fmt.Println(game1.ToString())
	//fmt.Println(game2.ToString())

	DataBases.SnowflakeFetch()

}

const rosterLimit = 25
const previewCut = 5

type DevelopmentProject struct {
	developer      string
	gameBuilder    FightingGame.IBuilder
	previewChannel chan FightingGame.FightingGame
	mutex          *sync.Mutex
	wg             *sync.WaitGroup
}

func developGame(devProject *DevelopmentProject) {
	defer devProject.wg.Done()
	for i := 0; i < rosterLimit; i++ {
		//Dedicate one developer for each character in the roster
		go developCharacter(devProject)
	}
}

func developCharacter(devProject *DevelopmentProject) {
	defer devProject.wg.Done()

	devProject.mutex.Lock()
	devProject.gameBuilder.AddCharacter(
		FightingGame.Character{
			Name: fmt.Sprintf("%s", DataBases.RandomSelectFighter()),
		},
	)
	devProject.mutex.Unlock()

	//Determine IF  A Preview Can Be Pushed To The Channel
	gameBuild, err := devProject.gameBuilder.Build()
	if err != nil {
		fmt.Println(err)
	}

	//If 5 Characters have been added to the roster, preview the game.
	var currentVersion = gameBuild.Version
	var charCnt = len(gameBuild.Roster)

	devProject.mutex.Lock()
	if charCnt != 0 && charCnt%previewCut == 0 && currentVersion == gameBuild.Version { //Check The Version. If It has changed that indicates a preview has been pushed.
		//Push The Game Build to the Channel
		devProject.gameBuilder.SetVersion(gameBuild.Version + 1)
		devProject.previewChannel <- gameBuild
	}
	devProject.mutex.Unlock()

}

type PreviewOutlet struct {
	outlet         string
	previewChannel chan FightingGame.FightingGame
	rwMutex        *sync.RWMutex
	wg             *sync.WaitGroup
}

func previewGame(outlet *PreviewOutlet) {
	defer outlet.wg.Done()
	defer close(outlet.previewChannel)
	for {
		gameBuild, ok := <-outlet.previewChannel
		if !ok {
			fmt.Println("No Futher Previews Available for ", outlet.outlet)
			break
		}
		outlet.rwMutex.RLock()
		fmt.Println(fmt.Sprintf("-=[PREVIEW (%v)]=-\n%v", outlet.outlet, gameBuild.ToString()))
		outlet.rwMutex.RUnlock()
	}
}

func startDevelopmentCycle() {

	var channel = make(chan FightingGame.FightingGame, 5)
	var wg = sync.WaitGroup{}
	var mutex = sync.Mutex{}
	var rwMutex = sync.RWMutex{}

	//=========================================================================================================
	//Development Projects
	//=========================================================================================================
	var developmentProjects = []DevelopmentProject{}

	//Tekken 8
	tekken8 := DevelopmentProject{
		developer: "Namco",
		gameBuilder: FightingGame.IBuilder(&FightingGame.Builder{}).
			SetTitle("Tekken 8"),
		previewChannel: channel,
		mutex:          &mutex,
		wg:             &wg,
	}
	developmentProjects = append(developmentProjects, tekken8)

	//Street Fighter 6
	sf6 := DevelopmentProject{
		developer: "Capcom",
		gameBuilder: FightingGame.IBuilder(&FightingGame.Builder{}).
			SetTitle("Street Fighter 6"),
		previewChannel: channel,
		mutex:          &mutex,
		wg:             &wg,
	}
	developmentProjects = append(developmentProjects, sf6)

	//Mortal Kombat 1
	mk1 := DevelopmentProject{
		developer: "Neither Realms Studios",
		gameBuilder: FightingGame.IBuilder(&FightingGame.Builder{}).
			SetTitle("Mortal Kombat 1"),
		previewChannel: channel,
		mutex:          &mutex,
		wg:             &wg,
	}
	developmentProjects = append(developmentProjects, mk1)

	//Fatal Fury
	fatalFury := DevelopmentProject{
		developer: "SNK",
		gameBuilder: FightingGame.IBuilder(&FightingGame.Builder{}).
			SetTitle("Fatal Fury"),
		previewChannel: channel,
		mutex:          &mutex,
		wg:             &wg,
	}
	developmentProjects = append(developmentProjects, fatalFury)

	//=========================================================================================================
	//Review Outlets
	//=========================================================================================================
	var previewOutlets = []PreviewOutlet{}

	//GameSpot
	gameSpot := PreviewOutlet{
		outlet:         "GameSpot",
		previewChannel: channel,
		rwMutex:        &rwMutex,
		wg:             &wg,
	}
	previewOutlets = append(previewOutlets, gameSpot)

	//IGN
	ign := PreviewOutlet{
		outlet:         "IGN",
		previewChannel: channel,
		rwMutex:        &rwMutex,
		wg:             &wg,
	}
	previewOutlets = append(previewOutlets, ign)

	//Game Informer
	gameInformer := PreviewOutlet{
		outlet:         "Game Informer",
		previewChannel: channel,
		rwMutex:        &rwMutex,
		wg:             &wg,
	}
	previewOutlets = append(previewOutlets, gameInformer)

	//PC Mag
	pcMag := PreviewOutlet{
		outlet:         "PC Mag",
		previewChannel: channel,
		rwMutex:        &rwMutex,
		wg:             &wg,
	}
	previewOutlets = append(previewOutlets, pcMag)

	//=========================================================================================================
	//Kick off Development/Preview Cycle
	//=========================================================================================================
	wg.Add((len(developmentProjects) * rosterLimit) + len(previewOutlets))

	for _, devProject := range developmentProjects {
		go developGame(&devProject)
	}

	for _, previewOutlet := range previewOutlets {
		go previewGame(&previewOutlet)
	}
	//Wait for all development projects to complete
	wg.Wait()

	GoBasics.PrintSection("Completed Games")
	for _, devProject := range developmentProjects {
		game, _ := devProject.gameBuilder.Build()
		fmt.Println(game.ToString())
	}
}
