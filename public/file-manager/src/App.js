import logo from './logo.svg';
import './App.css';
import React, {useState} from "react";
import {Router} from "react-router-dom";
import API, {setToken} from "./Infrastructure/network"

function App() {

  async function doUpload(e) {
    var formData = new FormData();
    setToken("$2a$10$rhNYojMa9s5R2NRLE33Tg.DXM/VnZcH/62E1nRIRKsxuZXi7VqTw.")
    formData.append("file", e.target.files[0]);
    let headers = {
      'Content-Type': 'multipart/form-data'
    }

    await API.post("upload", formData, {onUploadProgress: progressEvent =>{
        console.log(progressEvent)
      setProgress(
        parseFloat(((progressEvent.loaded/progressEvent.total)* 100).toFixed(2))
      )
    }})
  }


  let [progress, setProgress] = useState(0.0)


  return (

    <div className="App">
      <p>Progress {progress} %</p>
      <form id="uploadForm" encType="multipart/form-data"  change="uploadFile">
        <input type="file" id="file" name="file" onChange={doUpload} />
      </form>
    </div>
  );
}

export default App;
