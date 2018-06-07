import {api} from './api.js'

export async function userGetRequest (uuid) {
  return api.get('/user/' + uuid + '/', { withCredentials: true })
}

export async function userUpdateRequest (uuid, data) {
  return api.post('/user/' + uuid + '/update/', data, { withCredentials: true })
}
