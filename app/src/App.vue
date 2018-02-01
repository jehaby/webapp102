<template>
  <div class="container" id="app">
    <!--TODO: move somewhere to 'router' folder -->
    <!--TODO: fix home always active -->
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">

          <router-link class="navbar-item" to="/">home</router-link>

          <template v-if="loggedIn">
            <router-link class="navbar-item" to="/profile">profile</router-link>
            <router-link class="navbar-item" to="/ads/create">create ad</router-link>
            <button class="button" v-on:click="logout">Log out</button>
          </template>
          <template v-else>
            <router-link class="navbar-item" to="/login">login</router-link>
            <router-link class="navbar-item" to="/register">register</router-link>
          </template>

      </div>
    </nav>
    <!--router view shows content--> 
    <router-view></router-view>
    <error-message :message="errorMsg"></error-message>

  </div>
</template>

<script>
  import ErrorMessage from './components/ErrorMessage'

  require('./assets/sass/main.scss')

  export default {
    components: {ErrorMessage},
    name: 'app',
    computed: {
      errorMsg () { return this.$store.state.errorMsg },
      loggedIn () { return this.$store.getters.loggedIn }
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
