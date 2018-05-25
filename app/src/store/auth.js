import Vue from 'vue'
import Vuex from 'vuex'
import { loginRequest } from '@/api/auth.js'

const LOGIN_REQUEST = 'LOGIN_REQUEST'
const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
const LOGOUT = 'LOGOUT'

Vue.use(Vuex)

export const auth = {
  state: {
    user: null,
    pending: false
  },
  getters: {
    loggedIn: state => {
      return !!state.user
    }
  },
  mutations: {
    [LOGIN_REQUEST] (state) {
      state.pending = true
    },
    [LOGIN_SUCCESS] (state, user) {
      state.user = user
      state.pending = false
    },
    [LOGOUT] (state) {
      state.user = null
    }
  },
  actions: {
    async login ({ commit }, user) {
      commit(LOGIN_REQUEST) // show spinner
      const resp = await loginRequest(user)
      commit(LOGIN_SUCCESS, resp.user)
      localStorage.setItem('user', JSON.stringify(resp.user))
    },
    async logout ({ commit }) {
      console.log('in logout')
      // TODO: send request to invalidate token
      localStorage.removeItem('user')
      commit(LOGOUT)
    }
  }
}
