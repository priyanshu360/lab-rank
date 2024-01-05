import { redirect } from '@sveltejs/kit';

export const load = ({cookies}) => {
  let user_not_signin = true

  let jwt = cookies.get("jwt-lab-rank")
  if (jwt != undefined) {
      user_not_signin = false
  } 
  
     if (jwt != undefined) {
      throw redirect(303, '/subjects') 
     }
  console.log(user_not_signin)

  return {
    user_not_signin
  };

}


export const actions = {
	create: async ({ fetch, cookies, request }) => {
		const data = await request.formData();
    console.log(data)
    const email = data.get("email")
    const password= data.get("password")
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
    const authData = await res.json()
    console.log(authData)
    cookies.set("jwt-lab-rank", authData.Jwt) // TODO : Set university and College 
    cookies.set("college-id",authData.CollegeID)
    cookies.set("university-id",authData.UniversityID)
    cookies.set("user-id",authData.UserID)
    
  }
  
};

