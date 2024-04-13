package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"
	"time"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"

	comJwt "studyRoomGo/common/jwt"

	pb "studyRoomGo/api/course/v1"
)

type CourseService struct {
	pb.UnimplementedCourseServer

	uc *biz.CourseUsecase
}

func NewCourseService(uc *biz.CourseUsecase) *CourseService {
	return &CourseService{uc: uc}
}

func (s *CourseService) Get(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	data, err := s.uc.FindById(ctx, req.Id)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dto := dto.Course{Course: data}
	relpy := dto.Reply()
	return &pb.DataReply{
		Code: 200,
		Data: relpy,
	}, nil
}

func (s *CourseService) Find(ctx context.Context, req *pb.FindRequest) (*pb.DataReply, error) {
	var filterKey = []string{
		"id", "name",
	}
	if !_func.InStrSplice(filterKey, req.Key) {
		return &pb.DataReply{
			Code: 400,
			Msg:  fmt.Sprintf("key not in %v", strings.Join(filterKey, ",")),
		}, nil
	}
	data, err := s.uc.FindBy(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dto := dto.Course{Course: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *CourseService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {

	data, err := s.uc.List(ctx)
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto := dto.Course{Course: item}
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CourseService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.ListReply, error) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	data, err := s.uc.Search(ctx, &biz.SearchCourseRequest{
		MerchantId: int64(merchantId),
		Status:     2,
		UseShop:    req.UseShop,
		Sort:       int(req.Sort),
	})
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto := dto.Course{Course: item}
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CourseService) MyStudentList(ctx context.Context, req *pb.MyStudentRequest) (*pb.MyStudentDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.ListStudent(ctx, &biz.SearchStudentRequest{
		MemberId: int64(memberId),
	})
	if err != nil {
		return &pb.MyStudentDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Student{}
	var list []*pb.MyStudentData
	for _, item := range data {
		d.Student = item
		list = append(list, d.Reply())
	}
	return &pb.MyStudentDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CourseService) MyStudentInfo(ctx context.Context, req *pb.MyStudentRequest) (*pb.MyStudentDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.FindStudentById(ctx, req.Id)
	if err != nil {
		return &pb.MyStudentDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if data.MemberID != int64(memberId) {
		return &pb.MyStudentDataReply{
			Code: 400,
			Msg:  "无权限查看该学员信息",
		}, nil
	}
	d := dto.Student{}
	d.Student = data
	return &pb.MyStudentDataReply{
		Code: 200,
		Data: d.Reply(),
	}, nil
}
func (s *CourseService) CreateMyStudent(ctx context.Context, req *pb.MyStudentRequest) (*pb.MyStudentDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	d := dto.Student{}

	data, err := s.uc.CreateStudent(ctx, memberId, &biz.SaveStudentRequest{
		Name:   req.Name,
		Phone:  req.Phone,
		Gender: int8(req.Gender),
		Photo:  req.Photo,
	})
	if err != nil {
		return &pb.MyStudentDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d.Student = data
	return &pb.MyStudentDataReply{
		Code: 200,
		Data: d.Reply(),
	}, nil
}

func (s *CourseService) UpdateMyStudent(ctx context.Context, req *pb.MyStudentRequest) (*pb.MyStudentDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	d := dto.Student{}

	data, err := s.uc.UpdateStudent(ctx, memberId, &biz.SaveStudentRequest{
		Id:        req.Id,
		StudentId: req.StudentId,
		Name:      req.Name,
		Phone:     req.Phone,
		Photo:     req.Photo,
	})
	if err != nil {
		return &pb.MyStudentDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d.Student = data
	return &pb.MyStudentDataReply{
		Code: 200,
		Data: d.Reply(),
	}, nil
}
func (s *CourseService) MyStudentCourse(ctx context.Context, req *pb.MyCourseRequest) (*pb.MyStudentCourseDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.ListStudentCourse(ctx, &biz.SearchStudentCourseRequest{
		MemberId: int64(memberId),
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	})

	if err != nil {
		return &pb.MyStudentCourseDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Student{}

	var list []*pb.MyStudentCourseData
	for _, item := range data {
		d.StudentCourse = item

		list = append(list, d.StudentCourseReply())
	}
	return &pb.MyStudentCourseDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CourseService) MyCourseScheduling(ctx context.Context, req *pb.MyCourseRequest) (*pb.MyCourseSchedulingDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	var status []int32
	if req.Status != "" {
		_ = json.Unmarshal([]byte(req.Status), &status)
	}
	data, err := s.uc.ListCourseScheduling(ctx, &biz.SearchSchedulingRequest{
		IsStudent: true,
		MemberId:  int64(memberId),
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		Status:    status,
		Sort:      int(req.Sort),
	})
	if err != nil {
		return &pb.MyCourseSchedulingDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Student{}

	var list []*pb.MyCourseSchedulingData
	for _, item := range data {
		d.CourseScheduling = item

		list = append(list, d.CourseSchedulingReply())
	}
	return &pb.MyCourseSchedulingDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

/// 老师端

func (s *CourseService) TeacherCourseScheduling(ctx context.Context, req *pb.MyCourseRequest) (*pb.MyCourseSchedulingDataListReply, error) {
	if ok, err := comJwt.CheckIsTeacher(ctx); !ok {
		return &pb.MyCourseSchedulingDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	var status []int32
	if req.Status != "" {
		_ = json.Unmarshal([]byte(req.Status), &status)
	}
	classId := int64(0)
	data, err := s.uc.ListCourseScheduling(ctx, &biz.SearchSchedulingRequest{
		IsTeacher: true,
		ClassId:   &classId,
		MemberId:  int64(memberId),
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		Status:    status,
		Sort:      int(req.Sort),
	})
	if err != nil {
		return &pb.MyCourseSchedulingDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Student{}

	var list []*pb.MyCourseSchedulingData
	for _, item := range data {
		d.CourseScheduling = item

		list = append(list, d.CourseSchedulingReply())
	}
	return &pb.MyCourseSchedulingDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CourseService) TeacherClassScheduling(ctx context.Context, req *pb.MyCourseRequest) (*pb.MyClassSchedulingDataListReply, error) {
	if ok, err := comJwt.CheckIsTeacher(ctx); !ok {
		return &pb.MyClassSchedulingDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	var status []int32
	if req.Status != "" {
		_ = json.Unmarshal([]byte(req.Status), &status)
	}
	data, err := s.uc.ListClassScheduling(ctx, &biz.SearchSchedulingRequest{
		IsTeacher: true,
		MemberId:  int64(memberId),
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
		Status:    status,
		Sort:      int(req.Sort),
	})
	if err != nil {
		return &pb.MyClassSchedulingDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Student{}

	var list []*pb.MyClassSchedulingData
	for _, item := range data {
		d.ClassScheduling = item

		list = append(list, d.ClassSchedulingReply())
	}
	return &pb.MyClassSchedulingDataListReply{
		Code: 200,
		Data: list,
	}, nil
}
func (s *CourseService) TeacherCourseSchedulingCheckIn(ctx context.Context, req *pb.CourseSchedulingCheckInRequest) (*pb.CheckInReply, error) {
	if ok, err := comJwt.CheckIsTeacher(ctx); !ok {
		return &pb.CheckInReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	scheduling, err := s.uc.FindCourseSchedulingDetailById(ctx, req.Id)
	oldData := *scheduling
	if err != nil {
		return &pb.CheckInReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	scheduling.CheckIn = int8(req.CheckIn)
	scheduling.Status = int8(req.Status)
	if oldData.Status != int8(req.Status) || oldData.CheckIn != int8(req.CheckIn) {

		if ok, err := s.uc.UpdateCourseScheduling(ctx, &scheduling.EdCoursesScheduling); !ok {
			return &pb.CheckInReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		s.uc.RecordLog(ctx, &oldData.EdCoursesScheduling, &scheduling.EdCoursesScheduling, int(memberId))

		now := time.Now()

		//微信通知
		wx, _ := wechat.NewMessage(wechat.MerchantToOfficalMerchantId(scheduling.MerchantID))
		if scheduling.CheckIn == 2 {
			err = wx.SendStudentCheckInSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckInSuccess{
				CourseName:  scheduling.StudentCourses.Title,
				StudentName: scheduling.Student.Name,
				AddressName: scheduling.Address.Name,
				CourseTime:  scheduling.StartDate.Format(time.DateOnly) + " " + scheduling.StartTime,
				CheckIntime: now.Format(time.DateTime),
			})
		}
		if scheduling.CheckIn == 3 {
			err = wx.SendStudentCheckOutSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckOutSuccess{
				StudentName:  scheduling.Student.Name,
				CheckOutDate: now.Format(time.DateOnly),
				CheckOutTime: now.Format(time.TimeOnly),
			})
		}
		s.uc.RecordCheckInLog(biz.StudentCourseCheckIn{
			Id:              int(scheduling.ID),
			StudentCourseId: int(scheduling.StudentCourseID),
			StudentId:       int(scheduling.StudentID),
			OldCheckIn:      int(oldData.CheckIn),
			NewCheckIn:      int(scheduling.CheckIn),
			Err:             err,
		}, int(memberId))
	}
	return &pb.CheckInReply{
		Code: 200,
	}, nil
}

func (s *CourseService) TeacherClassSchedulingCheckIn(ctx context.Context, req *pb.ClassSchedulingCheckInRequest) (*pb.CheckInReply, error) {
	if ok, err := comJwt.CheckIsTeacher(ctx); !ok {
		return &pb.CheckInReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	scheduling, err := s.uc.FindClassSchedulingById(ctx, req.Id)
	oldData := *scheduling
	if err != nil {
		return &pb.CheckInReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	now := time.Now()

	scheduling.Status = int8(req.Status)
	if oldData.Status != int8(req.Status) {
		if ok, err := s.uc.UpdateClassScheduling(ctx, scheduling); !ok {
			return &pb.CheckInReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}

	}

	for _, courseScheduling := range req.Course {
		scheduling, err := s.uc.FindCourseSchedulingDetailById(ctx, courseScheduling.Id)
		oldData := *scheduling

		if err != nil {
			return &pb.CheckInReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		if oldData.EdCoursesScheduling.CheckIn != int8(courseScheduling.CheckIn) || oldData.EdCoursesScheduling.Status != int8(req.Status) {
			scheduling.CheckIn = int8(courseScheduling.CheckIn)
			scheduling.Status = int8(req.Status)
			if ok, err := s.uc.UpdateCourseScheduling(ctx, &scheduling.EdCoursesScheduling); !ok {
				return &pb.CheckInReply{
					Code: 400,
					Msg:  err.Error(),
				}, nil
			}
			s.uc.RecordLog(ctx, &oldData.EdCoursesScheduling, &scheduling.EdCoursesScheduling, int(memberId))

			//微信通知
			wx, _ := wechat.NewMessage(wechat.MerchantToOfficalMerchantId(scheduling.MerchantID))
			if scheduling.CheckIn == 2 {
				_ = wx.SendStudentCheckInSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckInSuccess{
					CourseName:  scheduling.StudentCourses.Title,
					StudentName: scheduling.Student.Name,
					AddressName: scheduling.Address.Name,
					CourseTime:  scheduling.StartDate.Format(time.DateOnly) + " " + scheduling.StartTime,
					CheckIntime: now.Format(time.DateTime),
				})
			}
			if scheduling.CheckIn == 3 {
				_ = wx.SendStudentCheckOutSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckOutSuccess{
					StudentName:  scheduling.Student.Name,
					CheckOutDate: now.Format(time.DateOnly),
					CheckOutTime: now.Format(time.TimeOnly),
				})
			}
			s.uc.RecordCheckInLog(biz.StudentCourseCheckIn{
				Id:              int(scheduling.ID),
				StudentCourseId: int(scheduling.StudentCourseID),
				StudentId:       int(scheduling.StudentID),
				OldCheckIn:      int(oldData.CheckIn),
				NewCheckIn:      int(scheduling.CheckIn),
				Err:             err,
			}, int(memberId))
		}
	}

	return &pb.CheckInReply{
		Code: 200,
	}, nil
}
