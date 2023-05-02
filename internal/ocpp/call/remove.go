package call

func (s *Call) Remove(uniqueID string) {
	close(s.channels[uniqueID])
	delete(s.channels, uniqueID)
}
