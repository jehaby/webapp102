import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Login from '@/components/auth/Login'
import Register from '@/components/auth/Register'
import Profile from '@/components/Profile'
import CreateAd from '@/components/ad/CreateAd'
import ViewAd from '@/components/ad/ViewAd'

Vue.use(Router)

export default new Router({
//  mode: 'history',
//  fallback: false,
//  scrollBehavior: () => ({ y: 0 }),
  linkActiveClass: 'is-active',
  routes: [
    {
      path: '/login',
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
      name: 'Home',
      component: Home
    },
    {
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/ads/create',
      name: 'AdsCreate',
      component: CreateAd
    },
    {
      path: '/ads/:uuid',
      name: 'AdsView',
      component: ViewAd
    }

  ]
})
