import { useState } from 'react'
import { Navigate } from 'react-router-dom';
import {setCookie, getCookie} from "./Utils";

function Login() {
  let loginDef = false
  if (getCookie("loginKey") != ""){
    loginDef = true
  }
  const [isLoggedIn, changeLoggedIn] = useState(loginDef)
  // const [isLoggedIn, changeLoggedIn] = useState(getCookie("loginKey") == "")
  const [ev, changeErrorVisibility] = useState(false)
  
  const [email, setEmail] = useState('test');
  const [password, setPassword] = useState('1234');
  const handleChange = (event) => {
    setEmail(event.target.value);
  };
  const handlePassChange = (event) => {
    setPassword(event.target.value);
  };
  const handleClick = () => {
    if(email === "test" && password === "1234"){
      changeErrorVisibility(false)
      changeLoggedIn(true)
      setCookie("loginKey", "ok")
    } else {
      changeErrorVisibility(true)
    }
  };
  if (isLoggedIn) {
    return <Navigate to='/app' />
  }
  return (
    <>
      <div style={styles.center}>
          <h3 text-align="center">Login karo</h3>
          <div style={styles.centerLatest}>
            <input name="Username/Email" type="text" value={email} onChange={handleChange} style={styles.inputCss}/>
            <input name="Password" type="password" value={password} onChange={handlePassChange} style={styles.inputCss}/>
            {ev && <h4 style={{ color: 'red', margin: '2px 0px', fontSize: '14px' }}>Invalid Data. Please retry.</h4>}
          </div>
          <button onClick={handleClick} style={styles.button}> Login </button>
      </div>
    </>
  )
}

const styles = {
  button: {
    maxWidth: "200px",
  },
  inputCss: {
    border: "1px solid #b4975a",
    margin: "0px 20px 20px 20px",
    height: "20px",
    maxWidth: "250px",
  },
  center: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    height: "100vh",
    flexDirection: "column",
    background: 'linear-gradient(to right, #ffffff , #87CEEB)',
  },
  centerLatest: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    flexDirection: "column",
  }
}

export default Login
