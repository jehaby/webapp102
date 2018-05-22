import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/auth/Login'
import Register from '@/components/auth/Register'
import Profile from '@/components/Profile'

import ListAds from '@/components/ad/ListAds'
import MyAds from '@/components/ad/MyAds'
import CreateAd from '@/components/ad/CreateAd'
import EditAd from '@/components/ad/EditAd'
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
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/ads/view/:uuid',
      name: 'AdsView',
      component: ViewAd
    },
    {
      path: '/ads/edit/:uuid/',
      name: 'AdsEdit',
      component: EditAd
    },
    {
      path: '/ads/create',
      name: 'AdsCreate',
      component: CreateAd
    },
    {
      // path: '/ads/my',
      path: '/death404',
      name: 'MyAds',
      component: MyAds
    },
    {
      path: '/',
      name: 'Home',
      component: ListAds
    }

  ]
})
