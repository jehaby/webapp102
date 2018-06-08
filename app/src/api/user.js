import {api} from './api.js'

export async function userGetRequest (uuid) {
  return api.get(userPath(uuid), { withCredentials: true })
}

export async function userUpdateRequest (uuid, data) {
  return api.put(userPath(uuid), data, { withCredentials: true })
}

function userPath (uuid) {
  return '/users/' + uuid + '/'
}
