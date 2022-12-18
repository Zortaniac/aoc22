package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type rock struct {
	width int
	shape [][]bool
}

func day17() {
	content, err := ioutil.ReadFile("day17.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	jets := strings.TrimSpace(string(content))

	rocks := []rock{
		{
			4,
			[][]bool{{true, true, true, true}},
		},
		{
			3,
			[][]bool{{false, true, false}, {true, true, true}, {false, true, false}},
		},
		{
			3,
			[][]bool{{false, false, true}, {false, false, true}, {true, true, true}},
		},
		{
			1,
			[][]bool{{true}, {true}, {true}, {true}},
		},
		{
			2,
			[][]bool{{true, true}, {true, true}},
		},
	}

	chamber := [][]bool{{true, true, true, true, true, true, true}}

	currentRock := 0
	currentJet := 0
	for j := 0; j < 2022; j++ {
		shapePos := 2
		for i := 0; i < 3; i++ {
			jet := jets[currentJet]
			currentJet = (currentJet + 1) % len(jets)
			if jet == '>' {
				if shapePos+rocks[currentRock].width < 7 {
					shapePos++
				}
			} else { // '<'
				if shapePos > 0 {
					shapePos--
				}
			}
		}

		// check

		shape := rocks[currentRock].shape
		depth := 0
		sink := true
		for sink {
			jet := jets[currentJet]
			currentJet = (currentJet + 1) % len(jets)
			direction := 1
			if jet == '<' {
				direction = -1
			}
			if shapePos+rocks[currentRock].width+direction <= 7 && shapePos+direction >= 0 {
				canMove := true
				for n := depth; n > 0; n-- {
					// test horizontal movement
					if len(shape)-1-(depth-n) < 0 {
						break
					}
					for i := 0; i < rocks[currentRock].width; i++ {
						if chamber[n-1][shapePos+direction+i] && shape[len(shape)-1-(depth-n)][i] {
							canMove = false
							break
						}
					}
					if !canMove {
						break
					}
				}
				if canMove {
					shapePos += direction
				}
			}
			for n := depth; n >= 0; n-- {
				if len(shape)-1-(depth-n) < 0 {
					continue
				}
				for i := shapePos; i < shapePos+rocks[currentRock].width; i++ {
					if shape[len(shape)-1-(depth-n)][i-shapePos] && chamber[n][i] {
						solidifyRock(shapePos, depth, &rocks[currentRock], &chamber)
						sink = false
						break
					}
				}
				if !sink {
					break
				}
			}
			if !sink {
				break
			}
			if false {
				// print

				snapShot := make([][]bool, Min(10, len(chamber)))
				for s := 0; s < len(snapShot); s++ {
					snapShot[s] = make([]bool, 7)
					copy(snapShot[s], chamber[s])
				}
				solidifyRock(shapePos, depth+1, &rocks[currentRock], &snapShot)
				printChamber(&snapShot, j)
				//time.Sleep(time.Second)
			}
			depth++
		}

		currentRock = (currentRock + 1) % len(rocks)
		//printChamber(&chamber, j)
		//time.Sleep(time.Second)
		if j == 2022 {
			fmt.Println(len(chamber) - 1)
		}
	}
	fmt.Println(len(chamber) - 1)
}

func printChamber(chamber *[][]bool, round int) {
	fmt.Println("\033[2J")
	fmt.Println("Round:", round)
	fmt.Println("=======")
	for i := 0; i < Min(10, len(*chamber)); i++ {
		fmt.Print("|")
		for n := 0; n < 7; n++ {
			if (*chamber)[i][n] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
}

func solidifyRock(pos int, depth int, rock *rock, chamber *[][]bool) {
	var paddingLeft []bool
	for i := 0; i < pos; i++ {
		paddingLeft = append(paddingLeft, false)
	}
	var paddingRight []bool
	for i := pos + rock.width; i < 7; i++ {
		paddingRight = append(paddingRight, false)
	}
	for i := len(rock.shape) - 1; i >= 0; i-- {
		if depth > 0 {
			for n := pos; n < (*rock).width+pos; n++ {
				(*chamber)[depth-1][n] = rock.shape[i][n-pos] || (*chamber)[depth-1][n]
			}
			depth--
		} else {
			currentLevel := make([]bool, len(paddingLeft))
			copy(currentLevel, paddingLeft)
			currentLevel = append(append(currentLevel, rock.shape[i]...), paddingRight...)
			test := [][]bool{currentLevel}
			*chamber = append(test, *chamber...)
			depth--
		}
	}
}
