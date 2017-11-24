import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Login from '@/components/Login'
import Register from '@/components/Register'
import Profile from '@/components/Profile'
import CreateAd from '@/components/CreateAd'
import ViewAd from '@/components/ViewAd'

Vue.use(Router)

export default new Router({
//  mode: 'history',
//  fallback: false,
//  scrollBehavior: () => ({ y: 0 }),
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
