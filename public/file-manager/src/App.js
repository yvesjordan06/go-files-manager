import logo from './logo.svg';
import './App.css';
import React, {useEffect, useState} from "react";
import {Route, Switch} from 'react-router'
import API, {setToken} from "./Infrastructure/network"
import LoginPage from "./Page/Auth/login";
import GuardedRoute from "./Components/GuardedRoute";
import Dashboard from "./Page/Document/home";

function App() {



  return (

    <div className="App">


      <Switch >

      <Route path={"/auth/login/"} exact>
          <LoginPage/>
      </Route>

          <Switch location={'/'}>
        <GuardedRoute path={"/"}>
            <Dashboard/>
        </GuardedRoute>
          </Switch>

        <Route >
            <h1>Not Found</h1>
        </Route>

      </Switch>
    </div>
  );
}

export default App;
