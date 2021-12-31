package models

func GetHomePage(articles []ArticleView, pagenumber int64, maxpages int64) HomePage {
	var hasbefore, hasnext bool

	if (pagenumber - 1) >= 1 {
		hasbefore = true
	}

	if (pagenumber + 1) <= maxpages {
		hasnext = true
	}

	start := pagenumber - 2
	finish := pagenumber + 2

	if start < 1 {
		start = 1
	}

	if finish > maxpages {
		finish = maxpages
	}

	pages := make([]int64, 0)
	for i := start; i <= finish; i++ {
		pages = append(pages, i)
	}

	home := HomePage{
		Articles: articles,
		HasBefore: hasbefore,
		HasNext: hasnext,
		Older: pagenumber - 1,
		Newer: pagenumber + 1,
		Current: pagenumber,
		Pages: pages,
	}
	return home
}