export  function load({ params, cookies }) {
    console.log("Load function called!");
    const slug = params.slug
    let jwt = cookies.get("jwt-lab-rank")
    let user_not_signin = true
    if (jwt != undefined) {
        user_not_signin = false
    } 
      return {
           slug,
           user_not_signin
      };
   
  }