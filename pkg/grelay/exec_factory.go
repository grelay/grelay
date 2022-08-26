package grelay

func getGrelayExec(c Configuration) grelayExec {
	if c.withGo {
		return grelayExecWithGo{}
	}
	// TODO: When we create the default grelayExec, we need to change this return to default grelayExec.
	return grelayExecWithGo{}
}
