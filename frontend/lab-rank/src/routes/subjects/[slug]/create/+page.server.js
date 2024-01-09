
export const load = ({ locals, params, cookies }) => {
    let user_not_signin = true

    let jwt = cookies.get("jwt_lab_rank")
    let collegeID = locals.user.college_id
    let subjectID = params.slug
    if (jwt != undefined) {
        user_not_signin = false
    }

    console.log(user_not_signin)

    return {
        user_not_signin,
        collegeID,
        subjectID
    };

}