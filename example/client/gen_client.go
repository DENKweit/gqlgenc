// Code generated by github.com/infiotinc/gqlgenc, DO NOT EDIT.

package client

import (
	"context"
	"encoding/json"
	"example/somelib"

	"github.com/infiotinc/gqlgenc/client"
	"github.com/infiotinc/gqlgenc/client/transport"
)

type Client struct {
	Client *client.Client
}

type RoomFragment struct {
	Name string "json:\"name\""
}
type GetRoom_Room struct {
	Name string "json:\"name\""
}
type GetRoom struct {
	Room *GetRoom_Room "json:\"room\""
}
type GetRoomNonNull_RoomNonNull struct {
	Name string "json:\"name\""
}
type GetRoomNonNull struct {
	RoomNonNull GetRoomNonNull_RoomNonNull "json:\"roomNonNull\""
}
type GetRoomFragment struct {
	Room *RoomFragment "json:\"room\""
}
type GetMedias_Image struct {
	Size int64 "json:\"size\""
}
type GetMedias_Video struct {
	Duration int64 "json:\"duration\""
}
type GetMedias_Medias struct {
	Typename string           "json:\"__typename\""
	Image    *GetMedias_Image "json:\"-\""
	Video    *GetMedias_Video "json:\"-\""
}

func (t *GetMedias_Medias) UnmarshalJSON(data []byte) error {
	type ΞAlias GetMedias_Medias
	var r ΞAlias

	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	*t = GetMedias_Medias(r)

	switch r.Typename {
	case "Image":
		var a GetMedias_Image
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Image = &a
	case "Video":
		var a GetMedias_Video
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Video = &a
	}

	return nil
}

type GetMedias struct {
	Medias []GetMedias_Medias "json:\"medias\""
}
type GetBooks_Textbook struct {
	Courses []string "json:\"courses\""
}
type GetBooks_ColoringBook struct {
	Colors []string "json:\"colors\""
}
type GetBooks_Books struct {
	Typename     string                 "json:\"__typename\""
	Title        string                 "json:\"title\""
	Textbook     *GetBooks_Textbook     "json:\"-\""
	ColoringBook *GetBooks_ColoringBook "json:\"-\""
}

func (t *GetBooks_Books) UnmarshalJSON(data []byte) error {
	type ΞAlias GetBooks_Books
	var r ΞAlias

	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	*t = GetBooks_Books(r)

	switch r.Typename {
	case "ColoringBook":
		var a GetBooks_ColoringBook
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.ColoringBook = &a
	case "Textbook":
		var a GetBooks_Textbook
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Textbook = &a
	}

	return nil
}

type GetBooks struct {
	Books []GetBooks_Books "json:\"books\""
}
type SubscribeMessageAdded_MessageAdded struct {
	ID string "json:\"id\""
}
type SubscribeMessageAdded struct {
	MessageAdded SubscribeMessageAdded_MessageAdded "json:\"messageAdded\""
}
type CreatePost_Post struct {
	ID   string "json:\"id\""
	Text string "json:\"text\""
}
type CreatePost struct {
	Post CreatePost_Post "json:\"post\""
}
type UploadFile_UploadFile struct {
	Size int64 "json:\"size\""
}
type UploadFile struct {
	UploadFile UploadFile_UploadFile "json:\"uploadFile\""
}
type UploadFiles_UploadFiles struct {
	Size int64 "json:\"size\""
}
type UploadFiles struct {
	UploadFiles []UploadFiles_UploadFiles "json:\"uploadFiles\""
}
type UploadFilesMap_Somefile struct {
	Size int64 "json:\"size\""
}
type UploadFilesMap_UploadFilesMap struct {
	Somefile UploadFilesMap_Somefile "json:\"somefile\""
}
type UploadFilesMap struct {
	UploadFilesMap UploadFilesMap_UploadFilesMap "json:\"uploadFilesMap\""
}

const GetRoomDocument = `query GetRoom ($name: String!) {
	room(name: $name) {
		name
	}
}
`

func (Ξc *Client) GetRoom(ctх context.Context, name string) (*GetRoom, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoom
		res, err := Ξc.Client.Query(ctх, "GetRoom", GetRoomDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomNonNullDocument = `query GetRoomNonNull ($name: String!) {
	roomNonNull(name: $name) {
		name
	}
}
`

func (Ξc *Client) GetRoomNonNull(ctх context.Context, name string) (*GetRoomNonNull, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoomNonNull
		res, err := Ξc.Client.Query(ctх, "GetRoomNonNull", GetRoomNonNullDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomFragmentDocument = `query GetRoomFragment ($name: String!) {
	room(name: $name) {
		... RoomFragment
	}
}
fragment RoomFragment on Chatroom {
	name
}
`

func (Ξc *Client) GetRoomFragment(ctх context.Context, name string) (*GetRoomFragment, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoomFragment
		res, err := Ξc.Client.Query(ctх, "GetRoomFragment", GetRoomFragmentDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomCustomDocument = `query GetRoomCustom ($name: String!) {
	room(name: $name) {
		name
	}
}
`

func (Ξc *Client) GetRoomCustom(ctх context.Context, name string) (*somelib.CustomRoom, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data somelib.CustomRoom
		res, err := Ξc.Client.Query(ctх, "GetRoomCustom", GetRoomCustomDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetMediasDocument = `query GetMedias {
	medias {
		__typename
		... on Image {
			size
		}
		... on Video {
			duration
		}
	}
}
`

func (Ξc *Client) GetMedias(ctх context.Context) (*GetMedias, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data GetMedias
		res, err := Ξc.Client.Query(ctх, "GetMedias", GetMediasDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetBooksDocument = `query GetBooks {
	books {
		__typename
		title
		... on Textbook {
			courses
		}
		... on ColoringBook {
			colors
		}
	}
}
`

func (Ξc *Client) GetBooks(ctх context.Context) (*GetBooks, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data GetBooks
		res, err := Ξc.Client.Query(ctх, "GetBooks", GetBooksDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const SubscribeMessageAddedDocument = `subscription SubscribeMessageAdded {
	messageAdded(roomName: "test") {
		id
	}
}
`

type MessageSubscribeMessageAdded struct {
	Data       *SubscribeMessageAdded
	Error      error
	Extensions transport.RawExtensions
}

func (Ξc *Client) SubscribeMessageAdded(ctх context.Context) (<-chan MessageSubscribeMessageAdded, func()) {
	Ξvars := map[string]interface{}{}

	{
		res := Ξc.Client.Subscription(ctх, "SubscribeMessageAdded", SubscribeMessageAddedDocument, Ξvars)

		ch := make(chan MessageSubscribeMessageAdded)

		go func() {
			for res.Next() {
				opres := res.Get()

				var msg MessageSubscribeMessageAdded
				if len(opres.Errors) > 0 {
					msg.Error = opres.Errors
				}

				err := opres.UnmarshalData(&msg.Data)
				if err != nil && msg.Error == nil {
					msg.Error = err
				}

				msg.Extensions = opres.Extensions

				ch <- msg
			}

			if err := res.Err(); err != nil {
				ch <- MessageSubscribeMessageAdded{
					Error: err,
				}
			}
			close(ch)
		}()

		return ch, res.Close
	}
}

const CreatePostDocument = `mutation CreatePost ($input: PostCreateInput!) {
	post(input: $input) {
		id
		text
	}
}
`

func (Ξc *Client) CreatePost(ctх context.Context, input PostCreateInput) (*CreatePost, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"input": input,
	}

	{
		var data CreatePost
		res, err := Ξc.Client.Mutation(ctх, "CreatePost", CreatePostDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFileDocument = `mutation UploadFile ($file: Upload!) {
	uploadFile(file: $file) {
		size
	}
}
`

func (Ξc *Client) UploadFile(ctх context.Context, file transport.Upload) (*UploadFile, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"file": file,
	}

	{
		var data UploadFile
		res, err := Ξc.Client.Mutation(ctх, "UploadFile", UploadFileDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFilesDocument = `mutation UploadFiles ($files: [Upload!]!) {
	uploadFiles(files: $files) {
		size
	}
}
`

func (Ξc *Client) UploadFiles(ctх context.Context, files []*transport.Upload) (*UploadFiles, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"files": files,
	}

	{
		var data UploadFiles
		res, err := Ξc.Client.Mutation(ctх, "UploadFiles", UploadFilesDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFilesMapDocument = `mutation UploadFilesMap ($files: UploadFilesMapInput!) {
	uploadFilesMap(files: $files) {
		somefile {
			size
		}
	}
}
`

func (Ξc *Client) UploadFilesMap(ctх context.Context, files UploadFilesMapInput) (*UploadFilesMap, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"files": files,
	}

	{
		var data UploadFilesMap
		res, err := Ξc.Client.Mutation(ctх, "UploadFilesMap", UploadFilesMapDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}