import logo from './logo.svg';
import './App.css';
import React from "react";
import {Router} from "react-router-dom";
import API from "./Infrastructure/network"
async function doUpload(e) {
  var formData = new FormData();
  formData.append("file", e.target.files[0]);
  let headers = {
    'Content-Type': 'multipart/form-data'
  }

  await API.post("upload", formData, {headers})
}

function App() {



  return (

    <div className="App">
      <form id="uploadForm" encType="multipart/form-data"  change="uploadFile">
        <input type="file" id="file" name="file" onChange={doUpload} />
      </form>
    </div>
  );
}

export default App;
