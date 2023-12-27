export const load = ({cookies}) => {
    let user_not_signin = true
  
    let jwt = cookies.get("jwt-lab-rank")
    if (jwt != undefined) {
        user_not_signin = false
    } 
    
    console.log(user_not_signin)
  
    return {
      user_not_signin
    };
  
  }