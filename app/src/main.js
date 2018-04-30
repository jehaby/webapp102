// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import { createStore } from './store'
import App from './App'
import router from './router'
import apolloProvider from './apollo'

Vue.config.productionTip = false

const store = createStore()

/* eslint-disable no-new */
new Vue({
  el: '#app',
  provide: apolloProvider.provide(),
  // apolloProvider,
  router,
  store,
  template: '<App/>',
  components: { App }
})
