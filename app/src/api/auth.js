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
  if (decoded.exp === undefined) {
    throw new UserException('no exp in response: ' + decoded)
  }

  return {'user': decoded.user, 'exp': decoded.exp * 1000, 'token': resp.data}
}

export async function loginRequest (user) {
  return requestWithUser('/auth/login/', user)
}

export function registerRequest (user) {
  return requestWithUser('/auth/register/', user)
}

export async function logoutRequest () {
  return api.get('/auth/logout/', { withCredentials: true })
}

export async function refreshToken () {
  return api.get('/auth/refresh/', { withCredentials: true })
}

export async function passwordResetRequest (data) {
  return api.post('/auth/resetRequest/', data)
}

export async function passwordResetAction (data) {
  return requestWithUser('/auth/resetAction/', data)
}
