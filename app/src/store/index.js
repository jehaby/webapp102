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
      message: {text: '', type: ''}
    },
    mutations: {
      message (state, msg) {
        state.message = msg
      }
    },
    actions: {
      error ({ commit, dispatch }, msg) {
        dispatch('showMsg', {text: msg, type: 'error'})
      },
      showMsg ({ commit }, msg) {
        // TODO: handle several calls in short period of time
        commit('message', msg)
        setTimeout(() => { commit('message', {text: '', type: ''}) }, 3000)
      }
    }
  })
}
