// Code generated by go generate; DO NOT EDIT.

package image

import (
	"fmt"

	"sacer/go/entry"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/tools"
	"time"
)

func (e *Image) Type() string {
	return "image"
}

func (e *Image) Parent() entry.Entry {
	return e.parent
}

func (e *Image) File() *file.File {
	return e.file
}

func (e *Image) Id() int64 {
	return e.date.Unix()
}

func (e *Image) Timestamp() string {
	return e.date.Format(tools.Timestamp)
}

func (e *Image) Hash() string {
	return tools.ToB16(e.date)
}

func (e *Image) HashShort() string {
	return tools.ShortenHash(e.Hash())
}

func (e *Image) Date() time.Time {
	return e.date
}

func (e *Image) Info() info.Info {
	return e.info
}

func (e *Image) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Image) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return tools.Normalize(e.info.Title(lang))
}

func (e *Image) MediaObject() bool {
	return e.Type() != "audio" && entry.IsBlob(e)
}

func (e *Image) ObjectType() string {
	if e.MediaObject() {
		return "mob"
	}
	return "tob"
}

func (e *Image) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Image) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Image) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Image) Section() string {
	return e.Parent().Section()
}

func (e *Image) Perma(lang string) string {
	if e.parent.Type() == "set" {
		return e.parent.Perma(lang)
	}

	name := e.Hash()
	if slug := e.Slug(lang); slug != "" {
		name = fmt.Sprintf("%v-%v", slug, e.Hash())
	}

	switch e.Section() {
	case "kine":
		return fmt.Sprintf(
			"/%v/%v/%v",
			tools.KineName[lang],
			e.Date().Format("06-01"),
			fmt.Sprintf("%v-%v", e.Date().Format("02"), name),
		)
	case "index":
		if e.Type() != "image" {
			return fmt.Sprintf("%v#%v", e.parent.Perma(lang), tools.Normalize(e.Title(lang)))
		}
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), name)
}
