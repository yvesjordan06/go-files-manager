import axios from "axios";


let baseURL

if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
  baseURL = 'http://localhost:7777/api'
} else {
  baseURL = '/api'
}


const API = axios.create(
    {
        baseURL,

    }
)
function setToken(token) {
  console.log("Setting token "+token)
  API.defaults.headers.common = {...API.defaults.headers.common, 'Authorization': `Bearer ${token}`}
}


function unsetToken() {
  console.log("Unsetting token ")
  API.defaults.headers.common = {...API.defaults.headers.common, 'Authorization': undefined}
}
export default API
export {setToken}

