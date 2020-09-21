package dto_request

type CourseRequestDto struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CourseContents []CourseContentRequestDto `json:"course_contents" from:"course_contents"`
}
