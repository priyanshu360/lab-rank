
import { redirect } from '@sveltejs/kit';
import { format } from "date-fns";


export const load = async ({cookies}) => {
 const jwt = cookies.get("jwt-lab-rank")
  if (jwt != undefined) {
   throw redirect(303, '/subjects') 
  } 
  var universities = []
  console.log("hello signup load")
  const fetchUniversities = async () => {
    try {
      console.log("server fetchUniversity")
      const response = await fetch("http://127.0.0.1:8080/university/names");
      const data = await response.json();
      universities = data.Message;
      console.log(universities);
    } catch (error) {
      console.error("Error fetching universities:", error);
    }
  };

  
  let user_not_signin = true
  await fetchUniversities();
  return {
    universities,
    user_not_signin
  };

}

export const actions = {
	create: async ({ fetch, request }) => {
		const data = await request.formData();
    console.log(data)
    const college_id = data.get("college_id");
    const email = data.get("email");
    const contact_no = data.get("contact_no");
    const dob = format(new Date(data.get("dob")), "yyyy-MM-dd'T'HH:mm:ssXXX");
    const university_id = data.get("university_id");
    const name = data.get("name");
    const user_name = data.get("user_name");
    const password = data.get("password");
    console.log(college_id, email)
    fetch("http://127.0.0.1:8080/auth/signup", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    college_id,
    email,
    contact_no,
    dob,
    university_id,
    name,
    user_name,
    password,
    // Add other variables here if needed
  }),
})
  .then((response) => {
    console.log(response.ok);
  if (response.ok) {
    // throw redirect(303, '/login') 
  }
  })
   	throw redirect(303, '/login');  // cookies.set("jwt-lab-rank", jwt.Message)
  }

};

