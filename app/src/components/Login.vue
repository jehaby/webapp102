<template>
  <div>
    <h5> login </h5>
    <form>
      <p>
        name
        <input v-model="user.name" type="text"/>
      </p>
      <p>
        password
        <input v-model="user.password" type="password"/>
      </p>
      <p>
        <input type="button" v-on:click="login" value="login"/>
      </p>
    </form>
  </div>
</template>

<script>
  import { loginRequest } from './../api/auth.js'

  export default {
    name: 'Login',
    data () {
      return {
        user: {
          name: '',
          password: ''
        }
      }
    },
    methods: {
      async login () {
        // TODO: form validation
        let resp = {}
        try {
          resp = await loginRequest({...this.user})
        } catch (e) {
          // TODO: better errors
          return this.$store.dispatch('error', 'Login failed')
        }
        this.$store.commit('setUser', resp.user)
        this.$store.commit('setJwtToken', resp.token)
        this.$router.push('/profile')
      }
    }
  }
</script>

<style scoped>

</style>
