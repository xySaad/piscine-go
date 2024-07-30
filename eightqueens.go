package piscine

func EightQueens() {

	type Coordinates struct {
		x int
		y int
	}

	queensPositions := [8]Coordinates{
		{x: 1, y: 1},
	}

	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			println("Checking: ", x, y)

			for queens := 1; queens <= 8; queens++ {
				println("Watch with queen: ", queens)
				println("==> {", queensPositions[queens-1].x, ",", queensPositions[queens-1].y, "}")
				isSafe := true

				if queens != y {
					if queensPositions[queens-1].x == x || queensPositions[queens-1].y == y {
						isSafe = false
					}

					if isSafe {
						queensPositions[y-1].x = x
						queensPositions[y-1].y = y
						println("Assigned: ", x, y)
					}
				}

			}
		}
	}
	println("--------------\n")
	for i := range queensPositions {
		println("{", queensPositions[i].x, ",", queensPositions[i].y, "}")
	}
}
