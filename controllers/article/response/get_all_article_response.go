package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type ArticleGetAll struct {
	Articles []ArticleResponse `json:"articles"`
}

func SliceFromUseCase(articles *[]entities.Article) *ArticleGetAll {
	allArticles := make([]ArticleResponse, len(*articles))
	for i, _article := range *articles {
		allArticles[i] = ArticleResponse{
			ID:      _article.ID,
			AdminID: _article.AdminID,
			Author:  _article.Admin.FullName,
			Title:   _article.Title,
			Content: _article.Content,
			Picture: _article.Picture,
		}
	}

	return &ArticleGetAll{
		Articles: allArticles,
	}
}
