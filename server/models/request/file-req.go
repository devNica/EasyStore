package request

type FileRequestModel struct {
	Filetype string
	Filesize int
	Buffer   []byte
}
