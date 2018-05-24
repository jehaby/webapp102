import Vue from 'vue'
import Vuex from 'vuex'
import { loginRequest } from '@/api/auth.js'

const LOGIN_REQUEST = 'LOGIN_REQUEST'
const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
const LOGOUT = 'LOGOUT'

Vue.use(Vuex)

export const auth = {
  state: {
    jwtToken: '',
    user: {}
  },
  getters: {
    loggedIn: state => {
      return state.user.name !== undefined
    }
  },
  mutations: {
    user (state, user) {
      state.user = user
    },
    jwtToken (state, token) {
      state.jwtToken = token
    },

    [LOGIN_REQUEST] (state) {
      state.pending = true
    },
    [LOGIN_SUCCESS] (state) {
      state.isLoggedIn = true
      state.pending = false
    },
    [LOGOUT] (state) {
      state.isLoggedIn = false
    }
  },
  actions: {
    async login ({ commit }, creds) {
      commit(LOGIN_REQUEST) // show spinner

      try {
        let resp = await loginRequest({ ...this.user })
        console.log(resp)
      } catch (e) {
        console.log('error in auth/login action: ', e)
      }

      return new Promise(resolve => {
        setTimeout(() => {
          localStorage.setItem('token', 'JWT')
          commit(LOGIN_SUCCESS)
          resolve()
        }, 1000)
      })
    },
    logout ({ commit }) {
      localStorage.removeItem('token')
      commit(LOGOUT)
    }
  }
}
