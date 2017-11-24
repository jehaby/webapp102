import axios from 'axios'

export const api = axios.create({
  baseURL: 'http://localhost:8899/api/v0/'
})
