package __5

type ATM struct {
	moneys []int
	values []int
}

func Constructor() ATM {
	s := ATM{
		moneys: make([]int, 5),
		values: make([]int, 5),
	}
	s.values[0] = 20
	s.values[1] = 50
	s.values[2] = 100
	s.values[3] = 200
	s.values[4] = 500
	return s
}

func (s *ATM) Deposit(banknotesCount []int) {
	for i, cnt := range banknotesCount {
		s.moneys[i] += cnt
	}
}

func (s *ATM) remove(moneys []int) {
	for i, cnt := range moneys {
		s.moneys[i] -= cnt
	}
}

func (s *ATM) Withdraw(amount int) []int {
	moneys := make([]int, 5)
	for i := 4; i >= 0; i-- {
		if s.moneys[i] == 0 {
			continue
		}

		if amount < s.values[i] {
			continue
		}

		cnt := amount / s.values[i]
		if s.moneys[i] > cnt {
			moneys[i] = cnt
		} else {
			moneys[i] = s.moneys[i]
		}

		amount -= s.values[i] * moneys[i]
		if amount == 0 {
			s.remove(moneys)
			return moneys
		}
	}
	return []int{-1}
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
