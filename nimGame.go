package demo

func canWinNim(n int) bool {
	//just think about that: when n is less than 3;no matter how many stones the first remove from 1 to 3,the first can win;
	// when n=4; the first must lose; because no matter how many stones the first removed, the left must less than or equal to 3, so we can say who tackle with the last 4 stone must lose
	// when 5<=n<=7; the first can make the left to be 4 stones, so we can say the first who tackle with 5 to 7 of stones can win this game;
	// when n=8; no matter how many stones the first remove,the second must tackle the left 5 to 7 stones and the second can win this game;
	// thinking iterately, we can make a conclusion that when n mod 4 completely, the first must lose, or this guy can win.
	if n%4 == 0 {
		return false
	}
	return true
}
