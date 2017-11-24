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
    setUser (state, user) {
      state.user = user
    },
    setJwtToken (state, token) {
      state.token = token
    }
  }
}
