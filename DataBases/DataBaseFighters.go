package DataBases

import (
	"math/rand"
	"time"
)

func GetFighterDb() [104]string {
	return fightersDb
}

var fightersDb = [104]string{
	"Ryu",
	"Ken",
	"Chun-Li",
	"Guile",
	"Zangief",
	"Dhalsim",
	"Blanka",
	"E. Honda",
	"Balrog",
	"Vega",
	"Sagat",
	"M. Bison",
	"Sub-Zero",
	"Scorpion",
	"Reptile",
	"Kitana",
	"Jax",
	"Sonya",
	"Kano",
	"Johnny Cage",
	"Raiden",
	"Liu Kang",
	"Shang Tsung",
	"Kung Lao",
	"Smoke",
	"Kabal",
	"Nightwolf",
	"Stryker",
	"Shao Kahn",
	"Motaro",
	"Kintaro",
	"Baraka",
	"Jade",
	"Skarlet",
	"Rain",
	"Kotal Kahn",
	"Erron Black",
	"Reptile",
	"Reiko",
	"Quan Chi",
	"Shinnok",
	"Bo Rai Cho",
	"Kenshi",
	"Takeda",
	"Kung Jin",
	"Jacqui Briggs",
	"Cassie Cage",
	"Kotal Kahn",
	"Erron Black",
	"Rashid",
	"Karin",
	"Kazuya Mishima",
	"Paul Phoenix",
	"Marshall Law",
	"King",
	"Yoshimitsu",
	"Nina Williams",
	"Anna Williams",
	"Lei Wulong",
	"Jin Kazama",
	"Xiaoyu",
	"Hwoarang",
	"Steve Fox",
	"Kuma",
	"Panda",
	"Jack",
	"Roger",
	"Alex",
	"Christie Monteiro",
	"Lee Chaolan",
	"Terry Bogard",
	"Andy Bogard",
	"Joe Higashi",
	"Geese Howard",
	"Kim Kaphwan",
	"Mai Shiranui",
	"Blue Mary",
	"Ryo Sakazaki",
	"Robert Garcia",
	"Yuri Sakazaki",
	"King",
	"Chang Koehan",
	"Choi Bounge",
	"Jin Chonrei",
	"Jin Chonshu",
	"Krauser",
	"Laurence Blood",
	"Tung Fu Rue",
	"Franco Bash",
	"Li Xiangfei",
	"Rick Strowd",
	"Ryuji Yamazaki",
	"Bob Wilson",
	"Hon Fu",
	"Blue Mary",
	"Bayonetta",
	"Bowser",
	"Bowser Jr.",
	"Captain Falcon",
	"Chrom",
	"Cloud",
	"Corrin",
	"Daisy",
	"Dark Pit",
}

func ReadFighterDb(wait int) string {
	//Simulate reading from a database
	time.Sleep(time.Duration(wait) * time.Second)
	return RandomSelectFighter()
}

func RandomSelectFighter() string {
	randomIndex := rand.Intn(len(GetFighterDb()))
	fighter := GetFighterDb()[randomIndex]
	return fighter
}
