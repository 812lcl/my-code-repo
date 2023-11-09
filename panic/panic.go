package panic

func CrashNotRecover() {
	RaceCondition()
	MuxFixRaceCondition()
	ChanToFixRaceCondition()
	CondToFixRaceCondition()
	PanicInGoroutine()
	NilPanic()
	ExitPanic()
	GoExit()
	StackOverflow()
}
