import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8899/api/v0/'
})

const jwtDecode = require('jwt-decode')

export function getUser () {
  api.get()
    .then(response => {
      console.log(response)
    })
    .catch(response => {
      console.log(response)
    })
}

export function loginRequest (user) {
  console.log('logging in request', user)
  api.post('/auth/login/', user)
    .then(response => {
      console.log(jwtDecode(response.data))
      // TODO: set session
      console.log(response)
    })
    .catch(response => {
      console.log(response)
    })
}

export function registerRequest (user) {
  console.log('logging in request', user)
  api.post('/auth/register/', {
    name: user.name,
    email: user.email,
    password: user.password
  })
    .then(response => {
      console.log(jwtDecode(response.data))
      console.log(response)
    })
    .catch(response => {
      console.log(response)
    })
}

export function register (user) {
  console.log('registering', user)
}
