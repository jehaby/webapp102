import {api} from './api.js'

export async function userGetRequest (uuid) {
  return api.get('/users/' + uuid + '/', { withCredentials: true })
}

export async function userUpdateRequest (uuid, data) {
  return api.post('/users/' + uuid + '/update/', data, { withCredentials: true })
}
