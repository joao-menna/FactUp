package utils

func GetPostOffset(limit, page int) int {
	if limit < 1 {
		limit = 1
	}

	if limit > 30 {
		limit = 30
	}

	if page < 0 {
		page = 0
	}

	return limit * page
}
