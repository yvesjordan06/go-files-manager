import React  from "react";
import {Redirect, Route} from "react-router";
import PropTypes from 'prop-types';

function GuardedRoute ({children, ...rest}){
    return (
        <Route {...rest} render={(props) => (
            !!localStorage.getItem("token") ? <React.Fragment children={children} {...props}/> : <Redirect to={"auth/login" } />
        )} />
    )
}

GuardedRoute.propTypes = {
    path: PropTypes.string.isRequired,
    exact: PropTypes.bool
}



export default GuardedRoute