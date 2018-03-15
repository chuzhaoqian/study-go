package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["study-go/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseCategoryRelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseFavoriteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseLessonLearnController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseMemberController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"] = append(beego.GlobalControllerRouter["study-go/controllers:CourseTeacherController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveAppointmentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"] = append(beego.GlobalControllerRouter["study-go/controllers:LiveVideoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:LoginController"] = append(beego.GlobalControllerRouter["study-go/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NavigationController"] = append(beego.GlobalControllerRouter["study-go/controllers:NavigationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NavigationController"] = append(beego.GlobalControllerRouter["study-go/controllers:NavigationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NavigationController"] = append(beego.GlobalControllerRouter["study-go/controllers:NavigationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NavigationController"] = append(beego.GlobalControllerRouter["study-go/controllers:NavigationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NavigationController"] = append(beego.GlobalControllerRouter["study-go/controllers:NavigationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:NotifyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:NotifyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:NotifyController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:NotifyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:NotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:NotifyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:OrderController"] = append(beego.GlobalControllerRouter["study-go/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:OrderController"] = append(beego.GlobalControllerRouter["study-go/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:OrderController"] = append(beego.GlobalControllerRouter["study-go/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:OrderController"] = append(beego.GlobalControllerRouter["study-go/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:OrderController"] = append(beego.GlobalControllerRouter["study-go/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:PermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:PermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:PermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:PermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:PermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"] = append(beego.GlobalControllerRouter["study-go/controllers:RolePermissionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleUserController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleUserController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleUserController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleUserController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:RoleUserController"] = append(beego.GlobalControllerRouter["study-go/controllers:RoleUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["study-go/controllers:SettingController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["study-go/controllers:SettingController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["study-go/controllers:SettingController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["study-go/controllers:SettingController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:SettingController"] = append(beego.GlobalControllerRouter["study-go/controllers:SettingController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserNotifyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study-go/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["study-go/controllers:UserProfileController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
