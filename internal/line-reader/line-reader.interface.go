package linereader

type LineReader interface {
	WriteStdin([]byte) (int, error)
	Readline() (string, error)
	Close() error
}
