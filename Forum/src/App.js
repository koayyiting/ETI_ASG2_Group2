
import { useState } from 'react';
// import './App.css';
import Login from './Components/Loginpage/Login';
import Posts from './Components/Posts_page/Posts';
function App() {
  const [email_on,set_email_on]=useState(null);
  const [name_on,set_name_on]=useState(null);
  return (
    <div className="">
      {email_on===null && name_on===null &&  <Login set_email_on={set_email_on} set_name_on={set_name_on}  /> }
      {email_on!==null &&name_on !==null && <Posts email_on={email_on} name_on={name_on} />}
    </div>
  );
}

export default App;
