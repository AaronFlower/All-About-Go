package indexer

type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {

}

type Merger interface {

}