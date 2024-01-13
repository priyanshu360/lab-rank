
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, locals }) => {
    const jwt = cookies.get("jwt_lab_rank")
    const userID = locals.user.id

    let user_not_signin = true
    if (jwt != undefined) {
        user_not_signin = false
    } else {
        // throw redirect(303, "/signup")
    }

    let submission = []
    const fetchSubmission = async () => {
        try {
            console.log("server fetchUniversity")
            const response = await fetch(`http://127.0.0.1:8080/submission/user/${userID}`);
            const data = await response.json();
            submission = data.Message;
            console.log(submission);
        } catch (error) {
            console.error("Error fetching universities:", error);
        }
    };


    await fetchSubmission();
    return {
        user_not_signin,
        submission
    };

}