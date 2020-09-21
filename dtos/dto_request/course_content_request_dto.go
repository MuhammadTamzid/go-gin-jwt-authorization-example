package dto_request

type CourseContentRequestDto struct {
	Title string `join:"title" binding:"required"`
	Description string `join:"description" binding:"required"`
	LectureUrl string `join:"lecture_url" from:"lecture_url"`
}
