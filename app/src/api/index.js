import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8899'
})

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
  console.log('logging in', user)
}

export function register (user) {
  console.log('registering', user)
}
