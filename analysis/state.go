package analysis

type State struct {
	// Map of document URIs to text contents
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri string, text string) {
	// TODO: handle incremental updates
	s.Documents[uri] = text
}
