import { useEffect, useState } from 'react'

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

  useEffect(()=>{console.log(apiResp)},[apiResp]);

  const handleClick = async () => {
    console.log("Register API Called")
    const newDate = dob + "T00:00:00Z";
      const response = await fetch("http://127.0.0.1:8080/user", {
        mode: "no-cors",
        method: "POST",
        headers: {
          "Content-Type": "application/json",
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
      console.log(response)
    // var urlencoded = new URLSearchParams();
    // urlencoded.append("college_id", collegeId);
    // urlencoded.append("access_id", access);
    // urlencoded.append("email", email);
    // urlencoded.append("contact_no", contact);
    // urlencoded.append("dob", newDate);
    // urlencoded.append("university_id", universityId);
    // urlencoded.append("name", userName);
    // urlencoded.append("password", password);
    // fetch('https://fakestoreapi.com/products',{
    //         method:"POST",
    //         body:JSON.stringify(
    //             {
    //                 title: email,
    //                 price: 13.5,
    //                 description: 'lorem ipsum set',
    //                 image: 'https://i.pravatar.cc',
    //                 category: 'electronic'
    //             }
    //         )
    //     })
    //         .then(res=>res.json())
    //         .then(json=>console.log(json))
    fetch('http://127.0.0.1:8080/user',{
      method: 'POST',
      mode: 'no-cors',
      headers: {
        // 'Accept': 'application/json',
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
    .then(response => {console.log(response);return response.json()})
    .then(json => setApiResp(json.body))
    .catch(error => console.error(error));
  };
  
  console.log(apiResp);

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
