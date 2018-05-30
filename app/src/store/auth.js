import Vue from 'vue'
import Vuex from 'vuex'
import { loginRequest, logoutRequest } from '@/api/auth.js'

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
    async logout ({ commit, dispatch }) {
      console.log('in logout')
      try {
        await logoutRequest()
        localStorage.removeItem('user')
        commit(LOGOUT)
      } catch (e) {
        dispatch('error', 'error happened')
        console.log('got error doing logout request ', e)
      }
    }
  }
}
