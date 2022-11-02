package main

import (
	"log"
	"net/http"
)

type SuppliersPageData struct {
	Suppliers []Supplier
}

type Supplier struct {
	Name     string
	Vendor   string
	LogoUrl  string
	Services []string
	Url      string
	Docs     []SupplierDoc
}

type SupplierDoc struct {
	Name    string // ISO27k, SOC2
	Logo    string
	Url     string
	Comment string
}

func suppliers(w http.ResponseWriter, r *http.Request) {

	data := SuppliersPageData{
		Suppliers: []Supplier{
			{
				Name:     "Google Workspace",
				Vendor:   "Google LLC",
				LogoUrl:  "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/Google_Workspace_Logo.svg/460px-Google_Workspace_Logo.svg.png",
				Services: []string{"IdP", "Mail", "Calendar"},
				Url:      "https://workspace.google.com/",
				Docs: []SupplierDoc{
					{
						Name:    "ISO/IEC 27001",
						Logo:    "logos/iso.png",
						Url:     "https://cloud.google.com/security/compliance/iso-27001/",
						Comment: "",
					}, {
						Name:    "SOC3",
						Logo:    "logos/soc.png",
						Url:     "https://cloud.google.com/security/compliance/soc-3/",
						Comment: "",
					},
				},
			},
			{
				Name:     "Notion",
				Vendor:   "Notion Labs, Inc.",
				LogoUrl:  "https://upload.wikimedia.org/wikipedia/commons/4/45/Notion_app_logo.png",
				Services: []string{"Task & Project Management"},
				Url:      "https://www.notion.so/",
				Docs: []SupplierDoc{
					{
						Name:    "SOC2",
						Logo:    "logos/soc.png",
						Url:     "https://www.notion.so/security",
						Comment: "",
					},
				},
			},
			{
				Name:     "GitHub",
				Vendor:   "GitHub, Inc.",
				LogoUrl:  "https://github.githubassets.com/images/modules/logos_page/GitHub-Logo.png",
				Services: []string{"Version Control System"},
				Url:      "https://github.com/",
				Docs: []SupplierDoc{
					{
						Name:    "SOC2",
						Logo:    "logos/soc.png",
						Url:     "https://github.com/security",
						Comment: "",
					},
					{
						Name:    "ISO/IEC 27001:2013",
						Logo:    "logos/iso.png",
						Url:     "https://github.com/security",
						Comment: "",
					},
					{
						Name:    "Cloud Security Alliance",
						Logo:    "logos/csa.png",
						Url:     "https://github.com/security",
						Comment: "",
					},
				},
			},
		},
	}

	err := tmpl.ExecuteTemplate(w, "suppliers.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
