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
    }
  },
  created () {
    if (localStorage['user']) {
      this.$store.commit('LOGIN_SUCCESS', JSON.parse(localStorage['user']))
    }
  }
}
</script>

<style>

</style>
