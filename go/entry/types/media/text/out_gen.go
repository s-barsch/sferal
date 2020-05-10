// Code generated by go generate; DO NOT EDIT.

package text

import (
	"fmt"

	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Text) Type() string {
	return "text"
}

func (e *Text) Parent() entry.Entry {
	return e.parent
}

func (e *Text) File() *file.File {
	return e.file
}

func (e *Text) Id() int64 {
	return e.date.Unix()
}

func (e *Text) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Text) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Text) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Text) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Text) Date() time.Time {
	return e.date
}

func (e *Text) Info() info.Info {
	return e.info
}

func (e *Text) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func (e *Text) IsBlob() bool {
	return entry.IsBlob(e)
}

func (e *Text) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Text) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Text) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Text) Section() string {
	return e.Parent().Section()
}

func (e *Text) Perma(lang string) string {
	slug := e.Slug(lang)
	if slug != "" {
		return fmt.Sprintf("%v/%v-%v", e.parent.Path(lang), slug, e.Hash())
	}
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Hash())
}
