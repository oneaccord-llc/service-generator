package generator

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	ProjectName string
}

func NewGenerator(projectName string) *Generator {
	return &Generator{
		ProjectName: projectName,
	}
}

func (g *Generator) Generate(templatesFS embed.FS) error {
	templatesDir := "templates"
	destDir := g.ProjectName

	// Open the embedded filesystem
	fsys, err := fs.Sub(templatesFS, templatesDir)
	if err != nil {
		return err
	}

	// Create the destination directory
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(g.ProjectName, "migrations"), 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(g.ProjectName, "sql"), 0755); err != nil {
		return err
	}

	// Walk through the embedded filesystem and copy files to the destination
	if err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		srcPath := filepath.Join(templatesDir, path)

		destPath := filepath.Join(destDir, strings.TrimSuffix(path, ".template"))

		// Check if it's a directory, create it in the destination
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Copy the file from the embedded filesystem to the destination
		srcFile, err := templatesFS.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return err
		}
		if err := g.ReplaceProjectName(destPath, "<PROJECT_NAME>", destDir); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (g *Generator) ReplaceProjectName(filePath, projectName, destDir string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	modifiedData := strings.ReplaceAll(string(fileData), projectName, filepath.Base(destDir))

	if err := os.WriteFile(filePath, []byte(modifiedData), 0644); err != nil {
		return err
	}

	return nil
}
