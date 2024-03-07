package supabase

import (
	"mime/multipart"
	"os"

	supabase_storage_uploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type supabaseStorage struct {
	client *supabase_storage_uploader.Client
}

func Init() Interface {
	supClient := supabase_storage_uploader.New(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_TOKEN"),
		os.Getenv("SUPABASE_BUCKET"),
	)

	return &supabaseStorage{
		client: supClient,
	}
}

func (s *supabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	link, err := s.client.Upload(file)
	if err != nil {
		return link, err
	}

	return link, nil
}

func (s *supabaseStorage) Delete(link string) error {
	err := s.client.Delete(link)
	if err != nil {
		return err
	}

	return nil
}
