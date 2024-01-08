
import { redirect } from '@sveltejs/kit';

export const load = async ({ locals, cookies, fetch }) => {
    const jwt = cookies.get("jwt_lab_rank")
    const universityId = locals.user.university_id

    console.log("event", locals.user)

    const collegeId = locals.user.college_id

    let user_not_signin = true
    if (jwt != undefined) {
        user_not_signin = false
    } else {
        // throw redirect(303, "/signup")
    }

    console.log(collegeId, universityId)

    let subjects = []
    const fetchSubjects = async () => {
        try {
            const response = await fetch(`http://127.0.0.1:8080/subject/${universityId}`);
            const data = await response.json();
            subjects = data.Message;
            console.log("subjects in server.js ",subjects);
        } catch (error) {
            console.error("Error fetching Subjects:", error);
        }
    };
    await fetchSubjects();

    var universities = []
    const fetchUniversities = async () => {
      try {
        console.log("server fetch Universities")
        const response = await fetch("http://127.0.0.1:8080/university/names");
        const data = await response.json();
        universities = data.Message;
      } catch (error) {
        console.error("Error fetching universities:", error);
      }
    };
    await fetchUniversities();

    return {
        user_not_signin,
        subjects,
        universities,
    };

}