
export const load = async ({cookies}) => {
    const jwt = cookies.get("jwt-lab-rank")
    const universityId = cookies.get("university-id")
    const collegeId = cookies.get("college-id")
    
     let user_not_signin = true
     if (jwt != undefined) { 
        user_not_signin = false
     }

     let subjects = []
     const fetchSubjects = async () => {
        try {
          console.log("server fetchUniversity")
          const response = await fetch(`http://127.0.0.1:8080/subject/${universityId}`);
          const data = await response.json();
          subjects = data.Message;
          console.log(subjects);
        } catch (error) {
          console.error("Error fetching universities:", error);
        }
      };
    
      
      await fetchSubjects();
     return {
       user_not_signin,
       subjects
     };
   
   }