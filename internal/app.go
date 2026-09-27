package internal

import (
	"fmt"
	"time"

	"github.com/skvdmt/chrome"
)

const (
	TEST_PAGE_URL = "https://skvdmt.ru/"
)

// App Приложение.
type App struct {
	driver *chrome.Driver
}

// NewApp Конструктор.
func NewApp() (*App, error) {
	// Создание драйвера.
	d, err := chrome.NewDriver(
		chrome.WithRemoveUserDataDirAfterClose(),
	)
	if err != nil {
		return nil, err
	}
	return &App{
		driver: d,
	}, nil
}

// Start Старт.
func (a *App) Start() error {
	// Открытие браузера.
	if err := a.driver.Open(); err != nil {
		return err
	}
	defer func() {
		if err := a.close(); err != nil {
			panic(err)
		}
	}()
	// Переход на страницу.
	if err := a.driver.Navigate(TEST_PAGE_URL); err != nil {
		return err
	}
	// Тесты.
	tests := []struct {
		name     string
		selector string
		expected string
	}{
		{"Main header", "#main_header", "Dmitry Skidanov"},
		{"Profession", "#profession", "Full stack engineer. Apps for high load."},
		{"Second header", "#second_header", "Main technologies"},
		{"Name Technology Go", "#technology_name_go", "Go"},
		{"Name Technology Postgres", "#technology_name_postgres", "Postgres"},
		{"Name Technology Docker", "#technology_name_docker", "Docker"},
		{"Name Technology REST API", "#technology_name_rest_api", "REST API"},
		{"Name Technology gRPC", "#technology_name_grpc", "gRPC"},
		{"Name Technology Git", "#technology_name_git", "Git"},
		{"Name Technology CI/CD", "#technology_name_cicd", "CI/CD"},
		{"Name Technology JavaScript", "#technology_name_javascript", "JavaScript"},
		{"Name Technology Vue", "#technology_name_vue", "Vue"},
		{"Examples header", "#examples_header", "Examples"},
		{"Example title Chrome", "#example_name_chrome", "Golang Chrome driver"},
		{"Example title JWT", "#example_name_jwt", "Golang JSON Web Tokens"},
		{"Example title Auth", "#example_name_auth", "Authentication"},
		{"Example title Home", "#example_name_home", "Homepage"},
		{"Example title Telegram bot", "#example_name_tgbot", "Telegram bot"},
		{"Example title Chess game", "#example_name_chess", "Chess game"},
		{"Favorites header", "#favorites_header", "Favorites"},
		{"Favorites first subheader", "#favorites_first_subheader", "Software for development"},
		{"Favorites software Visual Studio Code", "#favorites_software_visual_studio_code", "Visual Studio Code"},
		{"Favorites software GoLand", "#favorites_software_goland", "GoLand"},
		{"Favorites software WebStorm", "#favorites_software_webstorm", "WebStorm"},
		{"Favorites software DataGrip", "#favorites_software_datagrip", "DataGrip"},
		{"Favorites software Bruno", "#favorites_software_bruno", "Bruno"},
		{"Favorites software Swagger", "#favorites_software_swagger", "Swagger"},
		{"Favorites software Vite", "#favorites_software_vite", "Vite"},
		{"Favorites second subheader", "#favorites_second_subheader", "Golang external libraries i like"},
		{"Favorites lib labstack/echo", "#favorites_lib_labstack_echo", "labstack/echo"},
		{"Favorites lib grpc/grpc-go", "#favorites_lib_grpc_grpc-go", "grpc/grpc-go"},
		{"Favorites lib jackc/pgx", "#favorites_lib_jackc_pgx", "jackc/pgx"},
		{"Favorites lib gorilla/websocket", "#favorites_lib_gorilla_websocket", "gorilla/websocket"},
		{"Link Name GitHub", "#link_name_github", "GitHub repositories"},
		{"Link URL GitHub", "#link_url_github", "github.com/skvdmt"},
		{"Link Name Docker", "#link_name_docker", "Docker Hub registry"},
		{"Link URL Docker", "#link_url_docker", "hub.docker.com/u/skvdmt"},
		{"Link Name Telegram", "#link_name_telegram", "Telegram"},
		{"Link URL Telegram", "#link_url_telegram", "t.me/skidanovdima"},
		{"Link Name Email", "#link_name_email", "Email"},
		{"Link URL Email", "#link_url_email", "skvdmt@yandex.ru"},
		{"Copyright", "#copyright", fmt.Sprintf("Dmitry Skidanov — full stack engineer %d", time.Now().Year())},
		{"Location", "#location", "Russian Federation, Moscow"},
	}
	// Тестирование.
	for _, t := range tests {
		if err := a.Test(t.name, t.selector, t.expected); err != nil {
			return err
		}
	}
	return nil
}

// Test Тестирование.
func (a *App) Test(name, selector, expected string) error {
	fmt.Printf("%s test\n", name)
	got, err := a.driver.WaitNodeText(selector)
	if err != nil {
		return err
	}
	if got != expected {
		return fmt.Errorf(
			`error: node %s; text expected: "%s"; got "%s"`,
			selector,
			expected,
			got,
		)
	}
	fmt.Println("OK")
	return nil
}

// close Закрытие ресурсов.
func (a *App) close() error {
	return a.driver.Close()
}
