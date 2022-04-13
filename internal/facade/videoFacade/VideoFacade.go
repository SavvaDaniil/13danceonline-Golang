package videoFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/repository/videoRepository"
	"danceonline/internal/viewmodel/videoViewModel"
	"errors"
	"os"
	"strconv"
)

func GetVideoTutorialLessonViewModel(id_of_tutorial int) (*videoViewModel.VideoTutorialLessonViewModel, error) {

	video, err := videoRepository.FindByIdOfTutorial(id_of_tutorial)
	if err != nil {
		return nil, err
	}

	if video == nil {
		return nil, nil
	}

	var videoSrc *string
	if video.Hash != nil {
		videoSrc = getVideoSrc(video.Id, *(video.Hash))
	}

	var videoTutorialLessonViewModel videoViewModel.VideoTutorialLessonViewModel = videoViewModel.VideoTutorialLessonViewModel{
		Id:             video.Id,
		Id_of_tutorial: id_of_tutorial,
		VideoSrc:       *videoSrc,
	}

	return &videoTutorialLessonViewModel, nil
}

func GetVideoCourseLessonViewModel(id_of_course int, number_of_lesson int) (*videoViewModel.VideoCourseLessonViewModel, error) {

	videos, err := videoRepository.ListAllConnectedToCourse(id_of_course)
	if err != nil {
		return nil, err
	}

	if videos == nil {
		return nil, errors.New("video_not_find")
	}

	if number_of_lesson <= 0 {
		number_of_lesson = 1
	}
	var countOfLessons int = len(videos)
	if number_of_lesson > countOfLessons {
		return nil, errors.New("number_of_lesson_out_of_range")
	}
	var video entity.Video = videos[number_of_lesson-1]

	var idsOfVideos []int
	for _, video := range videos {
		idsOfVideos = append(idsOfVideos, video.Id)
	}

	var videoSrc *string
	if video.Hash != nil {
		videoSrc = getVideoSrc(video.Id, *(video.Hash))
	}

	var videoCourseLessonViewModel videoViewModel.VideoCourseLessonViewModel = videoViewModel.VideoCourseLessonViewModel{
		Id:             video.Id,
		Id_of_course:   id_of_course,
		NumberOfLesson: number_of_lesson,
		IdsOfVideos:    idsOfVideos,
		VideoSrc:       *videoSrc,
	}

	return &videoCourseLessonViewModel, nil
}

func getVideoSrc(id_of_video int, hash string) *string {
	var path string = "static/XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" + strconv.Itoa(id_of_video) + "_" + hash + "/playlist.m3u8"
	infoOfExistFile, err := os.Stat("./" + path)
	if os.IsNotExist(err) || infoOfExistFile.IsDir() {
		return nil
	}
	return &path
}
