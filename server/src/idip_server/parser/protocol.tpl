// this is auto-genrated file,
// Don't modify this file manually

package idip_server

type IDIP_COMMON_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
}

{{range $s:=.}}
{{if $s.Id}}
// {{$s.Desc}}
type {{$s.Id}} struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
{{range $e:=.Entry}}	{{$e.Name}} {{$e.Type}} // {{$e.Desc}}
{{end}}
	} `json:"body"`

}
{{else}}
// {{$s.Desc}}
type {{$s.Name}} struct {
{{range $e:=.Entry}}	{{$e.Name}} {{$e.Type}} // {{$e.Desc}}
{{end}}
}
{{end}}
{{end}}