import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8899/api/v0/'
})

const jwtDecode = require('jwt-decode')

export async function loginRequest (user) {
  let resp = await api.post('/auth/login/', user)
  return jwtDecode(resp.data)
}

export async function registerRequest (user) {
  let resp = await api.post('/auth/register/', user)
  return jwtDecode(resp.data)
}

export function register (user) {
  console.log('registering', user)
}
