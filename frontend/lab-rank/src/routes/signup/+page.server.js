
import { redirect } from '@sveltejs/kit';

export const load = ({cookies}) => {
 const jwt = cookies.get("jwt-lab-rank")
  if (jwt != undefined) {
   throw redirect(303, '/subjects') 
  } 
  var universities = []
  const fetchUniversities = async () => {
    try {
      const response = await fetch("http://localhost:8080/university/names");
      const data = await response.json();
      universities = data.Message;
      console.log(universities);
    } catch (error) {
      console.error("Error fetching universities:", error);
    }
  };

  
  

  return {
    universities
  };

}