import { createBrowserRouter } from "react-router-dom";
import Login from "../screens/Login";
import Register from "../screens/Register";
import Home from "../screens/Home";
import Splash from "../screens/Splash";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Splash />,
    errorElement: <Login/>,
    children: [
      {
        path: "/auth/login",
        element: <Login />,
      },
      {
        path: "/app",
        element: <Home />,
      },
      {
        path: "/auth/register",
        element: <Register />,
      },
    ]
  },
]);

export default router;