// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"study-go/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/*", &controllers.OptionsController{}, "options:Options")

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/admin_menu",
			beego.NSInclude(
				&controllers.AdminMenuController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/course_category",
			beego.NSInclude(
				&controllers.CourseCategoryController{},
			),
		),

		beego.NSNamespace("/course_category_rel",
			beego.NSInclude(
				&controllers.CourseCategoryRelController{},
			),
		),

		beego.NSNamespace("/course_favorite",
			beego.NSInclude(
				&controllers.CourseFavoriteController{},
			),
		),

		beego.NSNamespace("/course_lesson",
			beego.NSInclude(
				&controllers.CourseLessonController{},
			),
		),

		beego.NSNamespace("/course_lesson_learn",
			beego.NSInclude(
				&controllers.CourseLessonLearnController{},
			),
		),

		beego.NSNamespace("/course_member",
			beego.NSInclude(
				&controllers.CourseMemberController{},
			),
		),

		beego.NSNamespace("/course_teacher",
			beego.NSInclude(
				&controllers.CourseTeacherController{},
			),
		),

		beego.NSNamespace("/live",
			beego.NSInclude(
				&controllers.LiveController{},
			),
		),

		beego.NSNamespace("/live_appointment",
			beego.NSInclude(
				&controllers.LiveAppointmentController{},
			),
		),

		beego.NSNamespace("/live_course",
			beego.NSInclude(
				&controllers.LiveCourseController{},
			),
		),

		beego.NSNamespace("/live_video",
			beego.NSInclude(
				&controllers.LiveVideoController{},
			),
		),

		beego.NSNamespace("/navigation",
			beego.NSInclude(
				&controllers.NavigationController{},
			),
		),

		beego.NSNamespace("/notify",
			beego.NSInclude(
				&controllers.NotifyController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),

		beego.NSNamespace("/permission",
			beego.NSInclude(
				&controllers.PermissionController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/role_permission",
			beego.NSInclude(
				&controllers.RolePermissionController{},
			),
		),

		beego.NSNamespace("/role_user",
			beego.NSInclude(
				&controllers.RoleUserController{},
			),
		),

		beego.NSNamespace("/setting",
			beego.NSInclude(
				&controllers.SettingController{},
			),
		),


		beego.NSNamespace("/user_notify",
			beego.NSInclude(
				&controllers.UserNotifyController{},
			),
		),

		beego.NSNamespace("/user_profile",
			beego.NSInclude(
				&controllers.UserProfileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
