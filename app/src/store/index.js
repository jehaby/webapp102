import Vue from 'vue'
import Vuex from 'vuex'
import { loginRequest } from './../api/index.js'

Vue.use(Vuex)

export function createStore () {
  return new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    state: {
      user: {},
      errorMsg: ''
    },
    mutations: {
      increment (state) {
        state.count++
      },
      loginSuccess (state, user) {
        console.log('success login', user)
        state.user = user
        this.$router.push('/profile')
      },
      registerSuccess (state, response) {
        console.log('success login', response)
      },
      registerFail (state, response) {
        console.log('FAIL!!!', response)
      },
      error (state, msg) {
        state.errorMsg = msg
      }
    },
    actions: {
      login ({ commit, state, dispatch }, user) {
        loginRequest(
          user,
          (user) => commit('loginSuccess', user),
          (msg) => dispatch('error', msg)
        )
      },
      error ({ commit }, msg) {
        // TODO: handle several calls in short period of time
        commit('error', msg)
        setTimeout(() => { commit('error', '') }, 3000)
      }
    }
  })
}
