package model

type Folder struct {
	Id int32
	Name string
	CoverUrl string
}

type Audio struct {
	Id int32
	FolderId int32
	Name string
	Url string
}

type Genre struct {
	Id int32
	Name string
}

type Author struct {
	Id int32
	Name string
}