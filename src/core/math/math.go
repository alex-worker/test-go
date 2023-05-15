package math

func CalcFPS(startTicks uint64, endTicks uint64) uint64 {
	deltaTicks := endTicks - startTicks
	if deltaTicks == 0 {
		return 0
	}
	fps := 1000 / deltaTicks
	return fps
}
