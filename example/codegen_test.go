package example

import (
	"context"
	"example/client"
	client2 "github.com/DENKweit/gqlgenc/client"
	"github.com/DENKweit/gqlgenc/client/transport"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestSubscription(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	ch, stop := gql.SubscribeMessageAdded(ctx)
	defer stop()

	ids := make([]string, 0)

	for msg := range ch {
		if msg.Error != nil {
			t.Fatal(msg.Error)
		}

		ids = append(ids, msg.Data.MessageAdded.ID)
	}

	assert.Len(t, ids, 3)
}

func isPointer(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

func TestQuery(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	room, _, err := gql.GetRoom(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", room.Room.Name)
	assert.True(t, isPointer(room.Room), "room must be a pointer")
}

func TestQueryNonNull(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	room, _, err := gql.GetRoomNonNull(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", room.RoomNonNull.Name)
	assert.False(t, isPointer(room.RoomNonNull), "room must not be a pointer")
}

func TestQueryCustomType(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	room, _, err := gql.GetRoomCustom(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Room: test", room.String())
}

func TestQueryFragment(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	res, _, err := gql.GetRoomFragment(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", res.Room.Name)
}

func TestQueryUnion(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	res, _, err := gql.GetMedias(ctx)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, res.Medias, 2)

	assert.Equal(t, int64(100), res.Medias[0].Image.Size)
	assert.Equal(t, int64(200), res.Medias[1].Video.Duration)
}

func TestQueryInterface(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	res, _, err := gql.GetBooks(ctx)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, res.Books, 2)

	assert.Equal(t, "Some textbook", res.Books[0].Title)
	assert.Equal(t, []string{"course 1", "course 2"}, res.Books[0].Textbook.Courses)

	assert.Equal(t, "Some Coloring Book", res.Books[1].Title)
	assert.Equal(t, []string{"red", "blue"}, res.Books[1].ColoringBook.Colors)
}

func TestMutationInput(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td, _ := splitcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	res, _, err := gql.CreatePost(ctx, client.PostCreateInput{
		Text: "some text",
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "some text", res.Post.Text)
}

func uploadcli(ctx context.Context) (*client2.Client, func()) {
	return clifactory(ctx, func(ts *httptest.Server) (transport.Transport, func()) {
		tr := httptr(ctx, ts.URL)
		tr.UseFormMultipart = true

		return tr, nil
	})
}

func createUploadFile(t *testing.T) (transport.Upload, int64, func()) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = f.WriteString("some content")
	if err != nil {
		t.Fatal(err)
	}

	err = f.Sync()
	if err != nil {
		t.Fatal(err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	up := transport.NewUpload(f)

	return up, 12, func() {
		os.Remove(f.Name())
	}
}

func TestMutationUploadFile(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td := uploadcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	up, l, rm := createUploadFile(t)
	defer rm()

	res, _, err := gql.MyUploadFile(ctx, up)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, l, res.UploadFile.Size)
}

func TestMutationUploadFiles(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td := uploadcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	up, l, rm := createUploadFile(t)
	defer rm()

	res, _, err := gql.MyUploadFiles(ctx, []*transport.Upload{&up})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, l, res.UploadFiles[0].Size)
}

func TestMutationUploadFilesMap(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cli, td := uploadcli(ctx)
	defer td()

	gql := &client.Client{
		Client: cli,
	}

	up, l, rm := createUploadFile(t)
	defer rm()

	res, _, err := gql.MyUploadFilesMap(ctx, client.UploadFilesMapInput{
		Somefile: up,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, l, res.UploadFilesMap.Somefile.Size)
}
