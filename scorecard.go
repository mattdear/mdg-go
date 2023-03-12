package main

import "strconv"

type Scorecard struct {
	ones          int
	twos          int
	threes        int
	fours         int
	fives         int
	sixes         int
	topBonus      int
	topScore      int
	twoThreeMatch int
	threeMatch    int
	fourMatch     int
	fiveMatch     int
	threeLine     int
	fourLine      int
	extras        int
	bottomScore   int
	totalScore    int
}

func (scorecard *Scorecard) DisplayScorecard() string {

	output := "\n|      MDG Scorecard     |\n" +
		"| Top Section -----------|\n" +
		"| 1  | Ones          | " + strconv.Itoa(scorecard.ones) + " |\n" +
		"| 2  | Twos          | " + strconv.Itoa(scorecard.twos) + " |\n" +
		"| 3  | Threes        | " + strconv.Itoa(scorecard.threes) + " |\n" +
		"| 4  | Fours         | " + strconv.Itoa(scorecard.fours) + " |\n" +
		"| 5  | Fives         | " + strconv.Itoa(scorecard.fives) + " |\n" +
		"| 6  | Sixes         | " + strconv.Itoa(scorecard.sixes) + " |\n" +
		"| -  | Bonus         | " + strconv.Itoa(scorecard.topBonus) + " |\n" +
		"| Bottom Section --------|\n" +
		"| 7  | 2/3 Match     | " + strconv.Itoa(scorecard.twoThreeMatch) + " |\n" +
		"| 8  | 3 Match       | " + strconv.Itoa(scorecard.threeMatch) + " |\n" +
		"| 9  | 4 Match       | " + strconv.Itoa(scorecard.fourMatch) + " |\n" +
		"| 10 | 5 Match       | " + strconv.Itoa(scorecard.fiveMatch) + " |\n" +
		"| 11 | 3 Line        | " + strconv.Itoa(scorecard.threeLine) + " |\n" +
		"| 12 | 4 Line        | " + strconv.Itoa(scorecard.fourLine) + " |\n" +
		"| 13 | Extra         | " + strconv.Itoa(scorecard.extras) + " |\n" +
		"| Totals ----------------|\n" +
		"| -  | Top Score     | " + strconv.Itoa(scorecard.topScore) + " |\n" +
		"| -  | Bottom Score  | " + strconv.Itoa(scorecard.bottomScore) + " |\n" +
		"| -  | Total Score   | " + strconv.Itoa(scorecard.totalScore) + " |\n"
	return output
}

func (scorecard *Scorecard) calcTopScore() {
	scorecard.topScore = scorecard.ones + scorecard.twos + scorecard.threes + scorecard.fours + scorecard.fives + scorecard.sixes
	if scorecard.topScore >= 63 {
		scorecard.topBonus = 35
	}
	scorecard.topScore = scorecard.topScore + scorecard.topBonus
}

func (scorecard *Scorecard) calcBottomScore() {
	scorecard.bottomScore = scorecard.twoThreeMatch + scorecard.threeMatch + scorecard.fourMatch + scorecard.fiveMatch + scorecard.threeLine + scorecard.fourLine + scorecard.extras
}

func (scorecard *Scorecard) calcTotalScore() {
	scorecard.totalScore = scorecard.topScore + scorecard.bottomScore
}
