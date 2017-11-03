package main

import (
	"fmt"
	"flag"
	"time"
	"math/rand"
	"image"
	"os"
	"image/png"
	"image/color"
	"strings"
)


// Deafult Values
const (
	MAX_X = 400
	LINES = 400
	RULE = 30
	WRAP = true
	SEED = 0
	VERSION = "0.5b"
)

type config struct {
	filename        string  // filename
	x               int	// population
	lines           int	// generations
	randomizestart  bool	// start with randomized set?
	rndborn		int	// percentage of a cell in randomized mode to be alive
	wrap_mode       bool	// Wrap around?
	wall_alive      bool	// Is there an always living cell behind the wall? (Only if Wrap is off)
	randomness_seed int	// Seed for randomness
	shift		int	// shift evey step x-wise
	rulenumber      uint8	// RuleNumber 0-255
	progress	bool	// Show Progress
	invert 		bool	// Invert first line
	genesis		string	// Genesis String e.g. "1001"
}

func (c *config)parseParameters() {
	rule := 0
	version := false

	flag.Usage = func() {
		fmt.Printf("\nThis Programm generates PNG-images of 1dimensional Cellular Automata. It can only generate the \n" +
			"first 256 Automata described by Stephen Wolfram. \n\n" +
			"Try following Options: \n" +
			"\"%s -pop 500 -generations 500 -rule 30\" for a 500x500 Images created by Rule 30  \n" +
			"\"%s -pop 500 -generations 500 -rule 73 -genesis 001000100\" for a different starting block  \n" +
			"\"%s -rnd -file xyz.png\" for a different image file  \n" +
			"(c) Alfred Eichenseher 2017\n" +
			"Have Fun\n\n", os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&c.filename, "file", "ca5.png", "Resulting PNG filename")
	flag.IntVar(&c.x, "pop", MAX_X, "Length (X-Axis)")
	flag.IntVar(&c.lines, "generations", LINES, "Genereations (Length on Y Axis)")
	flag.IntVar(&c.shift, "shift", 0, "shifts each new generation")
	flag.BoolVar(&c.randomizestart, "rnd", false, "Randomize first generation")
	flag.BoolVar(&c.wrap_mode, "wrap", WRAP, "Wrap edges around")
	flag.BoolVar(&c.wall_alive, "living walls", false, "Force Living cell behind the borders")
	flag.IntVar(&c.randomness_seed, "randseed", SEED, "Seed Randomness")
	flag.IntVar(&rule, "rule", RULE, "Rule for Genereting (1-256)")
	flag.BoolVar(&c.progress, "progress", false, "Show Progress")
	flag.StringVar(&c.genesis, "genesis", "010", "Genesis Block. \"010\" is the default block")
	flag.BoolVar(&c.invert, "invert", false, "Invert first Generation")
	flag.IntVar(&c.rndborn, "rnd_born", 50, "Percentage(0-100) if state in a randomized first generation is actice (Only works with rnd)")
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()
	c.rulenumber = (uint8(rule % 255))
	//fmt.Println("config: ", c)
	if version {
		fmt.Println("Version: ", VERSION)
	}
}

func init() {
	fmt.Println("Visualize 1D CAs. Run with '-h' to see flags. \n(Version 0.5a)\n(c) Alfred Eichenseher 2017")
	fmt.Println("Starting up ....")
}

func initializeGeneration(gen []int, c config) {

	genesisRaw := strings.Split(c.genesis,"")
	fmt.Printf("genesis: %v \n\n",len(genesisRaw))
	genesis := make([]int, len(genesisRaw), len(genesisRaw))

	for k,v :=range genesisRaw {
		if v == "1" {
			genesis[k] = 1
		} else {
			genesis[k] = 0
		}
	}

	fmt.Printf("genesis: %v\n",genesis)

	genStartPoint := (c.x / 2) - (len(genesis) / 2)
	genEndPoint := genStartPoint + len(genesis)

	getBorn := 0
	if c.randomness_seed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(int64(c.randomness_seed))
	}
	for i:=0;i<c.x;i++ {
		if c.randomizestart  {
			getBorn = rand.Intn(100) + 1
			//fmt.Println(i, getBorn)
			if getBorn < c.rndborn {
				gen[i] = 1
			}
		} else if i == (c.x/2) {
			gen[i] = 1
		}
		if len(genesis) > 1 {
			if i >= genStartPoint && i < genEndPoint {
				//fmt.Println(i,genStartPoint,i-genStartPoint)
				gen[i] = genesis[i-genStartPoint]
			}
		}
	}
	//fmt.Println(gen)

	if c.invert {
		for i := 0; i < c.x; i++ {
			gen[i] = (gen[i] + 1) % 2
		}
	}
}

func printgen(g []int, line int, image *image.RGBA, conf *config) {
	for i := 0; i < conf.x; i++ {
		if g[i] == 0 {
			//fmt.Print(".")
			image.Set(i, line, color.Gray{0})
		} else {
			//fmt.Print("O")
			image.Set(i, line, color.Gray{212})

		}
	}
	//fmt.Print(" \n")
}


func applyrule(one []int, two[]int, c *config) {
	var pos100, pos010, pos001 int
	var v_pos100, v_pos010, v_pos001 int
	var newborn int
	shift := c.shift
	for i := 0; i < c.x; i++ {
		newborn = 0
		pos010 = i
		pos100 = i - 1
		pos001 = i + 1

		if c.wrap_mode {
			if i == 0 {
				pos100 = c.x - 1
			}
			if i == (c.x - 1) {
				pos001 = 0
			}
			v_pos010 = one[pos010]
			v_pos100 = one[pos100]
			v_pos001 = one[pos001]
		} else {
			v_pos010 = one[pos010]
			if i == 0 {
				if c.wall_alive {
					v_pos100 = 1
				}
				v_pos001 = one[pos001]
			} else {
				v_pos100 = one[pos100]
			}
			if i == (c.x - 1) {
				v_pos100 = one[pos100]
				if c.wall_alive {
					v_pos001 = 1
				}
			} else {
				v_pos001 = one[pos001]
			}

		}

		if (c.rulenumber & 1 == 1) {
			if isActive(v_pos100,v_pos010,v_pos001,0,0,0) {
				newborn = 1
			}
		}
		if (c.rulenumber & 2 == 2) {
			if isActive(v_pos100, v_pos010, v_pos001, 0, 0, 1) {
				newborn = 1
			}
		}
		if (c.rulenumber & 4 == 4) {
			if isActive(v_pos100,v_pos010,v_pos001,0,1,0) {
				newborn = 1
			}

		}
		if (c.rulenumber & 8 == 8) {
			if isActive(v_pos100,v_pos010,v_pos001,0,1,1) {
				newborn = 1
			}

		}
		if (c.rulenumber & 16 == 16) {
			if isActive(v_pos100,v_pos010,v_pos001,1,0,0) {
				newborn = 1
			}

		}
		if (c.rulenumber & 32 == 32) {
			if isActive(v_pos100,v_pos010,v_pos001,1,0,1) {
				newborn = 1
			}

		}
		if (c.rulenumber & 64 == 64) {
			if isActive(v_pos100,v_pos010,v_pos001,1,1,0) {
				newborn = 1
			}

		}
		if (c.rulenumber & 128 == 128) {
			if isActive(v_pos100,v_pos010,v_pos001,1,1,1) {
				newborn = 1
			}

		}
		two[(c.x + i +shift) % c.x] = newborn
	}

}

func isActive(p1,p2,p3,a,b,c int) bool{
	if p1 == a && p2 == b && p3 == c {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	conf := new(config)
	conf.parseParameters()

	genOne := make([]int, conf.x, conf.x)
	genTwo := make([]int, conf.x, conf.x)

	initializeGeneration(genOne, *conf)

	myImage := image.NewRGBA(image.Rect(0, 0, conf.x, conf.lines))
	outputFile, err := os.Create(conf.filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer outputFile.Close()

	arr := [2][]int{0:genOne,1:genTwo}


	printgen(genOne, 0, myImage, conf)
	step := 0
	for i:=1;i< conf.lines; i++ {
		applyrule(arr[(i + 1) % 2], arr[i % 2], conf)
		printgen(arr[i % 2], i, myImage, conf)

		if conf.progress {
			if (i % int((conf.lines / 100)) ) == 0 {
				step++
				fmt.Printf("progress: %v %% \n", step)
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime elapsed: %v \n", elapsed)

	fmt.Println("Encoding File ... please wait.")
	error := png.Encode(outputFile, myImage)

	if error != nil {
		fmt.Println("ERR:", error)
	}
	fmt.Println("Generated Images saved as:",conf.filename)
	// Files gets closed by defer
}