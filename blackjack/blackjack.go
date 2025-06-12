package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
// | card  | value | card    | value |
// | :---: | :---: | :-----: | :---: |
// |  ace  |  11   | eight   |   8   |
// |  two  |   2   | nine    |   9   |
// | three |   3   |  ten    |  10   |
// | four  |   4   | jack    |  10   |
// | five  |   5   | queen   |  10   |
// |  six  |   6   | king    |  10   |
// | seven |   7   | *other* |   0   |
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
//   - If you have a pair of aces you must always split them.
//   - If you have a Blackjack (two cards that sum up to a value of 21),
//     and the dealer does not have an ace, a figure or a ten then you automatically win.
//     If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
//   - If your cards sum up to a value within the range [17, 20] you should always stand.
//   - If your cards sum up to a value within the range [12, 16] you should always stand
//     unless the dealer has a 7 or higher, in which case you should always hit.
//   - If your cards sum up to 11 or lower you should always hit.
func FirstTurn(card1, card2, dealerCard string) string {
	total := ParseCard(card1) + ParseCard(card2)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case total == 21:
		switch {
		case ParseCard(dealerCard) >= 10:
			return "S"
		default:
			return "W"
		}
	case total >= 17 && total <= 20:
		return "S"
	case total >= 12 && total <= 16:
		switch {
		case ParseCard(dealerCard) >= 7:
			return "H"
		default:
			return "S"
		}
	default:
		return "H"

	}

}
