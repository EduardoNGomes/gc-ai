package providers

type SelectAgentSpy struct {
	option Provider
}

func (s *SelectAgentSpy) SetAgent(p Provider) {
	s.option = p
}

func (s *SelectAgentSpy) SelectedOption() Provider {
	return s.option
}

func NewSelectAgentSpy() *SelectAgentSpy {
	return &SelectAgentSpy{}
}
