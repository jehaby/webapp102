import Vue from 'vue'
import Vuex from 'vuex'
import { loginRequest, logoutRequest } from '@/api/auth.js'

const LOGIN_REQUEST = 'LOGIN_REQUEST'
const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
const LOGOUT = 'LOGOUT'

Vue.use(Vuex)

export const auth = {
  state: {
    auth: null,
    pending: false
  },
  getters: {
    loggedIn: state => {
      return !!state.auth
    }
  },
  mutations: {
    [LOGIN_REQUEST] (state) {
      state.pending = true
    },
    [LOGIN_SUCCESS] (state, auth) {
      state.auth = auth
      state.pending = false
    },
    [LOGOUT] (state) {
      state.auth = null
    }
  },
  actions: {
    async login ({ commit }, user) {
      commit(LOGIN_REQUEST) // show spinner
      const resp = await loginRequest(user)
      const auth = { user: resp.user, exp: resp.exp }
      commit(LOGIN_SUCCESS, auth)
      localStorage.setItem('auth', JSON.stringify(auth))
    },
    async logout ({ commit, dispatch }) {
      console.log('in logout')
      try {
        await logoutRequest()
        localStorage.removeItem('auth')
        commit(LOGOUT)
      } catch (e) {
        // TODO: think here; logout frontend maybe?
        dispatch('error', 'error happened')
        console.log('got error doing logout request ', e)
      }
    }
  }
}
