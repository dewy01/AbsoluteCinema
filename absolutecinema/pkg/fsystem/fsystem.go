package fsystem

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/afero"
)

type FileStorage interface {
	SaveMovieFile(movieID, filename string, file io.Reader) error
	SaveReservationFile(reservationID, filename string, file io.Reader) error

	RemoveMovieFile(movieID, filename string) error
	RemoveReservationFile(reservationID, filename string) error

	GetMovieFile(movieID, filename string) (io.ReadCloser, error)
	GetReservationFile(reservationID, filename string) (io.ReadCloser, error)
}

type Storage struct {
	fs       afero.Fs
	basePath string
}

func New(basePath string, fs afero.Fs) *Storage {
	return &Storage{
		fs:       fs,
		basePath: basePath,
	}
}

func (s *Storage) SaveMovieFile(movieID string, filename string, file io.Reader) error {
	dir := filepath.Join(s.basePath, "movies", movieID)
	return s.saveFile(dir, filename, file)
}

func (s *Storage) SaveReservationFile(reservationID string, filename string, file io.Reader) error {
	dir := filepath.Join(s.basePath, "reservations", reservationID)
	return s.saveFile(dir, filename, file)
}

func (s *Storage) RemoveMovieFile(movieID string, filename string) error {
	path := filepath.Join(s.basePath, "movies", movieID, filename)
	return s.fs.Remove(path)
}

func (s *Storage) RemoveReservationFile(reservationID string, filename string) error {
	path := filepath.Join(s.basePath, "reservations", reservationID, filename)
	return s.fs.Remove(path)
}

func (s *Storage) GetMovieFile(movieID string, filename string) (io.ReadCloser, error) {
	path := filepath.Join(s.basePath, "movies", movieID, filename)
	file, err := s.fs.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *Storage) GetReservationFile(reservationID string, filename string) (io.ReadCloser, error) {
	path := filepath.Join(s.basePath, "reservations", reservationID, filename)
	file, err := s.fs.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *Storage) saveFile(dir string, filename string, file io.Reader) error {
	if err := s.fs.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	fpath := filepath.Join(dir, filename)
	dst, err := s.fs.Create(fpath)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
