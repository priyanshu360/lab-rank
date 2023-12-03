import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';

function Register() {
  const [email, setEmail] = useState('test@user.in');
  const [password, setPassword] = useState('1234');
  const [universityId, setUniversityId] = useState('e71c2a34-974f-42a6-a7f3-e533c7597f12');
  const [collegeId, setCollegeId] = useState('ec0835c8-c768-423b-8cb4-e71a7bab0a5c');
  const [access, setAccess] = useState('010919a4-b154-406a-b0d1-b35a47870107');
  const [contact, setContact] = useState('1234567890');
  const [userName, setUserName] = useState('Priyanshu Rajput');
  const [dob, setDob] = useState('2000-01-01');
  const [apiResp, setApiResp] = useState(null);
  const [error, setError] = useState(false);
  const navigate = useNavigate();
  useEffect(() => { if(apiResp)navigate('/auth/login') },[apiResp]);


  const handleClick = async () => {
    const newDate = dob + "T00:00:00Z";
    fetch('http://127.0.0.1:8080/user',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        "college_id": collegeId,
        "access_id": access,
        "email": email,
        "contact_no": contact,
        "dob": newDate,
        "university_id": universityId,
        "name": userName,
        "password": password
      })
    })
    .then(response => {
      console.log("Response raw", response.status == 200)
      if(response.status != 200)setError(true)
      return response.json()
    })
    .then(json => {
      console.log("API Response", /*response,*/ json)
      setApiResp(json)
      setError(false)
    })
    .catch(error => console.error(error));
  };

  return (
    <>
      <div style={styles.center}>
          <h3 text-align="center">Register User</h3>
          <div style={styles.centerLatest}>
            {/* <label for="college_id">Enter College ID</label> */}
            <EditTextWithLabel name="University ID" placeholder="Enter University ID" type="text" value={universityId} onChangeValue={(event)=>{setUniversityId(event.target.value)}}/>
            <EditTextWithLabel name="College ID" placeholder="Enter College ID" type="text" value={collegeId} onChangeValue={(event)=>{setCollegeId(event.target.value)}}/>
            <EditTextWithLabel name="Access ID" placeholder="Enter Access ID" type="text" value={access} onChangeValue={(event)=>{setAccess(event.target.value)}}/>
            <EditTextWithLabel name="Email" placeholder="Enter Email" type="text" value={email} onChangeValue={(event) => {setEmail(event.target.value)}}/>
            <EditTextWithLabel name="Password" placeholder="Enter Password" type="password" value={password} onChangeValue={(event) => {setPassword(event.target.value)}}/>
            <EditTextWithLabel name="Contact Number" placeholder="Contact Number" type="text" value={contact} onChangeValue={(event)=>{setContact(event.target.value)}}/>
            <EditTextWithLabel name="Name" placeholder="Enter Name" type="text" value={userName} onChangeValue={(event)=>{setUserName(event.target.value)}}/>
            <EditTextWithLabel name="DOB" placeholder="Date of Birth" type="date" value={dob} onChangeValue={(event)=>{setDob(event.target.value)}}/>
            {error && <p style={{color:'red'}}>Something Went Wrong</p>}
          </div>
          <button onClick={handleClick} style={styles.button}> Register Now </button>
      </div>
    </>
  )
}

const EditTextWithLabel = function({name, placeholder, type, value, onChangeValue}){
  return <input name={name} placeholder={placeholder} type={type} value={value} onChange={onChangeValue} style={styles.inputCss}/>
}


const styles = {
  button: {
    width: "100px",
    height: "32px",
    backgroundColor: "#04AA6D",
    border: "none",
    color: "white",
    borderRadius: "2px",
  },
  inputCss: {
    border: "1px solid #b4975a",
    margin: "0px 20px 20px 20px",
    padding: "0px 0px 0px 10px",
    height: "25px",
    width: "400px",
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

export default Register
