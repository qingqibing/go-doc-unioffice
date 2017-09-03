// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"strings"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/content_types"
)

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct {
	x *content_types.Types
}

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes() ContentTypes {
	ct := ContentTypes{x: content_types.NewTypes()}
	// add content type defaults
	ct.AddDefault("xml", "application/xml")
	ct.AddDefault("rels", "application/vnd.openxmlformats-package.relationships+xml")
	ct.AddDefault("png", "image/png")
	ct.AddDefault("jpeg", "image/jpeg")
	ct.AddDefault("jpg", "image/jpg")
	ct.AddDefault("wmf", "image/x-wmf")

	ct.AddOverride("/docProps/core.xml", "application/vnd.openxmlformats-package.core-properties+xml")
	ct.AddOverride("/docProps/app.xml", "application/vnd.openxmlformats-officedocument.extended-properties+xml")

	return ct
}

// X returns the inner raw content types.
func (c ContentTypes) X() *content_types.Types {
	return c.x
}

// AddDefault registers a default content type for a given file extension.
func (c ContentTypes) AddDefault(fileExtension string, contentType string) {
	def := content_types.NewDefault()
	def.ExtensionAttr = fileExtension
	def.ContentTypeAttr = contentType
	c.x.Default = append(c.x.Default, def)
}

// AddOverride adds an override content type for a given path name.
func (c ContentTypes) AddOverride(path, contentType string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	or := content_types.NewOverride()
	or.PartNameAttr = path
	or.ContentTypeAttr = contentType
	c.x.Override = append(c.x.Override, or)
}
