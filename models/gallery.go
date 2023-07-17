package models

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type Image struct {
	GalleryID int
	Path      string
	Filename  string
}

type GalleryService struct {
	DB *sql.DB

	ImagesDir string
}

func (service *GalleryService) Create(title string, userID int) (*Gallery, error) {
	gallery := Gallery{
		Title:  title,
		UserID: userID,
	}
	row := service.DB.QueryRow(`
	INSERT INTO galleries (title, user_id) 
	VALUES ($1, $2) RETURNING id`, gallery.Title, gallery.UserID)
	err := row.Scan(&gallery.ID)
	if err != nil {
		return nil, fmt.Errorf("create gallery: %w", err)
	}
	return &gallery, nil
}

func (service *GalleryService) ByID(id int) (*Gallery, error) {
	// TODO: implement
	gallery := Gallery{
		ID: id,
	}
	row := service.DB.QueryRow(`
	SELECT title, user_id
	FROM galleries
	WHERE id = $1`, gallery.ID)
	err := row.Scan(&gallery.Title, &gallery.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get gallery by id: %w", err)
	}
	return &gallery, nil
}

func (service *GalleryService) ByUserID(userID int) ([]Gallery, error) {
	rows, err := service.DB.Query(`
	SELECT id, title 
	FROM galleries
	WHERE user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("get galleries by user id: %w", err)
	}
	var galleries []Gallery
	for rows.Next() {
		gallery := Gallery{
			UserID: userID,
		}
		err = rows.Scan(&gallery.ID, &gallery.Title)
		if err != nil {
			return nil, fmt.Errorf("get galleries by user id: %w", err)
		}
		galleries = append(galleries, gallery)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("get galleries by user id: %w", err)
	}
	return galleries, nil
}

func (service *GalleryService) Update(gallery *Gallery) error {
	_, err := service.DB.Exec(`
	UPDATE galleries
	SET title = $2
	WHERE id = $1`, gallery.ID, gallery.Title)
	if err != nil {
		return fmt.Errorf("update gallery: %w", err)
	}
	return nil
}

func (service *GalleryService) Delete(id int) error {
	_, err := service.DB.Exec(`
	DELETE FROM galleries
	WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete gallery: %w", err)
	}
	return nil
}

func (service *GalleryService) Images(galleryID int) ([]Image, error) {
	// TODO: Implement this
	globPattern := filepath.Join(service.galleryDir(galleryID), "*")
	allFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, fmt.Errorf("retrieving gallery images: %w", err)
	}
	// var imagePaths []string
	var images []Image
	for _, file := range allFiles {
		if hasExtension(file, service.extensions()) {
			// imagePaths = append(imagePaths, file)
			images = append(images, Image{
				GalleryID: galleryID,
				Path:      file,
				Filename:  filepath.Base(file),
			})
		}
	}
	return images, nil
}

func (service *GalleryService) extensions() []string {
	return []string{".jpg", ".jpeg", ".png", ".gif"}
}

func (service *GalleryService) galleryDir(id int) string {
	imagesDir := service.ImagesDir
	if imagesDir == "" {
		imagesDir = "images"
	}
	return filepath.Join(imagesDir, fmt.Sprintf("%d", id))
}

func hasExtension(file string, extensions []string) bool {
	for _, ext := range extensions {
		file = strings.ToLower(file)
		ext = strings.ToLower(ext)
		if filepath.Ext(file) == ext {
			return true
		}
	}
	return false
}
