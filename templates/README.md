# Awesome Podcasts

> 😎 Curated list of awesome programming podcasts  [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

---

<details>

<summary>Table of Contents</summary>

{{range .Categories}}- [{{titled .Name}}](#{{dashed .Name}})
{{end}}
</details>

---

{{range .Categories}}## {{titled .Name}}

> {{.SubTitle}}

{{range .Pods}}- [{{.Name}}]({{.URL}}) - {{.Desc}}
{{end}}
{{end}}
