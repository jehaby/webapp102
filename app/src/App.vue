<template>
  <div id="app" class="section">
    <!--TODO: move somewhere to 'router' folder -->
    <!--TODO: fix home always active -->
    <nav class="navbar" role="navigation" aria-label="main navigation">

      <div class="navbar-brand">
        <router-link class="navbar-item" to="/">main</router-link>
      </div>

      <div class="navbar-menu">
        <div class="navbar-end">
        <template v-if="loggedIn">
          <router-link class="navbar-item" to="/ads/my">my ads</router-link>          
          <router-link class="navbar-item" to="/profile">profile</router-link>
          <router-link class="navbar-item" to="/ads/create">create ad</router-link>
          <button class="button navbar-item" v-on:click="logout">logout</button>
        </template>
        <template v-else>
          <router-link class="navbar-item" to="/login">login</router-link>
          <router-link class="navbar-item" to="/register">register</router-link>
        </template>
        </div>
      </div>          
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

require('./assets/sass/main.scss')

export default {
  components: { ErrorMessage },
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
    logout () {
      console.log(this.loggedIn, this.$store.state.user)
      this.$store.commit('user', {})
      this.$router.push('/')
    }
  }
}
</script>

<style>

</style>
