package file

import (
	"io"
	"os"
	"path/filepath"

	"/src/internal/service/file/proto/gen"
)

type FileService struct {
	gen.UnimplementedFileServiceServer
}

func (s *FileService) UploadFile(stream gen.FileService_UploadFileServer) error {
	var file *os.File

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&gen.UploadFileResponse{Message: "File uploaded successfully"})
		}
		if err != nil {
			return err
		}
		if file == nil {
			filePath := filepath.Join("uploads", req.GetFilename())
			file, err = os.Create(filePath)
			if err != nil {
				return err
			}
		}
	}
}
