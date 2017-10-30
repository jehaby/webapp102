import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Login from '@/components/Login'
import Register from '@/components/Register'

Vue.use(Router)

export default new Router({
//  mode: 'history',
//  fallback: false,
//  scrollBehavior: () => ({ y: 0 }),
  routes: [
    {
      path: '/loginRequest/',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/',
      name: 'Hello',
      component: HelloWorld
    }
  ]
})
