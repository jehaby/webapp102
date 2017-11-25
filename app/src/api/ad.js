import {api} from './api.js'

export async function createAd (ad, jwtToken) {
  console.log('post createAd: ', jwtToken)
  return api.post('/ads/', ad, {
    headers: {
      'Authorization': 'BEARER: ' + jwtToken
    }
  })
}

export async function getAd (uuid) {
  return api.get('/ads/' + uuid)
}
