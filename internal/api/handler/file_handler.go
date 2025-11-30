package handler

import (
	"net/http"
	"io"
	"database/sql"
	"elotus-home-test/internal/api/utils"
	"elotus-home-test/internal/services"
)

func UploadFile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)

		r.ParseMultipartForm(10 << 20)

		file, header, err := r.FormFile("file")
		if err != nil {
			utils.Error(w, "file is required", 422)
			return
		}
		defer file.Close()

		buff := make([]byte, 512)
		file.Read(buff)
		contentType := http.DetectContentType(buff)
		file.Seek(0, io.SeekStart)

		service := services.NewUploadService(db)

		resp, err := service.UploadFile(
			file,
			header.Filename,
			contentType,
			header.Size,
			userID,
		)

		if err != nil {
			utils.Error(w, err.Error(), 400)
			return
		}

		utils.Success(w, "upload success", resp)
	}
}
