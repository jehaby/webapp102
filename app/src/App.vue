<template>
  <div id="app" class="section">
    <!--TODO: move somewhere to 'router' folder -->
    <!--TODO: fix home always active -->
    <nav class="navbar" role="navigation" aria-label="main navigation">

      <div class="navbar-brand">
        <router-link class="navbar-item" to="/">main</router-link>
      </div>

      <nav-menu></nav-menu>

    </nav>

    <error-message :message="errorMsg"></error-message>    
    
    <!--router view shows content--> 
    <div class="container">
      <router-view></router-view>
    </div>

    <footer class="footer">
      <div class="container">
        <div class="content has-text-centered">
          <p>
            <strong>webapp102</strong>
          </p>
        </div>
      </div>
    </footer>    

  </div>
</template>

<script>
import ErrorMessage from './components/ErrorMessage'
import NavMenu from './NavMenu'
import { refreshToken } from '@/api/auth.js'

require('./assets/sass/main.scss')

export default {
  components: { ErrorMessage, NavMenu },
  name: 'app',
  computed: {
    errorMsg () {
      return this.$store.state.errorMsg
    },
    loggedIn () {
      return this.$store.getters.loggedIn
    }
  },
  methods: {
    async logout () {
      await this.$store.dispatch('logout')
      this.$router.push('/')
    },

    /**
     * Checks if 'auth' key exists in local storage; if true, then logs in user
     * and start routine which updates jwt token if necessary.
     */
    checkAuth () {
      if (!localStorage['auth']) {
        return
      }

      const auth = JSON.parse(localStorage['auth'])
      const exp = new Date(auth.exp)

      if ((exp - (new Date())) > tenMinutes) {
        this.$store.commit('LOGIN_SUCCESS', auth)
        this.startRefreshJwtRoutine(exp)
      }
    },

    async startRefreshJwtRoutine (exp) {
      const checkForRefresh = async function (exp) {
        if ((exp - (new Date())) < oneDay) {
          try {
            const resp = await refreshToken()
            console.log(resp)
          } catch (e) {
            // TODO: better err handling
            console.log('caught error in refresh token', e)
          }
        }
      }

      await checkForRefresh(exp)
      setInterval(async function () {
        await checkForRefresh(exp)
      }, tenMinutes)
    }

  },

  created () {
    this.checkAuth()
  }
}

const oneDay = 1000 * 60 * 60 * 24
const tenMinutes = 1000 * 60 * 10
</script>

<style>

</style>
