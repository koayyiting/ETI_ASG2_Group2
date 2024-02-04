// Login.js

import React, { useState } from "react";
import "./Login.css";
import axios from "axios";

const Login = (props) => {
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");

  const handleLogin = async () => {
    // Perform login logic here
    if (email === "" || name === "") {
      alert("fill all the fields");
      return;
    }

   

    const userData = {
      Email: email,
      Username: name,
    };

    // Use Axios to send a POST request
    axios
      .post("http://localhost:4090/createUser", userData)
      .then((response) => {
        // Handle the success response if needed
        console.log("Post request successful:", response.data);
      })
      .catch((error) => {
        // Handle the error if needed
        console.error("Error in post request:", error);
      });

    props.set_email_on(email);
    props.set_name_on(name);
    console.log("Email:", email);
    console.log("Name:", name);
    // Add your authentication logic or API calls here
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2>Login</h2>
        <div className="input-group">
          <label>Email:</label>
          <input
            type="email"
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="input-group">
          <label>Name:</label>
          <input
            type="text"
            placeholder="Enter your name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <button onClick={handleLogin}>Login</button>
      </div>
    </div>
  );
};

export default Login;
