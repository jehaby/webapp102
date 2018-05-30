import {api} from './api.js'

const jwtDecode = require('jwt-decode')

function UserException (message) {
  this.message = message
  this.name = 'UserException'
}

async function requestWithUser (method, data) {
  const resp = await api.post(method, data, {
    withCredentials: true
  })
  const decoded = jwtDecode(resp.data)
  if (decoded.user === undefined) {
    throw new UserException('no user in response: ' + decoded)
  }
  return {'user': decoded.user, 'token': resp.data}
}

export async function loginRequest (user) {
  return requestWithUser('/auth/login/', user)
}

export function registerRequest (user) {
  return requestWithUser('/auth/register/', user)
}

export async function logoutRequest () {
  return api.get('/auth/logout/', {
    withCredentials: true
  })
}
