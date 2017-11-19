import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8899/api/v0/'
})

const jwtDecode = require('jwt-decode')

async function requestAndDecode (method, data) {
  const resp = await api.post(method, data)
  return jwtDecode(resp.data)
}

export async function loginRequest (user) {
  return requestAndDecode('/auth/login', user)
}

export async function registerRequest (user) {
  return requestAndDecode('/auth/login', user)
}

export function register (user) {
  console.log('registering', user)
}
