package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Clever/wag/gen-go/client"
	"github.com/Clever/wag/gen-go/models"
	"github.com/Clever/wag/gen-go/server"
	"github.com/Clever/wag/swagger"

	"net/http"
	"net/http/httptest"
)

func init() {
	swagger.InitCustomFormats()
}

func TestBasicEndToEnd(t *testing.T) {
	s := setupServer()

	c := client.New(s.URL)

	bookID := int64(124)
	bookName := "Test"

	createdBook, err := c.CreateBook(
		context.Background(), &models.Book{ID: bookID, Name: bookName})
	assert.NoError(t, err)
	assert.Equal(t, bookID, createdBook.ID)
	assert.Equal(t, bookName, createdBook.Name)

	booksOutput, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.Equal(t, 1, len(booksOutput))
	assert.Equal(t, bookID, (booksOutput)[0].ID)
	assert.Equal(t, bookName, (booksOutput)[0].Name)

	singleBookOutput, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: bookID})
	assert.NoError(t, err)
	singleBook, ok := singleBookOutput.(*models.GetBookByID200Output)
	require.True(t, ok)
	assert.Equal(t, bookID, singleBook.ID)
	assert.Equal(t, bookName, singleBook.Name)

	// If we have a bookID == 2mod4 then it returns a 204
	otherBookID := int64(126)

	createdBook, err = c.CreateBook(
		context.Background(), &models.Book{ID: otherBookID, Name: bookName})
	assert.NoError(t, err)
	assert.Equal(t, otherBookID, createdBook.ID)
	assert.Equal(t, bookName, createdBook.Name)

	singleBookOutput, err = c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: otherBookID})
	assert.NoError(t, err)
	_, ok = singleBookOutput.(models.GetBookByID204Output)
	require.True(t, ok)
}

func TestUserDefinedErrorResponse(t *testing.T) {
	// The 404 generated by the code
	s := setupServer()

	c := client.New(s.URL)

	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 124})
	assert.Error(t, err)
	_, ok := err.(models.GetBookByID404Output)
	assert.True(t, ok)
}

func TestValidationErrorResponse(t *testing.T) {
	s := setupServer()
	c := client.New(s.URL)

	// Book ID should be a multiple of two
	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 123})
	assert.Error(t, err)
	_, ok := err.(models.DefaultBadRequest)
	assert.True(t, ok)
}

func TestClientSideError(t *testing.T) {
	c := client.New("badServer")

	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.Error(t, err)
	_, ok := err.(models.DefaultInternalError)
	assert.True(t, ok)
}

func TestHeaders(t *testing.T) {
	s := setupServer()

	bookID := int64(124)
	c := client.New(s.URL)
	_, err := c.CreateBook(context.Background(),
		&models.Book{ID: bookID, Name: "test"})
	assert.NoError(t, err)

	// Make a raw HTTP request (i.e. don't use the client) so we can check the headers
	resp, err := http.Get(fmt.Sprintf("%s/v1/books/%d", s.URL, bookID))
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestCustomStringValidation(t *testing.T) {
	s := setupServer()
	c := client.New(s.URL)

	bookID := int64(124)
	_, err := c.CreateBook(context.Background(),
		&models.Book{ID: bookID, Name: "test"})
	assert.NoError(t, err)

	badFormat := "nonMongoFormat"
	_, err = c.GetBookByID(context.Background(),
		&models.GetBookByIDInput{BookID: bookID, AuthorID: &badFormat})
	_, ok := err.(models.DefaultBadRequest)
	assert.True(t, ok)

	validFormat := "012345678901234567890123"
	_, err = c.GetBookByID(context.Background(),
		&models.GetBookByIDInput{BookID: bookID, AuthorID: &validFormat})
	assert.NoError(t, err)
}

type LastCallServer struct {
	lastState     string
	lastAvailable bool
	lastMaxPages  float64
	lastMinPages  int32
	lastAuthors   []string
}

func (d *LastCallServer) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, error) {
	d.lastState = *input.State
	d.lastAvailable = *input.Available
	d.lastMaxPages = *input.MaxPages
	d.lastMinPages = *input.MinPages
	d.lastAuthors = input.Authors
	return []models.Book{}, nil
}
func (d *LastCallServer) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (models.GetBookByIDOutput, error) {
	return nil, nil
}
func (d *LastCallServer) GetBookByID2(ctx context.Context, input *models.GetBookByID2Input) (*models.Book, error) {
	return nil, nil
}
func (d *LastCallServer) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (c *LastCallServer) HealthCheck(ctx context.Context) error {
	return nil
}

func TestDefaultValue(t *testing.T) {
	d := LastCallServer{}
	s := server.New(&d, ":8080")
	testServer := httptest.NewServer(s.Handler)
	c := client.New(testServer.URL)

	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.NoError(t, err)
	assert.Equal(t, "finished", d.lastState)
	assert.True(t, d.lastAvailable)
	assert.Equal(t, 500.5, d.lastMaxPages)
	assert.Equal(t, int32(5), d.lastMinPages)
}

func TestPassInArray(t *testing.T) {
	d := LastCallServer{}
	s := server.New(&d, ":8080")
	testServer := httptest.NewServer(s.Handler)
	c := client.New(testServer.URL)

	_, err := c.GetBooks(context.Background(),
		&models.GetBooksInput{Authors: []string{"author1", "author2"}})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(d.lastAuthors))
	assert.Equal(t, "author1", d.lastAuthors[0])
	assert.Equal(t, "author2", d.lastAuthors[1])
}

type MiddlewareContextTest struct {
	foundKey string
}

func (m *MiddlewareContextTest) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, error) {
	m.foundKey = ctx.Value(testContextKey{}).(string)
	return []models.Book{}, nil
}
func (m *MiddlewareContextTest) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (models.GetBookByIDOutput, error) {
	return nil, nil
}
func (m *MiddlewareContextTest) GetBookByID2(ctx context.Context, input *models.GetBookByID2Input) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareContextTest) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareContextTest) HealthCheck(ctx context.Context) error {
	return nil
}

type testContextKey struct{}

func testContextMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCtx := context.WithValue(r.Context(), testContextKey{}, "contextValue")
		h.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func TestSettingContextInMiddleware(t *testing.T) {
	controller := MiddlewareContextTest{}
	s := server.New(&controller, "")
	testServer := httptest.NewServer(testContextMiddleware(s.Handler))
	c := client.New(testServer.URL)
	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.NoError(t, err)
	assert.Equal(t, "contextValue", controller.foundKey)
}

type TimeoutController struct{}

func (m *TimeoutController) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, error) {
	var books []models.Book
	for i := 0; i < 1000; i++ {
		books = append(books, models.Book{Name: "testing"})
	}
	time.Sleep(100 * time.Millisecond)
	return books, nil
}
func (m *TimeoutController) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (models.GetBookByIDOutput, error) {
	return nil, nil
}
func (m *TimeoutController) GetBookByID2(ctx context.Context, input *models.GetBookByID2Input) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) HealthCheck(ctx context.Context) error {
	return nil
}

func TestTimeout(t *testing.T) {
	s := server.New(&TimeoutController{}, "")
	testServer := httptest.NewServer(testContextMiddleware(s.Handler))
	c := client.New(testServer.URL)

	// Without a timeout, no error
	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.NoError(t, err)

	// Add a per request context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err = c.GetBooks(ctx, &models.GetBooksInput{})
	assert.Error(t, err)
	assert.Equal(t, "context deadline exceeded", err.Error())
	end := time.Now()
	assert.True(t, end.Sub(start) < 80*time.Millisecond)

	// Try with a global client setting
	c = c.WithTimeout(10 * time.Millisecond)
	_, err = c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.Error(t, err)
	assert.Equal(t, "context deadline exceeded", err.Error())
}
