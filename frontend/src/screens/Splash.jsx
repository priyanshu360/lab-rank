import { useState, useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { getCookie } from './../Utils'

const Splash = function(){
  const [loading, changeLoading] = useState(true);
  const [user, setUser] = useState(null);
  const navigate = useNavigate();
  useEffect(() => {
    if (getCookie("loginKey") != ""){
      // verify Data from api and use it.
      setTimeout(() => {
        debugger;
        setUser({name: "User", authKey: "test"})
        changeLoading(false);
      }, 5000);
    } else {
      setTimeout(() => {
        setUser({name: undefined, authKey: undefined});
        if(!(window.location.href).includes('auth'))navigate('/auth/login')
        changeLoading(false);
      }, 1000);
    }
    // check cookie or local storeag
    // setuser({name: undefined, authKey: undefined});
    // setUser({name: vipul})
  },[]);


  // setTimeout(()=> {console.log("here");changeLoading(false)}, 2000);
  return loading
          ? <DefaultScreen />
          : <div height="100%" width="100%">
              <Outlet />
            </div>
}

const DefaultScreen = () => {
  return <div height="100%" width="100%"></div>
}


export default Splash;