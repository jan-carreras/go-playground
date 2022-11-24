package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	var name string
	var rootDirectory string
	flag.StringVar(&name, "name", "", "the name of the exercise: 1.3.45_stack")
	flag.StringVar(&rootDirectory, "root", os.Getenv("ROBERT_SEDGEWICK_PATH"), "root directory of robert-sedgewick exercises")
	flag.Parse()

	if name == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := createExercise(rootDirectory, name); err != nil {
		log.Fatalln(err)
	}
}

func createExercise(rootDirectory, name string) error {
	var chapter, module, exercise int
	var exerciseName string

	_, err := fmt.Sscanf(name, "%d.%d.%d_%s", &chapter, &module, &exercise, &exerciseName)
	if err != nil {
		return fmt.Errorf("unable to parse %q: %w", name, err)
	}

	chapterDir := path.Join(rootDirectory, fmt.Sprintf("chapter%d", chapter))
	if f, err := os.Stat(chapterDir); err != nil {
		return err
	} else if !f.IsDir() {
		return fmt.Errorf("chapterDir(%q) is not a directory", chapterDir)
	}

	moduleDir, err := findModule(chapterDir, module)
	if err != nil {
		return err
	}

	if err := checkExerciseExists(moduleDir, exercise); err != nil {
		return err
	}

	return createNewExercise(moduleDir, exercise, exerciseName)
}

func createNewExercise(moduleDir string, exercise int, exerciseName string) error {
	exercisePath := path.Join(moduleDir, fmt.Sprintf("%d_%s", exercise, exerciseName))
	err := os.Mkdir(exercisePath, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(exercisePath, "main.go"))
	if err != nil {
		return err
	}

	template := `package %s

func main() {
	
}
`
	_, err = f.WriteString(fmt.Sprintf(template, exerciseName))
	return err
}

func checkExerciseExists(moduleDir string, exercise int) error {
	files, err := os.ReadDir(moduleDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), fmt.Sprintf("%d_", exercise)) {
			return fmt.Errorf("this exericse already exists: %q", file.Name())
		}
	}
	return nil
}

func findModule(chapterDir string, module int) (moduleDir string, err error) {
	files, err := os.ReadDir(chapterDir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), fmt.Sprintf("%d_", module)) {
			return path.Join(chapterDir, file.Name()), nil
		}
	}

	return "", errors.New("module directory not found")

}
