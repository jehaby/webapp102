import {api} from './api.js'

export async function userGetRequest (uuid) {
  return api.get(userPath(uuid), { withCredentials: true })
}

export async function userUpdateRequest (uuid, data) {
  return api.put(userPath(uuid), data, { withCredentials: true })
}

export async function phoneDeleteRequest (userUUID, phoneUUID) {
  return api.delete(userPath(userUUID) + 'phones/' + phoneUUID + '/', { withCredentials: true })
}

export async function phoneCreateRequest (data) {
  return api.post(userPath(data.user_uuid) + 'phones/', data, { withCredentials: true })
}

function userPath (uuid) {
  return '/users/' + uuid + '/'
}
