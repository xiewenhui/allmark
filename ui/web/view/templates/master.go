// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

import (
	"fmt"
)

var masterTemplate = fmt.Sprintf(`<!DOCTYPE HTML>
<html lang="{{.LanguageTag}}">
<head>
	<meta charset="utf-8">
	<base href="{{ .BaseUrl }}">

	<title>{{.Title}}</title>

	<link rel="schema.DC" href="http://purl.org/dc/terms/">
	<meta name="DC.date" content="{{.CreationDate}}">

	{{if .GeoLocation }}
	{{if .GeoLocation.Coordinates}}
	<meta name="geo.position" content="{{.GeoLocation.Coordinates}}">
	{{end}}
	
	{{if .GeoLocation.PlaceName}}
	<meta name="geo.placename" content="{{.GeoLocation.PlaceName}}">	
	{{end}}
	{{end}}

	<link rel="alternate" type="application/rss+xml" title="RSS" href="/rss.xml">

	<link rel="shortcut icon" href="/theme/favicon.ico" />

	<link rel="stylesheet" href="/theme/deck.css" media="screen">
	<link rel="stylesheet" href="/theme/screen.css" media="screen">
	<link rel="stylesheet" href="/theme/print.css" media="print">
	<link rel="stylesheet" href="/theme/codehighlighting/highlight.css" media="screen, print">

	<script src="/theme/modernizr.js"></script>
</head>
<body>

{{ if .ToplevelNavigation}}
<nav class="toplevel">
	<form class="search" action="/search" method="GET">
		<input type="text" name="q" placeholder="search">
	</form>

	<ul>
	{{range .ToplevelNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

{{ if .BreadcrumbNavigation}}
<nav class="breadcrumb">
	<ul>
	{{range .BreadcrumbNavigation.Entries}}
	<li>
		<a href="{{.Path}}">{{.Title}}</a>
	</li>
	{{end}}
	</ul>
</nav>
{{end}}

<article class="{{.Type}} level-{{.Level}}">
%s
</article>

<footer>
	<nav>
		<ul>
			<li><a href="/search">Search</a></li>
			<li><a href="/tags.html">Tags</a></li>
			<li><a href="/sitemap.html">Sitemap</a></li>
			<li><a href="/feed.rss">RSS Feed</a></li>
		</ul>
	</nav>
</footer>

<script src="/theme/jquery.js"></script>
<script src="/theme/autoupdate.js"></script>
<script src="/theme/pdf.js"></script>
<script src="/theme/pdf-preview.js"></script>
<script src="/theme/codehighlighting/highlight.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<script src="/theme/deck.js"></script>
<script src="/theme/presentation.js"></script>

</body>
</html>`, ChildTemplatePlaceholder)