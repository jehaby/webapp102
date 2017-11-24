import Vue from 'vue'
import Vuex from 'vuex'
import { auth } from './auth.js'

Vue.use(Vuex)

export function createStore () {
  return new Vuex.Store({
    modules: {
      auth: auth
    },
    strict: process.env.NODE_ENV !== 'production',
    state: {
      errorMsg: ''
    },
    mutations: {
      error (state, msg) {
        state.errorMsg = msg
      }
    },
    actions: {
      error ({ commit }, msg) {
        console.log('dispatching error: ', msg)
        // TODO: handle several calls in short period of time
        commit('error', msg)
        setTimeout(() => { commit('error', '') }, 3000)
      }
    }
  })
}
