import {api} from './api.js'

export async function createAd (ad) {
  return api.post('/ads/', ad)
}

export async function getAd (uuid) {
  return api.get('/ads/' + uuid)
}
