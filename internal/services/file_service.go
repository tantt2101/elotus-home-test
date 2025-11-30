package services

import (
	"database/sql"
	"errors"
	"io"
	"os"

	"elotus-home-test/internal/structs"
)

type UploadService struct {
	DB *sql.DB
}

func NewUploadService(db *sql.DB) *UploadService {
	return &UploadService{DB: db}
}

func validateFile(fileName string, size int64, contentType string) error {

	if fileName == "" {
		return errors.New("file name is required")
	}

	if size <= 0 {
		return errors.New("file is empty")
	}

	if size > 8*1024*1024 {
		return errors.New("file exceeds max size 8MB")
	}

	allowed := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
		"image/webp": true,
		"image/gif":  true,
	}

	if !allowed[contentType] {
		return errors.New("unsupported file type")
	}

	return nil
}

func (s *UploadService) UploadFile(
	file io.Reader,
	fileName string,
	contentType string,
	size int64,
	userID string,
) (*structs.MediaFile, error) {

	if err := validateFile(fileName, size, contentType); err != nil {
		return nil, err
	}

	savePath := "/tmp/" + fileName
	out, err := os.Create(savePath)
	if err != nil {
		return nil, errors.New("cannot save file")
	}
	defer out.Close()

	io.Copy(out, file)
	result, err := s.DB.Exec(`
		INSERT INTO media_files (filename, content_type, size, path, user_id)
		VALUES (?, ?, ?, ?, ?)`,
		fileName, contentType, size, savePath, userID,
	)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return &structs.MediaFile{
		ID:          id,
		FileName:    fileName,
		ContentType: contentType,
		Size:        size,
		UserID:      userID,
		Path:		 path
	}, nil
}
