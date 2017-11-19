<template>
  <div id="app">
      <p>
        <router-link to="/">home</router-link>

        <template v-if="loggedIn">
          <router-link to="/profile">profile</router-link>
          <button v-on:click="logout">Log out</button>
        </template>
        <template v-else>
          <router-link to="/login">login</router-link>
          <router-link to="/register">register</router-link>
        </template>
      </p>
    <error-message :message="errorMsg"></error-message>
    <router-view/>

  </div>
</template>

<script>
  import ErrorMessage from './components/ErrorMessage'

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
        this.$store.commit('setUser', {})
        this.$router.push('/')
      }
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
