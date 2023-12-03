// import React, { useEffect, useState } from 'react';
// import {
//   createBrowserRouter,
//   RouterProvider,
//   useLocation,
//   useNavigate,
//   Navigate,
//   Outlet
// } from "react-router-dom";
// import Login from './Login';
// import Home from './Home';
// import Register from './Register';
// import { getCookie } from './Utils';

// const App = function() {
//   const [isLoading, setLoading] = useState(true);
//   const [user, setuser] = useState(null);
//   // const location = useLocation()
//   // const navigate = useNavigate();
//   useEffect(() => {
//     if(user && user.name){
//       console.log("this is true")
//       window.location.href = '/app';
//       // setLoading(false);
//     } else {
//       if(!window.location.href.includes('auth')){
//         window.location.href = '/auth/login';
//         // setLoading(false);
//       } else {
//         // setLoading(false);
//       }
//     }
//   },[user])

//   useEffect(() => {
//     if (getCookie("loginKey") != ""){
//       // verify Data from api and use it.
//       setTimeout(() => {
//         debugger;
//         setuser({name: "User", authKey: "test"})
//         setLoading(false);
//       }, 5000);
//     } else {
//       setTimeout(() => {
//         setuser({name: undefined, authKey: undefined});
//         setLoading(false);
//       }, 1000);
//     }
//     // check cookie or local storeag
//     // setuser({name: undefined, authKey: undefined});
//     // setUser({name: vipul})
//   },[]);

//   return (<div>
//     {isLoading ? <h1>Loading</h1> ://Splash
//     <RouterProvider router={router} />}
//   </div>)
// }

const App = () => {return <h1>Hello 2</h1> }

export default App;