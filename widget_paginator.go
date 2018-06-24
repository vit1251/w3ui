package w3ui

import "fmt"
import "strings"
import "html/template"
import "log"
import "strconv"
import "net/http"

type Paginator struct {
	PageSize uint
	Page uint
	ItemCount uint
}

type Page struct {
	Index uint
	Page uint
	URL string
	Active bool
}

func NewPaginator() ( _ *Paginator, _ error ) {
	p := &Paginator{
		PageSize: 25,
		Page: 1,
		ItemCount: 0,
	}
	return p, nil
}


func (p *Paginator) SetPageSize( pageSize uint ) {
  p.PageSize = pageSize
}


func (p *Paginator) SetPage( page uint ) {
  p.Page = page
}

func (p *Paginator) SetItemCount( item_count uint ) {
  p.ItemCount = item_count
}

func (p *Paginator) Limit() ( int ) {
  var result int

  result = int( p.PageSize )

  return result
}


func (p *Paginator) Offset() ( int ) {

  var result int

  if p.Page < 1 {
    p.Page = 1
  }
  result = int( p.Page * p.PageSize - p.PageSize )

  return result
}


func (p *Paginator) ParseRequest(r *http.Request) {

	log.Printf("Parse request\n")

	page := r.URL.Query().Get("page")
	if page != "" {
		page_number, err0 := strconv.ParseUint( page, 10, 0 )
		if err0 == nil {
			log.Printf("Request page is %v\n", page_number)
			p.Page = uint(page_number)
		}
	}
}


func (p *Paginator) Pages() []Page {

	var page_index uint = 0
	var item_count uint = 0

	pages := make([]Page, 0)

	for {
		page_number := page_index + 1
		p1 := Page{
			Index: page_index,
			Page: page_number,
			Active: p.Page == page_number,
		}
		pages = append(pages, p1)
		item_count += p.PageSize
		page_index += 1
		if item_count >= p.ItemCount {
			break
		}
	}

	return pages
}


type WidgetPaginator struct {
	Paginator *Paginator
}

func NewWidgetPaginator() (*WidgetPaginator, error) {
	return &WidgetPaginator{}, nil
}

func (w *WidgetPaginator) SetPaginator(p *Paginator) {
	w.Paginator = p
}

func (w *WidgetPaginator) Render() ( template.HTML ) {

	var items []string

	items = append(items, "<div class=\"w3-container w3-padding-24\">")
	items = append(items, "	<div class=\"w3-bar\">")

	pages := w.Paginator.Pages()
	for _, page := range pages {
		var pageURL string
		pageURL = page.URL
		if pageURL == "" {
			pageURL = fmt.Sprintf("?page=%d", page.Page)
		}
		var classNames []string
		classNames = append(classNames, "w3-bar-item")
		classNames = append(classNames, "w3-button")
		if page.Active {
			classNames = append(classNames, "w3-red")
		} else {
			classNames = append(classNames, "w3-black")
		}
		class := strings.Join(classNames, " ")
		items = append(items, fmt.Sprintf("<a href=\"%s\" class=\"%s\">%d</a>", pageURL, class, page.Page))
	}
	items = append(items, "	</div>")
	items = append(items, "</div>")

	content := strings.Join(items, "\n")

	return template.HTML(content)
}
