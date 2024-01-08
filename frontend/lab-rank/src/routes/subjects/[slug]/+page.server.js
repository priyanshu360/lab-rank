import { redirect } from '@sveltejs/kit';

export const load = ({ params, cookies }) => {
    let user_not_signin = true

    let jwt = cookies.get("jwt_lab_rank")
    let collegeID = cookies.get("college-id")
    let subjectID = params.slug
    if (jwt != undefined) {
        user_not_signin = false
    } else {
        throw redirect(303, "/signup")
    }

    console.log(user_not_signin)

    return {
        user_not_signin,
        collegeID,
        subjectID
    };

}