package biz

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// CourseUsecase is a Course usecase.
type CourseUsecase struct {
	repo                 CourseRepo
	shopRepo             ShopRepo
	studentRepo          StudentRepo
	studentCourseRepo    StudentCourseRepo
	courseSchedulingRepo CourseSchedulingRepo
	classSchedulingRepo  ClassSchedulingRepo
	log                  *log.Helper
}

// NewCourseUsecase new a Course usecase.
func NewCourseUsecase(repo CourseRepo, studentCourseRepo StudentCourseRepo, studentRepo StudentRepo, courseSchedulingRepo CourseSchedulingRepo, classSchedulingRepo ClassSchedulingRepo, shopRepo ShopRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{
		repo: repo,

		shopRepo:             shopRepo,
		studentRepo:          studentRepo,
		studentCourseRepo:    studentCourseRepo,
		courseSchedulingRepo: courseSchedulingRepo,
		classSchedulingRepo:  classSchedulingRepo,
		log:                  log.NewHelper(logger),
	}
}

type Course struct {
	*model.Course
	UseShopStr string
}

func (uc *CourseUsecase) FindById(ctx context.Context, id int64) (*Course, error) {
	data, err := uc.repo.FindByID(ctx, id)

	shopNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListAll(ctx)
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}

	var shopNameArr []string
	shopArr := strings.Split(data.UseShop, "|")
	if shopArr[0] == "shopId" {
		for i := 1; i < len(shopArr); i++ {
			if shopArr[i] != "" {
				id, _ := strconv.Atoi(shopArr[i])
				if id != 0 {
					shopNameArr = append(shopNameArr, shopNameMap[int64(id)])
				} else {
					shopNameArr = append(shopNameArr, "通用")
				}
			}

		}
	}

	if err != nil {
		return nil, err
	}
	return &Course{
		Course:     data,
		UseShopStr: strings.Join(shopNameArr, ","),
	}, nil
}

func (uc *CourseUsecase) FindBy(ctx context.Context, key string, value string) (*Course, error) {
	data, err := uc.repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &Course{Course: data}, nil

}

func (uc *CourseUsecase) List(ctx context.Context) ([]*Course, error) {
	data, err := uc.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var list []*Course
	for _, item := range data {
		list = append(list, &Course{Course: item})
	}
	return list, nil
}

type SearchCourseRequest struct {
	CId          []string
	MerchantId   int64
	TeachingType int32
	CourseType   int32
	Status       int32
	UseShop      string
	Sort         int
}

func (uc *CourseUsecase) Search(ctx context.Context, req *SearchCourseRequest) ([]*Course, error) {
	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Course

	shopNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListAll(ctx)
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}

	for _, item := range data {
		var shopNameArr []string
		shopArr := strings.Split(item.UseShop, "|")
		if shopArr[0] == "shopId" {
			for i := 1; i < len(shopArr); i++ {
				if shopArr[i] != "" {
					id, _ := strconv.Atoi(shopArr[i])
					if id != 0 {
						shopNameArr = append(shopNameArr, shopNameMap[int64(id)])
					} else {
						shopNameArr = append(shopNameArr, "通用")
					}
				}
			}
		}

		list = append(list, &Course{
			Course:     item,
			UseShopStr: strings.Join(shopNameArr, ","),
		})
	}
	return list, nil
}

func (uc *CourseUsecase) BuyCourseValidate(ctx context.Context, req *CreateBuyCourseOrderRequest) (ok bool, err error) {

	return true, nil
}

type CreateStudentCourseRequest struct {
	Idx          int32 //创建时的索引
	MerchantId   int64
	ShopId       int64
	MemberId     int64
	CourseId     int64
	SaleId       int64
	StudentId    int64
	Qty          int32 //数量
	OrderNo      string
	IsActive     bool
	ActiveAmount float64
}

func (uc *CourseUsecase) CreateStudentCourse(ctx context.Context, req *CreateStudentCourseRequest) (id int64, err error) {

	saleInfo, err2 := uc.repo.FindSaleByID(ctx, req.SaleId)
	if err2 != nil {
		log.Errorf("Service GetSrPackageCard error:%s \r\n", err2)
		return 0, err
	}
	createBy := uint(req.MemberId)
	model := &model.StudentCourse{
		EdStudentCourses: gen.EdStudentCourses{
			ID:              id,
			MerchantID:      req.MerchantId,
			ShopID:          req.ShopId,
			StudentID:       req.StudentId,
			ActualAmount:    req.ActiveAmount,
			CourseID:        saleInfo.CourseID,
			Title:           saleInfo.Course.Title,
			CourseType:      saleInfo.Course.CourseType,
			TeachingType:    saleInfo.Course.TeachingType,
			CourseSize:      saleInfo.Course.CourseSize,
			StartDate:       saleInfo.Course.StartDate,
			EndDate:         saleInfo.Course.EndDate,
			CourseSaleID:    saleInfo.ID,
			CourseSaleType:  saleInfo.Type,
			Rule:            saleInfo.Rule,
			Price:           saleInfo.Price,
			OriPrice:        saleInfo.OriPrice,
			ActiveAt:        &time.Time{},
			InvalidAt:       &time.Time{},
			ValidDay:        saleInfo.ValidDay,
			ActiveLimit:     saleInfo.ActiveLimit,
			Sort:            0,
			CourseSalemanID: id,
			Remark:          "",
			CreateBy:        &createBy,
		},
	}
	if saleInfo.Type == 99 {
		model.Rule = saleInfo.Rule
	}
	model.Status = 1
	t1, _ := time.ParseInLocation(time.DateTime, "1991-07-31 00:00:00", time.Local)
	model.ActiveAt = &t1
	model.InvalidAt = &t1

	type rule struct {
		M int
		N int
	}
	var ruleData rule
	json.Unmarshal([]byte(saleInfo.Rule), &ruleData)
	model.Remain = ruleData.M
	model.PerMin = ruleData.N

	_, err = uc.studentCourseRepo.Save(ctx, model)
	if err != nil {
		log.Errorf("CreateStudentCourse Insert error:%s \r\n", err)
		return 0, err
	}
	return model.ID, nil
}

type SaveStudentRequest struct {
	Id        int64
	StudentId int64
	Name      string
	Phone     string
	Gender    int8
	Photo     string
}

func (uc *CourseUsecase) CreateStudent(ctx context.Context, memberId int64, student *SaveStudentRequest) (*model.MyStudent, error) {
	model := &model.MyStudent{
		EdStudentRelateMember: gen.EdStudentRelateMember{
			MemberID: memberId,
		},
		Student: &gen.EdStudentInfo{
			Name:   student.Name,
			Phone:  student.Phone,
			Gender: student.Gender,
			Photo:  &student.Photo,
		},
	}

	_, err := uc.studentRepo.Save(ctx, model)
	if err != nil {
		log.Errorf("CreateStudent Insert error:%s \r\n", err)
		return model, err
	}
	return model, err
}
func (uc *CourseUsecase) UpdateStudent(ctx context.Context, memberId int64, student *SaveStudentRequest) (*model.MyStudent, error) {
	model := &model.MyStudent{
		EdStudentRelateMember: gen.EdStudentRelateMember{
			MemberID: memberId,
			ID:       student.Id,
		},
		Student: &gen.EdStudentInfo{
			ID:     int64(student.StudentId),
			Name:   student.Name,
			Phone:  student.Phone,
			Gender: student.Gender,
			Photo:  &student.Photo,
		},
	}

	_, err := uc.studentRepo.Update(ctx, model)
	if err != nil {
		log.Errorf("CreateStudent Insert error:%s \r\n", err)
		return model, err
	}
	return model, err
}

func (uc *CourseUsecase) FindStudentById(ctx context.Context, id int64) (*model.MyStudent, error) {
	data, err := uc.studentRepo.FindByID(ctx, id)
	return data, err
}

func (uc *CourseUsecase) ListStudent(ctx context.Context, req *SearchStudentRequest) ([]*model.MyStudent, error) {
	list, err := uc.studentRepo.ListBy(ctx, req)
	return list, err
}

func (uc *CourseUsecase) ListStudentCourse(ctx context.Context, req *SearchStudentCourseRequest) ([]*model.StudentCourse, error) {
	list, err := uc.studentCourseRepo.ListBy(ctx, req)
	return list, err
}
func (uc *CourseUsecase) ListCourseScheduling(ctx context.Context, req *SearchSchedulingRequest) ([]*model.CourseScheduling, error) {
	list, err := uc.courseSchedulingRepo.ListBy(ctx, req)
	return list, err
}

func (uc *CourseUsecase) FindCourseSchedulingById(ctx context.Context, id int64) (*gen.EdCoursesScheduling, error) {
	data, err := uc.courseSchedulingRepo.FindByID(ctx, id)
	return data, err
}

func (uc *CourseUsecase) FindCourseSchedulingDetailById(ctx context.Context, id int64) (*model.CourseScheduling, error) {
	data, err := uc.courseSchedulingRepo.FindDetailByID(ctx, id)
	return data, err
}
func (uc *CourseUsecase) UpdateCourseScheduling(ctx context.Context, model *gen.EdCoursesScheduling) (bool, error) {
	ok, err := uc.courseSchedulingRepo.Update(ctx, model)
	return ok, err
}

func (uc *CourseUsecase) RecordCheckInLog(course StudentCourseCheckIn, by int) (bool, error) {
	ok, err := uc.courseSchedulingRepo.RecordCheckInLog(course, by)
	return ok, err
}

func (uc *CourseUsecase) RecordLog(ctx context.Context, oldModel *gen.EdCoursesScheduling, model *gen.EdCoursesScheduling, by int) (bool, error) {
	ok, err := uc.courseSchedulingRepo.RecordLog(oldModel, model, by)
	return ok, err
}
func (uc *CourseUsecase) ListClassScheduling(ctx context.Context, req *SearchSchedulingRequest) ([]*model.ClassScheduling, error) {
	list, err := uc.classSchedulingRepo.ListBy(ctx, req)
	return list, err
}

func (uc *CourseUsecase) FindClassSchedulingById(ctx context.Context, id int64) (*gen.EdClassScheduling, error) {
	data, err := uc.classSchedulingRepo.FindByID(ctx, id)
	return data, err
}

func (uc *CourseUsecase) UpdateClassScheduling(ctx context.Context, model *gen.EdClassScheduling) (bool, error) {
	ok, err := uc.classSchedulingRepo.Update(ctx, model)
	return ok, err
}
