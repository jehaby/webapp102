import Vue from 'vue'
import Vuex from 'vuex'

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
    }
  }
}
