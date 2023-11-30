import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import {setCookie, getCookie} from "./Utils";
import { Navigate } from 'react-router-dom';

function App() {
  const [isloggedout, changeLoggedOut] = useState(getCookie("loginKey") == "")
  const [count, setCount] = useState(0)
  const handleClick = () => {
    setCookie("loginKey", "")
    changeLoggedOut(true)
  };
  if(isloggedout){
    return <Navigate to='/' />
  }
  return (
    <>
      <div>
        <div>
          <button onClick={handleClick} style={{margin:"5px", background:"#cdef00"}}> Log out </button>
        </div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
