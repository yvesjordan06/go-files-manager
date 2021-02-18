import logo from './logo.svg';
import './App.css';
import React, {useState} from "react";
import {Route, Switch} from 'react-router'
import API, {setToken} from "./Infrastructure/network"
import LoginPage from "./Page/Auth/login";
import GuardedRoute from "./Components/GuardedRoute";
import Dashboard from "./Page/Document/home";

function App() {

  async function doUpload(e) {
    var formData = new FormData();
    setToken("$2a$10$oTCydB01hCuDl6Dx86oTVe1RF23s.j3Kn0BBOl3XB3Bmhy/VS8Hxa")
    formData.append("file", e.target.files[0]);
    let headers = {
      'Content-Type': 'multipart/form-data'
    }

    await API.post("files/upload", formData, {onUploadProgress: progressEvent =>{
        console.log(progressEvent)
      setProgress(
        parseFloat(((progressEvent.loaded/progressEvent.total)* 100).toFixed(2))
      )
    }})
  }


  let [progress, setProgress] = useState(0.0)


  return (

    <div className="App">
    {/*  <p>Progress {progress} %</p>
      <form id="uploadForm" encType="multipart/form-data"  change="uploadFile">
        <input type="file" id="file" name="file" onChange={doUpload} />
      </form>*/}

      <Switch >

      <Route path={"/auth/login/"} exact>
          <LoginPage/>
      </Route>

        <GuardedRoute path={"/"} exact>
            <Dashboard/>
        </GuardedRoute>

        <Route >
            <h1>Not Found</h1>
        </Route>

      </Switch>
    </div>
  );
}

export default App;
