import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export function createStore () {
  return new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    state: {
      user: {},
      errorMsg: ''
    },
    getters: {
      loggedIn: state => {
        return state.user.name !== undefined
      }
    },
    mutations: {
      increment (state) {
        state.count++
      },
      setUser (state, user) {
        state.user = user
      },
      registerFail (state, response) {
        console.log('FAIL!!!', response)
      },
      logout (state) {
        state.user = {}
      },
      error (state, msg) {
        state.errorMsg = msg
      }
    },
    actions: {
      error ({ commit }, msg) {
        // TODO: handle several calls in short period of time
        commit('error', msg)
        setTimeout(() => { commit('error', '') }, 3000)
      }
    }
  })
}
