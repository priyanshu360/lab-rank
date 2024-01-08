import { redirect } from '@sveltejs/kit';


export async function load({ params, cookies }) {
  console.log("Load function called!");
  const slug = params.slug
  let jwt = cookies.get("jwt_lab_rank")
  let user_not_signin = true
  if (jwt != undefined) {
    user_not_signin = false
  } else {
    // throw redirect(303, "/signup")
  }

  const response = await fetch(
    `http://127.0.0.1:8080/problem?id=${slug}`
  );
  const responseData = await response.json();

  // const problem_title = responseData.Message.title;
  // const problem_file = atob(responseData.Message.problem_file);
  // console.log(problem_title, problem_file)
  return {
    slug,
    user_not_signin,
    responseData
  };

}


export const actions = {
  create: async ({ fetch, cookies, request }) => {
    const data = await request.formData();
    console.log(data)
    const email = data.get("email")
    const password = data.get("password")
    const res = await fetch("http://127.0.0.1:8080/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });
    const jwt = await res.json()
    console.log(jwt.Message)
    cookies.set("jwt_lab_rank", jwt.Message)
  }

};
