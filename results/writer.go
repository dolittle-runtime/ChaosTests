package results

// Writer writes results to a Collector
type Writer struct {
	collector *Collector
	name      string
	kind      string
}

// NewFor creates a new child Writer for the provided name and kind
func (w *Writer) NewFor(name, kind string) *Writer {
	return &Writer{
		collector: w.collector,
		name:      w.name + "." + name,
		kind:      kind,
	}
}

// Log writes a log message with format and arguments like the "fmt" package
func (w *Writer) Log(format string, a ...interface{}) {
	w.collector.log(w.name, w.kind, format, a...)
}
